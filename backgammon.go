package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.New(color.BgCyan)
	fmt.Println("Welcome to Backgammon!")
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	board := Board{}
	board = board.initiateBoard()
	var err string
	/*for board.game == true {
		board = board.rollDice()
		board.render()
		if board.uses == -1 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("move *character* *num*")
			text, _ := reader.ReadString('\n')
			stuff := strings.Fields(text)
			if len(stuff) == 3 && stuff[0] == "move" {
				board, err = board.move(stuff[1], stuff[2])
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {

		}
	}*/
	for board.game == true {
		board = board.rollDice()
		board.render()
		if board.uses == 0 {
			for board.adice1 != -1 || board.adice2 != -1 {
				reader := bufio.NewReader(os.Stdin)
				fmt.Println("move *character* *num*")
				text, _ := reader.ReadString('\n')
				stuff := strings.Fields(text)
				diceint, interr := strconv.Atoi(stuff[2])
				if interr != nil {
					fmt.Println("You didn't give a number!")
					board.render()
					continue
				}
				if len(stuff) == 3 && stuff[0] == "move" {
					board, err = board.move(stuff[1], diceint)
					if err != "" {
						fmt.Println(err)
					}
				}
				fmt.Println()
				if board.adice1 != -1 || board.adice2 != -1 {
					board.render()
				}
			}
		} else {
			for board.uses != 0 {
				reader := bufio.NewReader(os.Stdin)
				fmt.Println("move *character*")
				text, _ := reader.ReadString('\n')
				stuff := strings.Fields(text)
				if len(stuff) == 2 && stuff[0] == "move" {
					board, err = board.move(stuff[1], board.dice1)
					if err != "" {
						fmt.Println(err)
					}
				}
				fmt.Println()
				if board.uses != 0 {
					board.render()
				}
			}
		}
		if board.turn == 0 {
			board.turn = 1
		} else {
			board.turn = 0
		}
	}
}
