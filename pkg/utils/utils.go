package utils

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/op/go-logging"
)

// Constants
const CHANNEL_MAX_SIZE = 100
var Logger *logging.Logger = nil

func InitLogging() {
	format := logging.MustStringFormatter(
		`%{color}%{module:-8s} ▶ %{level:+5s} %{id:04d}%{color:reset} %{message}`,
	)
	logging.SetFormatter(format)
	Logger = logging.MustGetLogger("clientgo")
}

// Returns the first error that is not nil
func CheckAllErrors(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

// Returns the string rep of the current time
func GetTimeString() string {
	t := time.Now()
	timeString := fmt.Sprintf("%s %s, %s:%s:%s",
		t.Month(),
		fmt.Sprintf("%02d", t.Day()),
		fmt.Sprintf("%02d", t.Hour()),
		fmt.Sprintf("%02d", t.Minute()),
		fmt.Sprintf("%02d", t.Second()),
	)
	return timeString
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomNHash(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}