package test

import (
	// "fmt"
	// "sync/atomic"

	// v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "github.com/netsys/triggers2/kubehandler/pkg/utils"
	// "github.com/netsys/triggers2/kubehandler/pkg/state"
)

func (c *Controller) printNodeUtilSpecs() {
	nodes, err := c.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if (err != nil) {
		c.Logger.Debugf("%v", err)
	}
	for _, n := range nodes.Items {
		c.Logger.Infof("==%v==\n", n.Name)
		c.Logger.Infof("%v, %v", n.Status.Capacity["cpu"], n.Status.Capacity["memory"])
	}
}


// prints the utilization for nodes
func (c *Controller) printNodeMetrics() {

	metrics, err := c.MetClient.MetricsV1beta1().NodeMetricses().List(metav1.ListOptions{});
	if (err != nil) {
		c.Logger.Debugf("%v", err)
	}
	for _, metric := range metrics.Items {
		c.Logger.Infof("%v", metric.ObjectMeta.Name)
		c.Logger.Infof("%v", metric.Usage)
	}
}

// prints the utilization for pods
func (c *Controller) printPodMetrics() {

	metrics, err := c.MetClient.MetricsV1beta1().PodMetricses("").List(metav1.ListOptions{});
	if (err != nil) {
		c.Logger.Debugf("%v", err)
	}
	for _, metric := range metrics.Items {

		for _, container := range metric.Containers {
			c.Logger.Infof("%v\n", container.Name)
			c.Logger.Infof("%v\n", container.Usage["cpu"])
		}

		/**
		// Gets the PodMetricsObject, but doesn't seem tooooooo useful....T.T
		met, er := c.MetClient.MetricsV1beta1().PodMetricses("").Get(metric.Name, metav1.GetOptions{})
		if er!= nil {
			c.Logger.Debugf("%v", err)
		}
		c.Logger.Infof("%v", met)
		*/
	}
}


