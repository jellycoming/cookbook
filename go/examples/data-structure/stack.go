package data_structure

import "sync"

// 栈
type Stack struct {
	lock  sync.Mutex
	nodes []*Node
}

func NewStack() *Stack {
	s := new(Stack)
	s.nodes = make([]*Node, 0)
	return s
}

// 入栈
func (s *Stack) Push(node *Node) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.nodes = append(s.nodes, node)
}

// 弹出栈顶元素，当栈为空时返回nil
func (s *Stack) Pop() *Node {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.nodes) == 0 {
		return nil
	}
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0 : len(s.nodes)-1]
	return node
}

// 读取栈顶元素但不弹出，当栈为空时返回nil
func (s *Stack) Top() *Node {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.nodes) == 0 {
		return nil
	}
	return s.nodes[len(s.nodes)-1]
}

// 当前栈大小
func (s *Stack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.nodes)
}

// 清除栈中所有元素，变为空栈
func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.nodes = make([]*Node, 0)
}

// 判断栈是否为空
func (s *Stack) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.nodes) == 0
}
