import * as fs from "fs";

function main() {
  const data = fs.readFileSync("data.txt", "utf8");
  const array = data.split("\n");

  let elves = [0];
  let i = 0;

  for (const item of array) {
    if (!item) {
      i++;
      elves.push(0);
      continue;
    }

    elves[i] += Number(item);
  }
  elves.sort((a, b) => b - a);

  const topThreeChads = elves.slice(0, 3);

  const sum = topThreeChads.reduce((acc, prev) => (acc += prev), 0);

  console.log(`top 3 chad sum calories ${sum}`);
}

main();
