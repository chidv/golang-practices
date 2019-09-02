package main

import "fmt"

func main() {
	slic := []int{1,2,3,4,5,6,7,8,9,10}

	for _, v := range slic {
		fmt.Println(v)
	}

	fmt.Printf("%#T\n", slic)

	//Slicing the slice
	slic1 := slic[0:5]
	slic2 := slic[1:6]
	slic3 := slic[2:7]

	fmt.Println(slic1)
	fmt.Println(slic2)
	fmt.Println(slic3)

	//Appending Slices
	org := []int {42,43,44,45,46,47,48,49,50,51}
	org = append(org, 52)
	fmt.Println(org)

	org = append(org,53,54,55)
	fmt.Println(org)

	newslic := []int {56,57,58,59,60}
	org = append(org,newslic...)
	fmt.Println(org)


	//Deleting the elements in slice using append
	//42,43,44,45,46,47,48,49,50,51 52 53 54 55 56 57 58 59 60
	//New  - 42,43,44,48,49,50,51
	org = append(org[:3], org[6:10]...)
	fmt.Println("After removing elements\n", org)


	//Making Slices
	made := make([]string, 50,50)
	made = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,}

	fmt.Println("Newly Made Slice\n",made)
	fmt.Println(len(made))
	fmt.Println(cap(made))

	for i:= 0; i < len(made); i++ {
		fmt.Println(made[i])
	}

	//Slice of Slice
	fmt.Println("Slice of Slices")
	slicsqr := [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, slic1:= range slicsqr {
		fmt.Println("Record :",i+1)
		for _, slic2 := range slic1 {
			fmt.Printf("\t %v\n",slic2)
		}
	}
}