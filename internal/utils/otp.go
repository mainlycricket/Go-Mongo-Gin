package utils

import (
	"math/rand/v2"
	"strconv"
)

func generateOtp(length int) string {
	var otp string

	for range length {
		digit := rand.IntN(10)
		otp += strconv.Itoa(digit)
	}

	return otp
}
