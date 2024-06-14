package app

import (
	"github.com/kunanon-c/url-shortner/internal/handler"
	"github.com/kunanon-c/url-shortner/internal/repository"
)
import "github.com/gin-gonic/gin"

func StartApp() {
	g := gin.Default()

	repo := repository.Repository{}
	err := repo.Init()
	if err != nil {
		panic(err)
	}

	h := handler.Handler{
		Repo: repo,
	}

	// router
	g.POST("/save", h.SaveURL)

	// g.GET("/goto/:shorten", h.Redirect)
	// example localhost:8080/goto/80775
	// go get Long URl pair of 80775 and redirect to that

	// g.GET("/all") list every shorten URL pair, pagination if possible

	// g.PATCH("/patch/:shorten") // change long URL of this shorten
	// g.DELETE("/delete/:shorten") // delete this shorten

	g.GET("/ping", h.Ping)
	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
