# terraform-cascade

```
docker build -t cascade .
```

```
cd samples/basic/
docker run -it --rm -v $(pwd):/w -v $(pwd)/samples/tmp:/tmp -w /w cascade init --cascade-recursive
docker run -it --rm -v $(pwd):/w -v $(pwd)/samples/tmp:/tmp -w /w cascade apply --cascade-recursive
```
