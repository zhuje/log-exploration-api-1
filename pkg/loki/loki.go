package loki

import (
	"fmt"
	"github.com/ViaQ/log-exploration-api/pkg/logs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

// TODO -- replace localhost by access service components via kubernetes api
// https://kubernetes.io/docs/tasks/administer-cluster/access-cluster-api/
const URL = "http://localhost:3100/"

type LokiRepository struct {
	lokiClient *http.Client
	log        *zap.Logger
}

func NewLokiRepository(log *zap.Logger) (logs.LogsProvider, error) {
	c := &http.Client{Timeout: time.Duration(1) * time.Second}
	repository := &LokiRepository{
		log:        log,
		lokiClient: c,
	}
	return repository, nil
}

func (l LokiRepository) GetLogs() []byte {
	//c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := l.lokiClient.Get(URL)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body

	//controller := &LogsController{
	//	log:          log,
	//	logsProvider: logsProvider,
	//}
	//
	//r := router.Group("loki")
	//r.GET("filter",
	//}))
	//
	//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
}



// LOGS CONTROLLER -- routes incoming http request

type LogsController struct {
	logsProvider logs.LogsProvider
	log          *zap.Logger
}

func NewLogsController(log *zap.Logger, logsProvider logs.LogsProvider, router *gin.Engine) *LogsController {
	controller := &LogsController{
		log:          log,
		logsProvider: logsProvider,
	}

	r := router.Group("loki")
	r.GET("/filter",  l.lokiClient.Get(URL))

	return controller
}

func (controller *LogsController) FilterEntityLogs(gctx *gin.Context) {

	gctx.JSON(http.StatusOK, gin.H{"Logs": "To Be Implemented"})

}








// LogsProvider interface -- methods that need to be implemented

func (l LokiRepository) FilterLogs(params logs.Parameters) ([]string, error) {
	panic("implement me")
}

func (l LokiRepository) FilterContainerLogs(params logs.Parameters) ([]string, error) {
	panic("implement me")
}

func (l LokiRepository) FilterLabelLogs(params logs.Parameters, labelList []string) ([]string, error) {
	panic("implement me")
}

func (l LokiRepository) FilterNamespaceLogs(params logs.Parameters) ([]string, error) {
	panic("implement me")
}

func (l LokiRepository) FilterPodLogs(params logs.Parameters) ([]string, error) {
	panic("implement me")
}

// Logs	-- is called by the controller and returns all logs,
func (l LokiRepository) Logs(params logs.Parameters) ([]string, error) {
	panic("implement me")
}


func (l LokiRepository) CheckReadiness() bool {
	panic("implement me")
}





