package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var players int
	var dices int

	for {
		fmt.Print("Enter the number of players: ")
		_, playerErr := fmt.Scanf("%d", &players)
		if playerErr != nil {
			fmt.Println("Error:", playerErr)
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter the number of dice: ")
		_, dicesErr := fmt.Scanf("%d", &dices)
		if dicesErr != nil {
			fmt.Println("Error:", dicesErr)
		} else {
			break
		}
	}

	fmt.Println()

	playerRolls := make([][]int, players)
	scores := make([]int, players)
	round := 1

	for idx := range playerRolls {
		playerRolls[idx] = rollDice(dices)
	}
	fmt.Println("Round -", round)
	fmt.Println("Value :", playerRolls)

	for {
		filteredArray, newScores := evaluateRolls(playerRolls, scores)
		scores = newScores
		fmt.Println("Evaluation", filteredArray)
		fmt.Println("Score", newScores)
		fmt.Println()

		notEmpty := 0
		for _, v := range filteredArray {
			if len(v) > 0 {
				notEmpty++
			}
		}
		if notEmpty <= 1 {
			break
		}

		for idx := range playerRolls {
			playerRolls[idx] = rollDice(len(filteredArray[idx]))
		}
		round++
		fmt.Println("Round -", round)
		fmt.Println("Value :", playerRolls)
	}

	Winner := whoIsWin(scores)

	fmt.Println("The Winner : ", Winner)
}

func whoIsWin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxIndex := 0
	maxValue := arr[0]

	for idx, num := range arr {
		if num > maxValue {
			maxValue = num
			maxIndex = idx
		}
	}

	return maxIndex + 1
}

func rollDice(dice int) []int {
	allRolls := make([]int, dice)

	for idx := range allRolls {
		allRolls[idx] = rand.Intn(6) + 1
	}
	return allRolls
}

func evaluateRolls(arr [][]int, scores []int) ([][]int, []int) {
	playerRolls := make([][]int, len(arr))
	copy(playerRolls, arr)
	newScores := make([]int, len(scores))
	copy(newScores, scores)

	for idx := range playerRolls {
		temp := []int{}
		except := 0

		for idx2 := range playerRolls[idx] {
			if playerRolls[idx][idx2] == 1 {
				except++
				continue
			}
			if playerRolls[idx][idx2] == 6 {
				newScores[idx]++
			} else {
				temp = append(temp, playerRolls[idx][idx2])
			}
		}

		for i := 0; i < except; i++ {
			temp = append(temp, 1)
		}

		playerRolls[idx] = temp
	}

	return playerRolls, newScores
}
