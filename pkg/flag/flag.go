package main

import (
    "flag"
    "fmt"
    "os"
)

// https://www.jianshu.com/p/f9cf46a4de0e
// 实际中应该用更好的变量名
var (
    h bool

    v, V bool
    t, T bool
    q    *bool

    s string
    p string
    c string
    g string
)

func init() {
    flag.BoolVar(&h, "h", false, "this help")

    flag.BoolVar(&v, "v", false, "show version and exit")
    flag.BoolVar(&V, "V", false, "show version and configure options then exit")

    flag.BoolVar(&t, "t", false, "test configuration and exit")
    flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

    // 另一种绑定方式
    q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

    // 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
    flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
    flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
    flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
    flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

    // 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
    flag.Usage = usage
}

func main() {
    flag.Parse()   // 解析命令行输入的参数，当遇到单独的一个“-”或不是“-”开始时，会停止解析

    if h {    // 根据参数值，进入分支程序，执行相关功能
        flag.Usage()
    }
}

func usage() {
    fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
    flag.PrintDefaults()
}