package main

import (
	"fmt"
)

// 策略模式
// 定义一系列的算法，把每一个算法封装起来，并且使它们可相互替换，以使某个对象的某个行为在不同的场景中有不同的实现算法。

type TensorProto struct {
}

// 策略接口
type VectorizationStrategy interface {
	Inputs(in string) (map[string]*TensorProto, error)
}

// 定义模型结构体实现策略接口
type DeepNNStrategy struct {
}

func (dnn *DeepNNStrategy) Inputs(in string) (map[string]*TensorProto, error) {
	return nil, nil
}

type DeepFMStrategy struct {
}

func (dfm *DeepFMStrategy) Inputs(in string) (map[string]*TensorProto, error) {
	return nil, nil
}

// 定义VectorizationContext结构体，用来做策略筛选
type VectorizationContext struct {
	strategy VectorizationStrategy
}

func NewVectorizationContext(sign string) *VectorizationContext {
	vc := new(VectorizationContext)
	switch sign {
	case "DeepNN":
		vc.strategy = &DeepNNStrategy{}
	case "DeepFM":
		vc.strategy = &DeepFMStrategy{}
	}
	return vc
}

func (vc *VectorizationContext) Vectorize(in string) (map[string]*TensorProto, error) {
	return vc.strategy.Inputs(in)
}

// 使用策略模式
func main() {
	vc := NewVectorizationContext("DeepNN")
	fmt.Println(vc.Vectorize("raw request"))

	vc = NewVectorizationContext("DeepFM")
	fmt.Println(vc.Vectorize("raw request"))
}
