# 描述
从 OKLink 下载源代码（默认链为 XLayer）

# 安装
```shell
go install github.com/cuiweixie/xlayer-dlsc
```

# 用法
```shell
xlayer-dlsc -h
```

```aiignore
xlayer-dlsc 使用说明：
  -address string
        合约地址（必填）
  -chain string
        链的简称（默认值："xlayer"）
  -out string
        输出目录（默认值：合约地址）
```

示例：下载 XLayer 上的 USDT0（合约地址：0x779ded0c9e1022225f8e0630b35a9b54be713736）
```shell
xlayer-dlsc -out test_output -address 0x779ded0c9e1022225f8e0630b35a9b54be713736
```