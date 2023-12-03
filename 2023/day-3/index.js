const fs = require("fs");

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const matrix = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  const isNumber = (str) => !isNaN(parseInt(str));

  const isSymbol = (char) => {
    return !(isNumber(char) || char === ".");
  };

  const hasNearbySymbol = (matrix, yIndex, xIndex, nrLength) => {
    const startY = Math.max(yIndex - 1, 0);
    const endY = Math.min(yIndex + 1, matrix.length - 1);
    const startX = Math.max(xIndex - 1, 0);
    const endX = Math.min(xIndex + nrLength, matrix[yIndex].length - 1);

    for (let y = startY; y <= endY; y++) {
      for (let x = startX; x <= endX; x++) {
        if (isSymbol(matrix[y][x])) {
          return true;
        }
      }
    }

    return false;
  };

  const partNumbers = matrix.flatMap((line, yIndex) => {
    let partNumbersInLine = [];
    let startNumberIndex = -1;
    let currentNumberAsStr = "";
    for (let x = 0; x <= line.length - 1; x++) {
      const currentChar = matrix[yIndex][x];
      if (!isNumber(currentChar)) {
        continue;
      }
      // Is number
      currentNumberAsStr += currentChar;
      startNumberIndex = startNumberIndex > -1 ? startNumberIndex : x;

      if (x + 1 < line.length && isNumber(matrix[yIndex][x + 1])) {
        // If we still have line left and the next char is a number
        continue;
      }

      // End of line or end of number, check around it
      if (
        hasNearbySymbol(
          matrix,
          yIndex,
          startNumberIndex,
          x - startNumberIndex + 1
        )
      ) {
        partNumbersInLine = [
          ...partNumbersInLine,
          parseInt(currentNumberAsStr),
        ];
      }
      startNumberIndex = -1;
      currentNumberAsStr = "";
    }

    return partNumbersInLine;
  });

  const sum = partNumbers.reduce((sum, current) => sum + current, 0);

  return sum;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const matrix = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  const isNumber = (str) => !isNaN(parseInt(str));
  const isGear = (str) => str === "*";

  const nearbyGears = (matrix, yIndex, xIndex, nrLength) => {
    const startY = Math.max(yIndex - 1, 0);
    const endY = Math.min(yIndex + 1, matrix.length - 1);
    const startX = Math.max(xIndex - 1, 0);
    const endX = Math.min(xIndex + nrLength, matrix[yIndex].length - 1);

    let gears = [];
    for (let y = startY; y <= endY; y++) {
      for (let x = startX; x <= endX; x++) {
        if (isGear(matrix[y][x])) {
          gears = [...gears, y + "," + x];
        }
      }
    }

    return gears;
  };

  const gearIndexToNumbersThatAreClose = {};

  matrix.forEach((line, yIndex) => {
    let startNumberIndex = -1;
    let currentNumberAsStr = "";
    for (let x = 0; x <= line.length - 1; x++) {
      const currentChar = matrix[yIndex][x];
      if (!isNumber(currentChar)) {
        continue;
      }
      // Is number
      currentNumberAsStr += currentChar;
      startNumberIndex = startNumberIndex > -1 ? startNumberIndex : x;

      if (x + 1 < line.length && isNumber(matrix[yIndex][x + 1])) {
        // If we still have line left and the next char is a number
        continue;
      }

      // End of line or end of number, check around it
      const nearbyGearsList = nearbyGears(
        matrix,
        yIndex,
        startNumberIndex,
        x - startNumberIndex + 1
      );
      if (nearbyGearsList.length > 0) {
        nearbyGearsList.forEach((gear) => {
          if (!gearIndexToNumbersThatAreClose[gear]) {
            gearIndexToNumbersThatAreClose[gear] = [];
          }
          gearIndexToNumbersThatAreClose[gear] = [
            ...gearIndexToNumbersThatAreClose[gear],
            parseInt(currentNumberAsStr),
          ];
        });
      }
      startNumberIndex = -1;
      currentNumberAsStr = "";
    }
  });

  const gearRatios = Object.values(gearIndexToNumbersThatAreClose)
    .filter((partNumbers) => partNumbers.length === 2)
    .map((partNumbers) => partNumbers[0] * partNumbers[1]);

  const sum = gearRatios.reduce((sum, current) => sum + current, 0);

  return sum;
};

const expectedTestAnswer = [4361, 467835];

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
