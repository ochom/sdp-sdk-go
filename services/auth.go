package services

import (
	"encoding/json"
	"fmt"

	"github.com/ochom/gttp"
)

// Authenticate requests auth token from server
func Authenticate(username, password string) (string, error) {
	data := map[string]string{
		"username": username,
		"password": password,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	url := getURL() + "/auth/login"

	res, status, err := gttp.NewRequest(url, headers, body).Post()
	if err != nil {
		return "", err
	}

	if status != 200 {
		return "", fmt.Errorf("status code: %d, %s", status, string(res))
	}

	var resp map[string]any
	if err := json.Unmarshal(res, &resp); err != nil {
		return "", err
	}

	val, ok := resp["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return fmt.Sprintf("%v", val), nil
}
