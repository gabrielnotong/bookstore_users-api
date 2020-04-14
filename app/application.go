package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() // will create go routine for each received request
)

func StartApplication() {
	mapUrls()
	_ = router.Run()
}
