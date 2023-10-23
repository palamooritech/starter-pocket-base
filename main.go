package main

import (
	pb_adapter "31arthur/starter-pocket-base/pkg/adapters/pocket-base"
	domain_core "31arthur/starter-pocket-base/pkg/domain/domain-core"
)

func main() {
	//creates a instance/reference to the pocket-base initialization.
	app := pb_adapter.PBNewAPIServer()

	//binds the pocketbase app to the Interface, so that the functions defined in the interface could be accessed with the pocketbase as the reciever
	server := domain_core.BinderPBInterface(app)

	//calls the run server
	server.Run()

}
