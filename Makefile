# +----------------------------------
# | Build
# +----------------------------------
# linux_amd64/release/idpctl-linux-amd64:
# 	GOOS=linux GOARCH=amd64 common/build.sh $@ ./idpctl/cmd

TAG := latest
IMAGE_REPOSITORY = achoistic98
CONTEXT = .

TRACE_MANAGER_IMAGE_NAME = home-idp-trace-manager
TRACE_MANAGER_DOCKERFILE = docker/Dockerfile.trace-manager

RBAC_MANAGER_IMAGE_NAME = home-idp-rbac-manager
RBAC_MANAGER_DOCKERFILE = docker/Dockerfile.rbac-manager

INSTALL_MANAGER_IMAGE_NAME = home-idp-install-manager
INSTALL_MANAGER_DOCKERFILE = docker/Dockerfile.install-manager

DEPLOY_MANAGER_IMAGE_NAME = home-idp-deploy-manager
DEPLOY_MANAGER_DOCKERFILE = docker/Dockerfile.deploy-manager

GATEWAY_IMAGE_NAME = home-idp-gateway
GATEWAY_DOCKERFILE = docker/Dockerfile.gateway


.PHONY: trace
trace:
	docker build -t $(IMAGE_REPOSITORY)/$(TRACE_MANAGER_IMAGE_NAME):$(TAG) -f $(TRACE_MANAGER_DOCKERFILE) $(CONTEXT)

.PHONY: rbac
rbac:
	docker build -t $(IMAGE_REPOSITORY)/$(RBAC_MANAGER_IMAGE_NAME):$(TAG) -f $(RBAC_MANAGER_DOCKERFILE) $(CONTEXT)

.PHONY: install
install:
	docker build -t $(IMAGE_REPOSITORY)/$(INSTALL_MANAGER_IMAGE_NAME):$(TAG) -f $(INSTALL_MANAGER_DOCKERFILE) $(CONTEXT)

.PHONY: deploy
deploy:
	docker build -t $(IMAGE_REPOSITORY)/$(DEPLOY_MANAGER_IMAGE_NAME):$(TAG) -f $(DEPLOY_MANAGER_DOCKERFILE) $(CONTEXT)

.PHONY: gateway
gateway:
	docker build -t $(IMAGE_REPOSITORY)/$(GATEWAY_IMAGE_NAME):$(TAG) -f $(GATEWAY_DOCKERFILE) $(CONTEXT)

.PHONY: delete
delete:
	helm uninstall -n idp-system home-idp-postgresql &
	helm uninstall -n idp-system home-idp-cd &
	helm uninstall -n idp-system home-idp-harbor &
	kubectl delete secrets -n idp-system --all &
	kubectl delete job -n idp-system --all &
	kubectl delete crd -n idp-system --all &
	kubectl delete pv -n idp-system --all &
	kubectl delete pvc -n idp-system --all &
	kubectl delete pods -n idp-system --all &
	kubectl delete statefulset -n idp-system --all &
	kubectl delete pvc -n idp-system --all &
	kubectl delete pods -n idp-system --all &
	wait
