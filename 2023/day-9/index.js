const fs = require("fs");

const differencesBetweenNumbers = (numbers) => {
  const diffs = [];
  numbers.forEach((n, i) => {
    if (i > 0) diffs[i - 1] = numbers[i] - numbers[i - 1];
  });
  return diffs;
};

const findDifferences = (allNumbers, numbers) => {
  if (numbers.every((d) => d === 0)) {
    return allNumbers;
  }
  const differences = differencesBetweenNumbers(numbers);
  return findDifferences([...allNumbers, differences], differences);
};

const findNextHistoryValue = (rowsOfNumbers) => {
  return rowsOfNumbers
    .slice(0, -1)
    .reverse()
    .reduce((nextHistoryValueOfLastRow, numberRow) => {
      const lastValueInRow = numberRow[numberRow.length - 1];
      return lastValueInRow + nextHistoryValueOfLastRow;
    }, 0);
};

const findPreviousHistoryValue = (rowsOfNumbers) => {
  return rowsOfNumbers
    .slice(0, -1)
    .reverse()
    .reduce((previousHistoryValueOfLastRow, numberRow) => {
      return numberRow[0] - previousHistoryValueOfLastRow;
    }, 0);
};

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const histories = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(" ").map((s) => parseInt(s)));

  const differences = histories.map((history) =>
    findDifferences([history], history)
  );

  const nextValues = differences.map(findNextHistoryValue);

  const sum = nextValues.reduce((sum, current) => sum + current, 0);

  return sum;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const histories = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(" ").map((s) => parseInt(s)));

  const differences = histories.map((history) =>
    findDifferences([history], history)
  );

  const previousValues = differences.map(findPreviousHistoryValue);

  const sum = previousValues.reduce((sum, current) => sum + current, 0);

  return sum;
};

const expectedTestAnswer = [114, 2];

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
