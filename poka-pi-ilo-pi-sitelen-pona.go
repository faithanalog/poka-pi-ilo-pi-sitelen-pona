package poka_pi_ilo_pi_sitelin_pona

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RenderSettings struct {
	WorkerHost string
	WorkerPort int
}

var DefaultRenderSettings = &RenderSettings{WorkerHost: "127.0.0.1", WorkerPort: 3002}

func RenderTokiPona(tokipona string, cfg *RenderSettings) ([]byte, error) {
	if cfg == nil {
		cfg = DefaultRenderSettings
	}

	url := fmt.Sprintf("http://%s:%d/", cfg.WorkerHost, cfg.WorkerPort)

	res, err := http.Post(url, "text/plain", strings.NewReader(tokipona))
	if err != nil {
		return nil, err
	}

	b64png, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	png, err := base64.StdEncoding.DecodeString(string(b64png))
	if err != nil {
		return nil, err
	}

	return png, nil
}
