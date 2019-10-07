package communication

import (
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go/jsonapi"
	"net/http"
)

type Article struct {
	ID    string `json:"-"`
	Title string `json:"title"`
}

func (a Article) GetID() string {
	return a.ID
}

func Run() {
	a := Article{
		ID:    "id_a",
		Title: "title_a",
	}
	j, _ := jsonapi.Marshal(a)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/vnd.api+json", j)
	})
	_ = router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
