package book

import (
	"github.com/EikoNakashima/test-go3/db"
	"github.com/sample/sample-go-api/entity"
)

type Service struct{}

type book entity.Book

type Parameter struct {
	title    string
	category int
	author   string
}

// 検索
func (s Service) Search(title string, category int, author string) ([]Book, error) {

	// DB接続
	db := db.GetDB()

	// 本モデルから作成
	var book []Book

	// パラメータセット
	p := Parameter{title, category, author}

	// DB接続確認
	if err := db.Take(&book).Error; err != nil {
		return nil, err
	}

	// 本検索クエリ
	tx := db
	tx = tx.Find(&book)

	// タイトル
	if p.title != "" {
		tx = tx.Where("title = ?", p.title).Find(&book)
	}

	// カテゴリ
	if p.category != 0 {
		tx = tx.Where("category = ?", p.category).Find(&book)
	}

	// 著者
	if p.author != "" {
		tx = tx.Where("author = ?", p.author).Find(&book)
	}

	return book, nil
}
