# terraform-cascade

[![asciicast](https://asciinema.org/a/598560.svg)](https://asciinema.org/a/598560)

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
