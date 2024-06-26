// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	http2 "github.com/santekno/learn-golang-restful/internal/delivery/http"
	"github.com/santekno/learn-golang-restful/internal/repository"
	"github.com/santekno/learn-golang-restful/internal/repository/mysql"
	"github.com/santekno/learn-golang-restful/internal/usecase"
	"github.com/santekno/learn-golang-restful/internal/usecase/article"
	"github.com/santekno/learn-golang-restful/pkg/database"
	"github.com/santekno/learn-golang-restful/pkg/environment"
	"github.com/santekno/learn-golang-restful/pkg/router"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() (*http.Server, error) {
	string2, err := environment.Load()
	if err != nil {
		return nil, err
	}
	db, err := database.New(string2)
	if err != nil {
		return nil, err
	}
	articleStore := mysql.New(db)
	usecase := article.New(articleStore)
	delivery := http2.New(usecase)
	httprouterRouter := router.NewRouter(delivery)
	server := router.NewServer(httprouterRouter)
	return server, nil
}

// injector.go:

var articleSet = wire.NewSet(mysql.New, wire.Bind(new(repository.ArticleRepository), new(*mysql.ArticleStore)), article.New, wire.Bind(new(usecase.ArticleUsecase), new(*article.Usecase)), http2.New)
