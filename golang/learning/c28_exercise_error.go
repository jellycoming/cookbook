package main

import "fmt"

type ErrNegativeSqrt float64

// 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。
// 因为 e 为 error 类型，在 fmt.Sprint(e) 时会调用 e.Error() 来输出错误信息，这样就造成了死循环。
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", e)
}

func SqrtE(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := float64(1)
		for i := 0; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
			fmt.Println(z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(SqrtE(2))
	fmt.Println(SqrtE(-2))
}
