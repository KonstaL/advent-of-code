import fs from "fs";

async function main() {
  const file = await fs.promises.readFile("input.txt", "utf-8");

  const cals = file
    .split("\n")
    .map((line) => parseInt(line.trim()))
    .reduce(
      (acc, curr) => {
        if (Number.isNaN(curr)) {
          acc.push(0);
        } else {
          acc[acc.length - 1] += curr;
        }
        return acc;
      },
      [0]
    );

  cals.sort((a, b) => b - a);

  let maxCals = 0;
  for (let i = 0; i < 3; i++) {
    maxCals += cals[i];
  }

  console.log(maxCals);
}

main();
