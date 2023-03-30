package dto

import "errors"

var (
	// ErrUsernameRequired is returned when username is not provided
	ErrUsernameRequired = errors.New("username is required")

	// ErrPasswordRequired is returned when password is not provided
	ErrPasswordRequired = errors.New("password is required")

	// ErrCpIDRequired is returned when cpID is not provided
	ErrCpIDRequired = errors.New("cpID is required")

	// ErrCpNameRequired is returned when cpName is not provided
	ErrCpNameRequired = errors.New("cpName is required")

	// ErrSenderIDRequired is returned when senderID is not provided
	ErrSenderIDRequired = errors.New("senderID is required")

	// ErrRequestIDRequired is returned when requestID is not provided
	ErrRequestIDRequired = errors.New("requestID is required")

	// ErrRecipientRequired is returned when recipient is not provided
	ErrRecipientRequired = errors.New("recipient is required")

	// ErrMessageRequired is returned when message is not provided
	ErrMessageRequired = errors.New("message is required")

	// ErrCallbackURLRequired is returned when callbackURL is not provided
	ErrCallbackURLRequired = errors.New("callbackURL is required")

	// ErrOfferCodeRequired is returned when offerCode is not provided
	ErrOfferCodeRequired = errors.New("offerCode is required")

	// ErrPhoneRequired is returned when phone is not provided
	ErrPhoneRequired = errors.New("phone is required")
)
