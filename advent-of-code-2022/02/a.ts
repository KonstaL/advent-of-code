import fs from "fs";

async function main() {
  const file = await fs.promises.readFile("input.txt", "utf-8");

  const roundScores = file
    .split("\n")
    .map((line) => line.split(" "))
    .map((line) => [translateArgs(line[0]), translateArgs(line[1])])
    .map((line) => getRoundScore(line));

  const totalScore = roundScores.reduce((acc, curr) => acc + curr, 0);
  console.log(totalScore);
}

function getRoundScore(line: string[]) {
  let score = 0;
  const [a, b] = line;

  if (b === "R") score += 1;
  if (b === "P") score += 2;
  if (b === "S") score += 3;

  if (
    (a === "R" && b === "P") ||
    (a === "P" && b === "S") ||
    (a === "S" && b === "R")
  ) {
    score += 6;
  }
  if (a === b) score += 3;
  return score;
}

function translateArgs(arg: string) {
  if (arg === "A" || arg === "X") return "R";
  if (arg === "B" || arg === "Y") return "P";
  if (arg === "C" || arg === "Z") return "S";
  throw new Error("Invalid argument");
}

main();
