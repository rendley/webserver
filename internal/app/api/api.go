package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendley/webserver/storage"
	"github.com/sirupsen/logrus"
)

type API struct {
	// unexported field
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	// Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	// Trying to configure Router
	api.configureRouterField()
	//Trying to config Storage
	if err := api.configureStorageField(); err != nil {
		return (err)
	}
	// На этапе валидного завершения стратуем http-server
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
