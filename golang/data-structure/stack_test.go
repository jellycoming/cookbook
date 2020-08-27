package data_structure

import (
	"os"
	"testing"
)

var s *Stack

//func init() {
//	if s == nil {
//		s = NewStack()
//	}
//}

func TestStackPush(t *testing.T) {
	var tests = []struct {
		f    func(node *Node)
		arg  *Node
		size int
	}{
		{s.Push, &Node{Value: 100}, 1},
		{s.Push, &Node{Value: 200}, 2},
		{s.Push, &Node{Value: 300}, 3},
		{s.Push, &Node{Value: 400}, 4},
		{s.Push, &Node{Value: 500}, 5},
	}
	for _, test := range tests {
		test.f(test.arg)
		if s.Size() != test.size {
			t.Errorf("stack.size=%d, want %d", s.Size(), test.size)
		}
	}
}

func TestStackPop(t *testing.T) {
	var tests = []struct {
		f    func() *Node
		size int
	}{
		{s.Pop, s.Size() - 1},
		{s.Pop, s.Size() - 2},
		{s.Pop, s.Size() - 3},
		{s.Pop, s.Size() - 4},
		{s.Pop, s.Size() - 5},
		{s.Pop, 0},
	}
	for _, test := range tests {
		test.f()
		if s.Size() != test.size {
			t.Errorf("stack.size=%d, want %d", s.Size(), test.size)
		}
	}
}

func TestStackTop(t *testing.T) {
	node := s.Top()
	if node != nil {
		t.Errorf("empty stack.top=%v, want nil", node)
	}
	s.Push(&Node{Value: 1})
	s.Push(&Node{Value: 2})
	node = s.Top()
	if s.Size() != 2 {
		t.Errorf("stack.size=%d, want %d", s.Size(), 2)
	}
	if node.Value != 2 {
		t.Errorf("top node.value=%d, want %d", node.Value, 2)
	}
}

func TestStackClear(t *testing.T) {
	s.Clear()
	if !s.Empty() {
		t.Errorf("stack.size=%d, want empty", s.Size())
	}
}

func TestMain(m *testing.M) {
	if s == nil {
		s = NewStack()
	}
	os.Exit(m.Run())
}
