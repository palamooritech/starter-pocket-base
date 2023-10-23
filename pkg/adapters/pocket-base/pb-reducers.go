package pb_adapter

import (
	typos "31arthur/starter-pocket-base/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type PocketBaseServer struct {
	PBServer *pocketbase.PocketBase
}

type PocketBaseServerReducers interface {
	CreateGETPath(string, typos.PBPathFunc) error
	CreatePOSTPath(string, typos.PBPathFunc) error
	Start() error
	FetchData() error
	PutData() error
}

// initializes the Pocketbase app along with the Echo server
func PBNewAPIServer() *PocketBaseServer {
	app := pocketbase.New()
	return &PocketBaseServer{PBServer: app}
}

func (p *PocketBaseServer) CreateGETPath(path_name string, handler_function typos.PBPathFunc) error {

	//we will have to serve the paths, before the echo server starts, the same as usual in mux or chi router.
	err := errors.New(p.PBServer.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		//Pocket-bases uses Echo Server, which is slightly different from the mux or chi routers. The Echo Serve along with the handlers, have templates and middlewares too. Read Echo Web Framework for more details
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    path_name,
			Handler: echo.HandlerFunc(handler_function),
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(p.PBServer),
			},
		})

		return nil
	}))

	if err != nil {
		return err
	}

	return nil

}

func (p *PocketBaseServer) CreatePOSTPath(path_name string, handler_function typos.PBPathFunc) error {
	//would repeat the same for post requests, except for the method.
	err := errors.New(
		p.PBServer.OnBeforeServe().Add(func(e *core.ServeEvent) error {
			e.Router.AddRoute(echo.Route{
				Method:  http.MethodPost,
				Path:    path_name,
				Handler: echo.HandlerFunc(handler_function),
				Middlewares: []echo.MiddlewareFunc{
					apis.ActivityLogger(p.PBServer),
				},
			})

			return nil
		}))

	if err != nil {
		return err
	}

	return nil
}

func (p *PocketBaseServer) Start() error {
	return p.PBServer.Start()
}
