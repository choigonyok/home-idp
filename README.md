# home-idp


<img src="https://img.shields.io/badge/15091-FFFFFF?style=flat&label=lines of code"/>

<img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white"/> <img src="https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black"/> <img src="https://img.shields.io/badge/Kubernetes-326CE5?style=flat&logo=kubernetes&logoColor=white"/> <img src="https://img.shields.io/badge/ArgoCD-EF7B4D?style=flat&logo=argo&logoColor=white"/> <img src="https://img.shields.io/badge/Harbor-60B932?style=flat&logo=harbor&logoColor=white"/> <img src="https://img.shields.io/badge/Kaniko-FFA600?style=flat&logo=kaniko&logoColor=white"/> <img src="https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white"/> <img src="https://img.shields.io/badge/ProtoBuf-4285F4?style=flat&logo=google&logoColor=white"/> <img src="https://img.shields.io/badge/Gitops-181717?style=flat&logo=github&logoColor=white"/>

### **Build your own simple internal developer platform(IDP) with home-idp**


## Features

* Easy application deployment to kubernetes cluster by typing dockerfile in dashboard
* Easy environment variables and files mount to application by typing in dashboard
* Monitoring application deployment process from building container image to applying manifest
* User access control based on roles and policies
* Visualization deployed kubernetes resources 

## Installation

You can install home-idp with [idpctl](https://github.com/choigonyok/home-idp?tab=readme-ov-file#install-with-idpctl) or [helm chart](https://github.com/choigonyok/home-idp?tab=readme-ov-file#install-with-idpctl).

### Install with idpctl

idpctl is CLI to install hoem-idp.

1. make kubernetes client to access API server
2. set CLI root/sub commanstall/uninstall/monitor/update idp in cd to inluster
3. build makefile to make home-idp container image and helm chart

```bash
wget https://github.com/choigonyok/home-idp/releases/latest/download/idpctl-1.0-linux-amd64.tar.gz
tar xzvf idpctl-1.0-linux-amd64.tar.gz
./idpctl install -f config.yaml
```

### Install with Helm chart

```
helm repo add choigonyok https://registry.choigonyok.com
helm repo update
helm install home-idp -n idp-system choigonyok/home-idp -f values.yaml
```

## Demo

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/wHIg_MWo9h0/0.jpg)](https://www.youtube.com/watch?v=wHIg_MWo9h0)