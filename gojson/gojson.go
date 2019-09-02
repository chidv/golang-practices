package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	bs, err := json.Marshal(users)
	if err == nil {
		fmt.Println(bs)
	}

	var usrs []user
	//Unmarshal
	err = json.Unmarshal(bs, &usrs)
	if err == nil {
		fmt.Println(usrs)
	}


	//Marshal directly to os.Stdout
	err = json.NewEncoder(os.Stdout).Encode(users)
}
