package utils

import (
	"math/rand"
	"time"
)

func GenerateTicketNumber() (ticketNumber string) {
	rand.Seed(time.Now().UnixNano())
	firstChar := 'A' + rune(rand.Intn('Z'-'A'+1))
	secondChar := '0' + rune(rand.Intn('9'-'0'+1))
	thirdChar := '0' + rune(rand.Intn('9'-'0'+1))
	fourthChar := '0' + rune(rand.Intn('9'-'0'+1))
	ticketNumber = string(firstChar) + string(secondChar) + string(thirdChar) + string(fourthChar)
	return
}
