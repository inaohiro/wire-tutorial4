// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/inaohiro-sandbox/wire-tutorial4/app/config"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/domain"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/database"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/database/mysql"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/infra/rest/controllers"
	database2 "github.com/inaohiro-sandbox/wire-tutorial4/app/interfaces/database"
	"github.com/inaohiro-sandbox/wire-tutorial4/app/interfaces/inmemorydb"
)

// Injectors from wire.go:

func initializeRealServer() (*rest.Server, func(), error) {
	serverConfig := config.HttpServerConfig()
	mysqlConfig := config.MysqlConfig()
	dataSource := mysql.NewDataSource(mysqlConfig)
	db, cleanup, err := database.Open(dataSource)
	if err != nil {
		return nil, nil, err
	}
	userRepository := database2.NewUserRepository(db)
	userService := domain.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	router := rest.NewRouter(userController)
	server := rest.NewServer(serverConfig, router)
	return server, func() {
		cleanup()
	}, nil
}

func initializeServer() (*rest.Server, func(), error) {
	serverConfig := config.HttpServerConfig()
	inMemoryUserRepository := inmemorydb.NewUserRepository()
	userService := domain.NewUserService(inMemoryUserRepository)
	userController := controllers.NewUserController(userService)
	router := rest.NewRouter(userController)
	server := rest.NewServer(serverConfig, router)
	return server, func() {
	}, nil
}
