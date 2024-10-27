package api

import "github.com/sirupsen/logrus"

type API struct {
	// unexported field
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
	}
}

func (api *API) Start() error {
	// Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port:", api.config.BindAddr)
	return nil
}
