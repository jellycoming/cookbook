package main

import (
	"fmt"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
)

// 策略模式
// 定义一系列的算法，把每一个算法封装起来，并且使它们可相互替换，以使某个对象的某个行为在不同的场景中有不同的实现算法。

// 在CTR&CVR在线预估场景中，需要对请求的原始特征数据进行向量化处理，返回 map[string]*framework.TensorProto 作为模型服务的输入。
// 而不同的模型向量化的过程与规则是有区别的，各个模型的向量化方法就可以封装成一个个策略，根据请求中的特定字段选择不同的策略来生成向量化的输入数据。

// 策略接口
type VectorizationStrategy interface {
	Inputs(in string) (map[string]*framework.TensorProto, error)
}

// 定义模型结构体实现策略接口
type DeepNNStrategy struct {
}

func (dnn *DeepNNStrategy) Inputs(in string) (map[string]*framework.TensorProto, error) {
	return nil, nil
}

type DeepFMStrategy struct {
}

func (dfm *DeepFMStrategy) Inputs(in string) (map[string]*framework.TensorProto, error) {
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

func (vc *VectorizationContext) Vectorize(in string) (map[string]*framework.TensorProto, error) {
	return vc.strategy.Inputs(in)
}

// 使用策略模式
func main() {
	vc := NewVectorizationContext("DeepNN")
	fmt.Println(vc.Vectorize("raw request"))

	vc = NewVectorizationContext("DeepFM")
	fmt.Println(vc.Vectorize("raw request"))
}
