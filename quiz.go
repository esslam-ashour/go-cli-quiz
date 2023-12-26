package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var cyan = color.New(color.FgCyan)
var introText = `
						░██████╗░░█████╗░░██████╗░██╗░░░██╗██╗███████╗░█████╗░██╗░░░░░██╗
						██╔════╝░██╔══██╗██╔═══██╗██║░░░██║██║╚════██║██╔══██╗██║░░░░░██║
						██║░░██╗░██║░░██║██║██╗██║██║░░░██║██║░░███╔═╝██║░░╚═╝██║░░░░░██║
						██║░░╚██╗██║░░██║╚██████╔╝██║░░░██║██║██╔══╝░░██║░░██╗██║░░░░░██║
						╚██████╔╝╚█████╔╝░╚═██╔═╝░╚██████╔╝██║███████╗╚█████╔╝███████╗██║
						░╚═════╝░░╚════╝░░░░╚═╝░░░░╚═════╝░╚═╝╚══════╝░╚════╝░╚══════╝╚═╝
`

func parseFlags() (string, int, bool) {
	problemsFile := flag.String("pfile", "problems.csv", "The name of the problems CSV file")
	quizDuration := flag.Int("duration", 30, "The time duration for the quiz")
	shuffleProblems := *flag.Bool("shuffle", false, "Option to shuffle questions")
	flag.Parse()

	return *problemsFile, *quizDuration, shuffleProblems
}

func readProblems(problemsFile string, shuffleProblems bool) ([][]string, error) {

	file, err := os.Open(problemsFile)
	if err != nil {
		return nil, fmt.Errorf("error while reading file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records: %v", err)
	}

	if shuffleProblems {
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	return records, nil
}

func askQuestions(problems [][]string, timerExpired chan bool, quizDuration int) (int, int) {
	correctCount := 0
	wrongCount := 0

	timer := time.NewTimer(time.Duration(quizDuration) * time.Second)

	for qid, record := range problems {
		question := record[0]
		answer := record[1]

		var userAnswer string
		fmt.Printf("Question %d: %s\n\n", qid+1, question)
		fmt.Println("Answer: ")

		go func() {
			fmt.Scanln(&userAnswer)
			timerExpired <- true
		}()

		userAnswer = strings.ToLower(strings.TrimSpace(userAnswer))

		select {
		case <-timer.C:
			fmt.Println("Time's up! Quiz is over.")
			return correctCount, wrongCount
		case <-timerExpired:
			if userAnswer == answer {
				correctCount++
			} else {
				wrongCount++
			}
		}

	}
	return correctCount, wrongCount
}

func startQuiz(problemsFile string, shuffleProblems bool, quizDuration int) {
	problems, err := readProblems(problemsFile, shuffleProblems)
	if err != nil {
		log.Fatal(err)
	}

	cyan.Println(">> Press any key to start the quiz <<")
	fmt.Scanln()

	timerExpired := make(chan bool)
	correct, wrong := askQuestions(problems, timerExpired, quizDuration)
	fmt.Printf("You got %d correct answers and %d wrong answers. That is a score of %d", correct, wrong, (correct/wrong)*100)
	fmt.Printf("Thanks for playing!")
}

func main() {
	cyan.Println(introText)
	problemsFile, quizDuration, shuffleProblems := parseFlags()
	startQuiz(problemsFile, shuffleProblems, quizDuration)
}
