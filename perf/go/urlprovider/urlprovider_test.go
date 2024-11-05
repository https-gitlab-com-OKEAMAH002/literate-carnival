package urlprovider

import (
	"context"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	perfgit "go.skia.org/infra/perf/go/git"
	"go.skia.org/infra/perf/go/git/gittest"
)

func TestProvider_Default(t *testing.T) {
	perfgit := getPerfGit(t)
	urlProvider := New(perfgit)
	params := map[string][]string{
		"param1": {"value1"},
		"param2": {"value2"},
		"param3": {"value3"},
	}
	queryurl := urlProvider.Explore(context.Background(), 1234, 5678, params, false)
	assert.NotNil(t, queryurl, "Url expected to be generated")
	queryIndex := strings.Index(queryurl, "&queries=")
	assert.NotEqual(t, -1, queryIndex)
	queryString := queryurl[queryIndex+9:]
	unescaped_queryString, err := url.QueryUnescape(queryString)
	assert.Nil(t, err)
	parsed_query, _ := url.ParseQuery(unescaped_queryString)

	assert.Equal(t, parsed_query.Get("param1"), "value1")
	assert.Equal(t, parsed_query.Get("param2"), "value2")
	assert.Equal(t, parsed_query.Get("param3"), "value3")
}

func TestProvider_Chromeperf_NoCustomization(t *testing.T) {
	perfgit := getPerfGit(t)
	urlProvider := New(perfgit)
	params := map[string][]string{
		"param1": {"value1"},
		"param2": {"value2"},
		"param3": {"value3"},
	}
	queryurl := urlProvider.Explore(context.Background(), 1234, 5678, params, false)
	assert.NotNil(t, queryurl, "Url expected to be generated")
	queryIndex := strings.Index(queryurl, "&queries=")
	assert.NotEqual(t, -1, queryIndex)
	queryString := queryurl[queryIndex+9:]
	unescaped_queryString, err := url.QueryUnescape(queryString)
	assert.Nil(t, err)
	parsed_query, _ := url.ParseQuery(unescaped_queryString)

	assert.Equal(t, parsed_query.Get("param1"), "value1")
	assert.Equal(t, parsed_query.Get("param2"), "value2")
	assert.Equal(t, parsed_query.Get("param3"), "value3")
}

func TestProvider_MultiGraph(t *testing.T) {
	perfgit := getPerfGit(t)
	urlProvider := New(perfgit)
	shortcutId := "shortcutId"
	queryurl := urlProvider.MultiGraph(context.Background(), 1234, 5678, shortcutId)
	assert.NotNil(t, queryurl, "Url expected to be generated")
	multiGraphIndex := strings.Index(queryurl, "/m/?")
	assert.NotEqual(t, -1, multiGraphIndex)
	parsed_query, _ := url.ParseQuery(queryurl)

	assert.Equal(t, parsed_query.Get("shortcut"), shortcutId)
}

func getPerfGit(t *testing.T) perfgit.Git {
	ctx, db, _, _, _, instanceConfig := gittest.NewForTest(t)
	git, err := perfgit.New(ctx, true, db, instanceConfig)
	require.NoError(t, err)
	return git
}
