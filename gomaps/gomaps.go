package main

import "fmt"


var newMap map[string][]string

func ADD(key string, val []string) {
	newMap[key] = val
}

func PRESENT(key string) bool {
	_,ok:= newMap[key]
	if (ok) {
		return true
	}
	return false
}

func main() {
	newMap = make(map[string][]string)
	
	xs := []string {`Shaken, not stirred`, `Martinis`, `Women`}
	xt := []string {`James Bond`, `Literature`, `Computer Science`}
	xu := []string {`Being evil`, `Ice cream`, `Sunsets`}

	ADD("bond_james", xs)
	ADD(`moneypenny_miss`, xt)
	ADD(`no_dr`, xu)

	for k,v := range newMap {
		fmt.Println(k,v)
	}

	//Add one more record in the map
	ADD("chidam", []string{`Games`, `Fiction Novels`, `Thriller Movies`})
	for k,v := range newMap {
		fmt.Println(k,v)
	}

	//Deleting a record from a map
	fmt.Println("Deleting a record from a Map.\nAfter Deleting")
	delete(newMap, "no_dr")
	for k,v := range newMap {
		fmt.Println(k,v)
	}
}