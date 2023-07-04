package _var

var (
	TestVar = "1"
)

// 测试 init 后被导入的变量值
func init() {
	TestVar = "3"
}
