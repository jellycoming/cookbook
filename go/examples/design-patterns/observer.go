// 观察者模式，又被称为发布-订阅（Publish/Subscribe）模式，属于行为型模式的一种。
// 在此种模式中，一个目标对象管理所有相依于它的观察者对象，并且在它本身的状态改变时主动发出通知。这通常通过调用各观察者所提供的方法来实现。
// 此种模式通常被用来实时事件处理系统。
package design_patterns

import (
	"fmt"
	"time"
)

// 消息体
type Message struct {
	content string
}

// 观察者接口
type Observer interface {
	Handle(msg *Message)
}

// 主题接口（被观察者接口）
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(msg *Message)
}

// 主题具体实现（具体被观察者）
type LogSubject struct {
	Observers map[Observer]bool
}

func (s *LogSubject) Attach(ob Observer) {
	s.Observers[ob] = true
}

func (s *LogSubject) Detach(ob Observer) {
	delete(s.Observers, ob)
}

func (s *LogSubject) Notify(msg *Message) {
	for ob := range s.Observers {
		ob.Handle(msg)
	}
}

func (s *LogSubject) Subscribe(c chan *Message) {
	for {
		msg := <-c
		s.Notify(msg)
	}
}

// 具体观察者A
type ObserverA struct {
}

func (a *ObserverA) Handle(msg *Message) {
	fmt.Printf("observer A handle msg: %s\n", msg.content)
}

// 具体观察者B
type ObserverB struct {
}

func (b *ObserverB) Handle(msg *Message) {
	fmt.Printf("observer B handle msg: %s\n", msg.content)
}

func NewLogSubject() *LogSubject {
	return &LogSubject{Observers: make(map[Observer]bool)}
}

func main() {
	ch := make(chan *Message)
	subject := NewLogSubject()
	go subject.Subscribe(ch)
	subject.Attach(&ObserverA{})
	subject.Attach(&ObserverB{})
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			ch <- &Message{content: "hello"}
		default:
		}
	}
}
