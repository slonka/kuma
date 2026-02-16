# Zone Proxy Deployment Model

* Status: accepted

Technical Story: https://github.com/kumahq/kuma/issues/9030

## Context and Problem Statement

Currently, zone proxies are **global-scoped** resources, meaning a single ZoneIngress or ZoneEgress instance handles traffic for all meshes in a zone.
This global nature creates fundamental limitations:

1. **Cannot issue MeshIdentity for zone egress**: MeshIdentity is mesh-scoped, but a global zone egress serves multiple meshes.
   This creates identity conflicts and prevents proper mTLS certificate issuance.
   See [MADR 090](090-zone-egress-identity.md) for detailed analysis.

2. **Cannot apply policies on zone proxies**: Kuma policies (MeshTrafficPermission, MeshTimeout, etc.) are mesh-scoped.
   A global zone proxy cannot be targeted by mesh-specific policies, limiting observability and traffic control for cross-zone communication.

3. **Limited observability scoping**: With a global zone proxy, metrics, access logs, and traces cannot be scoped to a specific mesh — all mesh traffic is mixed. Mesh-scoped zone proxies enable per-mesh observability via policies like MeshAccessLog and MeshMetric.

To resolve these limitations, zone proxies are being changed to **mesh-scoped** resources represented as Dataplane resources with specific tags.
This architectural change requires revisiting the deployment model for zone proxies.

**Scope of this document**: This MADR focuses on **deployment tooling** — how users deploy zone proxies via Helm, Konnect UI, and Terraform.

**Multi-mesh support**: The `meshes` list supports deploying zone proxies for multiple meshes in a single Helm release. Single-mesh is the simplest case — one entry in the list.

This document addresses the following questions:

1. Should we continue supporting `kuma.io/ingress-public-address` annotation?
2. What should be the default Helm installation behavior for zone proxies?

Note: Whether zone ingress and egress share a single deployment is addressed in a separate MADR. [^2]

### Decision Summary

| Tooling Decision | Choice |
|------------------|--------|
| Per-mesh Services | **Yes** - each mesh gets its own Service/LoadBalancer for mTLS isolation |
| Namespace placement | **kuma-system** |
| Deployment mechanism | **Helm-managed** (current pattern extended for mesh-scoped zone proxies) |
| Helm release structure | **Mesh subchart** (each mesh entry rendered by a subchart) |
| MADR 093 relaxation | **Yes** - allow multiple meshes per namespace, handle Workload collisions in controller |
| Additive migration | **Yes** - `meshes` config alongside existing `ingress`/`egress` keys |

| Question | Decision |
|----------|----------|
| 1. Support ingress-public-address? | **Yes** - keep as escape hatch |
| 2. Default Helm behavior? | `meshes: []` — explicit opt-in, no zone proxies deployed by default |

Note: Resource model (Dataplane representation, labels, tokens, workload identity) is in a separate MADR. [^1]
Note: Zone proxy deployment topology (shared vs separate ingress/egress) is addressed in a separate MADR. [^2]

### Document Structure

This document is organized in two parts:

1. **Tooling and User Flows** - Describes how users will deploy zone proxies using different tools (Konnect UI, Helm, Terraform). This covers the UX and configuration experience.

2. **Questions 1-2** - Answers deployment-related design questions. Each question analyzes options and recommends a decision.

## Design

### Tooling and User Flows

With zone proxies becoming mesh-scoped, users configure zone proxies per mesh via the `meshes` list in `values.yaml`. Each entry in `meshes` is rendered by a mesh subchart that manages the zone proxy Deployment(s), Service(s), HPA, PDB, ServiceAccount, and optionally the Mesh resource and default policies.

For single-mesh deployments, the `meshes` list has one entry. Multi-mesh deployments add additional entries.

**Full mesh entry schema**:

```yaml
meshes:
  - name: <mesh-name>            # Required. Name of the mesh this entry targets.
    createMesh: false             # Optional. Render a Mesh resource (unfederated zones only).
    createPolicies:               # Optional. Map of default policies to create with the mesh.
      MeshCircuitBreaker: true    # Default circuit breaker for all traffic.
      MeshRetry: true             # Default retry config (TCP + HTTP/GRPC).
      MeshTimeout: true           # Default timeouts for sidecars + gateways.
      TrafficPermission: true     # Allow-all traffic permission (requires createMeshDefaultRoutingResources=true).
      TrafficRoute: true          # Default load-balancing route (requires createMeshDefaultRoutingResources=true).

    # Zone proxy deployment — choose EITHER ingress/egress OR combinedProxies (not both).

    ingress:                      # Separate ingress deployment.
      enabled: true
      podSpec: {}
      hpa: {}
      pdb: {}
      service:
        name: ""                  # Optional override for auto-generated Service name.
      resources: {}

    egress:                       # Separate egress deployment.
      enabled: true
      podSpec: {}
      hpa: {}
      pdb: {}
      service:
        name: ""                  # Optional override for auto-generated Service name.
      resources: {}

    combinedProxies:              # Single deployment running both ingress + egress roles.
      enabled: true               # Mutually exclusive with ingress/egress above.
      podSpec: {}
      hpa: {}
      pdb: {}
      service:
        name: ""                  # Optional override for auto-generated Service name.
      resources: {}
```

Deployment contexts:

1. **Konnect (MinK)** - Global CP is managed, UI has full mesh visibility
2. **Self-hosted Global CP** - Zone CP deployed via Helm, limited mesh visibility at install time
3. **Unfederated Zone** - Standalone zone, no global CP
4. **Terraform** - Infrastructure-as-code with dependency management

#### Flow 1: Konnect UI (Mesh in Konnect)

##### Current State

```
┌─────────────────────────────────────────────────────────┐
│ Connect zone                                            │
├─────────────────────────────────────────────────────────┤
│                                                         │
│ ☑ Deploy Ingress                                        │
│ ☑ Deploy Egress                                         │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

Generates `values.yaml` with `ingress.enabled: true` / `egress.enabled: true` (explicit opt-in)

##### Proposed Flow

**Step 1: UI Enhancement**

```
┌─────────────────────────────────────────────────────────┐
│ Connect zone                                            │
├─────────────────────────────────────────────────────────┤
│                                                         │
│ Detected meshes:                                        │
│                                                         │
│ ☑ default     Ingress: [Yes ▼]  Egress: [Yes ▼]        │
│ ☑ payments    Ingress: [Yes ▼]  Egress: [Yes ▼]        │
│ ☐ backend     Ingress: [--- ▼]  Egress: [--- ▼]        │
│                                                         │
│ If no meshes exist:                                     │
│ "No meshes found — generate values to create the        │
│  default mesh?"                                         │
└─────────────────────────────────────────────────────────┘
```

All detected meshes are auto-filled as entries. Each mesh has its own ingress/egress toggles, mirroring the per-mesh structure in the `meshes` schema. Users can deselect meshes they don't need zone proxies for — deselected meshes are excluded from the generated `values.yaml`.

**Empty state**: If no mesh exists yet, the UI shows an empty state prompting the user to generate a `values.yaml` that creates a `default` mesh. The zone proxy can also be deployed targeting a mesh name that doesn't exist yet — it will wait and retry until the mesh is created.

**Step 2: Generated values.yaml**

The generated values.yaml uses the `meshes` list:

```yaml
kuma:
  controlPlane:
    mode: zone
    zone: zone-1
    kdsGlobalAddress: grpcs://us.mesh.sync.konghq.tech:443

  meshes:
    - name: default
      ingress:
        enabled: true
      egress:
        enabled: true
```

Note: Zone proxy can be deployed before the mesh exists. It will wait and retry until the mesh is created on the Global CP.

**Why this works for Konnect:**

- Konnect UI has API access to global CP
- Can fetch mesh list via `GET /meshes`
- Validation happens UI-side before generating values.yaml
- If multiple meshes exist in the zone, the UI can generate multiple entries in the `meshes` list

#### Flow 2: Self-Hosted Global CP (Helm)

##### Single-Mesh (Default)

```yaml
kuma:
  controlPlane:
    mode: zone
    zone: zone-1
    kdsGlobalAddress: grpcs://global-cp:5685

  meshes:
    - name: default
      ingress:
        enabled: true
      egress:
        enabled: true
```

- Helm renders the mesh subchart for each entry in `meshes`
- Zone proxy connects to CP, requests config for the configured mesh
- **If mesh doesn't exist**: CP returns error, zone proxy logs warning, retries
- **User experience**: Check zone proxy logs, see "mesh 'default' not found"

**Bootstrap validation**: Bootstrap already validates mesh existence for Dataplanes (returns HTTP 422 with `mesh: mesh "<name>" does not exist`).
With mesh-scoped zone proxy as Dataplane, this validation applies automatically.

Code references:
- Server validates mesh: [`pkg/xds/bootstrap/generator.go#L356-L366`](https://github.com/kumahq/kuma/blob/master/pkg/xds/bootstrap/generator.go#L356-L366)
- Server returns HTTP 422: [`pkg/xds/bootstrap/handler.go#L89-L96`](https://github.com/kumahq/kuma/blob/master/pkg/xds/bootstrap/handler.go#L89-L96)

This matches the "eventual consistency" model Kuma already uses.

##### Multi-Mesh Variant

For multi-mesh, add additional entries to the `meshes` list — each renders independently.

#### Flow 3: Unfederated Zone (Standalone)

##### Single-Mesh (Default)

```yaml
kuma:
  controlPlane:
    mode: zone
    zone: zone-1
    # No kdsGlobalAddress = unfederated

  meshes:
    - name: default
      createMesh: true
      ingress:
        enabled: true
      egress:
        enabled: true
```

**Mesh creation in unfederated zones**: Unlike federated zones (where the mesh is synced from the Global CP), unfederated zones create the mesh locally. The `createMesh` field controls this:

1. **`createMesh: true`**: The mesh subchart renders a Mesh resource. This is needed for unfederated zones where no Global CP syncs meshes.

2. **CP auto-creation**: When `skipMeshCreation: false` (the default), the CP creates the `default` mesh at startup (`EnsureDefaultMeshExists` in `pkg/defaults/mesh.go`). This only creates a mesh named `default` — non-default mesh names require `createMesh: true`.

3. **Conditional rendering**: The subchart only renders the Mesh resource when `createMesh: true` AND no `kdsGlobalAddress` is configured (unfederated). For federated zones, `createMesh` is ignored since meshes are synced from the Global CP.

For the common case (`name: default` + `skipMeshCreation: false`), the user can omit `createMesh` — the CP handles default mesh creation automatically.

#### Flow 4: Terraform

##### Key Advantage

Terraform has `depends_on` for ordering (mesh before zone proxy) and `templatefile()` for parameterizing mesh names. No custom provider resources are needed — zone proxies are deployed via the existing Helm chart using `helm_release`.

##### Single-Mesh (Default)

`main.tf`:

```hcl
resource "konnect_mesh" "payments" {
  name = "payments-mesh"
  # ...
}

resource "helm_release" "zone_proxy" {
  name       = "kuma-zone"
  repository = "https://kumahq.github.io/charts"
  chart      = "kuma"
  namespace  = "kuma-system"

  values = [
    templatefile("${path.module}/values.tftpl", {
      zone               = "zone-1"
      kds_global_address = "grpcs://us.mesh.sync.konghq.tech:443"
      mesh               = konnect_mesh.payments.name
    })
  ]

  depends_on = [konnect_mesh.payments]
}
```

`values.tftpl`:

```yaml
kuma:
  controlPlane:
    mode: zone
    zone: ${zone}
    kdsGlobalAddress: ${kds_global_address}

  meshes:
    - name: ${mesh}
      ingress:
        enabled: true
      egress:
        enabled: true
```

**Validation:** Same runtime model as Helm — zone proxy retries until the mesh exists on the Global CP. Terraform's `depends_on` ensures the mesh is created in Konnect/Global CP before the Helm release is applied, so in practice the mesh will already exist by the time the zone proxy starts.

##### Existing Mesh Variant

When the mesh already exists (e.g., `default` mesh), skip the `konnect_mesh` resource and reference the mesh by variable:

```hcl
variable "mesh_name" {
  default = "default"
}

resource "helm_release" "zone_proxy" {
  name       = "kuma-zone"
  repository = "https://kumahq.github.io/charts"
  chart      = "kuma"
  namespace  = "kuma-system"

  values = [
    templatefile("${path.module}/values.tftpl", {
      zone               = "zone-1"
      kds_global_address = "grpcs://us.mesh.sync.konghq.tech:443"
      mesh               = var.mesh_name
    })
  ]
}
```

#### Summary: Validation Strategies by Tool

| Tool | Can Validate Mesh? | Strategy |
|------|-------------------|----------|
| **Konnect UI** | Yes (API access) | Pre-populate mesh dropdown, validate before generating YAML |
| **Helm (federated)** | No (offline install) | Accept name, fail gracefully at runtime with clear logs |
| **Helm (unfederated)** | Yes (creates mesh) | `createMesh: true` in subchart, fail at install if mismatch |
| **Terraform** | Via `depends_on` | `helm_release` with `templatefile()`; `depends_on` ensures mesh exists before zone proxy |

#### Design Decisions

##### 1. Mesh Deletion Handling

**Existing Protection**: Kuma already prevents mesh deletion when Dataplanes are attached.
The mesh validator returns an error: `"unable to delete mesh, there are still some dataplanes attached"`.
See [`pkg/core/managers/apis/mesh/mesh_validator.go#L70-L88`](https://github.com/kumahq/kuma/blob/master/pkg/core/managers/apis/mesh/mesh_validator.go#L70-L88).

**Mesh-Scoped Zone Proxy Benefit**: With zone proxies represented as Dataplane resources, this protection applies automatically.
No additional implementation is needed for mesh deletion handling - the existing safeguard covers the new deployment model.

**Note**: Current ZoneIngress/ZoneEgress resources are NOT covered by this protection (they're global-scoped).
The move to mesh-scoped Dataplanes resolves this gap.

For single-mesh, cleanup is straightforward: `helm uninstall <release-name>`.

##### 2. Per-Mesh Services (Not Shared)

**Decision**: Each mesh gets its own Service/LoadBalancer.

**Rationale**: With mesh-scoped zone proxies:
- Each mesh has **different mTLS CA certificates**
- Zone egress must verify the correct mesh's CA
- Sharing a LoadBalancer would require SNI-based cert selection (complex)

**Per-mesh services provide**:
- Proper mTLS isolation
- Simpler Envoy configuration
- Independent scaling and failover
- Clear network boundaries

**Service naming**: Follows the existing pattern from [`deployments/charts/kuma/templates/egress-service.yaml`](https://github.com/kumahq/kuma/blob/master/deployments/charts/kuma/templates/egress-service.yaml) but with per-mesh naming: `<release>-<mesh>-zoneproxy` (e.g., `kuma-payments-mesh-zoneproxy`).
This prevents name collisions when multiple meshes are deployed.

**Name length**: The 63-character limit only applies to **Service names** (DNS label, RFC 1123). Deployments, HPAs, PDBs, and ServiceAccounts use DNS subdomain names (253 chars) so they are not a concern. The zone proxy naming pattern `<release>-<mesh>-zoneproxy` should validate the Service name at install time:

```
{{ if gt (len (include "kuma.zoneProxy.serviceName" .)) 63 }} {{ fail "zone proxy service name exceeds 63 characters; use zoneProxy.service.name to set a shorter name" }} {{ end }}
```

To handle long mesh names, expose a `service.name` override per mesh entry (matching the existing `ingress.service.name` pattern in `_helpers.tpl:65-68`):

```yaml
meshes:
  - name: my-very-long-mesh-name-that-exceeds-limits
    ingress:
      enabled: true
      service:
        name: zp-long-mesh-ingress  # Override when auto-generated name is too long
    egress:
      enabled: true
      service:
        name: zp-long-mesh-egress
```

**Cost implication**: More LoadBalancers = higher cloud cost.
Users can use NodePort or Ingress controllers to reduce LB count if needed.

##### 3. Migration Path from Global Zone Proxies

**Additive migration**: The `meshes` config is added alongside the existing `ingress`/`egress` keys in the chart.

- **Phase 1**: Add `meshes` support, old keys still functional
- **Phase 2**: Add deprecation warnings for old keys
- **Phase 3**: Remove old keys in future major release

**Traffic migration** (switching live traffic from global to mesh-scoped zone proxies) is out of scope for this MADR and will be addressed separately.

##### 4. Helm Release Structure: Mesh Subchart

Each entry in the `meshes` list is rendered by a **mesh subchart**. The subchart encapsulates all resources for a single mesh's zone proxy deployment.

**Subchart responsibilities** (per mesh entry):
- Deployment(s) — zone ingress, zone egress, or combined
- Service(s) — per-mesh LoadBalancer/NodePort
- HPA — horizontal pod autoscaler
- PDB — pod disruption budget
- ServiceAccount
- Mesh resource (when `createMesh: true`)
- Default policies (controlled by `createPolicies` map)

| Aspect | Analysis |
|--------|----------|
| **Simplicity** | One `helm install` deploys everything; `meshes` list is declarative |
| **Multi-mesh** | Natural extension — add entries to the list |
| **Upgrades** | Single `helm upgrade` updates all meshes' zone proxies |
| **Isolation** | Each mesh subchart renders independent resources; no cross-mesh interference |
| **GitOps** | Single `values.yaml` is the source of truth for all meshes in a zone |
| **Lifecycle** | Per-mesh lifecycle managed within the subchart (create/update/delete mesh entries independently) |

**`createPolicies`**: A true/false map controlling which default policies are created alongside the mesh. Derived from `skipCreatingInitialPolicies` in [`api/mesh/v1alpha1/mesh.proto:72-75`](https://github.com/kumahq/kuma/blob/master/api/mesh/v1alpha1/mesh.proto#L72-L75), implemented in [`pkg/defaults/mesh/mesh.go:58-102`](https://github.com/kumahq/kuma/blob/master/pkg/defaults/mesh/mesh.go#L58-L102). The five policies are:
- **MeshCircuitBreaker** — default circuit breaker for all traffic
- **MeshRetry** — default retry configuration (TCP + HTTP/GRPC)
- **MeshTimeout** — default timeouts for sidecars and gateways
- **TrafficPermission** — allow-all traffic permission (only created when `createMeshDefaultRoutingResources=true`)
- **TrafficRoute** — default load-balancing route (only created when `createMeshDefaultRoutingResources=true`)

**`combinedProxies`**: A map with the same shape as `ingress`/`egress` (enabled, podSpec, hpa, pdb, service, resources). Merges ingress and egress into a single Deployment. Mutually exclusive with `ingress`/`egress` — Helm validates and errors if a mesh entry defines both. No merging ambiguity: the combined deployment has its own self-contained config. References the separate deployment topology MADR [^2] for analysis of shared vs separate deployments.

**Subchart implementation**: The subchart lives within the main `kuma` chart under `charts/mesh/` (or as a dependency). It is not published separately — users interact with it purely through the `meshes` list in `values.yaml`.

```bash
# Install with mesh subchart entries
helm install kuma kumahq/kuma -n kuma-system -f values.yaml
```

```yaml
# values.yaml
meshes:
  - name: default
    ingress:
      enabled: true
    egress:
      enabled: true
```

To add a mesh later, update `values.yaml` and upgrade:

```bash
helm upgrade kuma kumahq/kuma -n kuma-system -f values.yaml
```

##### Considered Alternatives

1. **Single release with flat `zoneProxy` config** — the previous recommended approach: flat `zoneProxy.enabled`/`zoneProxy.mesh` in the main chart, no subchart. Rejected because it doesn't naturally extend to multi-mesh and doesn't encapsulate mesh lifecycle (mesh creation, default policies).

2. **Zone-proxy subchart** (not mesh subchart) — a standalone `kuma-zone-proxy` subchart for zone proxy releases. Rejected because it's too narrow (only handles zone proxies, not mesh lifecycle) and adds maintenance burden for a partial solution.

3. **Multi-mesh out of scope** — previous decision to leave multi-mesh to users. Reversed because the mesh subchart naturally supports it without additional complexity.

##### 5. Namespace Placement: kuma-system

All meshes' zone proxies are deployed in the `kuma-system` namespace. This is consistent with the MADR 093 relaxation (see [below](#6-relaxation-of-madr-093-one-mesh-per-namespace)) which allows multiple meshes per namespace.

**K8s Naming Constraint**: Service names are limited to 63 characters (DNS label, RFC 1123). Other resources (Deployments, HPAs, PDBs) use DNS subdomain names (253 chars) and are not a concern. With the naming pattern `<release>-<mesh>-zoneproxy`, long mesh names may exceed the Service name limit. Users can override via the per-entry `service.name` field (see Per-Mesh Services above).

**Multi-mesh example** — all zone proxies in `kuma-system`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zone-proxy-payments-ingress
  namespace: kuma-system
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zone-proxy-payments-egress
  namespace: kuma-system
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zone-proxy-backend-ingress
  namespace: kuma-system
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: zone-proxy-backend-egress
  namespace: kuma-system
```

| Aspect | Analysis |
|--------|----------|
| **Operations** | All Kuma components in one place |
| **RBAC** | Simple - one namespace to grant access |
| **Monitoring** | Single namespace to scrape metrics |
| **Multi-mesh** | MADR 093 relaxation allows coexistence; naming pattern avoids collisions |

**Note on sidecar injection**: Zone proxies run `kuma-dp` directly as standalone Deployments (the same way current ZoneIngress/ZoneEgress work). They connect to the CP via bootstrap, not via sidecar injection. The `kuma-system` namespace does **not** need sidecar injection enabled for zone proxies to function.

##### 6. Relaxation of MADR 093 (One Mesh Per Namespace)

> **This reverses [MADR 093](093-disallow-multiple-meshes-per-k8s-ns.md) (accepted).**

**Why**: The mesh subchart deploys zone proxies for multiple meshes into `kuma-system`. Requiring separate namespaces for each mesh's infrastructure components adds operational complexity with no benefit — zone proxies are infrastructure, not application workloads.

**What changes**: Zone proxies for different meshes coexist in `kuma-system`. More broadly, this is a general relaxation — multiple meshes per namespace are allowed everywhere, not just `kuma-system`.

**Collision handling**: The Workload controller fails with a clear error if a Workload name collision occurs across meshes. For zone proxies this is inherently avoided by the naming pattern `zone-proxy-<mesh>-<role>`, which guarantees unique Workload names per mesh.

**Scope**: General relaxation — allow multiple meshes per namespace everywhere. Handle collisions in the Workload controller rather than preventing the configuration.

### Question 1: Support kuma.io/ingress-public-address

#### Current Implementation

From `pkg/plugins/runtime/k8s/controllers/ingress_converter.go:24-71`, the address resolution follows this priority:

1. Pod annotations: `kuma.io/ingress-public-address` + `kuma.io/ingress-public-port`
2. Service LoadBalancer IP/Hostname
3. NodePort with node address (ExternalIP > InternalIP)

#### Options

| Option | Description |
|--------|-------------|
| **A. Keep annotation** | Continue supporting the annotation as an override mechanism |
| **B. Service-only** | Remove annotation, rely solely on Service configuration |

#### Analysis

Use cases for annotation override:
- NAT gateways where Service IP differs from externally accessible address
- Split DNS environments
- Cloud provider quirks where LoadBalancer metadata is incorrect
- On-premises environments with external load balancers

**Option A: Keep annotation (recommended)**
- Advantages:
  - Supports edge cases where Service address doesn't reflect reality
  - Low maintenance burden
  - Backward compatible
- Disadvantages:
  - Additional configuration option to document

**Option B: Service-only**
- Advantages:
  - Simpler model
  - Encourages proper Service configuration
- Disadvantages:
  - Breaks legitimate use cases
  - Forces workarounds in complex network topologies

#### Recommendation

**Option A: Keep annotation support** but document it as an escape hatch:
- Primary method: Configure Service (LoadBalancer/NodePort) correctly
- Annotation: Use only when Service address is not accessible from other zones
- Consider adding a deprecation warning in logs when annotation is used

### Question 2: Default Helm Installation Behavior

#### Recommendation

**`meshes: []`** — no zone proxies are deployed by default. Zone proxy deployment requires explicit opt-in by adding entries to the `meshes` list:

```yaml
controlPlane:
  mode: zone
  zone: zone-1

# No zone proxies by default
meshes: []

# Opt-in: add mesh entries
meshes:
  - name: default
    ingress:
      enabled: true
    egress:
      enabled: true
```

This avoids a broken state where zone proxy Deployments target meshes that don't exist. The `skipMeshCreation` flag is orthogonal — it controls whether the CP auto-creates the `default` Mesh at startup, and is independent of zone proxy deployment.

## Decision

### Tooling Decisions

1. **Mesh subchart**: Each entry in the `meshes` list is rendered by a mesh subchart that manages Deployment(s), Service(s), HPA, PDB, ServiceAccount, and optionally the Mesh resource and default policies.

2. **Per-mesh Services**: Each mesh gets its own Service/LoadBalancer for proper mTLS isolation.
   Sharing a LoadBalancer would require SNI-based cert selection which adds complexity.

3. **Namespace placement**: Deploy all meshes' zone proxies in `kuma-system` namespace.

4. **MADR 093 relaxation**: Allow multiple meshes per namespace everywhere (general relaxation, not scoped to zone proxies only). Handle Workload name collisions in the Workload controller with clear error messages rather than preventing the configuration.

5. **Additive migration**: The `meshes` config is added alongside existing `ingress`/`egress` keys. Old keys are deprecated and removed in a future major release.

6. **`combinedProxies` deployment option**: A map with the same shape as `ingress`/`egress` (enabled, podSpec, hpa, pdb, service, resources), merging ingress and egress into a single Deployment. Mutually exclusive with `ingress`/`egress` — Helm validates and errors if both are defined.

7. **Conditional mesh creation + `createPolicies` map**: `createMesh: true` renders a Mesh resource (for unfederated zones). `createPolicies` is a true/false map controlling which default policies (MeshCircuitBreaker, MeshRetry, MeshTimeout, TrafficPermission, TrafficRoute) are created with the mesh.

### Design Questions

1. **Keep kuma.io/ingress-public-address**: Support the annotation as an escape hatch for complex network topologies, but document Service-based configuration as the primary method.

2. **Helm defaults**: `meshes: []` — explicit opt-in, consistent with `ingress.enabled` / `egress.enabled`.
   The `skipMeshCreation` flag is orthogonal to zone proxy deployment.

Note: Zone proxy deployment topology (shared vs separate ingress/egress) is addressed in a separate MADR. [^2]

### Out of Scope (Deferred to Resource Model MADR)

The following topics are covered in a separate MADR:

- **Dataplane representation**: Fields, labels, `kuma.io/proxy-type` tag
- **Workload identity**: `kuma.io/workload` annotation and auto-generation pattern (`zone-proxy-<mesh>-<role>`)
- **Token model**: Zone tokens → DP tokens transition for authentication
- **Universal deployment**: VM/bare metal specifics with mesh-scoped Dataplane resources
- **Sidecar vs standalone**: Whether zone proxies should be sidecars to fake containers

## Notes

### Related Issues and MADRs

- [Kuma #15429](https://github.com/kumahq/kuma/issues/15429): Label-based MeshService matching (no inbound tags)
- [Kuma #15431](https://github.com/kumahq/kuma/issues/15431): Protocol stored in Inbound field, not tags
- [KM #9028](https://github.com/Kong/kong-mesh/issues/9028): Dataplane fields for zone proxies
- [KM #9029](https://github.com/Kong/kong-mesh/issues/9029): Policies on zone proxies
- [KM #9032](https://github.com/Kong/kong-mesh/issues/9032): MeshIdentity for zone egress
- MADR 090: Zone Egress Identity

### Key Files Reference

| Component | File Path |
|-----------|-----------|
| ZoneIngress proto | `api/mesh/v1alpha1/zone_ingress.proto` |
| ZoneEgress proto | `api/mesh/v1alpha1/zoneegress.proto` |
| Dataplane proto | `api/mesh/v1alpha1/dataplane.proto` |
| Mesh proto (skipCreatingInitialPolicies) | `api/mesh/v1alpha1/mesh.proto` |
| Default mesh policies | `pkg/defaults/mesh/mesh.go` |
| K8s Ingress converter | `pkg/plugins/runtime/k8s/controllers/ingress_converter.go` |
| K8s Egress converter | `pkg/plugins/runtime/k8s/controllers/egress_converter.go` |
| Helm ingress deployment | `deployments/charts/kuma/templates/ingress-deployment.yaml` |
| Helm egress deployment | `deployments/charts/kuma/templates/egress-deployment.yaml` |
| Helm values | `deployments/charts/kuma/values.yaml` |
| Annotations | `pkg/plugins/runtime/k8s/metadata/annotations.go` |

[^1]: Resource model MADR link will be backfilled when created.
[^2]: Zone proxy deployment topology MADR link will be backfilled when created.
