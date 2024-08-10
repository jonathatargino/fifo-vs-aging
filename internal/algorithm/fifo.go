package algorithm

const queueLength = 10

type Fifo struct{}

func contains(array []int, element int) bool {
	for _, el := range array {
		if element == el {
			return true
		}
	}

	return false
}

func (f Fifo) Run(pageReferences []int, framesNumber int) int {
	var queue [queueLength]int
	pageFaults := 0
	initialQueuePosition := 0

	for index, pageReference := range pageReferences {
		if index <= queueLength {
			queue[index] = pageReference
			pageFaults++
		}

		isPageinQueue := contains(queue[:], pageReference)

		if !isPageinQueue {
			queue[initialQueuePosition] = pageReference
			initialQueuePosition = (initialQueuePosition + 1) % queueLength

			pageFaults++
		}

	}

	return pageFaults
}
