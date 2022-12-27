package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question      string
	Option_a      string
	Option_b      string
	Option_c      string
	Option_d      string
	CorrectOption string
}

func init() {
	rand.Seed(time.Now().UnixNano())

}
func main() {
	var choice string
	var Score int = 0

	file, err := os.Open("Quiz.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	_, err = csvReader.Read()
	rec, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	var QuizList []Quiz
	for _, record := range rec {

		list := Quiz{
			Question:      record[0],
			Option_a:      record[1],
			Option_b:      record[2],
			Option_c:      record[3],
			Option_d:      record[4],
			CorrectOption: record[5],
		}
		QuizList = append(QuizList, list)
	}
	random := rand.Perm(10)
	for i := 0; i < len(QuizList); i++ {
		j := random[i]
		fmt.Printf("Question:%s\nOption A:%s\nOption B:%s\nOption C:%s\nOption D:%s\n", QuizList[j].Question, QuizList[j].Option_a, QuizList[j].Option_b, QuizList[j].Option_c, QuizList[j].Option_d)
		fmt.Println("Enter answer A,B,C,D")
		fmt.Scanf("%s", &choice)
		if choice == QuizList[j].CorrectOption || strings.ToUpper(choice) == QuizList[j].CorrectOption {
			Score++
		}
	}
	fmt.Printf("Your Exam Score is %d\n", Score)

}
