package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	//Database FileDesctiptor
	db *sql.DB
	// Subifield for repo interfacing (module user and arcticle)
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

// New returns a new Storage instance from the given configuration.
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

// Open initializes a connection to the database specified in the
// configuration. It returns an error if it is unable to connect, and
// otherwise logs a success message and sets the db field of the
// receiver.
func (storage *Storage) Open() error {
	//db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=qwerty dbname=restapi sslmode=disable")
	fmt.Println(storage.config.DatabaseURI)
	db, err := sql.Open("postgres", storage.config.DatabaseURI) // parse only api.toml add block [storage]
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

// Close closes the connection to the database that was opened with Open.
func (storage *Storage) Close() {
	storage.db.Close()
}

// Public Repo for Article
func (s *Storage) Article() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return nil
}

// Public Repo for User
func (s *Storage) User() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		storage: s,
	}
	return nil
}
