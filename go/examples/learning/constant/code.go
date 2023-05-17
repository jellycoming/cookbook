package constant

import "fmt"

var (
	IOS     = 1
	ANDROID = 2
)

func init() {
	fmt.Println("init pkg constant, file code.go")
}
