# Developer documentation

Hello, and welcome! If you're thinking about contributing to Kuma's code base, you came to the
right place. This document serves as guide/reference for technical
contributors. 

##  Dependencies

The following dependencies are all necessary. Please follow the installation instructions 
for any tools/libraries that you may be missing.

### Command line tool

- [`curl`](https://curl.haxx.se/)  
- [`git`](https://git-scm.com/)
- [`unzip`](http://infozip.sourceforge.net/UnZip.html)
- [`make`](https://www.gnu.org/software/make/)
- [`go`](https://golang.org/)
- [`ninja`](https://github.com/ninja-build/ninja)
- [`mise`](https://mise.jdx.dev)

### Helper commands

Throughout this guide, we will use `make` to run pre-defined tasks in the [Makefile](Makefile).
Use the following command to list out all the possible commands:

```bash
make help
```

### Installing dev tools

We use [mise](https://mise.jdx.dev) to manage dev tools for Kuma.
If you had mise installed previously we recommend running:

```bash
mise cache clear
rm -rf ~/.local/share/mise/{installs,plugins}/kubectl
```

If you're experiencing errors with installing `kubectl` or `container-structure-test`.
Please see https://github.com/jdx/mise/issues/3110 for details.

To install dev tools you can run:

```bash
make install
```

## Code checks

To run all code formatting, linting and vetting tools use the target:
```bash
make check
```

## Testing

We use Ginkgo as our testing framework. To run the existing test suite, you have several options:

For all tests, run:
```bash
make test
```
And you can run tests that are specific to a part of Kuma by appending the app name as shown below:
```bash
make test/kumactl
```
For even more specific tests you can specify the package you want to run tests from:
```bash
make test TEST_PKG_LIST=<pkgPath>
```
`pkgPath` is a package list selector for example: `./pkg/xds/...` will run all tests in the `pkg/xds` subtree.

There's a large set of integration tests that can be run with:

```bash
make test/e2e
```

These tests are big and slow, it is recommended to read [e2e-test-tips](docs/guides/e2e-test-tips.md) before running them.

## Building

To build all the binaries run:
```bash
make build
```

Like `make test`, you can append the app name to the target to build a specific binary. For example, here is how you would build the binary for only `kumactl`:
```bash
make build/kumactl
```
This could help expedite your development process if you only made changes to the `kumactl` files.

## Code live reload on k8s

With Skaffold we can utilize code live reload on k8s cluster. Everytime you make a change in code Kuma will be rebuild and reinstalled on cluster.

1. Run `make k3d/start`
2. Run `make dev/fetch-demo` to get the Kuma counter demo app
3. Run `skaffold dev`

## Debugging

Like any other go program Kuma can be debugged using [dlv](https://github.com/go-delve/delve).
In this section we'll go into how to trigger a breakpoint both in K8S and Universal.

### K8S

1. Run `make k3d/start`
2. Run `skaffold debug`
3. Run goland/vscode debugger with remote target on port `56268` (Skaffold will log exposed port for debugging)
4. Put a breakpoint where you want 
5. Enjoy!

### Universal

1. Add `4000` port in [UniversalApp](https://github.com/kumahq/kuma/blob/201413bdd532e92ff6e1fd017c4970073ba0c09f/test/framework/universal_app.go#L223) so that it's exposed, and the debugger can connect.
2. Remove "-w -s" from LDFLAGS [here](https://github.com/kumahq/kuma/blob/7398d8901798d5cf1c2715e036204fc3632ec45d/mk/build.mk#L2) so that debugging symbols are not stripped
3. Add a `time.Sleep` in a place where you want to debug the test
4. Run the tests `make -j test/e2e/debug EXTRA_GOFLAGS='-gcflags "all=-N -l"' E2E_PKG_LIST=./test/e2e_env/universal/...`
5. Wait to hit the `time.Sleep`
6. Figure out the `kuma-cp` container id by running: `docker ps | grep kuma-cp`
7. Exec into the container: `docker exec -it kuma-3_kuma-cp_3dkYrT bash`
8. Download the same go as in `go.mod` - e.g. `curl -o golang https://dl.google.com/go/go1.23.2.linux-arm64.tar.gz`
9. Extract using `tar xzf golang`
10. Install `dlv` [version that is closes](https://github.com/go-delve/delve/releases) to the `go.mod` version in the container, run: `go install github.com/go-delve/delve/cmd/dlv@vCLOSEST_DLV_VERSION`
11. Run `dlv --listen=:4000 --headless=true --api-version=2 --accept-multiclient attach 1`
12. Figure out the port on the host machine `docker ps | grep kuma-3_kuma-cp_3dkYrT`, look for the port forward for `4000`
13. Run goland/vscode debugger with remote target on port from point 12
14. Enjoy!

## Debugging E2E tests interactively

By using the JetBrains Goland IDE, you could also debug the E2E tests running in Kubernetes interactively. To do so, follow these steps:

1. Prepare the images by running `make build`, `make images` and `make docker/tag`
2. Start the Kubernetes clusters by running `make test/e2e/k8s/start`
3. Create a run configuration by clicking the "Play" button aside the package name in a test suite file (e.g. `test/e2e_env/kubernetes/kubernetes_suite_test.go`)
4. Quickly stop the execution to avoid running the tests, we only want the automatically generated run configuration
5. Open the "Run/Debug configuration" dialog and locate and edit the configuration created by using the settings mentioned later in this document
6. Locate to particular file containing the test case you want to debug, and mark the code block to `FDescribe` or `FIt` to run only that test case
7. **Run the configuration in debug mode by clicking the bug 🪲 icon and enjoy debugging**
8. Stop the clusters after debugging complete by executing `make test/e2e/k8s/stop`


Required settings for the run configuration:

1. Environment variables
   1. `KUMACTLBIN=<path-to-built-kumactl>` should point to the `kumactl` of your local build, e.g. `/Users/<username>/go/src/github.com/<username>/kuma/build/artifacts-darwin-arm64/kumactl/kumactl`
   2. `KUMA_K8S_TYPE=k3d` or `kind` if you also specified when starting the clusters 
   3. `K3D_NETWORK_CNI=flannel` only needed if you are using the `k3d` cluster type, which is the default
2. Go tool arguments
   1. Prepare the values for Envoy version and Kuma product version
      1. Envoy version: execute `ENVOY_VERSION=$(make build/info | grep Envoy | cut -d '=' -f 2)`
      2. Kuma product version: execute `KUMA_PRODUCT_VERSION=$(./tools/releases/version.sh | awk '{print $1}')`
   2. Set the Go tools arguments using output from this command: `echo "-ldflags='-X github.com/kumahq/kuma/pkg/version.Envoy=$ENVOY_VERSION -X github.com/kumahq/kuma/pkg/version.version=$KUMA_PRODUCT_VERSION'"` 

When you need to restart the debugging:

1. Rebuild the images if needed
   1. To rebuild images, execute `make build`, `make images` and `make docker/tag`
   2. To load the new images into the clusters, execute `make test/e2e/k8s/load/images/kuma-1` and `make test/e2e/k8s/load/images/kuma-2` 
2. Restart the clusters if needed
   1. To restart the clusters, execute `make test/e2e/k8s/stop` and then `make test/e2e/k8s/start`
   2. We usually don't need to restart the clusters when the debugged cases are sharing the cluster environment (under directory `test/e2e_env`) because they clean up the cluster correctly
   3. We need to restart the clusters when debugging the cases that do not share the cluster in normal runs or the cluster state is polluted for some reason.

## Running

### Kubernetes

Execute
```bash
make -j k3d/restart
```

To stop any existing Kuma K3D cluster, start a new K3D cluster, load images, deploy Kuma and Kuma counter demo. 

### GKE

You can test development versions by first pushing to `gcr.io`:

```bash
gcloud auth configure-docker
make images/push DOCKER_REGISTRY=gcr.io/proj-123456
```

then setting up `kubectl` to connect to your cluster and installing Kuma:

```bash
gcloud container clusters get-credentials cluster-name
kumactl install control-plane --registry=gcr.io/proj-123456 | kubectl apply -f -
```
