package dto

type Status string

var (
	Verified   Status = "verified"
	Unverified Status = "unverified"
)

type UserCreationSuccess struct {
	Status Status
	Data   any
}

type UserSignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
