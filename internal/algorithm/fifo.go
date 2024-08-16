package algorithm

type Fifo struct{}

func NewFifo() *Fifo {
	return &Fifo{}
}

func contains(array []int, element int) bool {
	for _, el := range array {
		if element == el {
			return true
		}
	}

	return false
}

func (f *Fifo) Run(pageReferences []int, framesNumber int) int {
	queue := make([]int, 0)
	pageFaults := 0
	initialQueuePosition := 0

	for index, pageReference := range pageReferences {
		if index < framesNumber {
			queue = append(queue, pageReference)
			pageFaults++
			continue
		}

		isPageinQueue := contains(queue[:], pageReference)

		if !isPageinQueue {
			queue[initialQueuePosition] = pageReference
			initialQueuePosition = (initialQueuePosition + 1) % framesNumber

			pageFaults++
		}

	}

	return pageFaults
}
