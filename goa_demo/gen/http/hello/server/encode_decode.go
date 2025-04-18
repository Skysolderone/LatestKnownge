// Code generated by goa v3.20.1, DO NOT EDIT.
//
// hello HTTP server encoders and decoders
//
// Command:
// $ goa gen goademo/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeSayResponse returns an encoder for responses returned by the hello say
// endpoint.
func EncodeSayResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSayRequest returns a decoder for requests sent to the hello say
// endpoint.
func DecodeSayRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			name string

			params = mux.Vars(r)
		)
		name = params["name"]
		payload := name

		return payload, nil
	}
}
