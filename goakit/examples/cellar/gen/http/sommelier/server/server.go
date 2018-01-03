// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier HTTP server
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package server

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/http"
	sommelier "goa.design/plugins/goakit/examples/cellar/gen/sommelier"
)

// Server lists the sommelier service endpoint HTTP handlers.
type Server struct {
	Pick http.Handler
}

// New instantiates HTTP handlers for all the sommelier service endpoints.
func New(
	e *sommelier.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) *Server {
	return &Server{
		Pick: NewPickHandler(e.Pick, mux, dec, enc),
	}
}

// Mount configures the mux to serve the sommelier endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountPickHandler(mux, h.Pick)
}

// MountPickHandler configures the mux to serve the "sommelier" service "pick"
// endpoint.
func MountPickHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/sommelier", f)
}

// NewPickHandler creates a HTTP handler which loads the HTTP request and calls
// the "sommelier" service "pick" endpoint.
func NewPickHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
) http.Handler {
	var (
		decodeRequest  = DecodePickRequest(mux, dec)
		encodeResponse = EncodePickResponse(enc)
		encodeError    = EncodePickError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept := r.Header.Get("Accept")
		ctx := context.WithValue(r.Context(), goahttp.ContextKeyAcceptType, accept)
		payload, err := decodeRequest(r)
		if err != nil {
			encodeError(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			encodeError(ctx, w, err)
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			encodeError(ctx, w, err)
		}
	})
}
