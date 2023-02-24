package serve

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Serve_File() {
	/*
		fileServer := http.FileServer(http.Dir("../.next"))
		http.Handle("/", fileServer)
		http.ListenAndServe(":8080", nil)
	*/

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../.next", true)))

	/*
		// Setup route group for the API
		api := router.Group("/api")
		{
			api.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}
	*/

	// Start and run the server
	router.Run(":5000")

}
