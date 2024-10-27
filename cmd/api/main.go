package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/rendley/webserver/internal/app/api"
)

var (
	configPath string
)

func init() {
	// Скажем что наше приложение будет на этапе запуска получать до конфиг файла из внешнего источника
	// path - флаг при запуске, какой использовать если не распарсится, описание
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")

}

func main() {
	// запуск go run cmd/api/main.go -path configs/api.toml
	flag.Parse() // без этой строчки будет пустая строка. Здесь происходит инициализация  переменной ConfigPath

	log.Println("It works now")
	// server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // deserialization .toml file
	if err != nil {
		log.Println("Can not find configs file. Using default values:", err)
	}

	// читаем из config/.toml or .env
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	// log.Fatal(server.Start()) // вместо обработки ошибок можно взять эту строку - одно и тоже
}
