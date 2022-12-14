import fs from "fs";

async function main() {
  const file: string = await fs.promises.readFile("input.txt", "utf-8");

  const overlappinRanges = file
    .split("\n")
    .map((line) => line.split(","))
    .map(([strA, strB]) => {
      const rangeA = strA.split("-").map(Number);
      const rangeB = strB.split("-").map(Number);

      return (
        isRangeOverlapping(rangeA, rangeB) || isRangeOverlapping(rangeB, rangeA)
      );
    })
    .map((isContained) => (isContained ? 1 : 0) as number)
    .reduce((acc, curr) => acc + curr, 0);

  console.log(overlappinRanges);
}

function isRangeOverlapping(a: number[], b: number[]) {
  if (b[0] >= a[0] && b[0] <= a[1]) return true;
  if (b[1] >= a[0] && b[1] <= a[1]) return true;

  return false;
}

main();
