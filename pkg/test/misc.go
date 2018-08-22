package test

import (
	"fmt"
	// prometheus_api "github.com/prometheus/client_golang/api"
	// prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	// "github.com/prometheus/common/model"
)

type Test struct{
	test func()
}


func (c *Controller) test() {
	temp := Test{
		test: func() {
			fmt.Printf("%v\n", "HELLO")
		},
	}
	temp.test()
}




