import fs from "fs";

async function main() {
  const file = await fs.promises.readFile("input.txt", "utf-8");

  const roundScores = file
    .split("\n")
    .map((line) => line.split(" "))
    .map((line) => translateArgs(line[0], line[1]))
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

function translateArgs(arg1: string, arg2: string) {
  let a = "R";

  if (arg1 === "B") a = "P";
  if (arg1 === "C") a = "S";

  if (arg2 === "X" && a === "R") return [a, "S"];
  if (arg2 === "X" && a === "P") return [a, "R"];
  if (arg2 === "X" && a === "S") return [a, "P"];

  if (arg2 === "Y") return [a, a];

  if (arg2 === "Z" && a === "R") return [a, "P"];
  if (arg2 === "Z" && a === "P") return [a, "S"];
  if (arg2 === "Z" && a === "S") return [a, "R"];

  throw new Error("Invalid argument");
}

main();
