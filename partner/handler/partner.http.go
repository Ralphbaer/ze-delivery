package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	commonHTTP "github.com/Ralphbaer/ze-delivery/common/net/http"
	repo "github.com/Ralphbaer/ze-delivery/partner/repository"
	uc "github.com/Ralphbaer/ze-delivery/partner/usecase"
)

// PartnerHandler represents a handler which deal with Partner resource operations
type PartnerHandler struct {
	UseCase *uc.PartnerUseCase
}

// GetByID returns a partner by its ID
// swagger:operation GET /partners/{id} partners GetByID
// Returns an partner by its id
// ---
// parameters:
//  -  name: id
//     in: path
//     type: string
//     description: The id of the partner
//     required: true
//
// responses:
//   '200':
//     description: Success Operation
//     schema:
//         "$ref": "#/definitions/Partner"
//   '404':
//     description: Not Found - Resource does not exists
//     schema:
//       "$ref": "#/definitions/ResponseError"
//     examples:
//       "application/json":
//         code: 404
//         message: message
func (handler *PartnerHandler) GetByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		partnerID := mux.Vars(r)["id"]
		partner, err := handler.UseCase.Get(r.Context(), partnerID)
		if partner == nil {
			log.Printf(fmt.Sprintf("partner with ID %s not found.", partnerID))
			commonHTTP.WithError(w, err)
			return
		}

		if err != nil {
			log.Println(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		log.Printf("Successfully found Partner: %+v", partner)

		commonHTTP.OK(w, partner)
	})
}

// GetNearestPartner returns the nearest partner given coordinates longitude and latitude
// swagger:operation GET /partners/nearest partners GetNearestPartner
// Returns a partner given coordinates longitude and latitude
// ---
// parameters:
//  -  name: long
//     in: query
//     type: integer
//     description: Longitude
//     required: true
//  -  name: lat
//     in: query
//     type: integer
//     description: Latitude
//     required: true
//
// responses:
//   '200':
//     description: Success Operation
//     schema:
//         "$ref": "#/definitions/Partner"
//   '404':
//     description: Not Found - Resource does not exists
//     schema:
//       "$ref": "#/definitions/ResponseError"
//     examples:
//       "application/json":
//         code: 404
//         message: message
func (handler *PartnerHandler) GetNearestPartner() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		schemaDecoder := schema.NewDecoder()

		var q repo.PartnerQuery
		if err := schemaDecoder.Decode(&q, r.URL.Query()); err != nil {
			log.Println(err.Error())
			commonHTTP.BadRequest(w, err.Error())
			return
		}

		partner, err := handler.UseCase.GetNearestPartner(r.Context(), &q)
		if err != nil {
			log.Println(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		log.Printf("Successfully found Partner: %+v", partner)

		commonHTTP.OK(w, partner)
	})
}


// Create creates a new Partner in the repository
// swagger:operation POST /partners partners Create
// Register a new Partner into database
// ---
// parameters:
//  -  name: input
//     in: body
//     type: string
//     description: The payload
//     required: true
//     schema:
//       "$ref": "#/definitions/CreatePartnerInput"
//
// security:
//  - Definitions: []
//
// responses:
//   '201':
//     description: Success Operation
//     schema:
//       "$ref": "#/definitions/Partner"
//   '400':
//     description: Invalid Input - Input has invalid/missing values
//     schema:
//       "$ref": "#/definitions/ValidationError"
//     examples:
//       "application/json":
//         code: 400
//         message: message
//   '409':
//     description: Conflict - partner document already taken
//     schema:
//       "$ref": "#/definitions/ResponseError"
//     examples:
//       "application/json":
//         code: 409
//         message: message
func (handler *PartnerHandler) Create(p interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := p.(*uc.CreatePartnerInput)

		log.Printf("Trying to create Partner with payload %+v", payload)

		partner, err := handler.UseCase.Create(r.Context(), payload)
		if err != nil {
			log.Println(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		log.Printf("Successfully created Partner. ID: %s", partner.ID)
		w.Header().Set("Location", fmt.Sprintf("%s/partner/partners/%s", r.Host, partner.ID))
		w.Header().Set("Content-Type", "application/json")

		commonHTTP.Created(w, partner)
	})
}