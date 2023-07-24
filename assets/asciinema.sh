#!/usr/bin/env bash -e

# This is the code structure (from a basic example)
cd samples/basic
tree -P backend.tf .

# Each depth level has a dependency on the previous one. For example:
cat dev/base/data.tf
cat dev/base/vpc.tf

# In this case it simulates that the VPC needs first the account where it has to be created
cat base/accounts.tf

# Let's run terraform init through cascade recursively
go run ../.. init --cascade-recursive

# Now, let's run apply
go run ../.. apply --cascade-recursive --auto-approve

# Let's see the created "infra"
tree -a -I .terraform /tmp/cascade

# And the state files
tree -a /tmp/cascade/.terraform/
