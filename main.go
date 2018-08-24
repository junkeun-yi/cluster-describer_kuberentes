package main

import (
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/client"
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/controller"
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/functions"
	"github.com/junkeun-yi/cluster-describer_kuberentes/pkg/config"
)


// Runs the controller and starts the server
func main() {

	kubeconfig := client.GetKubeConfig()
	prom := client.GetPrometheus(config.PrometheusURL)

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