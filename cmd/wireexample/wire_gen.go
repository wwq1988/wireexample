// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/wwq1988/wireexample/pkg/conf"
	"github.com/wwq1988/wireexample/pkg/mux"
	"github.com/wwq1988/wireexample/pkg/repo"
	"github.com/wwq1988/wireexample/pkg/server"
)

// Injectors from wire.go:

func Create() (*server.Server, error) {
	confConf, err := conf.New()
	if err != nil {
		return nil, err
	}
	repoRepo, err := repo.New(confConf)
	if err != nil {
		return nil, err
	}
	handler := mux.New(repoRepo)
	serverServer, err := server.New(confConf, handler, repoRepo)
	if err != nil {
		return nil, err
	}
	return serverServer, nil
}
