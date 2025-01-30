package utils

import (
	"math/rand"
	"strconv"
)

var (
	mx = 9
	mn = 0
)

func GenerateOTP(length int) string {
	otp := ""
	for i := 0; i < length; i++ {
		randomIntegerwithinRange := rand.Intn(mx-mn) + mn
		otp += strconv.Itoa(randomIntegerwithinRange)
	}

	return otp
}
