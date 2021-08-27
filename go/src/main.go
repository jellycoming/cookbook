package main

import (
	ds "cookbook/go/src/data-structure"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/url"
	"strconv"

	//"sort"
	"time"
)

type In struct {
	ad []*Ad
}
type Ad struct {
	Score float64
	Name  string
	Ecpm  uint32
	Cvr   float64
	*Label
	M map[uint32]uint32
	L []int
	*ds.Node
}

func (ad *Ad) String() string {
	return fmt.Sprintf("Ad:%s,Score:%f,Ecpm:%d,Cvr:%f\n", ad.Name, ad.Score, ad.Ecpm, ad.Cvr)
}

type Label struct {
	Version string
}

// 按score降序排序
type ScoreSortedAd []*Ad

func (ad ScoreSortedAd) Len() int {
	return len(ad)
}

func (ad ScoreSortedAd) Less(i, j int) bool {
	return ad[i].Score > ad[j].Score
}

func (ad ScoreSortedAd) Swap(i, j int) {
	ad[i], ad[j] = ad[j], ad[i]
}

type EcpmSortedAd []*Ad

func (ad EcpmSortedAd) Len() int {
	return len(ad)
}

func (ad EcpmSortedAd) Less(i, j int) bool {
	return ad[i].Ecpm > ad[j].Ecpm
}

func (ad EcpmSortedAd) Swap(i, j int) {
	ad[i], ad[j] = ad[j], ad[i]
}

// 生成min到max之间的随机float64类型小数
func RandFloat64(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Float64()*(max-min)
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func Test1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}

func Test2() (r int) {
	defer func(r int) {
		fmt.Println(r)
		r = r + 2
	}(r)
	return 2
}

func Test3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)
	return 2
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}

func s() []int {
	return nil
}

func gateway() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println(c.RemoteAddr())
	for {
		t := time.Now()
		//_, err := io.WriteString(c, t.Format("2006-01-02 00:00:00\n"))
		_, err := io.WriteString(c, t.Format("20060102000000\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func hmacSha256(data string, secret string) string {
	now := time.Now().UnixNano() / 1e6
	fmt.Println(now)
	stringToSign := fmt.Sprintf("%d\n%s", now, secret)
	fmt.Println(stringToSign)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	buf := h.Sum(nil)
	return url.QueryEscape(base64.StdEncoding.EncodeToString(buf))
}

func testAd(ad *Ad) {
	if ad == nil {
		fmt.Println("ad is nil")
	}
	if ad == nil || ad.Label == nil {
		fmt.Println("test Ad")
	}
	if v := ad.M[123]; v > 0 {
		fmt.Println(">0")
	} else {
		fmt.Println("<0")
		fmt.Println(v)
	}
}

func letter(ch1, ch2 chan int32) {
	for {
		char, ok := <-ch1
		if !ok {
			break
		}
		fmt.Printf("%c", char)
	}
}

func sender() {
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

func Modify(ad Ad) {
	ad.Name = "hello"
	ad.Score = 0.1
	ad.Label = nil
	ad.M[1] = 2
	ad.L[0] = 1
}

func dog(dogCh, catCh, quit chan struct{}) {
	for {
		select {
		case <-dogCh:
			fmt.Println("dog")
			catCh <- struct{}{}
		case <-quit:
			fmt.Println("quit dog")
			return
		}
	}
}

func cat(catCh, fishCh, quit chan struct{}) {
	for {
		select {
		case <-catCh:
			fmt.Println("cat")
			fishCh <- struct{}{}
		case <-quit:
			fmt.Println("quit cat")
			return
		}
	}
}

func fish(fishCh, end, quit chan struct{}) {
	for {
		select {
		case <-fishCh:
			fmt.Println("fish")
			end <- struct{}{}
		case <-quit:
			fmt.Println("quit fish")
			return
		}
	}
}

func main() {
	var a uint32 = 4
	var b uint32 = 5
	println(a - b)
	// 4+(-5)
	// 00000000 00000000 00000000 00000100

	// 10000000 00000000 00000000 00000101 原码
	// 11111111 11111111 11111111 11111010 反码
	// 11111111 11111111 11111111 11111011 补码

	// 11111111 11111111 11111111 11111111
	i, _ := strconv.ParseUint("11111111111111111111111111111111", 10, 32)
	println(i)
	var m map[int]int
	println(m[1])
	//m[1]=1
	//dogCh := make(chan struct{})
	//catCh := make(chan struct{})
	//fishCh := make(chan struct{})
	//endCh := make(chan struct{})
	//quitCh := make(chan struct{})
	//go dog(dogCh, catCh, quitCh)
	//go cat(catCh, fishCh, quitCh)
	//go fish(fishCh, endCh, quitCh)
	//for i := 0; i < 10; i++ {
	//	dogCh <- struct{}{}
	//	<-endCh
	//}
	//close(quitCh)

	//sender()
	//ad1 := &Ad{Name: "ad1", Score: 1.1, Ecpm: 50, Cvr: 1.11}
	//ad2 := &Ad{Name: "ad2", Score: 1.2, Ecpm: 20, Cvr: 1.12}
	//ad3 := &Ad{Name: "ad3", Score: 1.3, Ecpm: 30, Cvr: 1.13}
	//ad4 := &Ad{Name: "ad4", Score: 1.4, Ecpm: 40, Cvr: 1.14}
	//ads := make([]*Ad, 0, 10)
	//ads = append(ads, ad3)
	//ads = append(ads, ad1)
	//ads = append(ads, ad2)
	//ads = append(ads, ad4)
	//in := &In{ad: ads}
	//fmt.Printf("cap:%d, len:%d\n", cap(ads), len(ads))
	//l := len(ads)
	//ads = ads[:l]
	//fmt.Printf("cap:%d, len:%d\n", cap(ads), len(ads))
	//es := EcpmSortedAd(ads)
	//sort.Sort(es)
	//fmt.Println(ads)
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//fmt.Println(ch1==ch2)
	//d := make([]*Ad, len(ads))
	//copy(d, ads)
	//fmt.Println(d)
	//fmt.Printf("%v\n", in.ad)
	//ss := ScoreSortedAd(ads)
	//sort.Sort(ss)
	//in.ad = ss
	//fmt.Printf("%v", in.ad)
	//fmt.Println(d)
	//fmt.Println(ads)
	//
	//for i, ad := range d {
	//	fmt.Printf("%p\n", ad)
	//	fmt.Printf("%p\n", ads[i])
	//}
	//fmt.Println(hmacSha256("", "SECf23b18492b3191245c7b7c00305825a6b6353671c5f1ba59fd5858e1dad74735"))
	//gateway()
	//testAd(&Ad{})
	//sender()
	//num := 6
	//for index := 0; index < num; index++ {
	//	resp, _ := http.Get("https://www.baidu.com")
	//	fmt.Println(resp.StatusCode)
	//	_, _ = ioutil.ReadAll(resp.Body)
	//}
	//fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())
	//fmt.Println(19/100)
	//fmt.Println(19%100)
	//a := make(map[string]int, 5)
	//a["a"] = 1
	//a["b"] = 2
	//fmt.Println(a["c"])
	//fmt.Println(math.Cbrt(0.1 * 1.1 * 1))
	//ad := Ad{
	//	Score: 0.02,
	//	Name:  "hi",
	//	Label: &Label{Version: "1.0"},
	//	M:     map[uint32]uint32{1: 1},
	//	L:     []int{2},
	//}
	//fmt.Println(ad)
	//Modify(ad)
	//fmt.Println(ad)
	//fmt.Println(ad.Label)
	//ch := make(chan int)
	//go func(){
	//	time.Sleep(5*time.Second)
	//	fmt.Println("in goroutine")
	//	ch <- 1
	//}()
	//<- ch
	//fmt.Println("in main")
	//var s1 []int
	//s2 := make([]int, 0, 4)
	//s2 = append(s2, 1)
	//s1 = s2
	//fmt.Printf("%d %d\n", len(s1), cap(s1))
	//s2 = append(s2, 2)
	//fmt.Printf("%v %p\n", s1, &s1)
	//fmt.Printf("%v %p\n", s2, &s2)
	//fmt.Printf("%p %p", unsafe.Pointer(&s1[0]), unsafe.Pointer(&s2[0]))
	//var a uint = 1
	//var b uint = 2
	//fmt.Println(a-b)
}
