# golang

## 常见错误

### 浮点数丢失精度

对浮点数做运算有时候会丢失精度

```golang
var a float64 = 7.65
// 相加后精度丢失，b=7.6499999999999995
var b float64 = float64(1.01) + float64(6.64)
// 结果为 false
fmt.Println(a == b)
```

解决方案：需要保证精度时，比如支付，使用 [decimal.Decimal](https://pkg.go.dev/github.com/shopspring/decimal) 类型
