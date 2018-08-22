package test

import (
	"time"
	// "fmt"
	// "sync/atomic"
	"math/rand"

	// v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "github.com/netsys/triggers2/kubehandler/pkg/utils"
	// "github.com/netsys/triggers2/kubehandler/pkg/state"
)

func (c *Controller) printNodes() {
	nodes, nodeGetErr := c.Client.CoreV1().Nodes().List(metav1.ListOptions{});
	if (nodeGetErr != nil) {
		c.Logger.Debug(nodeGetErr);
	}
	for _, node := range nodes.Items {
		c.Logger.Infof("\n\n\n%v\n\n\n", node)
	}
}

func (c *Controller) printPods() {
	c.Logger.Infof("##################");
	c.Logger.Infof("#................#");
	c.Logger.Infof("##......POD......");
	c.Logger.Infof("#................#");
	c.Logger.Infof("##################");
	pods, err := c.Client.CoreV1().Pods("default").List(metav1.ListOptions{});
	if (err != nil) {
		c.Logger.Debug(err);
	}
	for i, pod := range pods.Items {
		c.Logger.Infof("========================================");
		c.Logger.Infof("Pod %v\n", i);
		c.Logger.Infof("%v", pod);
		c.Logger.Infof("========================================");
	}
	c.Logger.Infof("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n\n\n");
}

func (c *Controller) printDeployments() {
	c.Logger.Infof("##################");
	c.Logger.Infof("#................#");
	c.Logger.Infof("##..DEPLOYMENT...#");
	c.Logger.Infof("#................#");
	c.Logger.Infof("##################");
	deployments, err := c.Client.AppsV1beta2().Deployments("default").List(metav1.ListOptions{});
	if (err != nil) {
		c.Logger.Debug(err);
	}
	for i, dep := range deployments.Items {
		c.Logger.Infof("========================================");
		c.Logger.Infof("Deployment %v\n", i);
		c.Logger.Infof("%v", dep);
		c.Logger.Infof("========================================");
	}
	c.Logger.Infof("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^\n\n\n");
}

func (c *Controller) timedEnforce(t time.Duration) {
	ticker := time.NewTicker(t)

	for {
		select {
			case <- ticker.C:
				c.randomDelete()
		}
	}
}

/**
func (c *Controller) failAllPods() {
	nodes, err := c.Client.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}
	for _, n := range nodes.Items {
		c.failAllPodsInNode(n.Name, n.Namespace)
	}
}
*/

/**
func (c *Controller) failAllPodsInNode(metaname, namespace string) {
	forceDeleteOption := metav1.NewDeleteOptions(0)
	node, nodeGetErr := c.Client.CoreV1().Nodes().Get(metaname, metav1.GetOptions{})
	if nodeGetErr != nil {
		c.Logger.Debug(nodeGetErr)
	}
	pods, podGetErr := c.Client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if podGetErr != nil {
		c.Logger.Debug(podGetErr)
	}
	for _, pod := range pods.Items {
		if pod.
	}
} 	
*/

func (c *Controller) testPodRedeploy() {
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	for _, p := range pods.Items {
		result, getErr := c.Client.CoreV1().Pods("default").Get(p.Name, metav1.GetOptions{})
		if getErr != nil {
			c.Logger.Debug(getErr)
		}
		c.Logger.Infof("%v", result.Name)

		c.Logger.Infof("%v", result.Spec.NodeName)

		c.Logger.Infof("%T", result.Spec.NodeName)

		c.Logger.Infof("%T", result)

		result.Spec.NodeName = "ip-172-20-39-172.us-west-1.compute.internal"


		_, updateErr := c.Client.CoreV1().Pods("default").Update(result)

		if updateErr != nil {
			c.Logger.Debug(updateErr)
		}

		return
	}
}


func (c *Controller) randomDelete() {
	/**
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	//configPods := make([]*v1.Pod, len(pods.Items))
	podNames := make(string, len(pods.Items))

	for _, pod := range pods.Items {
		if pod.Namespace == "default" {
			podNames = append(podNames, pod.Name)
			//configPods = append(configPods, &pod)
		}
	}
	
	randIndex = rand.
	*/

	podDeleteOption := metav1.NewDeleteOptions(0)
	
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	for _, pod := range pods.Items {
		temp := rand.Int31n(10)
		if temp < 3 {
			name := pod.Name
			if pod.Namespace == "default" {
				err = c.Client.CoreV1().Pods(pod.Namespace).Delete(name, podDeleteOption)
			}
			if err != nil {
				c.Logger.Debug(err)
				panic(err)
			}
			c.Logger.Infof("deleted pod: %v", name)
			return
		}
	}

	defer c.randomDelete()
	return
	
	/**
	podDeleteOption := metav1.NewDeleteOptions(0)

	err := c.Client.CoreV1().Pods("frontend").Delete("frontend-685d7ff496-8q8n9", podDeleteOption)
	if err != nil {
		c.Logger.Debug(err)
	}
	*/
}
