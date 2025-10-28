# description
dowonload source code from oklink(default xlayer)
# install 
```shell
go install github.com/cuiweixie/xlayer-dlsc
```
# usage
```shell
xlayer-dlsc -h
```
```aiignore
Usage of xlayer-dlsc:
  -address string
        contract address (required)
  -chain string
        chain short name (default "xlayer")
  -out string
        output directory (default: address value)
```

example, download usdt0 on xlayer(0x779ded0c9e1022225f8e0630b35a9b54be713736):
```shell
xlayer-dlsc -out test_output -address 0x779ded0c9e1022225f8e0630b35a9b54be713736
```
