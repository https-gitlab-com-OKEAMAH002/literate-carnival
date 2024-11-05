// fiddle is the web server for fiddle.
package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/fiddlek/go/store/mocks"
	"go.skia.org/infra/fiddlek/go/types"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/scrap/go/scrap"
	scrapMocks "go.skia.org/infra/scrap/go/scrap/mocks"
)

const (
	code = "void draw(SkCanvas* canvas) {}"
	hash = "1234567890"
)

var (
	errMyMockError = fmt.Errorf("my mock error")
)

func TestScrapHandler_HappyPath(t *testing.T) {

	r := httptest.NewRequest("GET", "/scrap/svg/@smiley", nil)
	w := httptest.NewRecorder()

	// Mock the scrapClient.
	scrapClientMock := &scrapMocks.ScrapExchange{}
	scrapClientMock.On("Expand", testutils.AnyContext, scrap.SVG, "@smiley", scrap.CPP, mock.Anything).Run(func(args mock.Arguments) {
		_, err := args.Get(4).(io.Writer).Write([]byte(code))
		require.NoError(t, err)
	}).Return(nil)
	defer scrapClientMock.AssertExpectations(t)
	scrapClient = scrapClientMock

	// Mock the fiddle store.
	store := &mocks.Store{}
	expectedOptions := types.Options{
		Width:  256,
		Height: 256,
		Source: 1,
	}
	store.On("Put", code, expectedOptions, (*types.Result)(nil)).Return(hash, nil)
	defer store.AssertExpectations(t)
	fiddleStore = store

	// Make the request through a chi.Router so the URL paths get parsed and
	// routed correctly.
	router := chi.NewRouter()
	addHandlers(router)

	router.ServeHTTP(w, r)

	// Should redirect to the newly created fiddle.
	require.Equal(t, 307, w.Code)
	require.Equal(t, fmt.Sprintf("/c/%s", hash), w.Header().Get("Location"))
}

func TestScrapHandler_ScrapExchangeFails_ReturnsInternalServerError(t *testing.T) {

	r := httptest.NewRequest("GET", "/scrap/svg/@smiley", nil)
	w := httptest.NewRecorder()

	// Mock the scrapClient.
	scrapClientMock := &scrapMocks.ScrapExchange{}
	scrapClientMock.On("Expand", testutils.AnyContext, scrap.SVG, "@smiley", scrap.CPP, mock.Anything).Return(errMyMockError)
	defer scrapClientMock.AssertExpectations(t)
	scrapClient = scrapClientMock

	// Make the request through a chi.Router so the URL paths get parsed and
	// routed correctly.
	router := chi.NewRouter()
	addHandlers(router)

	router.ServeHTTP(w, r)

	require.Equal(t, 500, w.Code)
}

func TestScrapHandler_FiddleStorePutFails_ReturnsInternalServerError(t *testing.T) {

	r := httptest.NewRequest("GET", "/scrap/svg/@smiley", nil)
	w := httptest.NewRecorder()

	// Mock the scrapClient.
	scrapClientMock := &scrapMocks.ScrapExchange{}
	scrapClientMock.On("Expand", testutils.AnyContext, scrap.SVG, "@smiley", scrap.CPP, mock.Anything).Run(func(args mock.Arguments) {
		_, err := args[4].(io.Writer).Write([]byte(code))
		require.NoError(t, err)
	}).Return(nil)
	defer scrapClientMock.AssertExpectations(t)
	scrapClient = scrapClientMock

	// Mock the fiddle store.
	store := &mocks.Store{}
	expectedOptions := types.Options{
		Width:  256,
		Height: 256,
		Source: 1,
	}
	store.On("Put", code, expectedOptions, (*types.Result)(nil)).Return("", errMyMockError)
	defer store.AssertExpectations(t)
	fiddleStore = store

	// Make the request through a chi.Router so the URL paths get parsed and
	// routed correctly.
	router := chi.NewRouter()
	addHandlers(router)

	router.ServeHTTP(w, r)

	require.Equal(t, 500, w.Code)
}

func TestScrapHandler_WithInvalidType_ReturnsBadRequest(t *testing.T) {

	r := httptest.NewRequest("GET", "/scrap/NotAType/@smiley", nil)
	w := httptest.NewRecorder()

	// Make the request through a chi.Router so the URL paths get parsed and
	// routed correctly.
	router := chi.NewRouter()
	addHandlers(router)

	router.ServeHTTP(w, r)

	require.Equal(t, 404, w.Code)
}
