if __name__ == "__main__":
    with open("data.txt") as file:
        elves = [0]
        i = 0
        for line in file:
            if line == "\n":
                i += 1
                elves.append(0)
                continue
            elves[i] += int(line)
    elves.sort(reverse=True)
    topElves = elves[:3]

    print(f"3 chad calories are: {sum(topElves)}")
