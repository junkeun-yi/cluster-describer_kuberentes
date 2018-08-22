package main

import (
	
	"github.com/junkeun-yi/kubernetestest/pkg/state"
	"github.com/junkeun-yi/kubernetestest/pkg/utils"
	"github.com/junkeun-yi/kubernetestest/pkg/test"
	"github.com/op/go-logging"
)


// Runs the controller and starts the server
func main() {

	kubeconfig := state.GetKubeConfig();

	utils.InitLogging()

	prom := state.GetPrometheus("http://abe9bbf9d8fa511e8844006e3ea88fe6-1987808729.us-west-1.elb.amazonaws.com:9090")

	// Initialise a controller
	var control = test.Controller{
		MetClient: state.GetMetricsClientOutOfCluster(kubeconfig),
		Client: state.GetClientOutOfCluster(kubeconfig),
		Prometheus: prom,
		PromAPI: state.GetPrometheusApi(&prom),
		Logger: logging.MustGetLogger("control"),
	}

	control.Run()

}