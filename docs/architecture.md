# Architecture

## Flow Chart

![FlowChart](https://home-idp-choigonyok.s3.ap-northeast-2.amazonaws.com/1+(1).png)

This image is about how user can deploy their applicaiton on kubernetes cluster with home-idp.

* User writes `dockerfile` in home-idp dashboard.
* gateway calls `rbac-manager` procedure to store dockerfile data in database.
* gateway pushes dockerfile on `home-idp github repo`.
* github sends `push webhook` to gateway
* gateway receives dockerfile push webhook, parses pushed dockerfile data, and calls `deploy-manager` procedure to deploy `kaniko` and `git-sync` container to build container image.
* git-sync containers clone `source codes` and `dockerfile` from different repo.
* kaniko conatiner builds `container image`, pushes image to `harbor registry`.
* harbor registry send `push webhook` to gateway.
* gateway parses pushed image name and tag, updated manifest, and pushes manifest to home-idp github repo.
* github sends `push webhook` to gateway.
* gateway receives `push webhook`, and send `webhook` to argocd for `synchronization`.
* argocd receives `webhook` and deploy manifest pushed in `home-idp gihtub repo`.

## Services

home-idp has 6 services, `ui`, `gateway`, `rbac-manager`, `install-manager`, `deploy-manager`, `trace-manager`.

* `ui`: dashboard service
* `gateway`: manages webhooks, forwards request to other services, manages github repo, checks JWT token validation
* `rbac-manager`: stores and responses home-idp datas.
* `install-manager`: install dependencies like argocd, harbor, postgres helm chart at initial runtime.
* `deploy-manager`: deploys pods like user application, kaniko, configmap, secrets.
* `trace-manager`: stores and responses traces and spans.

Each services have plugins to implement their funciton. And plugins are reusable.

![image](https://github.com/user-attachments/assets/fae64e40-054d-4292-b2ce-d9aa0d1f89a0)

![image](https://github.com/user-attachments/assets/20129215-0127-41e0-a3d8-a3a4c5214581)