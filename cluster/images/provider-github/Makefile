# ====================================================================================
# Setup Project

PLATFORMS := linux_amd64 linux_arm64 darwin_arm64
include ../../../build/makelib/common.mk

# ====================================================================================
#  Options
#IMAGE = $(BUILD_REGISTRY)/$(DOCKER_REPOSITORY)/provider-github-$(ARCH)
IMAGE = $(BUILD_REGISTRY)/provider-github-$(ARCH)
OSBASEIMAGE = scratch
include ../../../build/makelib/image.mk

# ====================================================================================
# Targets


img.build:
	@$(INFO) echo $(DOCKER_REPOSITORY)
	@$(INFO) echo $(BUILD_REGISTRY)
	@$(INFO) echo $(COMMON_SELF_DIR)
	@$(INFO) docker build $(IMAGE)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -R ../../../package $(IMAGE_TEMP_DIR) || $(FAIL)
	@echo $(IMAGE_TEMP_DIR)  $(OSBASEIMAGE)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|BASEIMAGE|$(OSBASEIMAGE)|g' Dockerfile || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|VERSION|$(VERSION)|g' package/crossplane.yaml || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|DOCKER_REGISTRY|$(DOCKER_REGISTRY)|g' package/crossplane.yaml || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|DOCKER_REPOSITORY|$(DOCKER_REPOSITORY)|g' package/crossplane.yaml || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && find package -type f -name '*.yaml' -exec cat {} >> 'package.yaml' \; -exec printf '\n---\n' \; || $(FAIL)
	@docker build $(BUILD_ARGS) \
		--build-arg ARCH=$(ARCH) \
		--build-arg TINI_VERSION=$(TINI_VERSION) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)
	@$(OK) docker build $(IMAGE)
