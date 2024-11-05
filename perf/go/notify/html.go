package notify

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"go.skia.org/infra/go/now"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/perf/go/alerts"
	"go.skia.org/infra/perf/go/clustering2"
	"go.skia.org/infra/perf/go/git/provider"
	"go.skia.org/infra/perf/go/ui/frame"
)

const (
	newRegressionHTML = `<b>Alert</b><br><br>
<p>
	A Perf Regression ({{.Cluster.StepFit.Status}}) has been found at:
</p>
<p style="padding: 1em;">
	<a href="{{.URL}}/g/t/{{.Commit.GitHash}}">{{.URL}}/g/t/{{.Commit.GitHash}}</a>
</p>
<p>
  For:
</p>
<p style="padding: 1em;">
  <a href="{{ .CommitURL }}">{{ .CommitURL }}</a>
</p>
<p>
	With {{.Cluster.Num}} matching traces.
</p>
<p>
   And direction {{.Cluster.StepFit.Status}}.
</p>
<p>
	From Alert <a href="{{.URL}}/a/?{{ .Alert.IDAsString }}">{{ .Alert.DisplayName }}</a>
</p>
`
	regressionMissingHTML = `<b>Alert</b><br><br>
<p>
	A Perf Regression ({{.Cluster.StepFit.Status}}) can no longer be found at:
</p>
<p style="padding: 1em;">
	<a href="{{.URL}}/g/t/{{.Commit.GitHash}}">{{.URL}}/g/t/{{.Commit.GitHash}}</a>
</p>
<p>
	For:
</p>
<p style="padding: 1em;">
	<a href="{{ .CommitURL }}">{{ .CommitURL }}</a>
</p>
<p>
	With {{.Cluster.Num}} matching traces.
</p>
<p>
	And direction {{.Cluster.StepFit.Status}}.
</p>
<p>
	From Alert <a href="{{.URL}}/a/?{{ .Alert.IDAsString }}">{{ .Alert.DisplayName }}</a>
</p>
`
)

var (
	htmlTemplateNewRegression     = template.Must(template.New("newRegressionHTML").Parse(newRegressionHTML))
	htmlTemplateRegressionMissing = template.Must(template.New("regressionMissingHTML").Parse(regressionMissingHTML))
)

// HTMLFormatter implements Formatter.
type HTMLFormatter struct {
	commitRangeURITemplate string
}

// NewHTMLFormatter returns a new HTMLFormatter.
func NewHTMLFormatter(commitRangeURITemplate string) HTMLFormatter {
	return HTMLFormatter{
		commitRangeURITemplate: commitRangeURITemplate,
	}
}

// FormatNewRegression implements Formatter.
func (h HTMLFormatter) FormatNewRegression(ctx context.Context, commit, previousCommit provider.Commit, alert *alerts.Alert, cl *clustering2.ClusterSummary, URL string, frame *frame.FrameResponse) (string, string, error) {
	templateContext := &TemplateContext{
		URL:       URL,
		Commit:    commit,
		CommitURL: URLFromCommitRange(commit, previousCommit, h.commitRangeURITemplate),
		Alert:     alert,
		Cluster:   cl,
	}

	var b bytes.Buffer
	if err := htmlTemplateNewRegression.Execute(&b, templateContext); err != nil {
		return "", "", skerr.Wrapf(err, "format HTML body for a new regression")
	}
	subject := fmt.Sprintf("%s - Regression found for %s", alert.DisplayName, commit.Display(now.Now(ctx)))

	return b.String(), subject, nil
}

// FormatRegressionMissing implements Formatter.
func (h HTMLFormatter) FormatRegressionMissing(ctx context.Context, commit, previousCommit provider.Commit, alert *alerts.Alert, cl *clustering2.ClusterSummary, URL string, frame *frame.FrameResponse) (string, string, error) {
	templateContext := &TemplateContext{
		URL:       URL,
		Commit:    commit,
		CommitURL: URLFromCommitRange(commit, previousCommit, h.commitRangeURITemplate),
		Alert:     alert,
		Cluster:   cl,
	}

	var b bytes.Buffer
	if err := htmlTemplateRegressionMissing.Execute(&b, templateContext); err != nil {
		return "", "", skerr.Wrapf(err, "format HTML body for a regression that has gone missing")
	}
	subject := fmt.Sprintf("%s - Regression no longer found for %s", alert.DisplayName, commit.Display(now.Now(ctx)))
	return b.String(), subject, nil
}

var _ Formatter = HTMLFormatter{}
