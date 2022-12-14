import fs from "fs";

async function main() {
  const file: string = await fs.promises.readFile("input.txt", "utf-8");

  const containedRanges = file
    .split("\n")
    .map((line) => line.split(","))
    .map(([strA, strB]) => {
      const rangeA = strA.split("-").map(Number);
      const rangeB = strB.split("-").map(Number);

      return (
        isRangeContained(rangeA, rangeB) || isRangeContained(rangeB, rangeA)
      );
    })
    .map((isContained) => (isContained ? 1 : 0) as number)
    .reduce((acc, curr) => acc + curr, 0);

  console.log(containedRanges);
}

function isRangeContained(a: number[], b: number[]) {
  return a[0] <= b[0] && a[1] >= b[1];
}

main();
