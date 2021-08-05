// Package app Partner API.
//
// This guide describes all Zé Delivery Partner API and usecase.
//
//     Schemes: http, https
//     BasePath: /partner
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package app

import (
	lib "github.com/Ralphbaer/ze-delivery/common/net/http"
	"github.com/Ralphbaer/ze-delivery/partner/handler"
	uc "github.com/Ralphbaer/ze-delivery/partner/usecase"
	"github.com/gorilla/mux"
)

// NewRouter registers routes to the Server
func NewRouter(p *handler.PartnerHandler) *mux.Router {
	r := mux.NewRouter()
	config := NewConfig()

	r.Use(lib.WithCorrelationID)

	r.Handle("/partner/partners/nearest", p.GetNearestPartner()).Methods("GET")

	r.Handle("/partner/partners/{id}", p.GetByID()).Methods("GET")

	r.Handle("/partner/partners", lib.WithBody(new(uc.CreatePartnerInput), p.Create)).Methods("POST")
	
	// Common

	r.HandleFunc("/partner/ping", lib.Ping)

	// Documentation

	lib.DocAPI(config.SpecURL, "partner", "partner API Documentation", r)

	return r
}
