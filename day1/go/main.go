package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := []int{0}
	i := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, 0)
			i++
			continue
		}

		foodCalories, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		elves[i] += foodCalories
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	chadElves := elves[:3]

	total := 0
	for _, num := range chadElves {
		total += num
	}

	fmt.Println("Total calories of top 3 elves is:", total)
}
