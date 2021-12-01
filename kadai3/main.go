package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var questions = []string{
	"ruby",
	"python",
	"node",
	"go",
	"php",
	"nim",
	"docker",
	"rustc",
	"docker-compose",
	"scala",
	"rails",
	"git",
	"npm",
	"vim",
	"brew",
	"yarn",
	"pip",
	"dep",
	"aws",
	"gcloud",
	"drone",
}

var t int

func init() {
	flag.IntVar(&t, "t", 1, "制限時間（分）")
	flag.Parse()
}

func makeQuestion() string {
	rand.Seed(time.Now().UnixNano())
	q := questions[rand.Intn(len(questions))]
	return q
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func()  {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	fmt.Println("input end")
	return ch
}

func main() {
	fmt.Printf("タイピングゲームを始めます。制限時間は%v分\n", t)

	ch1 := input(os.Stdin)
	point := 0

	go func() {
		time.After(time.Duration(t) * time.Minute)
		fmt.Println("時間切れです！")
		fmt.Printf("pointは%vです\n", point)
		os.Exit(0)
	}()

	for {
		question := makeQuestion()
		fmt.Println(question)
		select {
		case my_ans := <- ch1:
			if my_ans == question {
				point++
				fmt.Println("Correct!")
				fmt.Println("===============")
			} else {
				fmt.Println("oh, mistake...")
				fmt.Println("===============")
			}
		}
	}
}