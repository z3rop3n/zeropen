package types

type User struct {
	Id            *string `json:"id"`
	Email         *string `json:"email"`
	FirstName     *string `json:"firstName"`
	LastName      *string `json:"lastName"`
	ProfilePicUrl *string `json:"profilePic"`
}

type OTP struct {
	Email *string `json:"email"`
	OTP   *string `json:"otp"`
}
