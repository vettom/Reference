# GCP, Kubernetes with Terraform

# Terraform
- [ ] Local TF install
    - Formt tf code with terraform fmt
    - Validate code with Terraform validate
    - Install [tfswitch](https://github.com/warrensbox/terraform-switcher)
    - Set terraform plugins directory outside of repo
    - Lint tf code with [tflint](https://github.com/terraform-linters/tflint/releases)
    - run tfsec on tf code

- [ ] State file
    - Store state file in S3 with Dynamodb lock
    - Import a resource to tfstate (Eg: instance)

## VPC tasks
- [ ] Using Tf create a vpc and subnet 
    - single subnet
    - secondary ip for K8s
    - Firewall rule allowing incoming conneciton on port 9000

- [ ] Create VPC module 
    - VPC name and subnets are variables
    - Create Cloud NAT and attach
    - Route outbound traffic via CloudNAT

## GKE
- [ ] Build standard GKE cluster
    - Use secondary ip for pod and services
    - Single node, pre-emptible node
    - Enable private node
    - Enable public endpoint

- [ ] Create GKE Cluster module 
    - All inputs should be variables
    - Delete defult nodepool and create new
    - Output: Clustername, endpoint

- [ ] Upgrade GKE cluster to latest version

# Quick cheap GKE cluster with Gcloud (less than 15p an hour)
```bash

# Open Cloud shell in browser and execute.
CLUSTER_NAME="playground-cluster1"
PROJECTNAME= ""
MACHINE_TYPE="e2-micro"
DISK_SIZE="30"
CLUSTER_VERSION="latest"
VPC=""
SUBNET_NAME=""
SUBNET="projects/$PROJECTNAME/regions/us-west1/subnetworks/$SUBNET_NAME"
MASTER_NETWORK="projects/$PROJECTNAME/global/networks/$VPC"
 
gcloud beta container --project $PROJECTNAME clusters create $CLUSTER_NAME --zone "us-west1-b" --no-enable-basic-auth --cluster-version $CLUSTER_VERSION --release-channel "regular" --machine-type $MACHINE_TYPE --image-type "COS_CONTAINERD" --disk-type "pd-balanced" --disk-size $DISK_SIZE --metadata disable-legacy-endpoints=true --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" --spot --num-nodes "1" --enable-ip-alias --network $MASTER_NETWORK --subnetwork $SUBNET --no-enable-intra-node-visibility --default-max-pods-per-node "110" --security-posture=standard --workload-vulnerability-scanning=disabled --no-enable-master-authorized-networks --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver --enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0 --binauthz-evaluation-mode​=DISABLED --no-enable-managed-prometheus --enable-shielded-nodes --node-locations "us-west1-b"
 ```