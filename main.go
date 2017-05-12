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
	router.LoadHTMLGlob("templates/*.tmpl")
	//router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/problem", func(c *gin.Context) {
		c.HTML(http.StatusOK, "problem.tmpl", gin.H{
			"statement": "Calculate Sum of \\(A + B\\)",
			"limit": []string{"\\(0 \\leq A \\leq 1000, A \\in \\Z \\)","\\(0 \\leq B \\leq 1000, A \\in \\Z \\)"},
			"testcase":  []C{{"In #1", "3 5"}, {"Out #1", "8"}},
		})
	})
	router.Run(":8080")
}
