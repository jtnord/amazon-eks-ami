#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

source /helpers.sh

mock::kubelet 1.27.0
wait::dbus-ready

nodeadm init --config-source file://config.yaml

assert::files-equal /etc/containerd/config.toml expected-containerd-config.toml
