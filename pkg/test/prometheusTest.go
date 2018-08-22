package test

import (

	"fmt"
	"context"
	"time"
	"strings"
	"strconv"
	// prometheus_api "github.com/prometheus/client_golang/api"
	// prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	// "github.com/prometheus/common/model"

)

const nodeCPU = "sum(rate(node_cpu_seconds_total{mode!='idle'}[2m])) by (instance) / count(node_cpu_seconds_total{mode='idle'}) by (instance)"
const nodeMemory = "sum(node_memory_Active_bytes) by (instance) / count(node_memory_Active_bytes) by (instance)"
const podCPU = "sum (rate (container_cpu_usage_seconds_total{image!='', name=~'k8s_POD.*', namespace!='kube-system', namespace!='monitoring'}[1m])) by (pod_name)"
const podMemory = "sum(container_memory_usage_bytes{name=~'k8s_POD.*', namespace!='kube-system', namespace!='monitoring'}) by (pod_name)"

func (c *Controller) QueryExample() {
	res, _ := c.PromAPI.Query(context.Background(), nodeCPU, time.Now())
	c.Logger.Infof("%v\n", res)
	res, _ = c.PromAPI.Query(context.Background(), nodeMemory, time.Now())
	c.Logger.Infof("%v\n", res)
	res, _ = c.PromAPI.Query(context.Background(), podCPU, time.Now())
	c.Logger.Infof("%v\n", res)
	res, _ = c.PromAPI.Query(context.Background(), podMemory, time.Now())
	c.Logger.Infof("%v\n", res)
}

func (c *Controller) QueryNodeCPU() {

	res, err := c.PromAPI.Query(context.Background(), nodeCPU, time.Now())
	if err != nil {
		panic(err.Error())
	}
	// c.Logger.Infof("%v\n", res)
	resStr := res.String()

	fmt.Printf("\n$%v$\n", resStr)

	metMap, _ := queryStringToMap(resStr)

	fmt.Printf("\n!%v!\n", metMap)
}

func queryStringToMap(resStr string) (map[string]float64, []string){
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
	resourceKeys := make([]string, size)

	for _, pair := range keyVal {
		if keyStart := strings.Index(pair, "="); keyStart != -1 {
			keyEnd := strings.Index(pair, "}")
			key := pair[keyStart+1:keyEnd]
			valStart := strings.Index(pair, "=> ") + 3
			valueStr := pair[valStart:len(pair)]
			valueStr = valueStr[0:strings.Index(valueStr, " ")]
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				panic(err.Error())
			}
			metMap[key] = value
			resourceKeys = append(resourceKeys, key)
		}
	}
	
	return metMap, resourceKeys
}
