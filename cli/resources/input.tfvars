/* 
This is the template for user input. You may uncomment and change the configuration and save it as your input. 
If you don't need the variable, leave it commented.
If your variable has the same value as the default value, you may also leave it commented. 
*/


/* The variables uncommented here are mandatory. */

//The full cluster name
cluster_name              = 

//project id
project                   = 

//Need to get from Infoblox
master_ipv4_cidr_block    = 

//The next three are the subnet details
subnetwork                = 
ip_range_pods             = 
ip_range_services         = 

//If your cluster will enable ASM. If yes, the meshid label will be added to the cluster
#asm                       = false

//The project that shares the VPC with your project with
#network_project_id        = "otl-vpc-shared"

//This is the region of the VPC. It does NOT indicate if your cluster is a regional cluster or zonal
#region                    = "us-east4"

//The VPC network
#network                   = "otl-us-east4"

//Whether or not to enable binary authorization on the cluster
#enable_binary_authorization = false

//Change to true if you want a regional cluster instead of zonal
#regional                  = false

//If you choose regional cluster, this is the variable to set the region of the cluster
#cluster_region            = "us-east4"

//The location (zone or region) this cluster has been created in. For zonal cluster, put zone info
//i.e., "us-east4-a". For regional cluster, put the region. i.e., "us-east4". 
location                  = 

//The GKE version can only be set when you choose not to subscribe to a release channel. If you do set the release_channel,
//make sure the gke_version = null. Release_channel value is capitalized.
#gke_version               = "1.22.8-gke.202"
#release_channel           = "UNSPECIFIED"

//auto_upgrade can only be enabled if you subscribe to a release channle. Different release channels
//have different requirements on auto upgrade requirements. Check GKE documentation for more details.
#auto_upgrade              = false

//Optional description of the cluster
#cluster_description       = ""

//Number of nodes. Note if you choose regional cluster, the number will x3 when creating nodes.
// i.e. For regional clusters, if you put node_count=1, there will be 3 nodes created.
// For zonal clusters, this is literal meaning of the nodes.
#node_count                = 4

//Change to true if you want nodepool autoscaling
//autoscaling              = false
//The two numbers are only meaningful when autosacling is set to true. Both numbers are per zone. 
//If you choose regional cluster, the number will be x3 when adding nodes.
#min_count                 = 1
#max_count                 = 3


//The vm type of the node. Check GKE documentation for different choices.
#node_machine_type         = "e2-standard-4"

//The OS of the node. As docker is no longer supported, the choices left will be ubuntu_containerd and cos_containerd.
#node_image                = "COS_CONTAINERD"

//The max pods per node settings should be changed based on the resource request per pod. 110 is the maximum number
#default_max_pods_per_node = 32
#node_max_pods             = 32

//The default node disk configuration
#node_disk_type            = "pd-standard"
#node_disk_gb              = 100

//The name of the nodepool. Do not change
#node_pool_name            = "primary"

//If you choose regional cluster, specify the zones of the region where the cluster compute nodes reside
#node_zones                = "us-east4-a,us-east4-b,us-east4-c"

//If you choose zonal cluster, specify the zone where the cluster computer nodes reside
#gcp_zone                  = "us-east4-a"

//If there is network policy provisioned on the cluster
#network_policy            = false

//These are the service account settings for Terraform to provision the cluster. The use_default_sa and use_custom_sa 
//are NOT mutually exclusive. The default sa has the format "tfxxx@email-uuid", if you turn use_default_sa to false, 
//and use_custom_sa to false, the sa is "tfxxx@email", which removes the uuid. It is useful when creating firewall rules 
//when you need the SA info beforehand. The service_account_run_as must be set when you set use_custom_sa as true. The 
//custom SA is normally assigned for privilege purposes. If you don't have special requirements, leave them as is. 
#use_custom_sa             = false
#use_default_sa            = false
#service_account_run_as    = ""

//Whether to masquerade traffic to the link-local prefix (169.254.0.0/16)
#ip_masq_link_local        = false

//List of strings in CIDR notation that specify the IP address ranges that do not use IP masquerading
#non_masquerade_cidrs      = ["172.16.0.0/12", "192.168.0.0/16"]

//Enables the installation of ip masquerading, which is usually no longer required when using aliasied IP addresses.
#configure_ip_masq         = false

//Allows the traffic from the list to access your cluster.
#authorized_networks = [{
#  cidr_block   = "10.0.0.0/8"
#  display_name = "ot_corp"
#}]

//For secret store CSI driver, list the gcp plugin image location
#gcp_plugin_image   = "artifactory.otxlab.net/ot2-paas/gcp/secretmanager-csi/secrets-store-csi-driver-provider-gcp/plugin:v1.1.0"

/*
The following are variables related to the paas-apps. If you don't need a module, you don't need to fill in the values.
*/

//Mandatory: The reserved ip name can be external or internal depends on the type of project. For ASM, please fill out the 
//ingress_ip_name and comment out ingress_ip. For Nginx ingress, please fill out ingress_ip and comment out ingress_ip_name.
ingress_ip_name           = 
ingress_ip                = 

//The cluster_short_name is the name without the project id part. i.e otl-csd-paas-eng-asmdemo => asmdemo
cluster_short_name = 

//ACM kustomization template. For lab projects without ASM, use kustomization.tftpl. For lab project with ASM, 
//use kustomization-asm.tftpl. For commercial project(with ASM), use kustomization-asm-commercial.tftpl
template_file             = 

//For ACM repository sync(pulling configuration), the git repo information.
#branch                    = "engineering"
//The directory format is "environments/your project id/your cluster name"
directory                 = 
//If you use token to access the repo, the format has to start "https://" and cannot be "git@gitlab".
#repository_url            = "git@gitlab.otxlab.net:platform-team/paas-anthos-configs.git"
//For ACM sync repo, the default connection is using ssh key. If you need to use access token, change it to "token".
#secret_type               = "ssh"      
//If you choose to use token to connect to sync repo, provide the token in string
#token                     = "Lae5m6gUikQB4-By6xU_"
//If you choose to use token to connect to sync repo, specify the username depends on the type of token.
#gituser                   = "randomname"


//For ACM, do you need policy controller?
#enable_policy_controller  = true

//For ASM ingressgateways, whether you want to reserve the ips through API or you prefer to directly assign them
#reserve_ip                = true

//For ASM istiogateways, if you know the IPs of the two gateways. If you selected reserve_ip
//to be true, you MUST provide the values here.
#addresses                 = [""]

//The service account name for Trident. Make sure it's unique per project and shorter than 30 characters.
sa_name                   = 

//For trident export rule, add the cidr of the node subnet, i.e. if node subnet is  "otl-us-east4-data-6",
//enter "10.226.80.0/21" here. You can add more cidrs, using ',' to separate them.

export_rule_cidr         = 

//Https proxy to reach out of the cluster, used by Trident and ACM sync repo
#proxy = "http://gcp-prox01-l001.otxlab.net:3128/"

//This is the project number, not project id
#network_project_number = "165961614692"

//For ASM overlay generation, if the cluster is commercial
#commercial                = false

//For (managed) ASM, which channel it subscribed to. The channel is in small letters.
#channel                   = "regular"

//For ASM, the name of ssl policy. The name has to be unique within the project.
#ssl_policy                = "ssl-policy"

//For ASM, the name of ingress security policy. The name has to be unique within the project.
#security_policy           = "ingress-security-policy"

//For ASM Policy constraints template, the strictness level and enforcement action. 
//strict_level has two options: "Low" or "High". 
#strict_level = "High"
//enforce_action has two options, "deny" or "dryrun".
#enforce_action = "deny"

//The upload repo of the ACM overlays(different set of varialbes from sync repo), the git repo information. 
//The repo project name under the gitlab base
#uploadrepo_project        = "platform-team/paas-anthos-configs"
//The account user email. To retrieve the email for a bot account, run
//curl --header "PRIVATE-TOKEN: replace with uploadtoken" "https://gitlab.otxlab.net/api/v4/user"
#gitlab_user_email         = "project14342_bot2@noreply.gitlab.otxlab.net"
//The base url for gitlab repo
#gitlab_base               = "https://gitlab.otxlab.net/"
//The token that matches the user email
#uploadtoken               = "hqeKxaJ6EUAt63-8Mh3n"
//The Path of pushing, the suggested format is "environments/your project id/your cluster name"
//Make sure there is no "/" at last 
remote_path               = 
//The branch to push the overlay files to
#upload_branch             = "engineering"


//For Prometheus bundle, put the host name based on your dns record. For example, grafana_host = "monitoring.asmdemo.otl-csd-paas-eng.otxlab.net"
grafana_host              = 
alertmanager_host         = 
prometheus_host           = 
//If you use ASM, set it to false. Otherwise leave it.
#use_nginx = true

# For usage in path references to New Relic components. The default value for data_center_name is otl if not explicitly set here. 
# Uncomment and change the value to otc if deployed in commercial 
# data_center_name     = "otc"

# For usage in path references to New Relic components. The default value for data_center_location is us if not explicitly set here. 
# Uncomment and change the value to one of these locations (ca, eu) if deployed in commercial 
# data_center_location = "ca"

# For Node labels, there is some default labels already included in the code. But fileds like environment, auditable_entity, business_service, and owner
# Plus any other labels can be customized here. Also this map values will overwrite the default values so you don't have modify code.
# additional_node_pools_labels = {}

//If you need a gpu node pool, here is an example:
/*gpu_node_pool        = {
  name               = "gpu-node-pool"
  version            = "1.27.3-gke.100"
  autoscaling        = false
  machine_type       = "n1-standard-2"
  node_locations     = "us-east4-a,us-east4-b,us-east4-c"
  accelerator_count  = 1
  accelerator_type   = "nvidia-tesla-t4"
  node_count         = 1
  gpu_driver_version = "DEFAULT"
  auto_repair        = true
  auto_upgrade       = false
}*/
# gpu_node_pool = null












