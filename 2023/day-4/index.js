const fs = require("fs");

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const cards = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [winning, yours] = line
        .split(":")[1]
        .split("|")
        .flatMap((l) => l.trim());

      const winningNumbers = winning.split(" ").map((str) => parseInt(str));
      const yourNumbers = yours.split(/\s+/).map((str) => parseInt(str));

      return [winningNumbers, yourNumbers];
    });

  const yourPoints = cards.map(([winning, yours]) => {
    const yourWinningNumbers = winning.filter((winningNumber) =>
      yours.includes(winningNumber)
    );
    return yourWinningNumbers.reduce((sum) => (sum === 0 ? 1 : sum * 2), 0);
  });

  const sum = yourPoints.reduce((sum, current) => sum + current, 0);

  return sum;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const cards = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [winning, yours] = line
        .split(":")[1]
        .split("|")
        .flatMap((l) => l.trim());

      const winningNumbers = winning.split(" ").map((str) => parseInt(str));
      const yourNumbers = yours.split(/\s+/).map((str) => parseInt(str));

      return [winningNumbers, yourNumbers];
    });

  const nrOfCards = new Array(cards.length).fill(1);

  cards.forEach(([winning, yours], index) => {
    const nrOfWins = winning.filter((winningNumber) =>
      yours.includes(winningNumber)
    ).length;

    for (let i = index + 1; i <= index + nrOfWins; i++) {
      nrOfCards[i] = nrOfCards[i] + nrOfCards[index];
    }

    return nrOfWins;
  });

  const sum = nrOfCards.reduce((sum, current) => sum + current, 0);

  return sum;
};

const expectedTestAnswer = [13, 30];

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
