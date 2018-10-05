package main

import (
	"fmt"
	sitelen "github.com/faithanalog/poka-pi-ilo-pi-sitelen-pona"
	"io"
	"io/ioutil"
	"log"
	"net/http"
    "os"
    "strconv"
)

func main() {
    renderSettings := *sitelen.DefaultRenderSettings

    envWorkerHost := os.Getenv("WORKER_HOST")
    if envWorkerHost != "" {
        renderSettings.WorkerHost = envWorkerHost
    }

    envWorkerPort := os.Getenv("WORKER_PORT")
    if envWorkerPort != "" {
        p, err := strconv.Atoi(envWorkerPort)
        if err != nil {
            renderSettings.WorkerPort = p
        }
    }

	http.HandleFunc("/sitelen-pona", func(res http.ResponseWriter, req *http.Request) {
		// Accept at most 4096 characters
		const maxInputSize = 1024 * 4

		if req.ContentLength > maxInputSize {
			// 413 payload too large
			res.WriteHeader(413)

			res.Write([]byte(fmt.Sprintf("Error: Maximum input size is %d bytes.\n", maxInputSize)))
		}

		// Sanitize the input to maxInputSize
		reader := io.LimitReader(req.Body, maxInputSize)

		tokipona, err := ioutil.ReadAll(reader)
		if err != nil {
			res.WriteHeader(500)
			return
		}

		output, err := sitelen.RenderTokiPona(string(tokipona), &renderSettings)
		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(output)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:3004", nil))
}
