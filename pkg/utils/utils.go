package utils

import (
	"time"
	"fmt"
	"math/rand"
	"strings"
	"strconv"

	"math"
	"gopkg.in/inf.v0"
)

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

// Changes the prometheus Query Result String to a map
// resStr formatted as repetitions of {resourceName="$(Name)"} => $(Value) @[$(timestamp)]
func QueryStringToMap(resStr string) map[string]float64 {
	// since the prometheus query results are from the same timestamp,
	// separate the result based on the timestamp
	timeStart := strings.Index(resStr, "@")
	timeEnd := strings.Index(resStr, "]") + 1
	timeStr := resStr[timeStart:timeEnd]

	// pairs of key:values of the pulled metrics
	keyVal := strings.Split(resStr, timeStr)

	// size of returned values
	size := strings.Count(resStr, "=>")
	metMap := make(map[string]float64, size)

	// string operations to add Key:Value pairs to METMAP
	for _, pair := range keyVal {
		if keyStart := strings.Index(pair, "="); keyStart != -1 {
			keyEnd := strings.Index(pair, "}")
			key := pair[keyStart+2:keyEnd-1]
			valStart := strings.Index(pair, "=> ") + 3
			valueStr := pair[valStart:]
			valueStr = valueStr[0:strings.Index(valueStr, " ")]
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				panic(err.Error())
			}
			metMap[key] = value
		}
	}

	return metMap
}


func ChangeQuantityToFloat(quantity inf.Dec) float64{
	unscaled, faulted := quantity.Unscaled()
	if !faulted {
		return math.MaxFloat64
	}
	scale := math.Pow(10, float64(quantity.Scale()))
	res := float64(unscaled) / scale
	return res
}
