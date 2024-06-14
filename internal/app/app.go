package app

import (
	"github.com/kunanon-c/url-shortner/internal/handler"
	"github.com/kunanon-c/url-shortner/internal/repository"
)
import "github.com/gin-gonic/gin"

func StartApp() {
	r := gin.Default()

	repo := repository.Repository{}
	err := repo.Init()
	if err != nil {
		panic(err)
	}

	h := handler.Handler{
		Repo: repo,
	}

	// router
	r.POST("/save", h.SaveURL)

	// r.GET("/goto/:shorten", h.Redirect)
	// example localhost:8080/goto/80775
	// go get Long URl pair of 80775 and redirect to that

	// r.GET("/all") list every shorten URL pair, pagination if possible

	// r.PATCH("/patch/:shorten") // change long URL of this shorten
	// r.DELETE("/delete/:shorten") // delete this shorten

	r.GET("/ping", h.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
