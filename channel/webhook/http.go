package webhook

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func sendJson(method string, url string, body JobWebHook) {
	j, err := json.Marshal(body)
	if err != nil {
		logrus.Errorf("json marshal failed, err: %v", err)
		return
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(j)))
	if err != nil {
		logrus.Errorf("new request failed, err: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	send(req)
}

func send(req *http.Request) {
	req.Header.Add("Referer", "go-job")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("http request failed, err: %v", err)
		return
	}
	defer autoClose(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("http request failed, status code: %v", resp.StatusCode)
		return
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Warnf("read response body failed, err: %v", err)
		return
	}
	logrus.Debugf("http request success, status code: %v, body: %s", resp.StatusCode, string(bytes))
}

func autoClose(rc io.ReadCloser) {
	_ = rc.Close()
}
