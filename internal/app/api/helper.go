package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/rendley/webserver/storage"
	"github.com/sirupsen/logrus"
)

// Пытаемся отконфижить API instance
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Пытаетмя отконфижить router

func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api!"))
	})

}

// Пытаемся отконфижить storage

func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
