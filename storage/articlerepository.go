package storage

import (
	"fmt"
	"log"

	"github.com/rendley/webserver/internal/app/models"
)

// instance of Article repository (model interface - работа со структурами модели через методы)
type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle string = "articles"
)

// Create Article in db
func (ar *ArticleRepository) Create(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES($1, $2, $3) RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// Delete Article by id
func (ar *ArticleRepository) DeleteByID(id int) (*models.Article, error) {
	article, ok, err := ar.FindArticleById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)
		_, err := ar.storage.db.Exec(query, id) // Exec return Result count string
		if err != nil {
			return nil, err
		}
	}
	return article, nil
}

// Get all Articles
func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)
	rows, err := ar.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := make([]*models.Article, 0)
	for rows.Next() {
		a := models.Article{}
		err := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		articles = append(articles, &a)
	}
	return articles, nil
}

// Get Article by id
func (ar *ArticleRepository) FindArticleById(id int) (*models.Article, bool, error) {
	articles, err := ar.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var articleFinded *models.Article
	for _, a := range articles {
		if a.ID == id {
			articleFinded = a
			founded = true
		}
	}
	return articleFinded, founded, nil
}

// Update Article

func (ar *ArticleRepository) UpdateArticle(a *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("UPDATE %s SET title=$1, author=$2, content=$3 WHERE id=$4 RETURNING id", tableArticle)
	if err := ar.storage.db.QueryRow(query, a.Title, a.Author, a.Content, a.ID).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}
