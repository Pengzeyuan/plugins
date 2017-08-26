// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// storage HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/cellar/design

package server

import (
	"context"
	"io"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
	"goa.design/plugins/goakit/examples/cellar/gen/storage"
)

// EncodeListResponse returns an encoder for responses returned by the storage
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(storage.StoredBottleCollection)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeShowResponse returns an encoder for responses returned by the storage
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*storage.StoredBottle)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the storage show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]

		return NewShowShowPayload(id), nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show storage
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) {
		switch res := v.(type) {
		case *storage.NotFound:
			enc := encoder(ctx, w)
			body := NewShowNotFoundResponseBody(res)
			w.WriteHeader(http.StatusNotFound)
			if err := enc.Encode(body); err != nil {
				encodeError(ctx, w, err)
			}
		default:
			encodeError(ctx, w, v)
		}
	}
}

// EncodeAddResponse returns an encoder for responses returned by the storage
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the storage add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = goa.MergeErrors(err, body.Validate())

		if err != nil {
			return nil, err
		}

		return NewAddBottle(&body), nil
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the
// storage remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the storage
// remove endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]

		return NewRemoveRemovePayload(id), nil
	}
}

// wineryToWineryResponseBodyNoDefault builds a value of type
// *WineryResponseBody from a value of type *storage.Winery.
func wineryToWineryResponseBodyNoDefault(v *storage.Winery) *WineryResponseBody {
	res := &WineryResponseBody{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// componentToComponentResponseBodyNoDefault builds a value of type
// *ComponentResponseBody from a value of type *storage.Component.
func componentToComponentResponseBodyNoDefault(v *storage.Component) *ComponentResponseBody {
	res := &ComponentResponseBody{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// wineryToWineryNoDefault builds a value of type *Winery from a value of type
// *storage.Winery.
func wineryToWineryNoDefault(v *storage.Winery) *Winery {
	res := &Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// componentToComponentNoDefault builds a value of type *Component from a value
// of type *storage.Component.
func componentToComponentNoDefault(v *storage.Component) *Component {
	res := &Component{
		Varietal:   v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// wineryRequestBodyToWinerySrcPtr builds a value of type *storage.Winery from
// a value of type *WineryRequestBody.
func wineryRequestBodyToWinerySrcPtr(v *WineryRequestBody) *storage.Winery {
	res := &storage.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		URL:     v.URL,
	}

	return res
}

// componentRequestBodyToComponentSrcPtr builds a value of type
// *storage.Component from a value of type *ComponentRequestBody.
func componentRequestBodyToComponentSrcPtr(v *ComponentRequestBody) *storage.Component {
	res := &storage.Component{
		Varietal:   *v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}
