package domain_core

import (
	helper "31arthur/starter-pocket-base/helpers"
	pb_adapter "31arthur/starter-pocket-base/pkg/adapters/pocket-base"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

// Define the struct of tyoe API Server, with the reference to the PocketBaseServerReducers - the interface.
type APIServer struct {
	PbStore pb_adapter.PocketBaseServerReducers
}

// Creates of the type APIServer, with the reference to the PocketBaseServerReducers - the interface.
func BinderPBInterface(binder pb_adapter.PocketBaseServerReducers) *APIServer {
	return &APIServer{
		PbStore: binder,
	}
}

func (s *APIServer) Run() {

	s.PbStore.CreateGETPath("/app", GetHandler)

	s.PbStore.CreatePOSTPath("/postValue", PostHandler)

	if err := s.PbStore.Start(); err != nil {
		log.Fatal(err)
	}

	// s.PbStore.FetchData()

}

// Sample Get Handler
func GetHandler(c echo.Context) error {
	return helper.WriteJSON(c.Response(), http.StatusOK, map[string]any{"payload": "success"})
}

// Sample Post Handler
func PostHandler(c echo.Context) error {
	searchValues := new(SearchData)
	if err := json.NewDecoder(c.Request().Body).Decode(searchValues); err != nil {
		return err
	}
	fmt.Println("searchValues", searchValues)
	return helper.WriteJSON(c.Response(), http.StatusOK, map[string]any{"payload": searchValues.Keyword})
}
