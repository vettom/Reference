# Get started with Kubernetes pay loads

These tasks can be completed on local installation of K8s
> 1 e2 micro instance is free per month in GCP. Can be used to host Vault, Atlantis etc at no cost

- [ ] Local K8S
    - Run local minikube/kind k8s
    - Enable ingress
    - Deploy simple web app
    - Enable ingress and validate.

- [ ] General K8S
    - Create configMap
    - Mount ConfigMap in pod
    - Create a secret
    - Mount secret as file
    - Create secret with multiple entires
    - mount all secrets at mount point in pod
    - Mount secret to Env variable in Pod

- [ ] Web app and Helm
    - Create simple Web app [sample flask app](https://github.com/vettom/Flask)
    - Push image to Docker Hub or any other registry
    - Create Helm chart to deploy app
        - Specific namespace
        - Add service
        - Add ingress
        - App version as value
        - Option to enable/disable ingress based on value input

- [ ] [Certificate manager](https://cert-manager.io/docs/   )
    - Install Certificate manager
    - Create Cert issuer (LetsEncrypt staging) with Http verification
    - Configure certmanager to issue certificate to ingresses created (via Annotation)
    - Some examples [devops-k8s-setup](https://gitlab.com/dennyvettom/devops-k8s-setup)

- [ ] [External Secrets](https://external-secrets.io/latest/)
    - Create some secret in Vault or any other secret manager.
    - Create secretStore in cluster
    - Create externalSeret that retrive secret and creates secret object

- [ ] [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets)
    - Install Sealed secret
    - Create a sealed secret with cluster CA Cert
    - Make secret available in pod

- [ ] [External DNS](https://github.com/kubernetes-sigs/external-dns)
    - Install external DNS
    - Configure DNS entry to be created for every ingress

# Create a repo in Github/Gitlab to host your applicaiton helm chart

- [ ] [Argo CD](https://argo-cd.readthedocs.io/en/stable/)
    - Install Argocd in to argo namespace. 
```bash
    helm repo add argo https://argoproj.github.io/argo-helm
    helm repo update
    helm install argo-cd -n argocd --create-namespace argo/argo-cd
```
    - Create secret to configure [Argo repository](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#repositories) access
    - Create a project that allow access  to your repo
    - Deploy your app using Argo

- [ ] Argo as app
    # Once installed, Argocd itlesf can be app in Argo. 
    - Create a project for Argo
    - Create app for Argo with ingress
    - Deploy/refresh Argo app from Argo console
    - Upgrade/downgrade Argocd version by passing value
    - Deploy new version of Flask app by updating version tag
    - Experiment with auto/manual deployment 

# Combine all together
- [ ] App Deploy
    - Prepare cluster with core apps
    - Install your web app
        - Ingress configured
        - DNS configured auto or manual
        - Has certificate
        - Gets Secret from external secret source
