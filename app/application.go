package app

import (
	"github.com/gabrielnotong/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default() // will create go routine for each received request
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application")
	_ = router.Run()
}
