package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type familyMember struct {
	relation string
	DOB      string
	Age      int
}

func main() {
	currentyear, _, _ := time.Now().Date()
	//var age []string
	member1 := familyMember{relation: "Father ", DOB: "1965-02-06"}
	member2 := familyMember{relation: "Mother ", DOB: "1970-12-16"}
	member3 := familyMember{relation: "Brother", DOB: "1983-06-14"}
	member4 := familyMember{relation: "Sister ", DOB: "1990-05-15"}
	memberlist := []familyMember{
		member1,
		member2,
		member3,
		member4,
	}
	for i := 0; i < len(memberlist); i++ {
		odd := strings.Split(memberlist[i].DOB, "-")
		birthyear, _ := strconv.Atoi(odd[0])
		age := currentyear - birthyear
		memberlist[i].Age = age

	}
	for _, Members := range memberlist {
		fmt.Println(Members)
	}

}

//age = append(age, member1.DOB)
//fmt.Println(age)
