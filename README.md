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

## Demo

[![zguUa8yQ48A](https://img.youtube.com/vi/zguUa8yQ48A/0.jpg)](https://www.youtube.com/watch?v=zguUa8yQ48A)

## Getting Start

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

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/875D7928-A6BB-4C13-8952-76B2002B08AE_2/xLM0ptgyagce1rYqXIvbhlKdfyUs1R4DMcLXfuKORKMz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/0FB4A8E0-2A6A-4254-8E5F-4A9926243FE5_2/nk4I0mWqxB8FMbK7FxQq3izrs2N9xZrvMKEuaVWv3OIz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/74817761-28B7-4F48-AD0F-66CAC5179AD8_2/rQaxWPrxrnhRVgLUAxPJpkR67ofK4JIxdLgcKlZ90xsz/Image.png)

GitHub App name

Homepage URL

Callback URL : API serverhost + /github/callback

Create Github App button

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/746D7684-349C-4697-BD3A-7C098E4B957F_2/wV0CuDwG0bN4bMBxjw6C4je6r17PzL0h6EyiHC4yytEz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/87BF64C3-647B-4DC1-8FBD-C022611F9F8E_2/MjnPeDUaQ2IByAYVlSjguOqynxbw7tGLYHM8ZXdQ0lkz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/3AF6D8B3-F3C9-44EC-B5C2-51C7A66D6FA1_2/VrFjjp0r37W8BDS2Dw2z3HQ7bn3u9czPJxIRqAcHV5Qz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/7CD9A57C-1487-45EA-A98D-6CB4E4351465_2/Qdykt4CMhClbyoB8e1frXiS2QmYPemuRvAGvfR0MexEz/Image.png)

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/FCB26870-1127-4397-A07C-B9780DEEB03A_2/MqLdBqbk66z51Mw0QRFHBLNR2HzIFC8MIdhMR6KsMTUz/Image.png)

Note

Expiration

**Select scopes**


- read:repo_hook
- repo

![image](https://res.craft.do/user/full/6deb5b3a-d995-5f97-e85b-e7c3c5f9702a/doc/84D7EE07-A44B-40CC-8665-922232CB5FB4/06237B6E-F384-478E-BA24-B16CD23637C4_2/jPLxB9dsyjzXANdubC4pADLLLOrNSxUet0FrAPswVJsz/Image.png)

kubectl create ns `global.namespace`

helm install idp [oci://registry.choigonyok.com/library/home-idp](oci://registry.choigonyok.com/library/home-idp) --version=1



| field | description | type | required |default |
|---|---|---|---|---|
|`global.prefix`|prefix of kubernetes resource names |string|no|`home-idp`|
|`global.namespace`|namespace where home-idp system deployed|string|no|`idp-system`|
|`global.image.tag`|version of home-idp|string|no|`latest`|
|`auth.adminPassword`|system admin password|string|yes|-|
|`git.username`|github username of administrator|string|yes|-|
|`git.email`|github user email of administrator|string|yes|-|
|`git.repository.name`|github repository name that will be created|string|no|`home-idp-repo`|
|`git.repository.token`|github token|string|yes|-|
|`git.oauth.clientId`|github app client id|string|yes|-|
|`git.oauth.clientSecret`|github app client secret|string|yes|-|
|`ìngress.ui.host`|client host|string|yes|-|
|`ìngress.ui.port`|client port|integer|no|80|
|`ìngress.ui.tls`|using HTTPS|boolean|no|false|
|`ìngress.gateway.host`|api host|string|yes|-|
|`ìngress.gateway.port`|api port|integer|no|80|
|`ìngress.gateway.tls`|using HTTPS|boolean|no|false|
|`ìngress.harbor.host`|registry host|string|yes|-|
|`ìngress.harbor.port`|registry port|integer|no|80|
|`ìngress.harbor.tls`|using HTTPS|boolean|no|false|
|`ìngress.harbor.password`|harbor admin password|string|yes|-|
|`ìngress.argocd.password`|argocd admin password|string|yes|-|
|`storage.database.type`|kind of database|string|no|`postgresql`|
|`storage.databaseport`|port of database|integer|no|5432|
|`storage.database.name`|name of database|string|no|`idp_db`|
|`storage.database.username`|database username|string|no|`postgres`|
|`storage.database.password`|database user password|string|yes|-|
|`storage.persistence.storageClass`|storageClass for database|string|yes|-|
|`storage.persistence.size`|size of persistence volume|string|no|`8Gi`|
|`replicas.ui`|replicas of ui statefulset|integer|no|1|
|`replicas.deployManager`|replicas of deployManager statefulset|integer|no|1|
|`replicas.installManager`|replicas of installManager statefulset|integer|no|1|
|`replicas.traceManager`|replicas of traceManager statefulset|integer|no|1|
|`replicas.rbacManager`|replicas of rbacManager statefulset|integer|no|1|
|`replicas.gateway`|replicas of gateway statefulset|integer|no|1|





```bash
helm uninstall -n $RELEASE_NAMESPACE $RELEASE_NAME
helm uninstall -n $RELEASE_NAMESPACE home-idp-harbor
helm uninstall -n $RELEASE_NAMESPACE home-idp-postgres
helm uninstall -n $RELEASE_NAMESPACE home-idp-cd
```



## License

* Apache-2.0 license