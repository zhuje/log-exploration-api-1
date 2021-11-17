package loki

import (
	"fmt"
	"github.com/ViaQ/log-exploration-api/pkg/logs"
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





