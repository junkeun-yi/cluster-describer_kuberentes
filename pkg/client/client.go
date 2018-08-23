package client

import (
	"flag"
	"path/filepath"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	prometheus_api "github.com/prometheus/client_golang/api"
	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// Gets the kubeconfig file and flags it
func GetKubeConfig() *string{
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return kubeconfig
}

// Creates a kubernetes out of cluster client with client-go
func GetClientOutOfCluster(kubeconfig *string) *kubernetes.Clientset{
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func GetMetricsClientOutOfCluster(kubeconfig *string) *metrics.Clientset{
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := metrics.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func GetPrometheus(prom_url string) *prometheus_api.Client{
	var config = prometheus_api.Config{
		Address: prom_url,
	}
	client, _ := prometheus_api.NewClient(config)
	return &client
}

func GetPrometheusApi(promCli *prometheus_api.Client) prometheus_v1.API{
	api := prometheus_v1.NewAPI(*promCli)
	return api
}
