package mergesort

import (
	"sync"

	"github.com/SuddenGunter/dsa-playground/skiena/4/queue"
)

var maxDepth = int(4)

func ConcurrentSort(src []float64) []float64 {
	if len(src) == 0 {
		return src
	}

	resultsListener := make(chan float64, maxDepth)
	results := make([]float64, 0, cap(src))
	go concurrentMergeSort(src, 1, resultsListener)
	for v := range resultsListener {
		results = append(results, v)
	}

	return results
}

func concurrentMergeSort(src []float64, depth int, results chan float64) {
	if depth >= maxDepth {
		for _, v := range Sort(src) {
			results <- v
		}
		close(results)

		return
	}

	if len(src) > 1 {
		middle := len(src) / 2
		leftResults := make(chan float64, len(src)/2)
		rightResults := make(chan float64, len(src)/2)

		go concurrentMergeSort(src[0:middle], depth+1, leftResults)
		go concurrentMergeSort(src[middle:], depth+1, rightResults)

		for v := range concurrentMerge(leftResults, rightResults) {
			results <- v
		}
		close(results)
	} else {
		results <- src[0]
		close(results)
	}
}

func concurrentMerge(left, right chan float64) chan float64 {
	res := make(chan float64, maxDepth)
	leftBuf := queue.NewQueue()
	rightBuf := queue.NewQueue()

	wg := &sync.WaitGroup{}
	listenToTheEnd := func(ch chan float64, buf *queue.Queue, wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range ch {
			buf.Enqueue(queue.Element{Data: v})
		}
	}

	wg.Add(1)
	go listenToTheEnd(left, leftBuf, wg)
	wg.Add(1)
	go listenToTheEnd(right, rightBuf, wg)

	go func() {
		wg.Wait()

		// dequeue until one of the queues is empty
		for !leftBuf.IsEmpty() && !rightBuf.IsEmpty() {
			leftElement, rightElement := leftBuf.Peek(), rightBuf.Peek()
			if leftElement.Data <= rightElement.Data {
				res <- leftBuf.Dequeue().Data
			} else {
				res <- rightBuf.Dequeue().Data
			}
		}

		// flush the other (not-empty) queue
		for !leftBuf.IsEmpty() {
			res <- leftBuf.Dequeue().Data
		}
		for !rightBuf.IsEmpty() {
			res <- rightBuf.Dequeue().Data
		}

		close(res)
	}()

	return res
}
