import fs from "fs";

async function main() {
  const file: string = await fs.promises.readFile("input.txt", "utf-8");

  const lines = file.split("\n");

  const groups: string[][] = [];

  let groupIndex = 0;
  for (let i = 0; i < lines.length; i++) {
    if (!groups[groupIndex]) groups.push([]);

    groups[groupIndex].push(lines[i]);
    if (i % 3 == 2) {
      groupIndex++;
    }
  }

  const badgePrioritySum = groups
    .map(([a, b, c]) => {
      const uniqueA = [...new Set(a)];
      const uniqueB = [...new Set(b)];
      const uniqueC = [...new Set(c)];

      const commonChars = uniqueA.filter(
        (char) => uniqueB.includes(char) && uniqueC.includes(char)
      );
      return commonChars[0];
    })
    .map((char) => charToPriority(char))
    .reduce((acc, curr) => acc + curr, 0);

  console.log(badgePrioritySum);
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
