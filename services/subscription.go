package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ochom/gttp"
	"github.com/ochom/sdk-go/dto"
)

// Subscribe requests subscription from server
func Subscribe(accessToken string, req *dto.SubscriptionRequest) (*dto.Response, error) {
	data := map[string]interface{}{
		"requestId":        req.RequestID,
		"requestTimeStamp": time.Now().Format("20060102150405"),
		"channel":          "SMS",
		"operation":        "ACTIVATE",
		"requestParam": map[string]interface{}{
			"data": []map[string]string{
				{
					"name":  "OfferCode",
					"value": req.OfferCode,
				},
				{
					"name":  "Msisdn",
					"value": req.Phone,
				},
				{
					"name":  "Language",
					"value": "1",
				},
				{
					"name":  "CpId",
					"value": req.CpID,
				},
			},
		},
	}

	headers := map[string]string{
		"X-Authorization": "Bearer " + accessToken,
		"Content-Type":    "application/json",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := getURL() + "public/SDP/activate"

	res, status, err := gttp.NewRequest(url, headers, body).Post()
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("status code: %d, %s", status, string(res))
	}

	return dto.NewResponse(status, "Okay", res), nil
}

// UnSubscribe requests un-subscription from server
func UnSubscribe(accessToken string, req *dto.SubscriptionRequest) (*dto.Response, error) {
	data := map[string]interface{}{
		"requestId":        req.RequestID,
		"requestTimeStamp": time.Now().Format("20060102150405"),
		"channel":          "SMS",
		"operation":        "DEACTIVATE",
		"requestParam": map[string]interface{}{
			"data": []map[string]string{
				{
					"name":  "OfferCode",
					"value": req.OfferCode,
				},
				{
					"name":  "Msisdn",
					"value": req.Phone,
				},
				{
					"name":  "CpId",
					"value": req.CpID,
				},
			},
		},
	}

	headers := map[string]string{
		"X-Authorization": "Bearer " + accessToken,
		"Content-Type":    "application/json",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := getURL() + "public/SDP/deactivate"

	res, status, err := gttp.NewRequest(url, headers, body).Post()
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("status code: %d, %s", status, string(res))
	}

	return dto.NewResponse(status, "Okay", res), nil
}
