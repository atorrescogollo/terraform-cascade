#!/bin/bash

set -eo pipefail

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
WORKING_DIR="${SCRIPT_DIR}/../basic/"

cd "${WORKING_DIR}"


find . -type f -name "backend.tf" -delete
find . -type f -name "*.tf" -not -name "backend.tf" -exec dirname {} \; | sort -du | while read -r dir; do
  echo "Generating backend for ${dir}"
  safe_dir="$(echo "${dir}" | sed -e 's|^\./||g' -e 's|/|_|g')"
  cat <<EOF > "${dir}/backend.tf"
terraform {
  backend "local" {
    path = "/tmp/.terraform/${safe_dir}.tfstate"
  }
}
EOF
done
