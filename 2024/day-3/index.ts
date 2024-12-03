const one = (input: string) => {
  const regex = /mul\(\d+,\d+\)/g;
  const instructions = input.match(regex);

  const answer = instructions?.reduce((sum, instruction) => {
    const factors = instruction.match(/\d+/g);
    if (!factors || factors.length < 2) throw new Error("Didn't find factors");
    const [factor1, factor2] = factors;
    return sum + parseInt(factor1) * parseInt(factor2);
  }, 0);

  return answer;
};

const two = (input: string) => {
  const regex = /mul\(\d+,\d+\)|do\(\)|don't\(\)/g;
  const instructions = input.matchAll(regex);

  const { sum: answer } = [...instructions].reduce(
    ({ sum, enabled }, instructionMatch) => {
      const instruction = instructionMatch[0];

      if (instruction === "do()") {
        return { sum, enabled: true };
      }

      if (instruction === "don't()") {
        return { sum, enabled: false };
      }

      if (enabled === false) {
        return { sum, enabled };
      }

      const factors = instruction.match(/\d+/g);
      if (!factors || factors.length < 2)
        throw new Error("Didn't find factors " + instruction);

      const [factor1, factor2] = factors;
      return { sum: sum + parseInt(factor1) * parseInt(factor2), enabled };
    },
    { sum: 0, enabled: true }
  );
  return answer;
};

const expectedTestAnswerPartOne = 161;
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

const expectedTestAnswerPartTwo = 48;

const inputTest2 = await Deno.readTextFile("test2.txt");
const answerTwoTest = two(inputTest2);
console.log("Answer two test", answerTwoTest);
console.log(
  `Answer is ${
    answerTwoTest === expectedTestAnswerPartTwo ? "correct" : "wrong"
  }`
);

const answerTwo = two(input);
console.log("Answer two", answerTwo);
