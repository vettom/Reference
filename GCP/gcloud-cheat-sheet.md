# QUick ref Gcloud commands

## Projects and Auth
```bash
	# View current gcloud config
	gcloud config list
	# Set a project
	gcloud config  set project <project name>
	# Set impersonation as default in config
	gcloud config set auth/impersonate_service_account admin-svcaccount@project.gserviceaccount
```

## Impersonaiton
```bash
	# List instances as impersonated user
	gcloud compute  instances list --impersonate-service-account=admin-svcaccount@project.gserviceaccount
	# Print access token
	gcloud auth print-access-token --impersonate-service-account=admin-svcaccount@project.gserviceaccount
```
