package models

type User struct {
	Username  string   `json:"Username"`
	Email     string   `json:"Email"`
	Uid       string   `json:"Uid"`
	Picture   *string  `json:"Picture"` //base64-encoded string
	Following []string `json:"Following"`
	Friends   []string `json:"Friends"`
}
