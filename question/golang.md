# golang

## 常见错误

len(中文字符串)返回的是byte数组的值，应该用len([]rune(s))

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

### int和float转换丢失精度

```golang
var i float64 = 558242949628978650
fmt.Println(int(i)) // 直接转int变为了558242949628978624
bs, _ := json.Marshal(i)
fmt.Println(string(bs)) // json编码变为了558242949628978600
var j int
json.Unmarshal(bs, &j)
fmt.Println(j) // json解码回int，保持558242949628978600
```
