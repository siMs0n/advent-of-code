const parseReports = (input: string): number[][] => {
  return input.split(/\r?\n/).map((line) => {
    return line.split(" ").map((l) => parseInt(l));
  });
};

const isReportSafe = (report: number[]) =>
  report.reduce<{
    ascending: boolean;
    prev: number;
    safe: boolean;
  }>(
    ({ ascending, prev, safe }, level, index) => {
      if (!safe) return { ascending, prev, safe };

      if (index === 0) return { ascending, prev: level, safe };

      if (Math.abs(prev - level) > 3 || prev === level) {
        return { ascending, prev: level, safe: false };
      }

      if (index === 1) {
        return { ascending: prev < level, prev: level, safe: true };
      }

      const isSafe = ascending ? prev < level : prev > level;

      return { ascending, prev: level, safe: isSafe };
    },
    { ascending: false, prev: -1, safe: true }
  );

const one = (input: string) => {
  const reports = parseReports(input);

  const answer = reports.reduce((acc, r) => {
    const { safe: reportIsSafe } = isReportSafe(r);

    return acc + (reportIsSafe ? 1 : 0);
  }, 0);

  return answer;
};

const two = (input: string) => {
  const reports = parseReports(input);

  const answer = reports.reduce((acc, r) => {
    const reportIsSafe = r.some((l, i) => {
      const reportWithoutLevelI = r.toSpliced(i, 1);
      return isReportSafe(reportWithoutLevelI).safe;
    });

    return acc + (reportIsSafe ? 1 : 0);
  }, 0);

  return answer;
};

const expectedTestAnswerPartOne = 2;
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

const expectedTestAnswerPartTwo = 4;

const answerTwoTest = two(inputTest);
console.log("Answer two test", answerTwoTest);
console.log(
  `Answer is ${
    answerTwoTest === expectedTestAnswerPartTwo ? "correct" : "wrong"
  }`
);

const answerTwo = two(input);
console.log("Answer two", answerTwo);
