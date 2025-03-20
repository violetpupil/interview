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

### json解码number到any，再编码数值变化

```golang
s := "558242949628978650"
var a any
json.Unmarshal([]byte(s), &a) // 底层类型为float64
bs, _ := json.Marshal(a) // 编码float64时，精度丢失
fmt.Println(string(bs)) // 558242949628978600
```
