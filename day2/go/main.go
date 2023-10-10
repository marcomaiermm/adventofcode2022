package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameObj struct {
	Name         string
	WinsAgainst  string
	LosesAgainst string
	Points       int
}

func getOutcome(choice string, opponentChoice string) int {

	casesMap := make(map[string]GameObj)

	casesMap["A"] = GameObj{
		// rock
		Name:         "A",
		WinsAgainst:  "C",
		LosesAgainst: "B",
		Points:       1,
	}

	casesMap["B"] = GameObj{
		// paper
		Name:         "B",
		WinsAgainst:  "A",
		LosesAgainst: "C",
		Points:       2,
	}

	casesMap["C"] = GameObj{
		// scissors
		Name:         "C",
		WinsAgainst:  "B",
		LosesAgainst: "A",
		Points:       3,
	}

	opponentShape := casesMap[opponentChoice]
	var playerShape GameObj
	points := 0

	switch choice {
	case "Z":
		playerShape = casesMap[opponentShape.LosesAgainst]
		points = playerShape.Points + 6
		fmt.Println(playerShape.Name, "wins against", opponentChoice, "->", points)
	case "Y":
		playerShape = opponentShape
		points = playerShape.Points + 3
		fmt.Println(playerShape.Name, "draw against", opponentChoice, "->", points)
	default:
		playerShape = casesMap[opponentShape.WinsAgainst]
		points = playerShape.Points
		fmt.Println(playerShape.Name, "loses against", opponentChoice, "->", points)
	}

	return points
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
