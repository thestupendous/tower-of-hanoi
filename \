package main

import "fmt"

type Stack struct {
	Sslice []int
}

func (st *Stack) Add(newDisk int) {
	st.Sslice = append(st.Sslice, newDisk)
}

func (st *Stack) Extract() int {
	lastIndex := len(st.Sslice)
	toReturn := st.Sslice[lastIndex]
	st.Sslice = st.Sslice[:lastIndex]
	return toReturn
}

func (st Stack) Last() int {
	lastIndex := len(st.Sslice)
	return st.Sslice[lastIndex]
}

func checkValid(fromPole, toPole Stack) bool {
	currentDisk := fromPole.Last()
	onDisk := toPole.Last()
	if currentDisk < onDisk {
		return true
	} else {
		return false
	}
}

func main() {
	var poleA, poleB, poleC Stack
	// setting A as source, B as intermediary, C as target pole
	var poleNumber map[string]Stack = make(map[string]Stack)
	poleNumber["A"] = poleA
	poleNumber["B"] = poleB
	poleNumber["C"] = poleC

	var noOfDisks int
	fmt.Scanf("%d", &noOfDisks)
	//TODO	validation of number of disks 3 <= n <= 20

	for i := noOfDisks; i > 0; i-- {
		poleA.Add(i)
	}
	var inputSource, inputDest int
	for true {
		fmt.Scanf("enter \"from to\" : %d %d", &inputSource, &inputDest)
		if !checkValid(poleNumber(inputSource), poleNumber(inputDest)) {
			fmt.Println("Invalid move")
			continue
		}

	}
}
