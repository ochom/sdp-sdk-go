package services

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ochom/sdk-go/dto"
)

const devURL = "https://dtsvc.safaricom.com:8480/api"
const prodURL = "https://dsvc.safaricom.com:9480/api"

func getURL() string {
	val := os.Getenv("DEPLOYMENT_MODE")

	if val == "production" {
		return prodURL
	}

	return devURL
}

// newHttpClient is a global http client
func newHttpClient() (*http.Client, error) {
	return &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}, nil
}

// requestDo
func requestDo(url string, headers map[string]string, body []byte) (*dto.Response, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := newHttpClient()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	return dto.NewResponse(200, bodyBytes), nil
}
