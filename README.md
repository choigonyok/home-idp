# home-idp

### **Build your own simple internal developer platform(IDP) with home-idp framework.**

* Build idpctl CLI

1. make kubernetes client to access API server
2. set CLI root/sub command to install/uninstall/monitor/update idp in cluster
3. build makefile to make home-idp container image and helm chart

## Requirements

1. home-idp uses `idpctl` CLI to install hoem-idp in your kubernetes cluster. To make idpctl access k8s API server, you should set appropriate kubeconfig file in `$HOME/.kube/config`

2. You should have kubernetes cluster home-idp application can be deployed in.

3. If you want to deploy with idpctl CLI, there are some dependencies.

* Keycloak
* Oauth2-proxy
* Database (PostgreSQL or MySQL)
* Redis
* S3 object storage (including MinIO)

```
if you use Helm chart, there are subcharts for every dependency.
```

## Installation


### **With idpctl CLI**
install idpctl CLI in you linux server with below link.

https://github.com/choigonyok/home-idp/releases/latest/download/idpctl-1.0-linux-amd64.tar.gz

```
tar xzvf idpctl-1.0-linux-amd64.tar.gz
./idpctl install -f config.yaml
```

### **With helm chart**

```
helm repo add choigonyok https://registry.choigonyok.com
helm repo update
helm install home-idp -n idp-system choigonyok/home-idp -f values.yaml
```


## Test Environment

### Kind

You can simply run home-idp with KinD kubernetes cluster.
KinD is kubernetes provisioning opensource software which install kubernetes cluster in docker container.
KinD should be installed first in your local environment.
And `idp-system` namespace should be created for home-idp.

```
# mac
brew install kind
kind create cluster
kubectl create ns idp-system
```

```
# window
choco install kind
kind create cluster
kubectl create ns idp-system
```

### Manifest

```
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: install-manager
  namespace: test-ns
spec:
  containers:
  - name: install-manager
    image: achoistic98/install-manager:latest
    ports:
    - containerPort: 5051
EOF
```