#!/usr/bin/env bash

set -o errexit
set -o pipefail

function envoy_version() {
  # Returns Envoy version by ENVOY_TAG:
  # - if ENVOY_TAG is a real git tag like 'v1.20.0' then the version is equal to '1.20.0' (without the first letter 'v').
  # - if ENVOY_TAG is a commit hash then the version will look like '1.20.1-dev-b16d390f'

  ENVOY_TAG=${ENVOY_TAG:-"v1.22.1"}
  ENVOY_VERSION=$(curl --silent --location "https://raw.githubusercontent.com/envoyproxy/envoy/${ENVOY_TAG}/VERSION.txt")

  # for envoy versions older than v1.22.0 file 'VERSION.txt' used to be called 'VERSION'
  if [[ "${ENVOY_VERSION}" == "404: Not Found" ]]; then
    ENVOY_VERSION=$(curl --silent --location --fail "https://raw.githubusercontent.com/envoyproxy/envoy/${ENVOY_TAG}/VERSION")
  fi

  if [[ "${ENVOY_TAG}" =~ ^v[0-9]*\.[0-9]*\.[0-9]*$ ]]; then
    echo "${ENVOY_VERSION}"
  else
    echo "${ENVOY_VERSION}-${ENVOY_TAG:0:8}"
  fi
}

function version_info() {
  # Kuma version info has the following format:
  # version git-tag git-commit build-date envoy-version

  # version is built as follows:
  # 1) If a git tag is present on the current commit, then the version is a git tag (without a starting v)
  # 2) If the branch is "release-X.Y" look at the existing tags and use either X.Y.0-<shortHash> if there's none or
  #     the latest tag with a patch version increased by one (.e.g if latest tag is X.Y.1 the version will be `X.Y.2-<shortHash>`)
  # 3) In non release branch use `0.0.0-$shortHash`

  # git-tag - if the current HEAD has a tag then use it otherwise it's "not-tagged"
  # git-commit - HEAD sha
  # build-date - local date if built on CI

  # Note: this format must be changed carefully, other scripts depend on it

  envoyVersion=$(envoy_version)

  if [[ -z "${CI}" ]]; then
    versionDate="local-build"
  else
    versionDate=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  fi

  if ! command -v git &> /dev/null; then
    shortHash="no-git"
    currentBranch="no-git"
    exactTag="no-git"
    describedTag="no-git"
  else
    if git diff --quiet && git diff --cached --quiet; then
      longHash=$(git rev-parse HEAD 2>/dev/null || echo "no-commit")
      shortHash=$(git rev-parse --short=9 HEAD 2> /dev/null || echo "no-commit")
      describedTag=$(git describe --tags 2>/dev/null || echo "none")
    else
      longHash="local-build"
      shortHash="local-build"
      describedTag="local-build"
    fi

    currentBranch=$(git rev-parse --abbrev-ref HEAD 2> /dev/null || echo "no-branch")
    exactTag=$(git describe --exact-match --tags 2> /dev/null || echo "not-tagged")
    # We only support tags of the format: "v?X.Y.Z(-<alphaNumericName>)?" all other tags will just be ignored and use the regular versioning scheme
    if [[ ${exactTag} =~ ^v?[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9]+)?$ ]]; then
      echo "${exactTag/^v//} $describedTag $longHash $versionDate $envoyVersion"
      exit 0
    fi

    if [[ ${currentBranch} == release-* ]]; then
        releasePrefix=${currentBranch//release-/}
        lastGitTag=$(git tag -l | grep -E "^v?${releasePrefix}\.[0-9]+$" | sed 's/^v//'| sort -V | tail -1)
        if [[ ${lastGitTag} ]]; then
          IFS=. read -r major minor patch <<< "${lastGitTag}"
          version="${major}.${minor}.$((++patch))-preview.v${shortHash}"
        else
          version="${releasePrefix}.0-preview.v${shortHash}"
        fi
    else
      version="0.0.0-preview.v${shortHash}"
    fi
  fi

  echo "$version $describedTag $longHash $versionDate $envoyVersion"
}

version_info