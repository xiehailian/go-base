对一些工作中遇到的golang的问题，总结成代码，以备不时之需。这里的每一个文件都要阐明一个问题，并且能够直接运行。

> * [Golang 工程经验](https://juejin.im/post/5a6873fb518825733e60a1ae)
> * [RPC](https://blog.csdn.net/liubenlong007/article/details/54692241)
> * [Bigcache](https://github.com/allegro/bigcache)
> * [Golang 性能监控](http://xiaorui.cc/2016/03/20/golang%E4%BD%BF%E7%94%A8pprof%E7%9B%91%E6%8E%A7%E6%80%A7%E8%83%BD%E5%8F%8Agc%E8%B0%83%E4%BC%98/)


### Go 工程实践

#### Go 开发环境

- [ ] 下载[Golang](https://studygolang.com/dl)
- [ ] 设置环境变量，GOROOT: C:\GO；GOPATH: C:\GoPath
- [ ] GoLand


#### Go 生产环境


#### Go 包管理
  
* GO111MODULE 环境变量
    * `auto`：默认，有go.mod使用go modules寻找依赖，否则使用旧的GOPATH和vendor机制
    * `on`: 使用go modules寻找依赖
    * `off`：不使用go modules寻找依赖


* 初始化包
    ``` bash
    mkdir go-base
    cd go-base
    go mod init go-base    // 初始化包
    ```

* go.mod
  ``` 

  module suanjing

  require (
	github.com/mattn/go-gtk v0.0.0-20181205025739-e9a6766929f6 // indirect  // indirect表明此包为间接依赖
	github.com/coreos/etcd v3.3.9 +incompatible                // 在版本信息后加上+incompatible就可以不需要指定/vN
    gopkg.in/fatih/pool.v2 v2.0.0                              // 对于2以上的版本，不加+incompatible，要指定/vN
    github.com/pquerna/ffjson v0.0.0-20180717144149-af8b230fcd20
    )

  replace (
	gitlab.baixing.cn/baixing_search/go-common => ../go-common    // 只能替换顶层依赖，不能替换间接依赖
	gitlab.baixing.cn/baixing_search/suanjing => ./
    )
  ```


