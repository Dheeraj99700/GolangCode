package main

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	Name      string
	Address   string
	City      string
	Mobileno  int
	Altmobile int
}

func main() {

	friend1 := Friend{
		Name:      "Dheeraj Chaudhari",
		Address:   "Shanti Nagar",
		City:      "Nashik",
		Mobileno:  8877548455,
		Altmobile: 8965448789,
	}

	friend2 := Friend{
		Name:      "Saurabh Tupe",
		Address:   "Vallabh Nagar",
		City:      "Mumbai",
		Mobileno:  9877664455,
		Altmobile: 3296548789,
	}
	friend3 := Friend{
		Name:      "Rajesh Wankhede",
		Address:   "Patil Nagar",
		City:      "Nanded",
		Mobileno:  7677664455,
		Altmobile: 8796548789,
	}
	friend4 := Friend{
		Name:      "Raj patil",
		Address:   "sahakar Nagar",
		City:      "Jalgoan",
		Mobileno:  7698764455,
		Altmobile: 8778525689,
	}
	friend5 := Friend{
		Name:      "Nachiket Chajjed",
		Address:   "Chajjed Nagar",
		City:      "Nashik",
		Mobileno:  7678932455,
		Altmobile: 8774125789,
	}
	friend6 := Friend{
		Name:      "Ashish Najan",
		Address:   "Najan Nagar",
		City:      "Nagar",
		Mobileno:  7677789222,
		Altmobile: 8796789669,
	}
	friend7 := Friend{
		Name:      "Omkar Wankhede",
		Address:   "om Nagar",
		City:      "Pune",
		Mobileno:  7678965655,
		Altmobile: 8791245689,
	}
	friend8 := Friend{
		Name:      "Shubham  Patil",
		Address:   "chaunan Nagar",
		City:      "Mumbai",
		Mobileno:  7678964455,
		Altmobile: 8721234589,
	}
	friend9 := Friend{
		Name:      "Rajat Chauhan",
		Address:   "Magarpatta",
		City:      "Pune",
		Mobileno:  7671235468,
		Altmobile: 8796874532,
	}
	friend10 := Friend{
		Name:      "Mulul Chaudhari",
		Address:   "shanti Nagar",
		City:      "Jalgoan",
		Mobileno:  9825566455,
		Altmobile: 9825566455,
	}
	friends := []Friend{
		friend1,
		friend2,
		friend3,
		friend4,
		friend5,
		friend6,
		friend7,
		friend8,
		friend9,
		friend10,
	}
	Slice, _ := json.Marshal(friends)

	a := []Friend{}
	err := json.Unmarshal([]byte(string(Slice)), &a)
	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}
	for _, value := range a {
		fmt.Println(value)
	}

}
