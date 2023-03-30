package services

import (
	"encoding/json"
	"time"

	"github.com/ochom/sdk-go/dto"
)

// SendPremiumSms sends premium sms
func SendPremiumSms(accessToken string, data *dto.PremiumSmsRequest) (*dto.Response, error) {
	payload := map[string]any{
		"requestId": data.RequestID,
		"channel":   "APIGW",
		"operation": "SendSMS",
	}

	requestParamData := []map[string]string{
		{
			"name":  "Msisdn",
			"value": data.Recipient,
		},
		{
			"name":  "Content",
			"value": data.Message,
		},
		{
			"name":  "OfferCode",
			"value": data.OfferCode,
		},
		{
			"name":  "CpId",
			"value": data.CpID,
		},
	}

	if data.LinkID != "" {
		requestParamData = append(requestParamData, map[string]string{
			"name":  "LinkID",
			"value": data.LinkID,
		})
	}

	payload["requestParam"] = map[string]any{
		"data": requestParamData,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Authorization": "Bearer " + accessToken,
		"Content-Type":    "application/json",
	}

	url := getURL() + "/public/SDP/sendSMSRequest"

	return requestDo(url, headers, body)
}

// SendBulkSms sends bulk sms
func SendBulkSms(accessToken string, data *dto.BulkSmsRequest) (*dto.Response, error) {
	payload := map[string]any{
		"timeStamp": time.Now().Format("20060102150405"),
		"dataSet": []map[string]any{
			{
				"uniqueId":          data.RequestID,
				"userName":          data.CpName,
				"channel":           "sms",
				"oa":                data.SenderID,
				"msisdn":            data.Recipient,
				"message":           data.Message,
				"actionResponseURL": data.CallbackURL,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"X-Authorization": "Bearer " + accessToken,
		"Content-Type":    "application/json",
	}

	url := getURL() + "/public/CMS/bulksms"

	return requestDo(url, headers, body)
}
