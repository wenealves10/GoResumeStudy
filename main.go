package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World")
	var nome string = "wene"

	lastname := "Oliveira"
	var arraysNames [4]string = [4]string{"name", "name12", "name324", "name234"}
	var sliceName []string = []string{"Wene", "Ismael", "Lucas", "Ariel"}

	fmt.Println(sliceName)
	fmt.Println(arraysNames)

	fmt.Println(nome, lastname)

	if lastname == "Oliveira" {
		fmt.Println("Verdadeiro")
	} else if lastname == "Juvenal" {
		fmt.Println("Falsão")
	} else {
		fmt.Println("Só um merda mesmo!")
	}

	if test := test(); test == "A" {
		fmt.Println("Otário a lerta é A")
	} else {
		fmt.Println("Você acerto kkkk")
	}

	switch nome {
	case "wene":
		fmt.Println("Sortudo, Programador fod*")
		fallthrough
	case "ismael":
		fmt.Println("Merda, Programador lixo!!!")
		break
	case "lucas":
		fmt.Println("Entre os dois kkkk, Programador meia boca")
	default:
		fmt.Println("Lixo!!, não é nada na vida!!")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// var testtwo bool = false

	// for testtwo == false {
	// 	fmt.Printf("%v", testtwo)
	// }

	testThree := 0

	for testThree <= 100 {
		fmt.Println("Three", testThree)

		testThree++
	}

	anExpression := false

	for ok := true; ok; ok = anExpression {
		fmt.Println("Execute Once")
	}

	for index, value := range sliceName {
		fmt.Println(index, value)
	}

	for index, value := range arraysNames {
		fmt.Println(index, value)
	}

	fmt.Println("sliceName", cap(sliceName), len(sliceName))
	fmt.Println("ArrayName", cap(arraysNames), len(arraysNames))

	soma := testFor(10, 30)

	fmt.Println(soma)

	err, value := testMult(1000)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(value)

	errorOne, addedValue := testReturnNamedMult(100, 200)

	if errorOne != nil {
		log.Panic(errorOne)
	}

	fmt.Println(addedValue)

	nameComplete := testParameterReturnFunc("Wene", "Alves")

	nameCompleteValue := nameComplete()

	fmt.Println(nameCompleteValue)

	function := func(test string, name string) {
		fmt.Println(test, name)
	}

	testParameterFunc(function)

	functionOne := func(test string) string {
		return strings.ToUpper(test)
	}

	nameCompleteOne := testParameterFuncAndReturnFunc(functionOne)

	nameCompleteOne("Oliveira")

	nameLowerCase := funcAnonymous()

	fmt.Println(nameLowerCase)

	funcMultParameters("w", "e", "n", "e")

	funcAnonymousOne := func() {
		fmt.Println("Hello")
	}

	funcAnonymousTwo := func() {
		fmt.Println("World")
	}

	funcMultParametersFuncs(funcAnonymousOne, funcAnonymousTwo)

	var texts string = "Wene Alves"

	funcAnonymousPointer(&texts)

	fmt.Println(texts)

	pointer := 1000

	var referencePointer *int = &pointer

	fmt.Println(referencePointer)
	fmt.Println(&referencePointer)
	fmt.Println(*referencePointer)

	*referencePointer = 100

	fmt.Println(*referencePointer, pointer)

	var referencePointerTwo *int = referencePointer

	*referencePointerTwo = 10000

	fmt.Println(referencePointer, referencePointerTwo, &pointer)

	*referencePointerTwo = 1000 * 10

	fmt.Println(pointer, *referencePointer)

	go testSleep("Wene Alves")
	go testSleep("Wene Sleep Goroutines")
	go func() {
		testSleep("Goroutines Anonymous")
	}()
	time.Sleep(10 * time.Second)

	textTwo := "Text Pointer"

	var textPointer *string = &textTwo

	go testSleepPointer(textPointer)

	time.Sleep(50 * time.Millisecond)

	*textPointer = "Text Pointer New"

	time.Sleep(100 * time.Millisecond)

}

func test() string {
	return "A"
}

func testFor(inter int, inter2 int) int {
	return inter + inter2
}

func testMult(value int) (error, int) {
	if value > 1000 {
		return errors.New("Value muito maior que 1000"), 0
	} else {
		return nil, value + 10
	}
}

func testReturnNamedMult(valueA int, valueB int) (somaError error, soma int) {
	if valueA > valueB {
		somaError = errors.New("valueA is more than valueB")
		soma = 0

	} else {
		somaError = nil
		soma = valueA + valueB
	}

	return
}

func testParameterReturnFunc(valueA string, valueB string) func() string {
	return func() string {
		return valueA + valueB
	}
}

func testParameterFunc(valueFunc func(string, string)) {
	valueFunc("Wene", "Alves")
}

func testParameterFuncAndReturnFunc(valueFunc func(string) string) func(string) {
	function := func(lastname string) {
		fmt.Println(valueFunc("Wene"), lastname)
	}

	return function
}

func funcAnonymous() (nameComplete string) {

	func(lastname string, firstname string) {
		nameComplete = strings.ToLower(lastname + firstname)
	}("Wene", "Alves")

	return

}

func funcMultParameters(value ...string) {
	for _, x := range value {
		fmt.Println(x)
	}
}

func funcMultParametersFuncs(funcs ...func()) {
	for _, x := range funcs {
		x()
	}
}

func funcAnonymousPointer(text *string) {
	*text = "Wene Alves de Oliveira"
}

func testSleep(text string) {
	for index := 0; index < 10; index++ {
		fmt.Printf("%d -- %s\n", index, text)
		time.Sleep(1 * time.Second)
	}
}

func testSleepPointer(text *string) {
	for index := 0; index < 100; index++ {
		fmt.Printf("%d ----- %s\n", index, *text)
		time.Sleep(1 * time.Millisecond)
	}
}
