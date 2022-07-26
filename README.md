# safe
Go stack Package

## 示例

```go
import "github.com/tinystack/safe"

// 安全的启动goroutine
// 当发生 panic 时会捕捉相关报错信息并通过 SetLogger 设置的 logger 实例输出报错信息和堆栈信息
Go(func() {
    panic("test panic")
})

// 安全的启动goroutine
// 当发生 panic 时会执行对应的recover函数
GoWithRecover(func() {
    panic("test panic")
}, func(err interface{}) {
    fmt.Println(err)
})
```

### API

- safe.SetLogger(l Logger)
- safe.Go(goroutine func())
- safe.GoWithRecover(goroutine func(), customRecover func(err interface{}))