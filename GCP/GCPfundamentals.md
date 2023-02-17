# GCP fundamentals

## Structure
Organisation Node
	|
	Folder
		|
		projects
			|
			Resources

Projects have 3 identifiable attributes
- Project ID
- Project name. (Updatable)
- Project number
Policys can be applied to any level as well as to some resources. Policies are applied top down 
Folders allows to group teams or departments and have their own access policies. Folders can contain multiple projects.

# IAM
Types of accounts
- Google account
- Google groups
- Service account
- Cloud Identity Domain

### Roles
- Basic. Wide, and simple, applies to all resources.
	- Owner. Everything
	- Editor
	- Viewer
	- Billing Admin. View and billing. Not modify.
- predefined
- Custom Roles. Can be assigned at Project or Org level only

### Service account
Assigned to resources. It has email as username and Cryptographic key for auth. Permissions are granted same as any other resource. 


## VPC
VPC is global and subnet can span across zones.

## Compute Engine
- Billed by second with minimum of minute
- Long running compute automatically gets sustained-use discount
- Committed use discount when commit to 1 year or 3 year.
- Spot VM / preemptible VM
	- Preemptible VM can run for max 24 hours

## Cloud DNS
Hosted DNS solution by Google

## Connecting to Google
- IP Sec tunnel via Cloud Router
- Direct Peering. (Direct connect)
- Carrier peering (partner service provider)
- Dedicated interconnect 
- Partner interconnect. 

# Storage
- Cloud Storage (Bucket)
	- standard
	- Nearline like once a month access
	- Coldline (access once in 90 days)
	- Archive storage (1s year)
- Cloud SQL
- Cloud Spanner (SQL)
- Firestore, No-SQL
	- Data stored as documents and collections.
- Cloud Bigtable (Big Data)
	- No SQL
	- Semi-structured or structured data
	- Machine learning data
	- Analytical structured data


# App Engine
### Standard
- no ssh access
- no write to local disk
- Certain apps and versions only
### Flexible env
- Is hosted Docker container
- ssh allowed
- Any version of software can be installed
- Calls to network without app engine.

# API MANAGEMENT TOOL
- Cloud end points
	- Uses distributed service proxy
	- OpenAPI specification
	- Supports App Engine, k8s engine, and Compute engine
- API gateway
	- Secure access to backend services through REST Api
- Apigee Edge
	- Focuss on solving BU problems like rate limit, quotas, and Analytics

# Cloud Run
- Run stateless containers
- serverless

# Cloud Source Repo
- Google managed Git
# Cloud Functions
- Small event based compute aka lambda



