package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameObj struct {
	Name        string
	WinsAgainst string
	Points      int
}

func getOutcome(choice string, opponentChoice string) int {

	casesMap := make(map[string]GameObj)
	casesMap["X"] = GameObj{
		// rock
		Name:        "A",
		WinsAgainst: "C",
		Points:      1,
	}
	casesMap["Y"] = GameObj{
		// paper
		Name:        "B",
		WinsAgainst: "A",
		Points:      2,
	}
	casesMap["Z"] = GameObj{
		// scissors
		Name:        "C",
		WinsAgainst: "B",
		Points:      3,
	}

	shape := casesMap[choice]

	isWin := shape.WinsAgainst == opponentChoice
	isDraw := shape.Name == opponentChoice

	if isWin {
		return shape.Points + 6
	}

	if isDraw {
		return shape.Points + 3
	}

	return shape.Points

}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameCase := strings.Split(line, " ")

		opponentChoice := gameCase[0]
		playerChoice := gameCase[1]
		gamePoints := getOutcome(playerChoice, opponentChoice)
		totalPoints += gamePoints
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	fmt.Println(totalPoints)
}
