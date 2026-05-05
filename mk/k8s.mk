# Shared Kubernetes cluster variables.
# Included before tool-specific files (kind.mk, k3d.mk).
# Include-guarded so it can be pulled in automatically by dependents.
ifndef _K8S_MK
_K8S_MK := 1

# Auto-include common.mk (needed for _is_digits)
include $(dir $(lastword $(MAKEFILE_LIST)))common.mk

# Cluster identifier.
# Accepted forms: "kuma" (default), digits (e.g. "2" -> "kuma-2"), "kuma-<N>"
# `KIND_CLUSTER_NAME` is kept as a compatibility alias for older workflows.
CLUSTER ?= kuma
ifdef KIND_CLUSTER_NAME
ifeq ($(origin CLUSTER), default)
CLUSTER := $(KIND_CLUSTER_NAME)
endif
endif

# Where kubeconfig files are stored
KUBECONFIG_DIR ?= $(HOME)/.kube

# Prevent inherited KUBECONFIG from leaking into recipes;
# tool-specific files export it per-target instead.
unexport KUBECONFIG

# --- E2E pre-warm experiment ---

KUMACTL_PREWARM_NAMESPACE ?= kube-system
KUMACTL_PREWARM_TIMEOUT ?= 120s
KUMA_INIT_IMAGE ?= $(DOCKER_REGISTRY)/kuma-init:$(BUILD_INFO_VERSION)

define KUMACTL_PREWARM_RECIPE
	$(Q)for node in $$($(KUBECTL) get nodes -o jsonpath='{.items[*].metadata.name}'); do \
	  pod_name=$$(printf 'kumactl-prewarm-%s' "$$node" | tr '[:upper:]' '[:lower:]' | tr -c 'a-z0-9-' '-'); \
	  echo "Pre-warming $$node with $(KUMA_INIT_IMAGE)"; \
	  printf '%s\n' \
	    'apiVersion: v1' \
	    'kind: Pod' \
	    'metadata:' \
	    "  name: $$pod_name" \
	    '  namespace: $(KUMACTL_PREWARM_NAMESPACE)' \
	    '  labels:' \
	    '    app.kubernetes.io/name: kumactl-prewarm' \
	    'spec:' \
	    "  nodeName: $$node" \
	    '  restartPolicy: Never' \
	    '  terminationGracePeriodSeconds: 0' \
	    '  tolerations:' \
	    '  - operator: Exists' \
	    '  containers:' \
	    '  - name: kumactl-prewarm' \
	    '    image: $(KUMA_INIT_IMAGE)' \
	    '    imagePullPolicy: IfNotPresent' \
	    '    command:' \
	    '    - /usr/bin/kumactl' \
	    '    - install' \
	    '    - transparent-proxy' \
	    '    - --dry-run' \
	    '    - --redirect-all-dns-traffic' \
	    '    - --verbose' | $(KUBECTL) apply -f -; \
	  if ! $(KUBECTL) wait --namespace $(KUMACTL_PREWARM_NAMESPACE) --for=jsonpath='{.status.phase}'=Succeeded --timeout $(KUMACTL_PREWARM_TIMEOUT) pod/$$pod_name; then \
	    $(KUBECTL) describe --namespace $(KUMACTL_PREWARM_NAMESPACE) pod/$$pod_name || true; \
	    $(KUBECTL) logs --namespace $(KUMACTL_PREWARM_NAMESPACE) pod/$$pod_name || true; \
	    $(KUBECTL) delete --namespace $(KUMACTL_PREWARM_NAMESPACE) pod/$$pod_name --ignore-not-found=true --wait=false; \
	    exit 1; \
	  fi; \
	  $(KUBECTL) delete --namespace $(KUMACTL_PREWARM_NAMESPACE) pod/$$pod_name --ignore-not-found=true --wait=false; \
	done
endef

# --- Validation ---

_k8s_cluster_valid := $(or \
  $(filter kuma,$(CLUSTER)),\
  $(and $(filter kuma-%,$(CLUSTER)),$(call _is_digits,$(patsubst kuma-%,%,$(CLUSTER)))),\
  $(call _is_digits,$(CLUSTER)))

$(if $(_k8s_cluster_valid),,$(error Invalid CLUSTER "$(CLUSTER)". Expected "kuma", digits, or "kuma-<digits>"))

# --- Derived variables ---

# Numeric suffix (used for MetalLB subnet offset, port allocation)
ifeq ($(CLUSTER),kuma)
  CLUSTER_NUMBER := 0
else ifneq ($(filter kuma-%,$(CLUSTER)),)
  CLUSTER_NUMBER := $(patsubst kuma-%,%,$(CLUSTER))
else
  CLUSTER_NUMBER := $(CLUSTER)
endif

# Canonical cluster name: "kuma" stays "kuma"; numbers become "kuma-<N>"
ifeq ($(CLUSTER_NUMBER),0)
  CLUSTER_NAME := kuma
else
  CLUSTER_NAME := kuma-$(CLUSTER_NUMBER)
endif

# Tool-specific kubeconfig paths keep kind and k3d clusters from clobbering
# each other's files when they share the same cluster name.
k8s_cluster_kubeconfig = $(KUBECONFIG_DIR)/$(1)-$(2).yaml

KIND_CLUSTER_KUBECONFIG := $(call k8s_cluster_kubeconfig,kind,$(CLUSTER_NAME))
K3D_CLUSTER_KUBECONFIG := $(call k8s_cluster_kubeconfig,k3d,$(CLUSTER_NAME))

# Compatibility alias: kong-mesh and older workflows reference KIND_KUBECONFIG
KIND_KUBECONFIG = $(KIND_CLUSTER_KUBECONFIG)

# Temp workspace for generated k8s manifests and caches
TMP_DIR_K8S ?= /tmp/.kuma-dev

# --- Docker network ---
# Shared by all cluster tools and universal-mode test containers.

DOCKER_NETWORK ?= kuma
DOCKER_NETWORK_OPTS = --opt com.docker.network.bridge.enable_ip_masquerade=true
ifdef IPV6
    DOCKER_NETWORK_OPTS += --ipv6 --subnet "fd00:fd12:3456::0/64"
endif

# --- Docker network: shared create target ---
# Tool-agnostic; usable from e2e targets that just need the network present.

.PHONY: k8s/docker/network/create
k8s/docker/network/create:
	$(Q)docker network inspect $(DOCKER_NETWORK) >/dev/null 2>&1 \
	  || docker network create --driver bridge $(DOCKER_NETWORK_OPTS) $(DOCKER_NETWORK) >/dev/null 2>&1 \
	  || docker network inspect $(DOCKER_NETWORK) >/dev/null 2>&1

endif # _K8S_MK
