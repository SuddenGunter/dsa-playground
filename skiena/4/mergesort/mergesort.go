package mergesort

import "github.com/SuddenGunter/dsa-playground/skiena/4/queue"

func Sort(src []float64) []float64 {
	result := make([]float64, len(src))

	copy(result, src)

	mergeSort(result)

	return result
}

func mergeSort(src []float64) {
	if len(src) > 1 {
		middle := len(src) / 2

		mergeSort(src[0:middle])
		mergeSort(src[middle:])

		merge(src)
	}
}

func merge(src []float64) {
	middle := len(src) / 2

	left := queue.NewQueue()
	for i := 0; i < middle; i++ {
		left.Enqueue(queue.Element{
			Data: src[i],
		})
	}

	right := queue.NewQueue()
	for i := middle; i < len(src); i++ {
		right.Enqueue(queue.Element{
			Data: src[i],
		})
	}

	counter := 0

	// dequeue until one of the queues is empty
	for !left.IsEmpty() && !right.IsEmpty() {
		leftElement, rightElement := left.Peek(), right.Peek()
		if leftElement.Data <= rightElement.Data {
			src[counter] = left.Dequeue().Data
		} else {
			src[counter] = right.Dequeue().Data
		}

		counter++
	}

	// flush the other (not-empty) queue
	for !left.IsEmpty() {
		src[counter] = left.Dequeue().Data
		counter++
	}
	for !right.IsEmpty() {
		src[counter] = right.Dequeue().Data
		counter++
	}
}
