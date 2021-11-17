package logscontroller

import (
	"github.com/ViaQ/log-exploration-api/pkg/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)


func NewLokiLogsController(log *zap.Logger, logsProvider logs.LogsProvider, router *gin.Engine) *LogsController {
	controller := &LogsController{
		log:          log,
		logsProvider: logsProvider,
	}

	r := router.Group("loki")
	r.GET("/filter",  controller.TestLogs)

	return controller
}

func (controller *LogsController) TestLogs(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, gin.H{"Logs": "To Be Implemented"})
}

