package events

type UserCreatedEvent struct {
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	PasswordDigest string `json:"password_digest"`
}
