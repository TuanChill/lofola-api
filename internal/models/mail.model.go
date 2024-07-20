package models

type DataOtpMail struct {
	Title  string
	To     string
	Name   string
	Otp    string
	Expire string
}

type EmailData struct {
	From    string
	To      string
	Subject string
	Body    string
}
