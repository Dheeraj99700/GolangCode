package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

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
	}

	friend2 := Friend{
		name:      "Saurabh Tupe",
		address:   "Vallabh Nagar",
		city:      "Mumbai",
		mobileno:  9877664455,
		altmobile: 3296548789,
	}
	friend3 := Friend{
		name:      "Rajesh Wankhede",
		address:   "Patil Nagar",
		city:      "Nanded",
		mobileno:  7677664455,
		altmobile: 8721234589,
	}
	friend4 := Friend{
		name:      "Raj patil",
		address:   "sahakar Nagar",
		city:      "Jalgoan",
		mobileno:  7698764455,
		altmobile: 8721234589,
	}
	friend5 := Friend{
		name:      "Nachiket Chajjed",
		address:   "Chajjed Nagar",
		city:      "Nashik",
		mobileno:  7678932455,
		altmobile: 8774125789,
	}
	friend6 := Friend{
		name:      "Ashish Najan",
		address:   "Najan Nagar",
		city:      "Nagar",
		mobileno:  7677789222,
		altmobile: 8796789669,
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
		name:      "Rajat Chauhan",
		address:   "Magarpatta",
		city:      "Pune",
		mobileno:  7671235468,
		altmobile: 8721234589,
	}
	friend10 := Friend{
		name:      "Mulul Chaudhari",
		address:   "shanti Nagar",
		city:      "Jalgoan",
		mobileno:  9825566455,
		altmobile: 8721234589,
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
	file, err := os.Create("friends.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	_ = w.Write([]string{"Name", "Address", "City", "Mobile No", "AltMobile"})

	var data [][]string
	for _, record := range friendslist {
		row := []string{record.name, record.address, record.city, strconv.Itoa(record.mobileno), strconv.Itoa(record.altmobile)}
		data = append(data, row)
	}
	w.WriteAll(data)

}
