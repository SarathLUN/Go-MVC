package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PageContentRender struct {
	Title   string
	PageCSS string
	PageJS  string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./public")

	homePage := PageContentRender{
		Title:   "Home page",
		PageCSS: "home.page.css",
		PageJS:  "home.page.js",
	}
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.page.html", homePage)
	})

	aboutPage := PageContentRender{
		Title:   "About page",
		PageCSS: "about.page.css",
		PageJS:  "about.page.js",
	}
	router.GET("/about-us", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.page.html", aboutPage)
	})
  err := router.Run(":3000")
	if err != nil {
		log.Println(err)
	}
}
