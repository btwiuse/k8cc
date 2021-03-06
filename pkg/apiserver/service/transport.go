package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// PUT     /api/v1/distcc/:namespce/:tag/lease/:user   gives the user a lease for a distcc
	// DELETE  /api/v1/distcc/:namespce/:tag/lease/:user   removes the lease for a distcc
	// PUT     /api/v1/client/:namespce/:tag/lease/:user   gives the user a lease for a client
	// DELETE  /api/v1/client/:namespce/:tag/lease/:user   removes the lease for a client

	r.Methods("PUT").Path("/api/v1/distcc/{namespace}/{tag}/lease/{user}").Handler(httptransport.NewServer(
		e.PutLeaseDistccEndpoint,
		decodeLeaseRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/api/v1/distcc/{namespace}/{tag}/lease/{user}").Handler(httptransport.NewServer(
		e.DeleteLeaseDistccEndpoint,
		decodeLeaseRequest,
		encodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/api/v1/client/{namespace}/{tag}/lease/{user}").Handler(httptransport.NewServer(
		e.PutLeaseClientEndpoint,
		decodeLeaseRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/api/v1/client/{namespace}/{tag}/lease/{user}").Handler(httptransport.NewServer(
		e.DeleteLeaseClientEndpoint,
		decodeLeaseRequest,
		encodeResponse,
		options...,
	))

	return r
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrCanceled:
		return http.StatusRequestTimeout
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func decodeLeaseRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		return nil, ErrBadRouting
	}
	tag, ok := vars["tag"]
	if !ok {
		return nil, ErrBadRouting
	}
	ns, ok := vars["namespace"]
	if !ok {
		return nil, ErrBadRouting
	}
	return leaseRequest{
		User:      user,
		Namespace: ns,
		Tag:       tag,
	}, nil
}
