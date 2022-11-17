# ====================================================================================
# Setup Project

PLATFORMS := linux_amd64 linux_arm64
include ../../../build/makelib/common.mk

# ====================================================================================
#  Options
#IMAGE = $(BUILD_REGISTRY)/$(DOCKER_REPOSITORY)/provider-github-controller-$(ARCH)
IMAGE = $(BUILD_REGISTRY)/provider-github-controller-$(ARCH)
include ../../../build/makelib/image.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) docker build $(IMAGE)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@$(INFO) docker build $(IMAGE)
	@$(INFO) Test ROOT_DIR: $(shell echo $(ROOT_DIR))
	@$(INFO) Test OUTPUT_DIR: $(shell echo $(OUTPUT_DIR))
	@cp $(OUTPUT_DIR)/bin/$(HOSTOS)_$(ARCH)/provider $(IMAGE_TEMP_DIR) || $(FAIL)
	@cd $(IMAGE_TEMP_DIR) && $(SED_CMD) 's|BASEIMAGE|$(OSBASEIMAGE)|g' Dockerfile || $(FAIL)
	@docker build $(BUILD_ARGS) \
		--build-arg ARCH=$(ARCH) \
		--build-arg TINI_VERSION=$(TINI_VERSION) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)
	@$(OK) docker build $(IMAGE)
