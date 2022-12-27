package main

import "fmt"

type Friend struct {
	name, address, city, friend8 string
	mobileno, altmobile          int
}

func main() {

	friend1 := Friend{
		name:      "Dheeraj Chaudhari",
		address:   "Shanti Nagar",
		city:      "Nashik",
		mobileno:  8877548455,
		altmobile: 8965448789,
		friend8:   "collegefriend",
	}

	friend2 := Friend{
		name:      "Saurabh Tupe",
		address:   "Vallabh Nagar",
		city:      "Mumbai",
		mobileno:  9877664455,
		altmobile: 3296548789,
		friend8:   "areafriend",
	}
	friend3 := Friend{
		name:     "Rajesh Wankhede",
		address:  "Patil Nagar",
		city:     "Nanded",
		mobileno: 7677664455,
		friend8:  "areafriend",
	}
	friend4 := Friend{
		name:     "Raj patil",
		address:  "sahakar Nagar",
		city:     "Jalgoan",
		mobileno: 7698764455,
		friend8:  "collegefriend",
	}
	friend5 := Friend{
		name:      "Nachiket Chajjed",
		address:   "Chajjed Nagar",
		city:      "Nashik",
		mobileno:  7678932455,
		altmobile: 8774125789,
		friend8:   "areafriend",
	}
	friend6 := Friend{
		name:      "Ashish Najan",
		address:   "Najan Nagar",
		city:      "Nagar",
		mobileno:  7677789222,
		altmobile: 8796789669,
		friend8:   "areafriend",
	}
	friend7 := Friend{
		name:      "Omkar Wankhede",
		address:   "om Nagar",
		city:      "Pune",
		mobileno:  7678965655,
		altmobile: 8791245689,
	}
	friend8 := Friend{
		name:      "Shubham  Patil",
		address:   "chaunan Nagar",
		city:      "Mumbai",
		mobileno:  7678964455,
		altmobile: 8721234589,
	}
	friend9 := Friend{
		name:     "Rajat Chauhan",
		address:  "Magarpatta",
		city:     "Pune",
		mobileno: 7671235468,
		friend8:  "collegefriend",
	}
	friend10 := Friend{
		name:     "Mulul Chaudhari",
		address:  "shanti Nagar",
		city:     "Jalgoan",
		mobileno: 9825566455,
		friend8:  "collegefriend",
	}
	friendslist := []Friend{
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
	fmt.Printf("Friend Name and MObile Number\n")
	for _, friend := range friendslist {
		fmt.Println(friend.name, friend.mobileno)
	}
	fmt.Println("\n\nFriend Name and MObile Number who have alternate number")
	for _, friend := range friendslist {
		if friend.altmobile != 0 {
			fmt.Println(friend.name, friend.mobileno, friend.altmobile)
		}
	}
	//map
	var nameandtown = make(map[string]string)
	for _, friend := range friendslist {
		nameandtown[friend.name] = friend.city
	}

	for key, value := range nameandtown {
		fmt.Printf("%v ---> %v \n", key, value)
	}

	//collegefriend
	for _, friend := range friendslist {
		if friend.friend8 == "collegefriend" {
			fmt.Println(friend)
		}
	}

}
