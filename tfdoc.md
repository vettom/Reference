## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_container_cluster.cluster](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/container_cluster) | resource |
| [google_project.project](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/project) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_authenticator_security_group"></a> [authenticator\_security\_group](#input\_authenticator\_security\_group) | GKE RBAC security group. must start with gke-security-groups | `string` | `"gke-security-groups@gomedia.io"` | no |
| <a name="input_cluster_resource_labels"></a> [cluster\_resource\_labels](#input\_cluster\_resource\_labels) | The GCE resource labels (a map of key/value pairs) to be applied to the cluster | `map(string)` | <pre>{<br>  "owner": "gomedia"<br>}</pre> | no |
| <a name="input_default_node_pool_node_count"></a> [default\_node\_pool\_node\_count](#input\_default\_node\_pool\_node\_count) | Number of nodes in default nodepool | `string` | `1` | no |
| <a name="input_description"></a> [description](#input\_description) | The description of the cluster | `string` | `""` | no |
| <a name="input_disable_http_load_balancing"></a> [disable\_http\_load\_balancing](#input\_disable\_http\_load\_balancing) | Enable http loadbalancing | `bool` | `false` | no |
| <a name="input_enable_gke_backup"></a> [enable\_gke\_backup](#input\_enable\_gke\_backup) | Enable GKE backup | `bool` | `false` | no |
| <a name="input_enable_network_policy"></a> [enable\_network\_policy](#input\_enable\_network\_policy) | Enable network policy addon | `bool` | `true` | no |
| <a name="input_enable_private_endpoint"></a> [enable\_private\_endpoint](#input\_enable\_private\_endpoint) | private endpoint if required | `bool` | `true` | no |
| <a name="input_enable_private_nodes"></a> [enable\_private\_nodes](#input\_enable\_private\_nodes) | true configured private cluster | `bool` | `true` | no |
| <a name="input_horizontal_pod_autoscaling"></a> [horizontal\_pod\_autoscaling](#input\_horizontal\_pod\_autoscaling) | Enable or disable horizontal pod autoscaling | `bool` | `false` | no |
| <a name="input_ip_range_pods"></a> [ip\_range\_pods](#input\_ip\_range\_pods) | The _name_ of the secondary subnet ip range to use for pods | `string` | n/a | yes |
| <a name="input_ip_range_services"></a> [ip\_range\_services](#input\_ip\_range\_services) | The _name_ of the secondary subnet range to use for services | `string` | n/a | yes |
| <a name="input_kubernetes_version"></a> [kubernetes\_version](#input\_kubernetes\_version) | The Kubernetes version of the masters. If set to 'latest' it will pull latest available version in the selected region. | `string` | `"latest"` | no |
| <a name="input_logging_services"></a> [logging\_services](#input\_logging\_services) | Logging services to be enabled. | `list(string)` | <pre>[<br>  "SYSTEM_COMPONENTS",<br>  "WORKLOADS"<br>]</pre> | no |
| <a name="input_maintenance_start_time"></a> [maintenance\_start\_time](#input\_maintenance\_start\_time) | Time window specified for daily or recurring maintenance operations in RFC3339 format | `string` | `"03:00"` | no |
| <a name="input_master_authorized_networks"></a> [master\_authorized\_networks](#input\_master\_authorized\_networks) | List of master authorized networks. If none are provided, disallow external access (except the cluster node IPs, which GKE automatically whitelists). | `list(object({ cidr_block = string, display_name = string }))` | <pre>[<br>  {<br>    "cidr_block": "10.12.0.0/14",<br>    "display_name": "GoMedia Office"<br>  }<br>]</pre> | no |
| <a name="input_masters_private_cidr"></a> [masters\_private\_cidr](#input\_masters\_private\_cidr) | private network to host the cluster in (required if private cluster) | `string` | `null` | no |
| <a name="input_monitoring_enabled_components"></a> [monitoring\_enabled\_components](#input\_monitoring\_enabled\_components) | List of components to monitor | `list(string)` | `[]` | no |
| <a name="input_name"></a> [name](#input\_name) | The name of the cluster (required) | `string` | n/a | yes |
| <a name="input_network_policy_config_disable"></a> [network\_policy\_config\_disable](#input\_network\_policy\_config\_disable) | The network policy config | `bool` | `false` | no |
| <a name="input_network_policy_provider"></a> [network\_policy\_provider](#input\_network\_policy\_provider) | The network policy provider. | `string` | `"CALICO"` | no |
| <a name="input_region"></a> [region](#input\_region) | The region to host the cluster in (optional if zonal cluster / required if regional) | `string` | `"europe-west1"` | no |
| <a name="input_remove_default_node_pool"></a> [remove\_default\_node\_pool](#input\_remove\_default\_node\_pool) | true to remove default nodepool | `bool` | `true` | no |
| <a name="input_subnetwork"></a> [subnetwork](#input\_subnetwork) | The subnetwork to host the cluster in (required) | `string` | n/a | yes |
| <a name="input_vpc_network"></a> [vpc\_network](#input\_vpc\_network) | Name of the VPC network to use | `string` | n/a | yes |
| <a name="input_zones"></a> [zones](#input\_zones) | The zones to host the cluster in | `list(string)` | <pre>[<br>  "europe-west1-b",<br>  "europe-west1-c",<br>  "europe-west1-d"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_clustername"></a> [clustername](#output\_clustername) | n/a |
