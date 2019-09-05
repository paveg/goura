package api

// Age represents user age
type Age int

// Weight represents user weight
type Weight float32

// Height represents user height
type Height float32

// Gender represents user gender
type Gender string

// Email represents user email address
type Email string

// UserID represents user identifier
type UserID string

// UserInfo represents user information
type UserInfo struct {
	Age    Age    `json:"age"`
	Weight Weight `json:"weight"`
	Height Height `json:"height"`
	Gender Gender `json:"gender"`
	Email  Email  `json:"email"`
	UserID UserID `json:"user_id"`
}
