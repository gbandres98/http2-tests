package endpoint

import (
	"net/http"

	"github.com/PaackEng/paackit/httpx"
	"github.com/PaackEng/paackit/response"
)

const (
	Path = "/test"
)

type TestEndpointDI struct {
	Middleware []httpx.Middleware
}

type testEndpoint struct {
	middleware []httpx.Middleware
}

func NewTestEndpoint(di TestEndpointDI) httpx.Service {
	return &testEndpoint{
		middleware: di.Middleware,
	}
}

func (u *testEndpoint) Path() string {
	return Path
}

func (u *testEndpoint) Method() httpx.Method {
	return httpx.MethodGet
}

func (u *testEndpoint) Handler() http.HandlerFunc {
	return httpx.WithMiddleware(u.handlerFunc, u.middleware...)
}

func (u *testEndpoint) handlerFunc(w http.ResponseWriter, r *http.Request) {

	_ = httpx.Encode(w, r, 200, response.New(nil, nil))
}
