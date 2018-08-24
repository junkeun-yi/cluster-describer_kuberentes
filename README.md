# kube-informer

To run:

1. Make sure a kuberentes cluster is running.
2. create all manifests by running 
    kubectl apply -f manifests/
3. get the prometheus external url by running

		kubectl get services -n monitoring prometheus-service -o wide
- create the prometheus url by formatting it as "http://(prometheus external url from above):9090"
- copy this url into pkg/config/varables.go.PrometheusURL
4. run the app by running

		go run main.go

TODO: Make Makefile, add vendor