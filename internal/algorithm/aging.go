package algorithm

type Aging struct {
	age              map[int]uint8
	pageFaultCounter int
	pagesAlive       int
}

func NewAging() *Aging {
	return &Aging{age: make(map[int]uint8)}
}

func (a *Aging) Run(pageReferences []int, framesNumber int) int {
	for _, page := range pageReferences {
		// It is full and Page is not mapped
		if !a.isAlive(page) && framesNumber <= a.pagesAlive {
			a.pageFaultCounter++
			a.removeOlder()
			a.setNewPageInAge(page)
			a.clock()
			continue
		}

		// It isn't full and Page is not mapped
		if !a.isAlive(page) && framesNumber > a.pagesAlive {
			a.pageFaultCounter++
			a.setNewPageInAge(page)
			a.clock()
			continue
		}

		// Page is mapped
		a.addInAge(page)
		a.clock()
	}

	return a.pageFaultCounter
}

func (a *Aging) isAlive(page int) bool {
	_, ok := a.age[page]
	return ok
}

func (a *Aging) removeOlder() {
	var older uint8 = 255
	var pageTarget int

	for page, age := range a.age {
		if age < older {
			pageTarget = page
			older = age
		}
	}

	delete(a.age, pageTarget)
}

func (a *Aging) addInAge(page int) {
	age := a.age[page]
	age = age | 128

	a.age[page] = age
}

func (a *Aging) setNewPageInAge(page int) {
	a.pagesAlive++
	a.age[page] = 128
}

func (a *Aging) clock() {
	for page, age := range a.age {
		a.age[page] = age >> 1
	}
}
