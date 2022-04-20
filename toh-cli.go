package main

import "fmt"

type Stack struct {
	Sslice []int
}

func (st *Stack) Init() {
	st.Sslice = make([]int, 0)
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

func (pole Stack) String() string {
	var str string
	for i := len(pole.Sslice) - 1; i > -1; i-- {
		str += fmt.Sprintf("\t")
		for j := 1; j <= pole.Sslice[i]; j++ {
			str += fmt.Sprintf("_")
		}
		str += "\n"
	}
	return str
}

func main() {
	var poleA, poleB, poleC Stack
	//	poleA = Stack{}
	//	poleB = Stack{}
	//	poleC = Stack{}
	poleA.Init()
	poleB.Init()
	poleC.Init()

	// setting A as source, B as intermediary, C as target pole
	var poleNumber map[string]*Stack = make(map[string]*Stack)
	poleNumber["A"] = &poleA
	poleNumber["B"] = &poleB
	poleNumber["C"] = &poleC
	//	fmt.Printf("%T %+v", poleNumber, poleNumber)

	var noOfDisks int
	fmt.Scanf("%d", &noOfDisks)
	//TODO	validation of number of disks 3 <= n <= 20

	//initailly adding disks to pole A
	for i := noOfDisks; i > 0; i-- {
		poleA.Add(i)
	}
	//	fmt.Println(poleA.Sslice, "top :", poleA.Last())

	var inputSource, inputDest string
	var pickedDisk int
	for true {
		for k, _ := range poleNumber {
			fmt.Println(k)
			fmt.Println(poleNumber[k])
		}
		fmt.Scanf("enter \"from to\" : %s %s", &inputSource, &inputDest)
		if !checkValid(*poleNumber[inputSource], *poleNumber[inputDest]) {
			fmt.Println("Invalid move")
			continue
		}
		//		_ = pickedDisk
		pickedDisk = (*poleNumber[inputSource]).Extract()
		(*poleNumber[inputDest]).Add(pickedDisk)

	}
}
