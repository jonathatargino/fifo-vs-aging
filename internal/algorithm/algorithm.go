package algorithm

type Algorithm interface {
	Run(pageReferences []int, framesNumber int) int
}
