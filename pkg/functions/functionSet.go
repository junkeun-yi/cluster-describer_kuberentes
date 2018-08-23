package functions

import (
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"

	prometheus_v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type FunctionSet struct {
	Client			*kubernetes.Clientset
	MetClient		*metrics.Clientset
	Prometheus		prometheus_v1.API
}