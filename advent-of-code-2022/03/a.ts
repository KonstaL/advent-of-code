import fs from "fs";

async function main() {
  const file: string = await fs.promises.readFile("input.txt", "utf-8");

  const commonChars = file
    .split("\n")
    .map((line) => {
      const compA = line.slice(0, line.length / 2);
      const compB = line.slice(line.length / 2);
      return [compA, compB];
    })
    .map(([compA, compB]) => {
      const uniqueA = [...new Set(compA)];
      const uniqueB = [...new Set(compB)];
      return [uniqueA, uniqueB];
    })
    .map(([compA, compB]) => {
      const common = [...compA].filter((char) => compB.includes(char));
      return common;
    });

  const priorities = commonChars.map((chars) => {
    const priorities = chars.map(charToPriority);
    return priorities;
  });

  const summedPriorities = priorities
    .map((priority) => priority.reduce((acc, curr) => acc + curr, 0))
    .reduce((acc, curr) => acc + curr, 0);

  console.log(summedPriorities);
}

function charToPriority(char: string) {
  if (char.length > 1) throw new Error("Invalid char");

  const asciiCode = char.charCodeAt(0);
  // Uppercase chars have +26 priority compared to lowercase chars
  if (asciiCode > 64 && asciiCode <= 90) return asciiCode - 64 + 26;
  if (asciiCode > 96 && asciiCode <= 122) return asciiCode - 96;

  throw new Error("Unsupported charr");
}

main();
