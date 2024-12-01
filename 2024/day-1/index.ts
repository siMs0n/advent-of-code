const parseNumbers = (input: string) => {
  const [leftList, rightList] = input.split(/\r?\n/).reduce<number[][]>(
    (acc, line) => {
      const [leftListAcc, rightListAcc] = acc;
      const firstNonNumberIndex = line
        .split("")
        .findIndex((c) => !Number.isInteger(parseInt(c)));
      const leftNumber = line.substring(0, firstNonNumberIndex);
      const firstRightNumber = line
        .slice(firstNonNumberIndex)
        .split("")
        .findIndex((c) => Number.isInteger(parseInt(c)));
      const rightNumber = line.substring(
        firstNonNumberIndex + firstRightNumber
      );

      return [
        [...leftListAcc, parseInt(leftNumber)],
        [...rightListAcc, parseInt(rightNumber)],
      ];
    },
    [[], []]
  );

  return [leftList, rightList];
};

const one = (input: string) => {
  const [leftList, rightList] = parseNumbers(input);

  leftList.sort();
  rightList.sort();

  const answer = leftList.reduce((acc, val, index) => {
    return acc + Math.abs(val - rightList[index]);
  }, 0);

  return answer;
};

const two = (input: string) => {
  const [leftList, rightList] = parseNumbers(input);
  const occurencesOfNumbers = new Map();
  rightList.forEach((n) => {
    const occurencesOfN = occurencesOfNumbers.get(n) ?? 0;
    occurencesOfNumbers.set(n, occurencesOfN + 1);
  });
  const similarityScore = leftList.reduce((sum, leftNumber) => {
    const occurencesInRightList = occurencesOfNumbers.get(leftNumber);
    return (
      sum + (occurencesInRightList ? leftNumber * occurencesInRightList : 0)
    );
  }, 0);

  return similarityScore;
};

const expectedTestAnswerPartOne = 11;
const inputTest = await Deno.readTextFile("test.txt");

const answerOneTest = one(inputTest);
console.log("Answer one test", answerOneTest);
console.log(
  `Answer is ${
    answerOneTest === expectedTestAnswerPartOne ? "correct" : "wrong"
  }`
);

const input = await Deno.readTextFile("input.txt");
const answerOne = one(input);
console.log("Answer one", answerOne);

const expectedTestAnswerPartTwo = 31;

const answerTwoTest = two(inputTest);
console.log("Answer two test", answerTwoTest);
console.log(
  `Answer is ${
    answerTwoTest === expectedTestAnswerPartTwo ? "correct" : "wrong"
  }`
);

const answerTwo = two(input);
console.log("Answer two", answerTwo);
