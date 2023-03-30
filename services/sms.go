package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ochom/gttp"
	"github.com/ochom/sdk-go/dto"
)

// SendPremiumSms sends premium sms
func SendPremiumSms(accessToken string, req *dto.PremiumSmsRequest) (*dto.Response, error) {
	data := map[string]any{
		"requestId": req.RequestID,
		"channel":   "APIGW",
		"operation": "SendSMS",
	}

	requestParamData := []map[string]string{
		{
			"name":  "Msisdn",
			"value": req.Recipient,
		},
		{
			"name":  "Content",
			"value": req.Message,
		},
		{
			"name":  "OfferCode",
			"value": req.OfferCode,
		},
		{
			"name":  "CpId",
			"value": req.CpID,
		},
	}

	if req.LinkID != "" {
		requestParamData = append(requestParamData, map[string]string{
			"name":  "LinkID",
			"value": req.LinkID,
		})
	}

	data["requestParam"] = map[string]any{
		"data": requestParamData,
	}

	headers := map[string]string{
		"X-Authorization": "Bearer " + accessToken,
		"Content-Type":    "application/json",
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := getURL() + "public/SDP/sendSMSRequest"
	res, status, err := gttp.NewRequest(url, headers, body).Post()
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("status code: %d, %s", status, string(res))
	}

	return dto.NewResponse(status, "Okay", res), nil
}

// SendBulkSms sends bulk sms
func SendBulkSms(accessToken string, req *dto.BulkSmsRequest) (*dto.Response, error) {
	data := map[string]any{
		"timeStamp": time.Now().Format("20060102150405"),
		"dataSet": []map[string]any{
			{
				"uniqueId":          req.RequestID,
				"userName":          req.CpName,
				"channel":           "sms",
				"oa":                req.SenderID,
				"msisdn":            req.Recipient,
				"message":           req.Message,
				"actionResponseURL": req.CallbackURL,
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

	url := getURL() + "public/CMS/bulksms"
	res, status, err := gttp.NewRequest(url, headers, body).Post()
	if err != nil {
		return nil, err
	}

	if status != 200 {
		return nil, fmt.Errorf("status code: %d, %s", status, string(res))
	}

	return dto.NewResponse(status, "Okay", res), nil
}
