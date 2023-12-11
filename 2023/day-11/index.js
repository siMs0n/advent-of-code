const fs = require("fs");

const manhattanDistance = (g1, g2) =>
  Math.abs(g2.x - g1.x) + Math.abs(g2.y - g1.y);

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const universeGrid = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  let galaxies = [];

  let y = 0;
  universeGrid.forEach((row) => {
    let containsGalaxy = false;
    row.forEach((spot, x) => {
      if (spot === "#") {
        galaxies = [...galaxies, { x, y }];
        containsGalaxy = true;
      }
    });
    y = containsGalaxy ? y + 1 : y + 2;
  });

  const galaxyXs = galaxies.map((g) => g.x);
  const emptyColumns = universeGrid[0]
    .map((_, i) => i)
    .filter((i) => !galaxyXs.includes(i));

  emptyColumns.forEach((x, i) => {
    galaxies = galaxies.map((g) => (g.x > x + i ? { x: g.x + 1, y: g.y } : g));
  });

  const distancePairs = galaxies.flatMap((g1) =>
    galaxies.map((g2) => manhattanDistance(g1, g2))
  );

  const sumOfDistanceParis = distancePairs.reduce(
    (sum, current) => sum + current,
    0
  );

  return sumOfDistanceParis / 2;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const universeGrid = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  let galaxies = [];
  const emptyDistance = 1000000;

  const columnDistances = universeGrid[0].map((_, i) => {
    return universeGrid.every((row) => row[i] === ".") ? emptyDistance : 1;
  });

  let y = 0;
  universeGrid.forEach((row) => {
    let containsGalaxy = false;
    let x = -1;
    row.forEach((spot, i) => {
      x = x + columnDistances[i];
      if (spot === "#") {
        galaxies = [...galaxies, { x, y }];
        containsGalaxy = true;
      }
    });
    y = containsGalaxy ? y + 1 : y + emptyDistance;
  });

  const distancePairs = galaxies.flatMap((g1) =>
    galaxies.map((g2) => manhattanDistance(g1, g2))
  );

  const sumOfDistanceParis = distancePairs.reduce(
    (sum, current) => sum + current,
    0
  );

  return sumOfDistanceParis / 2;
};

const expectedTestAnswer = [374, 8410];

const answerOneTest = one("test.txt");
console.log("Answer one test", answerOneTest);
console.log(
  `Answer is ${answerOneTest === expectedTestAnswer[0] ? "correct" : "wrong"}`
);
const answerOne = one("input.txt");
console.log("Answer one", answerOne);

const answerTwoTest = two("test.txt");
console.log("Answer two test", answerTwoTest);
console.log(
  `Answer is ${answerTwoTest === expectedTestAnswer[1] ? "correct" : "wrong"}`
);

const answerTwo = two("input.txt");
console.log("Answer two", answerTwo);
