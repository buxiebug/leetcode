package cases

import "fmt"

type LRUCache struct {
	capacity int
	vMap map[int]*Node
	head *Node
	end *Node
}

type Node struct {
	val int
	key int
	next *Node
	pre *Node
}

func ConstructorLRU(capacity int) LRUCache{
	head := &Node{
		key: -1,
		val: -1,
		next: nil,
		pre: nil,
	}
	end := &Node{
		key: -1,
		val: -1,
		next: nil,
		pre: nil,
	}
	head.next = end
	end.pre = head
	return LRUCache{
		head: head,
		end: end,
		capacity: capacity,
		vMap:make(map[int]*Node),
	}

}

func (this *LRUCache) Get(key int) int {
	v, ok := this.vMap[key]
	if ok {
		if v.pre == this.head {
			return v.val
		}
		v.pre.next = v.next
		v.next.pre = v.pre
		v.pre = this.head
		v.next = this.head.next
		this.head.next.pre = v
		this.head.next = v
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, val int) {
	node, ok := this.vMap[key]
	if !ok {
		node := &Node{
			key: key,
			val: val,
		}
		if len(this.vMap) >= this.capacity{
			deleteNode := this.end.pre
			deleteNode.pre.next = this.end
			this.end.pre = deleteNode.pre
			delete(this.vMap, deleteNode.key)
		}
		if this.head.next != this.end {
			firstNode := this.head.next
			firstNode.pre = node
			node.next = firstNode
			this.head.next = node
			node.pre = this.head
		} else {
			this.head.next = node
			this.end.pre = node
			node.next = this.end
			node.pre = this.head
		}

		this.vMap[key] = node
		return
	}
	node.val = val
	if node.pre == this.head {
		return
	}
	node.next.pre = node.pre
	node.pre.next = node.next

	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
	node.pre = this.head
}

func TestLRUCache146() {
	lru := ConstructorLRU(10)
	lru.Put(10, 13)
	lru.Put(3, 17)
	lru.Put(6, 11)
	lru.Put(10, 5)
	lru.Put(9, 10)
	lru.Get(13)
	lru.Put(2, 9)
	lru.Get(2)
	lru.Get(3)
	lru.Put(5, 25)
	lru.Get(8)
	lru.Put(9, 22)
	lru.Put(5, 5)
	lru.Put(1, 30)
	lru.Get(11)
	lru.Put(9, 12)
	lru.Get(7)
	lru.Get(5)
	lru.Get(8)
	lru.Get(9)
	lru.Put(4, 30)
	lru.Put(9, 3)
	lru.Get(9)
	lru.Get(10)
	lru.Get(10)
	lru.Put(6, 14)
	lru.Put(3, 1)
	lru.Get(3)
	lru.Put(10, 11)
	lru.Get(8)
	lru.Put(2, 14)
	lru.Get(1)
	lru.Get(5)
	lru.Get(4)
	lru.Put(11, 4)
	lru.Put(12, 24)
	lru.Put(5, 18)
	lru.Get(13)
	lru.Put(7, 23)
	lru.Get(8)
	lru.Get(12)
	lru.Put(3, 27)
	lru.Put(2, 12)
	lru.Get(5)
	lru.Put(2, 9)
	lru.Put(13, 4)
	lru.Put(8, 18)
	lru.Put(1, 7)
	fmt.Println(lru.Get(6))

}