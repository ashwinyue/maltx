// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package apiserver

import (
	"github.com/ashwinyue/maltx/internal/apiserver/biz"
	"github.com/ashwinyue/maltx/internal/apiserver/pkg/validation"
	"github.com/ashwinyue/maltx/internal/apiserver/store"
	"github.com/ashwinyue/maltx/internal/pkg/server"
	"github.com/ashwinyue/maltx/pkg/authz"
)

// Injectors from wire.go:

func InitializeWebServer(config *Config) (server.Server, error) {
	string2 := config.ServerMode
	db, err := ProvideDB(config)
	if err != nil {
		return nil, err
	}
	redisStore, err := ProvideRedis(config)
	if err != nil {
		return nil, err
	}
	datastore := store.NewStore(db, redisStore)
	v := authz.DefaultOptions()
	authzAuthz, err := authz.NewAuthz(db, v...)
	if err != nil {
		return nil, err
	}
	bizBiz := biz.NewBiz(datastore, authzAuthz)
	validator := validation.New(datastore)
	userRetriever := &UserRetriever{
		store: datastore,
	}
	serverConfig := &ServerConfig{
		cfg:       config,
		biz:       bizBiz,
		val:       validator,
		retriever: userRetriever,
		authz:     authzAuthz,
	}
	serverServer, err := NewWebServer(string2, serverConfig)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}
