# home-idp

<img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white"/> 
<img src="https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black"/> 

<img src="https://img.shields.io/badge/Envoy-AC6199?style=flat&logo=envoyproxy&logoColor=white"/>
<img src="https://img.shields.io/badge/Kubernetes-326CE5?style=flat&logo=kubernetes&logoColor=white"/> 
 <img src="https://img.shields.io/badge/ArgoCD-EF7B4D?style=flat&logo=argo&logoColor=white"/>
 <img src="https://img.shields.io/badge/Harbor-60B932?style=flat&logo=harbor&logoColor=white"/> 
<img src="https://img.shields.io/badge/Kaniko-FFA600?style=flat&logo=kaniko&logoColor=white"/> 
<img src="https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white"/>  

<img src="https://img.shields.io/badge/ProtoBuf-4285F4?style=flat&logo=google&logoColor=white"/> 
<img src="https://img.shields.io/badge/Gitops-181717?style=flat&logo=github&logoColor=white"/>

### **Build your own simple internal developer platform(IDP) with home-idp framework.**



* Build idpctl CLI

1. make kubernetes client to access API server
2. set CLI root/sub command to install/uninstall/monitor/update idp in cluster
3. build makefile to make home-idp container image and helm chart

## Requirements

1. home-idp uses `idpctl` CLI to install hoem-idp in your kubernetes cluster. To make idpctl access k8s API server, you should set appropriate kubeconfig file in `$HOME/.kube/config`

2. You should have kubernetes cluster home-idp application can be deployed in.

3. If you want to deploy with idpctl CLI, there are some dependencies.

* Database (PostgreSQL)
* Harbor (Registry)
* ArgoCD (CD)

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

### KinD Cluster

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
  name: home-idp-install-manager
  namespace: idp-system
spec:
  containers:
  - name: home-idp-install-manager
    image: achoistic98/install-manager:latest
    ports:
    - containerPort: 5051
  serviceAccountName: home-idp-install-manager
---
apiVersion: v1
kind: Pod
metadata:
  name: home-idp-deploy-manager
  namespace: idp-system
spec:
  hostNetwork: true
  containers:
  - name: deploy-manager
    securityContext:
      privileged: true
    image: achoistic98/deploy-manager:latest
    ports:
    - containerPort: 5104
  serviceAccountName: home-idp-deploy-manager
---
apiVersion: v1
kind: Pod
metadata:
  name: home-idp-gateway
  namespace: idp-system
  labels:
    app.kubernetes.io/name: home-idp-gateway
spec:
  containers:
  - name: home-idp-gateway
    image: achoistic98/gateway:latest
    ports:
    - containerPort: 5106
  serviceAccountName: home-idp-install-manager
---
apiVersion: v1
kind: Service
metadata:
  name: home-idp-gateway
  namespace: idp-system
spec:
  selector:
    app.kubernetes.io/name: home-idp-gateway
  ports:
    - protocol: TCP
      port: 80
      targetPort: 5106
EOF
```