# GoQuizCLI
<p align="center">
  <img style="text-align: center;" width="632" alt="Screenshot 2023-12-26 at 2 52 46 AM" src="https://github.com/esslam-ashour/GoQuizCLI/assets/61587419/9710f1c8-5764-4e64-a9db-d9e831586b3b">
</p>

A simple quiz CLI app made using Go.

## Usage:
First, `quiz.go` needs to be compiled, run the following command inside the directory:
    `go build`
    
Then, we can use the program as follows:
    `go run qo-quiz-cli [flags]`

The flags are:

    -pfile=<string>                  this is the name of the problems CSV file (default is "problems.csv")
    -duration=<int>                     this is the duration of the quiz in seconds (default is 30 seconds)
    -shuffle=<bool>                     this is the option to shuffle problems (default is false)

Example:

    go-quiz-cli -problems="quiz_problems.csv" -duration=130 -shuffle=true
      

## Skills I learned:


