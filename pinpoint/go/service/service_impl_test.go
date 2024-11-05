package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.skia.org/infra/pinpoint/go/service/mocks"
	"go.skia.org/infra/pinpoint/go/workflows"
	pb "go.skia.org/infra/pinpoint/proto/v1"
	temporal_mocks "go.temporal.io/sdk/mocks"
	"golang.org/x/time/rate"
)

func newTemporalMock(t *testing.T) (*mocks.TemporalProvider, *temporal_mocks.Client) {
	tcm := &temporal_mocks.Client{}
	tcm.Mock.Test(t)

	tpm := mocks.NewTemporalProvider(t)
	t.Cleanup(func() { tcm.AssertExpectations(t) })
	return tpm, tcm
}

func newWorkflowRunMock(t *testing.T, wid string) *temporal_mocks.WorkflowRun {
	wfm := &temporal_mocks.WorkflowRun{}
	wfm.Test(t)
	wfm.On("GetID").Return(wid)
	t.Cleanup(func() { wfm.AssertExpectations(t) })
	return wfm
}

func TestScheduleBisection_ValidRequest_ReturnJobID(t *testing.T) {
	counter := int32(0)
	tpm, tcm := newTemporalMock(t)
	tpm.On("NewClient").Return(tcm, func() {
		counter += 1
	}, nil)

	const fakeID = "fake-job-id"
	wfm := newWorkflowRunMock(t, fakeID)
	tcm.On("ExecuteWorkflow", mock.Anything, mock.Anything, workflows.CatapultBisect, mock.Anything).Return(wfm, nil)

	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Inf, 0))

	resp, err := svc.ScheduleBisection(ctx, &pb.ScheduleBisectRequest{
		StartGitHash: "fake-start",
		EndGitHash:   "fake-end",
	})
	assert.Equal(t, fakeID, resp.JobId)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, counter, "CleanUp should be called exactly once.")
}

func TestScheduleBisection_RateLimitedRequests_ReturnError(t *testing.T) {
	tpm, _ := newTemporalMock(t)
	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Every(time.Hour), 1))

	_, err := svc.ScheduleBisection(ctx, &pb.ScheduleBisectRequest{})
	// invalid request should be rate limited.
	assert.ErrorContains(t, err, "git hash is empty")

	resp, err := svc.ScheduleBisection(ctx, &pb.ScheduleBisectRequest{})
	assert.Nil(t, resp)
	assert.ErrorContains(t, err, "unable to fulfill")
}

func TestScheduleBisection_InvalidRequests_ShouldError(t *testing.T) {
	tpm, _ := newTemporalMock(t)

	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Inf, 0))

	resp, err := svc.ScheduleBisection(ctx, &pb.ScheduleBisectRequest{})
	assert.Nil(t, resp)
	assert.ErrorContains(t, err, "git hash is empty")

	// TODO(b/322047067): Add requests with invalid fields
}

func TestQueryBisection_ExistingJob_ShouldReturnDetails(t *testing.T) {
	tpm, _ := newTemporalMock(t)
	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Inf, 0))

	expect := func(req *pb.QueryBisectRequest, want *pb.BisectExecution, desc string) {
		resp, err := svc.QueryBisection(ctx, req)
		// TODO(b/322047067): remove this once implemented, err should be nil
		assert.ErrorContains(t, err, "not implemented")

		// TODO(b/322047067): the response should match the expected
		assert.Nil(t, resp, desc)
	}

	// TODO(b/322047067): Add more combinations of query request and fix expected responses.
	expect(&pb.QueryBisectRequest{
		JobId: "TBD ID",
	}, nil, "should return job status")

}

func TestQueryBisection_NonExistingJob_ShouldError(t *testing.T) {
	tpm, _ := newTemporalMock(t)
	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Inf, 0))

	resp, err := svc.QueryBisection(ctx, &pb.QueryBisectRequest{
		JobId: "non-exist ID",
	})
	// TODO(b/322047067): change this to correct error message
	assert.ErrorContains(t, err, "not implemented", "Error should indicate job doesn't exist.")
	assert.Nil(t, resp, "Non-existed Job ID shouldn't contain any response.")

	resp, err = svc.QueryBisection(ctx, &pb.QueryBisectRequest{})
	// TODO(b/322047067): change this to correct error message
	assert.ErrorContains(t, err, "not implemented", "Empty Job ID should error.")
	assert.Nil(t, resp, "Empty Job ID shouldn't contain any response.")
}

func TestCancelJob_InvalidInput_ReturnError(t *testing.T) {
	tpm, _ := newTemporalMock(t)
	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Every(time.Hour), 1))

	_, err := svc.CancelJob(ctx, &pb.CancelJobRequest{JobId: "job-id"})
	assert.ErrorContains(t, err, "bad request: missing Reason")

	_, err = svc.CancelJob(ctx, &pb.CancelJobRequest{Reason: "cancel reason"})
	assert.ErrorContains(t, err, "bad request: missing JobId")
}

func TestCancelJob_JobCancelFailed_ReturnError(t *testing.T) {
	tpm, tcm := newTemporalMock(t)
	tpm.On("NewClient").Return(tcm, func() {}, nil)

	tcm.On("CancelWorkflow", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("internal error"))

	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Every(time.Hour), 1))

	resp, err := svc.CancelJob(ctx, &pb.CancelJobRequest{JobId: "job-id", Reason: "cancel reason"})
	assert.Nil(t, resp)
	assert.ErrorContains(t, err, "Unable to cancel workflow")
	assert.ErrorContains(t, err, "internal error")
}

func TestCancelJob_JobCancelled_ReturnSucceed(t *testing.T) {
	tpm, tcm := newTemporalMock(t)
	tpm.On("NewClient").Return(tcm, func() {}, nil)

	tcm.On("CancelWorkflow", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	ctx := context.Background()
	svc := New(tpm, rate.NewLimiter(rate.Every(time.Hour), 1))

	resp, err := svc.CancelJob(ctx, &pb.CancelJobRequest{JobId: "job-id", Reason: "cancel reason"})
	assert.Nil(t, err)
	assert.Equal(t, resp.JobId, "job-id")
	assert.Equal(t, resp.State, "Cancelled")
}
