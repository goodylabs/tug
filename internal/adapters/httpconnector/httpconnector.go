package httpconnector

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goodylabs/tug/internal/ports"
)

type httpConnector struct{}

func NewHttpConnector() ports.HttpConnector {
	return &httpConnector{}
}

func (h *httpConnector) HttpGetReq(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET request to %s failed with status %s", url, resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
