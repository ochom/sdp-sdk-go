package dto

// BulkSmsRequest is the request body for the bulk sms endpoint
type BulkSmsRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CpID        string `json:"cpID"`
	CpName      string `json:"cpName"`
	SenderID    string `json:"senderID"`
	RequestID   string `json:"requestID"`
	Recipient   string `json:"recipient"`
	Message     string `json:"message"`
	CallbackURL string `json:"callbackURL"`
}

// Validate validates the bulk sms request
func (r *BulkSmsRequest) Validate() error {
	if r.Username == "" {
		return ErrUsernameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	if r.CpID == "" {
		return ErrCpIDRequired
	}
	if r.CpName == "" {
		return ErrCpNameRequired
	}
	if r.SenderID == "" {
		return ErrSenderIDRequired
	}
	if r.RequestID == "" {
		return ErrRequestIDRequired
	}
	if r.Recipient == "" {
		return ErrRecipientRequired
	}
	if r.Message == "" {
		return ErrMessageRequired
	}
	if r.CallbackURL == "" {
		return ErrCallbackURLRequired
	}

	return nil
}

// PremiumSmsRequest is the request body for the premium sms endpoint
type PremiumSmsRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CpID      string `json:"cpID"`
	RequestID string `json:"requestID"`
	OfferCode string `json:"offerCode"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
	LinkID    string `json:"linkID"`
}

// Validate validates the premium sms request
func (r *PremiumSmsRequest) Validate() error {
	if r.Username == "" {
		return ErrUsernameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	if r.CpID == "" {
		return ErrCpIDRequired
	}
	if r.RequestID == "" {
		return ErrRequestIDRequired
	}
	if r.OfferCode == "" {
		return ErrOfferCodeRequired
	}
	if r.Recipient == "" {
		return ErrRecipientRequired
	}
	if r.Message == "" {
		return ErrMessageRequired
	}

	return nil
}

// SubscriptionRequest is the request body for the subscription endpoint
type SubscriptionRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CpID      string `json:"cpID"`
	RequestID string `json:"requestID"`
	OfferCode string `json:"offerCode"`
	Phone     string `json:"phoneNumber"`
}

// Validate validates the subscription request
func (r *SubscriptionRequest) Validate() error {
	if r.Username == "" {
		return ErrUsernameRequired
	}
	if r.Password == "" {
		return ErrPasswordRequired
	}
	if r.CpID == "" {
		return ErrCpIDRequired
	}
	if r.RequestID == "" {
		return ErrRequestIDRequired
	}
	if r.OfferCode == "" {
		return ErrOfferCodeRequired
	}
	if r.Phone == "" {
		return ErrPhoneRequired
	}
	return nil
}

// Response is the response body for all endpoints
type Response struct {
	Success      bool   `json:"success"`
	StatusCode   int    `json:"statusCode"`
	StatusText   string `json:"statusText"`
	ResponseBody []byte `json:"responseBody"`
}

// NewSuccessResponse returns a new success response
func NewResponse(statusCode int, statusText string, responseBody []byte) *Response {
	return &Response{
		Success:      true,
		StatusCode:   statusCode,
		StatusText:   statusText,
		ResponseBody: responseBody,
	}
}
