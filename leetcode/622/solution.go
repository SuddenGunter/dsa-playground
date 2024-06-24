package solution

type MyCircularQueue struct {
	capacity, writeIndex, readIndex int
	storage                         []int32
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		storage:  make([]int32, k),
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.capacity > 0 {
		this.capacity--
		this.storage[this.writeIndex] = int32(value)
		this.writeIndex++
		if this.writeIndex >= len(this.storage) {
			this.writeIndex = 0
		}
		return true
	}

	return false
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.capacity == len(this.storage) {
		return false
	}

	this.readIndex++
	this.capacity++
	if this.readIndex >= len(this.storage) {
		this.readIndex = 0
	}

	return true
}

func (this *MyCircularQueue) Front() int {
	if this.capacity == len(this.storage) {
		return -1
	}

	return int(this.storage[this.readIndex])
}

func (this *MyCircularQueue) Rear() int {
	if this.capacity == len(this.storage) {
		return -1
	}

	if this.writeIndex == 0 {
		return int(this.storage[len(this.storage)-1]) // inserted last el and went to start
	}

	return int(this.storage[this.writeIndex-1])
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.capacity == len(this.storage)
}

func (this *MyCircularQueue) IsFull() bool {
	return this.capacity <= 0
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
