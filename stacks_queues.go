package main

import "fmt"

/*
				THE BASIC CONCEPT OF STACK AND QUEUE
	---- POP
	----PUSH
	----PEEK
	----ISEMPTY
	----ENQUEUE
	----DEQUEUE
*/

type Stack struct {
	item []int
}
type Queue struct {
	item []int
}

func (s *Stack) pop() int {
	popItem := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return popItem
}
func (s *Stack) push(n int) {
	s.item = append(s.item, n)
}
func (s *Stack) peek() int {
	peekItem := s.item[len(s.item)-1]
	return peekItem
}
func (s *Stack) isEmpty() bool {
	return len(s.item) == 0
}
func (q *Queue) enQueue(val int) {
	q.item = append(q.item, val)
}
func (q *Queue) deQueue() {
	q.item = q.item[1:]
}

func main() {
	stackInstance := Stack{}
	stackInstance.push(1)
	stackInstance.push(2)
	stackInstance.push(3)
	stackInstance.push(4)
	stackInstance.push(5)
	fmt.Println(stackInstance)
	fmt.Println(stackInstance.pop())
	fmt.Println(stackInstance.peek())
	fmt.Println(stackInstance.isEmpty())
	fmt.Println(stackInstance.pop())
	fmt.Println(stackInstance.peek())
	queueInstance := Queue{}
	queueInstance.enQueue(10)
	queueInstance.enQueue(20)
	queueInstance.enQueue(30)
	queueInstance.enQueue(40)
	queueInstance.enQueue(50)
	fmt.Println(queueInstance)
	queueInstance.deQueue()
	queueInstance.deQueue()
	fmt.Println(queueInstance)

}
