package models

type Users struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Uid      string `json:"Uid"`
	Picture  string `json:"Picture"` //base64-encoded string
}
