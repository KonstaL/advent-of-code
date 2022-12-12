import fs from "fs";

async function main() {
  const file = await fs.promises.readFile("a.txt", "utf-8");

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

  const maxCal = cals.reduce((acc, curr) => Math.max(acc, curr), 0);
  console.log(maxCal);
}

main();
