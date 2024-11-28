# Install with Helm chart

## Prerequisites

You need `Github Oauth App`, `Github Token`.

### Github Oauth App

In Github, you can create your oauth app with `Settings` -> `Developer settings` -> `Oauth Apps` -> `New Oauth App`.

Fill `Application name` with anything you want, `Homepage URL` should be client full URL, `Authorization callback URL` should be `$HOMEPAGE_URL/github/callback`.

Click `Update application`, and get inside created Oauth App.

Click `Generate a new client secret`, and copy `Client ID` and `Client secret`.

### Github Token

Move to `Settings` -> `Developer settings` -> `Personal access tokens` -> `Tokens (classic)` -> `Generate new token (classic)`.

Fill the `Note` with anything you want, select `repo`, `admin:repo_hook` scopes. And click `Generate token`

Copy token value.

## Deploy Helm chart

You need to write `override-valeus.yaml` to configure your home-idp. you can find `values.yaml` in [here](https://github.com/choigonyok/home-idp/tree/main/charts).

```bash
kubectl create ns $RELEASE_NAMESPACE
helm install $RELEASE_NAME -n $RELEASE_NAMESPACE oci://registry.choigonyok.com/library/home-idp --version=1.0 -f override-values.yaml
```

## values.yaml

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
|`ìngress.ui.annotations`|annotations of ui ingress|map|no|-|
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

## Remove Helm chart

There are some dependencies, so you have to remove them either.

```bash
helm uninstall -n $RELEASE_NAMESPACE $RELEASE_NAME
helm uninstall -n $RELEASE_NAMESPACE home-idp-harbor
helm uninstall -n $RELEASE_NAMESPACE home-idp-postgres
helm uninstall -n $RELEASE_NAMESPACE home-idp-cd
```