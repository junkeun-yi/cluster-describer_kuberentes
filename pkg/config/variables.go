package config

const (

	PrometheusURL = ""

	NodeCPUUsageQuery = "sum(rate(node_cpu_seconds_total{mode!='idle'}[2m])) by (instance) / count(node_cpu_seconds_total{mode='idle'}) by (instance)"
	NodeMemoryUsageQuery = "sum(node_memory_Active_bytes) by (instance) / count(node_memory_Active_bytes) by (instance)"
	PodCPUUsageQuery = "sum (rate (container_cpu_usage_seconds_total{image!='', pod_name!=''}[1m])) by (pod_name)"
	PodMemoryUsageQuery = "sum( rate (container_memory_usage_bytes{image!='', pod_name!=''}[1m])) by (pod_name)"

	PodCPUUsageExceptKubeSystemAndMonitoringNamespacesQuery = "sum (rate (container_cpu_usage_seconds_total{image!='', pod_name!='', namespace!='kube-system', namespace!='monitoring'}[1m])) by (pod_name)"
	PodMemoryUsageExceptKubeSystemAndMonitoringNamespacesQuery = "sum( rate (container_memory_usage_bytes{pod_name!='', namespace!='kube-system', namespace!='monitoring'}[1m])) by (pod_name)"

	NodeCPUCoresQuery = "sum(machine_cpu_cores) by (instance)"
	NodeMemoryBytesQuery = "sum(machine_memory_bytes) by (instance)"

	NodeHTTPRequestsPerMin = "sum( rate (http_requests_total[1m])) by (instance)"

)
