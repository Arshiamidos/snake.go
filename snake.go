package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	term "github.com/nsf/termbox-go"
)

const BOARD_SIZE = 20

func init() {

}
func reset() {
	term.Sync() // cosmestic purpose
}
func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()
	board := make([][]int, BOARD_SIZE)
	timer := time.Tick(1000 * time.Millisecond)
	input := make(chan int, 1)
	direction := "W" //S W N E
	snake := [][]int{
		{BOARD_SIZE / 2, BOARD_SIZE / 2},
	}
	goal := [][]int{
		{5, 5},
	}

	var wg sync.WaitGroup
	Setup(board)
	wg.Add(1) //setup waiter for goroutins

	go func() { //render base on input recieve
		for {
			select {
			case m := <-input:
				m = m
				ShowBoardSnakeGoal(board, snake, goal)
			case <-timer:
				ClearScreen()
				fmt.Println(snake)
				CalcSnakePosition(snake, direction)
				ShowBoardSnakeGoal(board, snake, goal)

			}
		}
	}()

	go func() { //wait for input in channel
		for {
			ev := term.PollEvent()
			input <- 1
			switch ev.Key {
			case term.KeyEsc:
				os.Exit(1)
			case term.KeyArrowDown:
				direction = "S"
				reset()
			case term.KeyArrowUp:
				direction = "N"
				reset()
			case term.KeyArrowLeft:
				direction = "W"
				reset()
			case term.KeyArrowRight:
				direction = "E"
				reset()
			}
			//snake = append(snake, []int{1, 1})

			//fmt.Println(snake)
			//time.Sleep(5 * time.Second)
		}
	}()

	wg.Wait() //wait for goroutins

}
func CalcSnakePosition(snake [][]int, direction string) {
	if direction == "W" {
		s := snake[0]
		s[0] = s[0] - 1
	} else if direction == "S" {
		s := snake[0]
		s[1] = s[1] + 1
	} else if direction == "E" {
		s := snake[0]
		s[0] = s[0] + 1
	} else if direction == "N" {
		s := snake[0]
		s[1] = s[1] + 1
	}
}
func ClearScreen() {
	fmt.Println("\033[2J")
}
func Setup(board [][]int) {
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = make([]int, BOARD_SIZE)
		for j := 0; j < BOARD_SIZE; j++ {
			board[i][j] = 0

		}
	}
}
func ShowBoardSnakeGoal(board [][]int, snake [][]int, goal [][]int) {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {

			if goal[0][0] == i && goal[0][1] == j && snake[0][0] == i && snake[0][1] == j {
				fmt.Print("🤡 ")
			} else if goal[0][0] == i && goal[0][1] == j {
				fmt.Print("❤️ ")
			} else {
				for k, s := range snake {
					if k == 0 && s[0] == i && s[1] == j {
						fmt.Print("👹 ")
					} else if k != 0 && s[k] == i && s[k] == j {
						fmt.Print("🔷 ")
					} else {
						fmt.Print("⏹ ")
					}

				}
			}

		}
		fmt.Print("\n")
	}
}
