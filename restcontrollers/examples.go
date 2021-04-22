package restcontrollers

import (
	log "github.com/sirupsen/logrus"
	"github.com/sradevski/protoc-gen-mock/stub"
	"net/http"
)

type ExamplesController struct {
	StubExamples []stub.Stub
}

func (c ExamplesController) GetHandlers() []RESTHandler {
	return []RESTHandler{
		{
			Name:    "GetExamples",
			Path:    "",
			Methods: []string{http.MethodGet},
			Handler: c.getExamplesHandler,
		},
	}
}

func (c ExamplesController) GetPath() string {
	return "/examples"
}

func (c ExamplesController) getExamplesHandler(writer http.ResponseWriter, request *http.Request) {
	log.Info("REST: received call to get example stubs")

	writeErr := writeResponse(writer, c.StubExamples)
	if writeErr != nil {
		writeErrorResponse(writer, http.StatusInternalServerError, writeErr.Error())
	}
}
