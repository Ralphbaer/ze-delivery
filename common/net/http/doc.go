package http

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// DocAPI adds the default documentation route to the API.
// Ex: /{serviceName}/docs
// And adds the swagger route too.
// Ex: /{serviceName}/swagger.yaml
func DocAPI(specURL, serviceName, title string, r *mux.Router) {

	docURL := fmt.Sprintf("/%s/docs", serviceName)
	defaultSpecURL := fmt.Sprintf("/%s/swagger.yaml", serviceName)

	if strings.TrimSpace(specURL) == "" {
		specURL = defaultSpecURL
	}

	r.Handle(defaultSpecURL, File("./gen/swagger.yaml"))

	opts := middleware.RedocOpts{
		Path:    docURL,
		SpecURL: specURL,
		Title:   title,
	}

	docs := middleware.Redoc(opts, nil)

	r.Handle(fmt.Sprintf("/%s/docs", serviceName), docs)
}
