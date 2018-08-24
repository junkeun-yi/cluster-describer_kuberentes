package functions

import (
	"context"
	"time"
	"fmt"

	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/utils"
	"strings"
	"strconv"
)


func (f FunctionSet) query(query string) map[string]float64{
	res, err := f.Prometheus.Query(context.Background(), query, time.Now())
	if err != nil {
		fmt.Printf("%v", err)
	}
	return utils.QueryStringToMap(res.String())
}

// Changes the prometheus Query Result String to a map
// resStr formatted as repetitions of {resourceName="$(Name)"} => $(Value) @[$(timestamp)]
func queryStringToMap(resStr string) map[string]float64 {
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
