// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// sommelier client HTTP transport
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/http"
)

// Client lists the sommelier service endpoint HTTP clients.
type Client struct {
	// Pick Doer is the HTTP client used to make requests to the pick endpoint.
	PickDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the sommelier service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PickDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Pick returns a endpoint that makes HTTP requests to the sommelier service
// pick server.
func (c *Client) Pick() endpoint.Endpoint {
	var (
		encodeRequest  = EncodePickRequest(c.encoder)
		decodeResponse = DecodePickResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPickRequest(v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}

		resp, err := c.PickDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("sommelier", "pick", err)
		}
		return decodeResponse(resp)
	}
}
