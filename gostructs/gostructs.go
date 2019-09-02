package main

import "fmt"

type person struct {
	last string
	first string
	age	int
}


type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheeler bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	var p1 person = person{last:"singh", first:"kabir", age:32}
	var p2 person = person{last:"valliappan", first:"chidambaram", age:30}

	//Slice of structs
	var persons []person=[]person{p1,p2}

	for _, p := range persons {
		fmt.Println(p.last, p.first, p.age)
	}

	//Map of structs
	fmt.Println("Map of Structs")
	pMap := make(map[string]person)
	pMap[p1.last] = p1
	pMap[p2.last] = p2

	for k,v := range pMap {
		fmt.Println(k,v)
	}

	//Embedding structs and promotions
	var s sedan = sedan {vehicle: vehicle{doors: 4, color: "red"}, luxury:false}
	var t truck = truck {vehicle: vehicle{doors:6, color: "black"}, fourWheeler: true}

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(s.doors)
	fmt.Println(t.doors)

	//Anonymous struct
	as := struct {
		name string
		age int
		place string
	} {name: "chidam", age: 30, place: "cbe"}

	fmt.Println(as.name, as.age, as.place)

}