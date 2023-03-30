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

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	url := getURL() + "/auth/login"
	res, err := requestDo(url, headers, body)
	if err != nil {
		return "", err
	}

	var response map[string]string
	err = json.Unmarshal(res.ResponseBody, &response)
	if err != nil {
		return "", err
	}

	val, ok := response["accessToken"]
	if !ok {
		return "", fmt.Errorf("access token not found")
	}

	return val, nil

}
