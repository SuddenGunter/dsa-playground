package main

type FixedCapacityStackOfStrings struct {
	storage []string
	pos     int
}

func NewFixedCapacityStackOfStrings(num int) *FixedCapacityStackOfStrings {
	return &FixedCapacityStackOfStrings{
		storage: make([]string, num),
		pos:     0,
	}
}

func (fs *FixedCapacityStackOfStrings) Push(val string) {
	fs.storage[fs.pos] = val
	fs.pos++
}

func (fs *FixedCapacityStackOfStrings) Pop() string {
	fs.pos--
	return fs.storage[fs.pos]
}

func (fs *FixedCapacityStackOfStrings) Empty() bool {
	return fs.pos == 0
}

func (fs *FixedCapacityStackOfStrings) Size() int {
	return fs.pos
}

func (fs *FixedCapacityStackOfStrings) IsFull() bool {
	return fs.pos == len(fs.storage)
}
