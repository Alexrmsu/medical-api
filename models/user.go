package models

type User struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}
