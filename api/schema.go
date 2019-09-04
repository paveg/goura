package api

// Age represents user age
type Age int

// Weight represents user weight
type Weight int

// Gender represents user gender
type Gender string

// Email represents user email address
type Email string

// UserInfo represents user information
type UserInfo struct {
	Age    Age    `json:"age"`
	Weight Weight `json:"weight"`
	Gender Gender `json:"gender"`
	Email  Email  `json:"email"`
}

// UserInfoResponse represents user information response
type UserInfoResponse struct {
	Age    Age
	Weight Weight
	Gender Gender
	Email  Email
}
