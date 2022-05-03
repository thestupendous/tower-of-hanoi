package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

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
	lastIndex := len(st.Sslice) - 1
	toReturn := st.Sslice[lastIndex]
	st.Sslice = st.Sslice[:lastIndex]
	return toReturn
}

func (st Stack) Last() int {
	//for empty stack, returning -1
	if len(st.Sslice) == 0 {
		return -1
	}
	lastIndex := len(st.Sslice) - 1
	return st.Sslice[lastIndex]
}

func checkValid(fromPole, toPole Stack) bool {
	if toPole.Last() < 0 { //meaning, any sized disk can be placed on empty pole
		return true
	}
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
			str += fmt.Sprintf("*")
		}
		str += "\n"
	}
	return str
}

func main() {
	logFile, _ := os.OpenFile("output-of-toh-cli.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetFlags(0) //because the timestamp prompt messes up cli output screen
	log.SetOutput(mw)
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
	//	log.Printf("%T %+v", poleNumber, poleNumber)

	var noOfDisks int
	fmt.Scanf("%d", &noOfDisks)
	//TODO	validation of number of disks 3 <= n <= 20

	//initailly adding disks to pole A
	for i := noOfDisks; i > 0; i-- {
		poleA.Add(i)
	}
	log.Println("starting game, ", noOfDisks, " added to pole A")
	//	log.Println(poleA.Sslice, "top :", poleA.Last())

	var inputSource, inputDest string
	var movesCounter, pickedDisk int
	won := false
	for true {
		for k, _ := range poleNumber {
			log.Println(k)
			log.Println(poleNumber[k])
		}
		fmt.Printf("enter \"from to\" : ")
		fmt.Scanf("%s %s", &inputSource, &inputDest)
		movesCounter++
		if !checkValid(*poleNumber[inputSource], *poleNumber[inputDest]) {
			log.Println("Invalid move")
			continue
		}
		//		_ = pickedDisk
		pickedDisk = (*poleNumber[inputSource]).Extract()
		(*poleNumber[inputDest]).Add(pickedDisk)

		//if user has won, meaning all disks at pole C
		if len(poleC.Sslice) == noOfDisks {
			won = true
			break
		}

	}
	if won {
		for k, _ := range poleNumber {
			log.Println(k)
			log.Println(poleNumber[k])
		}
		log.Println("\t\tHURRAY!! YOU'VE WON , you used ", movesCounter, " moves")
	}

}
