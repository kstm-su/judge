package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type C struct {
	Id   string
	Name string
}

func main() {
	router := gin.Default()
	//router.Static("/css", "./templates/css")
	router.Static("/css", "./templates/css")
	router.Static("/js", "./templates/js")
	router.LoadHTMLGlob("templates/*.tmpl")
	//router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/problems", func(c *gin.Context) {
		c.HTML(http.StatusOK, "problems.tmpl", gin.H{})
	})

	router.GET("/problem/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "problem.tmpl", gin.H{
			"statement": "Calculate Sum of \\(A + B\\)",
			"limit": []string{"\\(0 \\leq A \\leq 1000, A \\in \\Z \\)","\\(0 \\leq B \\leq 1000, A \\in \\Z \\)"},
			"testcase":  []C{{"In #1", "3 5"}, {"Out #1", "8"}},
		})
	})

	router.GET("/submission", func(c *gin.Context) {
		c.HTML(http.StatusOK, "submission.tmpl", gin.H{

		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.HTML(http.StatusOK, "status.tmpl", gin.H{

		})
	})

	router.GET("/api/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "api.tmpl", gin.H{

		})
	})
	router.Run(":8080")
}
