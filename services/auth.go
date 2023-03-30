package services

import (
	"encoding/json"
	"fmt"
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

	url := getURL() + "/auth/login"
	res, err := requestDo(url, nil, body)
	if err != nil {
		return "", err
	}

	val, ok := res.ResponseBody["token"]
	if !ok {
		return "", fmt.Errorf("access token not found")
	}

	return val.(string), nil

}
