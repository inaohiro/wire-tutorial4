package main

import (
	"log"
)

//
//type app struct {
//	server *rest.Server
//}
//
//func (a *app) Run() error {
//	return a.server.ListenAndServe()
//}
//
//func newApp() *app {
//	return &app{}
//}
//
//func main() {
//	app, cleanup, err := initializeApp()
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	defer cleanup()
//
//	if err := app.Run(); err != nil {
//		inmemorydb.Dump()
//		log.Fatal(err.Error())
//	}
//}

func main() {
	server, cleanup, err := initializeServer()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cleanup()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
