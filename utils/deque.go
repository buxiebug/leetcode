package utils
import "fmt"

type DequeNode struct {
	next *DequeNode
	pre  *DequeNode
	Val  interface{}
}

type Deque struct {
	head *DequeNode
	tail *DequeNode
	size int
}

func (d *Deque) Peek() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.head.Val
}

func (d *Deque) PeekLast() interface{} {
	if d.GetSize() == 0 {
		return nil
	}
	return d.tail.Val
}

func (d *Deque) GetSize() int {
	return d.size
}

func (d *Deque) Poll() *DequeNode {
	if d.GetSize() == 0 {
		return nil
	}
	rs := d.head
	next := d.head.next
	if next != nil {
		d.head = next
		next.pre = nil
	} else {
		d.head = nil
		d.tail = nil
	}
	d.size--
	return rs
}

func (d *Deque) PollLast() *DequeNode {
	if d.GetSize() == 0 {
		return nil
	}
	rs := d.tail
	pre := d.tail.pre
	if pre != nil {
		d.tail = pre
		pre.next = nil
	} else {
		d.head = nil
		d.tail = nil
	}
	d.size--
	return rs
}

func (d *Deque) AddLast(val interface{}) {
	node := &DequeNode{
		Val: val,
	}
	if d.GetSize() == 0 {
		d.head = node
		d.tail = node
		d.size++
		return
	}
	node.pre = d.tail
	d.tail.next = node
	d.tail = node
	d.size++
}


func TestDeque() {
	q := Deque{}
	fmt.Println(q.GetSize())
	q.AddLast(1)
	fmt.Println(q.GetSize())
	fmt.Println(q.Peek())
	fmt.Println(q.PeekLast())
	q.Poll()
	fmt.Println(q.GetSize())
	q.AddLast(2)
	q.AddLast(3)
	fmt.Println(q.GetSize())
	fmt.Println(q.Peek())
	fmt.Println(q.PeekLast())
}