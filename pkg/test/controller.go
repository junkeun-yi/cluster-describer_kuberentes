package test

import (
	// "fmt"
	"k8s.io/client-go/kubernetes"
	"github.com/op/go-logging"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	prometheus_api "github.com/prometheus/client_golang/api"
	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type Controller struct {
	Client			*kubernetes.Clientset
	MetClient		*metrics.Clientset
	Prometheus		prometheus_api.Client
	PromAPI			prometheus_v1.API
	Logger			*logging.Logger
}

func (c *Controller) Run() {

	// c.QueryExample()
	c.QueryNodeCPU()

	//c.test()

	// fmt.Printf("%T\n", c.Prometheus)
	// fmt.Printf("%T\n", c.PromAPI)
	
	// c.printNodeUtilSpecs()
	// c.printNodeMetrics()
	// c.printPodMetrics();

	// c.printNodes();
	// c.printPods();
	// c.printDeployments();

	// c.failAllPods()
	// c.testPodRedeploy()
	// c.timedEnforce(time.Second * 10)
}

