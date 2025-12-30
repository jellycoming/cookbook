package review

import "fmt"

//
// 用channel将分别打印`A`、`B`、`C`的三个函数连接起来，循环依次打印`ABC`
//
// Output:
//	ABC
//	ABC
//	ABC
//	ABC
//	ABC
//	quit C
//	quit B
//	quit A
//
func PrintLoop() {
	a := make(chan struct{})
	b := make(chan struct{})
	c := make(chan struct{})
	end := make(chan struct{})
	quit := make(chan struct{})
	go printA(a, b, quit)
	go printB(b, c, quit)
	go printC(c, end, quit)
	for i := 0; i < 5; i++ {
		a <- struct{}{}
		<-end
		fmt.Printf("\n")
	}
	close(quit)
}

func printA(in, out, quit chan struct{}) {
	for {
		select {
		case <-in:
			fmt.Print("A")
			out <- struct{}{}
		case <-quit:
			fmt.Println("quit A")
			return
		}
	}
}

func printB(in, out, quit chan struct{}) {
	for {
		select {
		case <-in:
			fmt.Print("B")
			out <- struct{}{}
		case <-quit:
			fmt.Println("quit B")
			return
		}
	}
}

func printC(in, out, quit chan struct{}) {
	for {
		select {
		case <-in:
			fmt.Print("C")
			out <- struct{}{}
		case <-quit:
			fmt.Println("quit C")
			return
		}
	}
}

// 交替打印数字与字母
//
// Output:
//	12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
//
func PrintLetter() {
	var num, char = 1, 'A'
	ch := make(chan int, 1)
	for i := 0; i < 28; i++ {
		select {
		case _ = <-ch:
			if char <= 'Z' {
				fmt.Printf("%c%c", char, char+1)
				char += 2
			}
		case ch <- i:
			if num <= 28 {
				fmt.Printf("%d%d", num, num+1)
				num += 2
			}
		}
	}
}