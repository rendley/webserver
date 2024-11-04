package api

import (
	_ "github.com/gorilla/mux"
	"github.com/rendley/webserver/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleByID).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("Delete")
	a.router.HandleFunc(prefix+"/articles", a.PostArticles).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods("POST")
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
