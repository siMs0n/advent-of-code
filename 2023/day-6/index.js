const fs = require("fs");

const solveWithPq = (p, q) => {
  const root1 = -p / 2 + Math.sqrt(Math.pow(p / 2, 2) - q);
  const root2 = -p / 2 - Math.sqrt(Math.pow(p / 2, 2) - q);
  return [root1, root2];
};

const getNrOfWays = (timeAllowed, recordDistance) => {
  const [root1, root2] = solveWithPq(timeAllowed, recordDistance);

  const time1 = Math.abs(root1);
  const time2 = Math.abs(root2);

  // If even, hold 1 ms longer/shorter to beat the record
  const shortestHoldTime = time1 % 1 != 0 ? Math.ceil(time1) : time1 + 1;
  const longestHoldTime = time2 % 1 != 0 ? Math.floor(time2) : time2 - 1;

  return longestHoldTime - shortestHoldTime + 1;
};

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const [times, distances] = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((l) =>
      l
        .split(":")[1]
        .trim()
        .split(/\s+/)
        .map((s) => parseInt(s))
    );

  // timeHold * (time - timeHold) = distance
  // x * (7 - x) ) = 9
  // 7x - x2 = 9
  // x2 +7x + 9 = 0
  // x = -1.7 eller -5.3 so everything between 2 and 5
  // Math.ceil(x1) Math.floor(x2)
  // solution = (s2 - s1) + 1

  const nrOfWays = times.map((t, i) => getNrOfWays(t, distances[i]));

  const product = nrOfWays.reduce((sum, current) => sum * current, 1);

  return product;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");
  const [time, distance] = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((l) => parseInt(l.split(":")[1].trim().split(/\s+/).join("")));

  const nrOfWays = getNrOfWays(time, distance);

  return nrOfWays;
};

const expectedTestAnswer = [288, 71503];

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
