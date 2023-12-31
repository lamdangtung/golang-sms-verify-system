package data

type OTPData struct {
	PhoneNumber string `json: "phone_number,omitempty" validate: "required,e164"`
}

type VerifyOTPData struct {
	User *OTPData `json: "user,omitempty" validate: "required"`
	Code string   `json: "code,omitempty" validate: "required,numeric"`
}
