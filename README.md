# home-idp


<img src="https://img.shields.io/badge/15182-FFFFFF?style=flat&label=lines of code"/>

<img src="https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white"/> <img src="https://img.shields.io/badge/React-61DAFB?style=flat&logo=react&logoColor=black"/> <img src="https://img.shields.io/badge/Kubernetes-326CE5?style=flat&logo=kubernetes&logoColor=white"/> <img src="https://img.shields.io/badge/ArgoCD-EF7B4D?style=flat&logo=argo&logoColor=white"/> <img src="https://img.shields.io/badge/Harbor-60B932?style=flat&logo=harbor&logoColor=white"/> <img src="https://img.shields.io/badge/Kaniko-FFA600?style=flat&logo=kaniko&logoColor=white"/> <img src="https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white"/> <img src="https://img.shields.io/badge/ProtoBuf-4285F4?style=flat&logo=google&logoColor=white"/> <img src="https://img.shields.io/badge/Gitops-181717?style=flat&logo=github&logoColor=white"/>

### **Build your own simple internal developer platform(IDP) with home-idp**

| Title         | Content                                 |
|--------------|--------------------------------------|
| [Features](#Features) | Main features of Home-IDP                    |
| [Architecture](#Architecture) | Flow chart of Home-IDP                      |
| [Installation](#Installation)  | How to install Home-IDP               |
| [Demo](#Demo) | Demo videos of Home-IDP              |
| [License](#License) | License of Home-IDP              |

## Features

* Easy application deployment to kubernetes cluster by typing dockerfile in dashboard
* Easy environment variables and files mount to application by typing in dashboard
* Monitoring application deployment process from building container image to applying manifest
* User access control based on roles and policies
* Visualization deployed kubernetes resources 

## Architecture
![FlowChart](https://home-idp-choigonyok.s3.ap-northeast-2.amazonaws.com/1+(1).png)

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

[![zguUa8yQ48A](https://img.youtube.com/vi/zguUa8yQ48A/0.jpg)](https://www.youtube.com/watch?v=zguUa8yQ48A)

## License

* Apache-2.0 license
