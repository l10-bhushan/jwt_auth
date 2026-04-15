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
