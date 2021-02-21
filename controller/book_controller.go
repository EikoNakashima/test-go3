package book

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	book "github.com/EikoNakashima/test-go3/service"
)

type Controller struct{}

// 検索 GET /books
func (pc Controller) Index(c *gin.Context) {
	// パラメータ取得
	title := c.Query("title")
	category, _ := strconv.Atoi(c.Query("category"))
	author := c.Query("author")

	// 検索処理
	var s book.Service
	p, err := s.Search(title, category, author)

	// 検索結果を返す
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}
