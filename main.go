package main

import (
	"github.com/junkeun-yi/kube-informer/pkg/client"
	"github.com/junkeun-yi/kube-informer/pkg/controller"
	"github.com/junkeun-yi/kube-informer/pkg/functions"
)


// Runs the controller and starts the server
func main() {

	kubeconfig := client.GetKubeConfig()
	prom := client.GetPrometheus("http://a9876c924a63311e8988a0692683aa03-352606435.us-west-1.elb.amazonaws.com:9090")

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