# Terraform Cascade
<img src="./assets/gopher.png" height="200">

## Demo
[![asciicast](https://asciinema.org/a/JPYlivXxoZvB5PvjNvOxVQb7O.svg)](https://asciinema.org/a/JPYlivXxoZvB5PvjNvOxVQb7O)

# Overview
**Terraform Cascade** is a terraform-like tool that allows you to manage multiple terraform projects.

It's made to be fully compatible with terraform, so you can use it as a drop-in replacement. However, it requires the **terraform binary to be available in the PATH**.

# Design
It works with a very opinionated design:
* Every project is inside a deep directory structure.
* To define a project, you only need to place a `backend.tf` file in that directory.
* In each layer, will be executed in the following order:
    1. **Current directory** (only when it has a `backend.tf` file)
    2. **Whole `base` directory** (with its layer)
    3. **Other directories** (with its layer)


# Usage

### Build
```
docker build -t cascade .
```

### Run example
```
cd samples/basic/ # Some sample project that has dependencies between layers
```

```
# Full dependency tree in order
docker run -it --rm -v $(pwd):/w -v $(pwd)/tmp:/tmp -w /w cascade init --cascade-recursive
docker run -it --rm -v $(pwd):/w -v $(pwd)/tmp:/tmp -w /w cascade apply --cascade-recursive --auto-approve
```

```
# Full dependency tree in parallel
docker run -it --rm -v $(pwd):/w -v $(pwd)/tmp:/tmp -w /w/dev cascade apply --cascade-recursive --auto-approve
```



### Generated infra
```
$ tree -a -I .terraform tmp/cascade
tmp/cascade
├── dev
│   ├── .account
│   ├── s3
│   │   └── .s3
│   └── vpc
│       ├── .vpc
│       └── eks
│           └── .eks
├── ops
│   ├── .account
│   └── vpc
│       └── .vpc
└── prod
    ├── .account
    └── vpc
        ├── .vpc
        └── eks
            └── .eks
```

### Generated terraform states
```
$ tree -a tmp/cascade/.terraform/
tmp/cascade/.terraform/
├── base.tfstate
├── dev_base.tfstate
├── dev_eks.tfstate
├── ops_base.tfstate
├── prod_base.tfstate
└── prod_eks.tfstate
```
