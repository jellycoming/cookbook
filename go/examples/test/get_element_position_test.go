package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//go test -bench=. -run=^a -v
//goos: darwin
//goarch: arm64
//pkg: cookbook/go/test
//BenchmarkGetElemPositionInSlice
//BenchmarkGetElemPositionInSlice-8       28263322               901.4 ns/op
//BenchmarkGetElemPositionInMap
//BenchmarkGetElemPositionInMap-8         160679486                8.399 ns/op

func BenchmarkGetElemPositionInSlice(b *testing.B) {
	t := NewT()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getElemPositionInSlice(t.Target, t.S)
	}
}

func BenchmarkGetElemPositionInMap(b *testing.B) {
	t := NewT()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getElemPositionInMap(t.Target, t.M)
	}
}

type T struct {
	S      []string
	M      map[string]int
	Target string
}

func NewT() *T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]string, 1000)
	m := make(map[string]int)
	for i := range s {
		elem := fmt.Sprintf("abced:%d", i)
		s[i] = elem
		m[elem] = i
	}
	str := fmt.Sprintf("abced:%d", r.Intn(1000))
	return &T{
		S:      s,
		M:      m,
		Target: str,
	}
}

func getElemPositionInSlice(str string, s []string) int {
	for i, e := range s {
		if e == str {
			return i
		}
	}
	return -1
}

func getElemPositionInMap(str string, m map[string]int) int {
	if i, ok := m[str]; ok {
		return i
	} else {
		return -1
	}
}
