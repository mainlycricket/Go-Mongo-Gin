package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type SendEmailInput struct {
	To      string
	From    string
	Subject string
	Body    string
}

func GetOtpEmailInput(to string) *SendEmailInput {
	otp := generateOtp(6)

	input := SendEmailInput{
		To:      to,
		From:    os.Getenv("EMAIL"),
		Subject: "OTP Verification Code",
		Body:    fmt.Sprintf("OTP verification code: %s", otp),
	}

	return &input
}

func SendEmail(input *SendEmailInput) error {
	if input == nil {
		return errors.New("nil input struct")
	}

	log.Println(input)
	log.Println("mimicing sending email...")
	return nil
}
