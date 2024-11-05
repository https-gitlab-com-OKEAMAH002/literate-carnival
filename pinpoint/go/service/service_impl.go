package service

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
	"go.skia.org/infra/pinpoint/go/read_values"
	"go.skia.org/infra/pinpoint/go/workflows"
	pb "go.skia.org/infra/pinpoint/proto/v1"
)

type server struct {
	pb.UnimplementedPinpointServer

	// Local rate limiter to only limit the traffic for migration temporarilly.
	limiter *rate.Limiter

	temporal TemporalProvider
}

const (
	// Those should be configurable for each instance.
	hostPort  = "temporal.temporal:7233"
	namespace = "perf-internal"
	taskQueue = "perf.perf-chrome-public.bisect"
)

type TemporalProvider interface {
	// NewClient returns a Temporal Client and a clean up function
	NewClient() (client.Client, func(), error)
}

type defaultTemporalProvider struct {
}

// NewClient implements TemporalProvider.NewClient
func (defaultTemporalProvider) NewClient() (client.Client, func(), error) {
	c, err := client.NewLazyClient(client.Options{
		HostPort:  hostPort,
		Namespace: namespace,
	})

	if err != nil {
		return nil, nil, skerr.Wrap(err)
	}
	return c, func() {
		c.Close()
	}, nil
}

func New(t TemporalProvider, l *rate.Limiter) *server {
	if l == nil {
		// 1 token every 30 minutes, this allow some buffer to drain the hot spots in the bots pool.
		l = rate.NewLimiter(rate.Every(30*time.Minute), 1)
	}
	if t == nil {
		t = defaultTemporalProvider{}
	}
	return &server{
		limiter:  l,
		temporal: t,
	}
}

// updateFieldsForCatapult converts specific catapult Pinpoint arguments
// to their skia Pinpoint counterparts
func updateFieldsForCatapult(req *pb.ScheduleBisectRequest) *pb.ScheduleBisectRequest {
	if req.Statistic != "" {
		req.AggregationMethod = req.Statistic
	}
	return req
}

func validate(req *pb.ScheduleBisectRequest) error {
	switch {
	case req.StartGitHash == "" || req.EndGitHash == "":
		return skerr.Fmt("git hash is empty")
	case !read_values.IsSupportedAggregation(req.AggregationMethod):
		return skerr.Fmt("aggregation method (%s) is not available", req.AggregationMethod)
	default:
		return nil
	}
}

func NewJSONHandler(ctx context.Context, srv pb.PinpointServer) (http.Handler, error) {
	m := runtime.NewServeMux()
	if err := pb.RegisterPinpointHandlerServer(ctx, m, srv); err != nil {
		return nil, skerr.Wrapf(err, "unable to register pinpoint handler")
	}
	return m, nil
}

func (s *server) ScheduleBisection(ctx context.Context, req *pb.ScheduleBisectRequest) (*pb.BisectExecution, error) {
	// Those logs are used to test traffic from existing services in catapult, shall be removed.
	sklog.Infof("Receiving bisection request: %v", req)
	if !s.limiter.Allow() {
		sklog.Infof("The request is dropped due to rate limiting.")
		return nil, skerr.Fmt("unable to fulfill the request due to rate limiting, dropping")
	}

	// TODO(b/318864009): Remove this function once Pinpoint migration is completed.
	req = updateFieldsForCatapult(req)

	if err := validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	c, cleanUp, err := s.temporal.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to connect to Temporal (%v).", err)
	}

	defer cleanUp()

	wo := client.StartWorkflowOptions{
		ID:                       uuid.New().String(),
		TaskQueue:                taskQueue,
		WorkflowExecutionTimeout: 12 * time.Hour, // arbitrary timeout to assure it's not going forever.
		RetryPolicy: &temporal.RetryPolicy{
			// We will only attempt to run the workflow exactly once as we expect any failure will be
			// not retry-recoverable failure.
			MaximumAttempts: 1,
		},
	}
	wf, err := c.ExecuteWorkflow(ctx, wo, workflows.CatapultBisect, &workflows.BisectParams{
		Request:    req,
		Production: true,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to start workflow (%v).", err)
	}
	return &pb.BisectExecution{JobId: wf.GetID()}, nil
}

// TODO(b/322047067)
//	embbed pb.UnimplementedPinpointServer will throw errors if those are not implemented.
// func (s *server) QueryBisection(ctx context.Context, in *pb.QueryBisectRequest) (*pb.BisectExecution, error) {
// }

func (s *server) LegacyJobQuery(ctx context.Context, req *pb.LegacyJobRequest) (*pb.LegacyJobResponse, error) {
	qresp, err := s.QueryBisection(ctx, &pb.QueryBisectRequest{
		JobId: req.GetJobId(),
	})
	if err != nil {
		// We don't skerr.Wrap here because we expect to populate err with status.code back to
		// the client, this is automatic conversion to REST API status code when this is exposed
		// via grpc-gateway.
		// Note this API is only intermediate and will be gone, this is not considered to be
		// best practise.
		return nil, err
	}

	// TODO(b/318864009): convert BisectExecution to LegacyJobResponse
	// This should be just a copy.
	resp := &pb.LegacyJobResponse{
		JobId: qresp.GetJobId(),
	}
	return resp, nil
}

func (s *server) CancelJob(ctx context.Context, req *pb.CancelJobRequest) (*pb.CancelJobResponse, error) {
	sklog.Infof("Receiving cancel job request: %v", req)
	if req.JobId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bad request: missing JobId")
	}

	if req.Reason == "" {
		return nil, status.Errorf(codes.InvalidArgument, "bad request: missing Reason")
	}

	c, cleanUp, err := s.temporal.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to connect to Temporal (%v).", err)
	}

	defer cleanUp()

	err = c.CancelWorkflow(ctx, req.JobId, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to cancel workflow (%v).", err)
	}
	return &pb.CancelJobResponse{JobId: req.JobId, State: "Cancelled"}, nil
}
