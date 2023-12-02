const fs = require("fs");

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");
  const maxColors = {
    red: 12,
    green: 13,
    blue: 14,
  };

  const isPullPossible = ({ nr, color }) => {
    return maxColors[color] >= nr;
  };

  const games = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [game, rest] = line.split(": ");
      const id = parseInt(game.split(" ")[1]);

      const gamePossible = rest.split("; ").every((round) => {
        const roundPossible = round
          .split(", ")
          .map((str) => {
            const [nrAsStr, color] = str.split(" ");
            return { nr: parseInt(nrAsStr), color };
          })
          .every(isPullPossible);

        return roundPossible;
      });

      return { id, gamePossible };
    });

  const sumOfGamesPossible = games.reduce(
    (sum, game) => (game.gamePossible ? sum + game.id : sum),
    0
  );

  return sumOfGamesPossible;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const findMax = (list) => {
    return list.reduce(
      (gameMax, roundMax) => {
        return {
          red: Math.max(gameMax.red, roundMax.red),
          green: Math.max(gameMax.green, roundMax.green),
          blue: Math.max(gameMax.blue, roundMax.blue),
        };
      },
      {
        red: 0,
        green: 0,
        blue: 0,
      }
    );
  };

  const powersOfCubes = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [_, rest] = line.split(": ");

      const rounds = rest.split("; ").map((round) => {
        const cubes = round
          .split(", ")
          .map((str) => {
            const [nrAsStr, color] = str.split(" ");
            return { nr: parseInt(nrAsStr), color };
          })
          .reduce(
            (max, cube) => {
              return { ...max, [cube.color]: cube.nr };
            },
            {
              red: 0,
              green: 0,
              blue: 0,
            }
          );

        return cubes;
      });

      const { red, green, blue } = findMax(rounds);
      return red * green * blue;
    });

  const sum = powersOfCubes.reduce((sum, current) => sum + current, 0);

  return sum;
};

const expectedTestAnswer = [8, 2286];

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
