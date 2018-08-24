package main

import (
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/client"
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/controller"
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/functions"
)


// Runs the controller and starts the server
func main() {

	kubeconfig := client.GetKubeConfig()
	prom := client.GetPrometheus("http://ac39526f5a7c911e8964d060f0b9aa92-8462892.us-west-1.elb.amazonaws.com:9090")

	var functionSet = functions.FunctionSet{
		MetClient: client.GetMetricsClientOutOfCluster(kubeconfig),
		Client: client.GetClientOutOfCluster(kubeconfig),
		Prometheus: client.GetPrometheusApi(prom),
	}

	// Initialise a controller
	var control = controller.Controller{
		FunctionSet: functionSet,
	}

	control.Run()

}