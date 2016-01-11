rands [![Build Status](https://travis-ci.org/issue9/rands.svg?branch=master)](https://travis-ci.org/issue9/rands)
======

rands 为一个随机字符串生成工具。
```go
// 生成一个长度为[8,10)之间的随机字符串，包含小写与数字字符
str := rands.String(8, 10, Lower, Digit)


// 生成一个带缓存功能的随机字符串生成器
r := rands.New(100, 5, 7, Lower, Digit, Punct)
str1 := r.String()
str2 := r.String()
```

### 安装

```shell
go get github.com/issue9/rands
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/rands)
[![GoDoc](https://godoc.org/github.com/issue9/rands?status.svg)](https://godoc.org/github.com/issue9/rands)


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
