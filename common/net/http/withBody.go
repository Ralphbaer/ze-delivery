package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
)

// DecodeHandlerFunc is a handler which works with withBody decorator.
// It receives a struct which was decoded by withBody decorator before.
// Ex: json -> withBody -> DecodeHandlerFunc
type DecodeHandlerFunc func(p interface{}) http.Handler

// PayloadContextValue is a wrapper type used to keep Context.WithValue safe
type PayloadContextValue string

// ConstructorFunc representing a constructor of any type.
type ConstructorFunc func() interface{}

// DecoderHandler decodes payload caming from requests.
type decoderHandler struct {
	handler      DecodeHandlerFunc
	constructor  ConstructorFunc
	structSource interface{}
}

func newOfType(s interface{}) interface{} {
	t := reflect.TypeOf(s)
	v := reflect.New(t.Elem())
	return v.Interface()
}

// ServerHTTP to satisfy contract
func (d *decoderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var s interface{}
	var m map[string]interface{}

	if d.constructor != nil {
		s = d.constructor()
	} else {
		s = newOfType(d.structSource)
	}

	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := decodeJSON(bytes.NewReader(bodyBytes), &m); err != nil {
		BadRequest(w, err)
		return
	}

	if err := decodeJSON(bytes.NewReader(bodyBytes), s); err != nil {
		BadRequest(w, err)
		return
	}

	if err := ValidateStruct(s); err != nil {
		BadRequest(w, err)
		return
	}

	ctx := context.WithValue(r.Context(), PayloadContextValue("fields"), m)
	d.handler(s).ServeHTTP(w, r.WithContext(ctx))
}

// WithBody return a new instance of decoderHandler using a source for constructor.
func WithBody(s interface{}, h DecodeHandlerFunc) http.Handler {
	return &decoderHandler{
		handler:      h,
		structSource: s,
	}
}

func decodeJSON(r io.Reader, s interface{}) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(s)
	if err != nil {
		return &ResponseError{
			Code:    400,
			Message: fmt.Sprintf("Request body contains badly-formed JSON (%v).", err.Error()),
		}
	}
	return nil
}
