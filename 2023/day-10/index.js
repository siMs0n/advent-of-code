const fs = require("fs");

const findStart = (grid) => {
  let position;
  grid.forEach((row, y) =>
    row.forEach((v, x) => {
      if (v === "S") {
        position = { x, y };
      }
    })
  );
  return position;
};

const getNextPosition = (grid, prevPosition, position) => {
  const pipe = grid[position.y][position.x];
  switch (pipe) {
    case "|":
      return {
        x: position.x,
        y: prevPosition.y < position.y ? position.y + 1 : position.y - 1,
      };
    case "-":
      return {
        x: prevPosition.x < position.x ? position.x + 1 : position.x - 1,
        y: position.y,
      };
    case "L":
      return {
        x: prevPosition.y < position.y ? position.x + 1 : position.x,
        y: prevPosition.x > position.x ? position.y - 1 : position.y,
      };
    case "J":
      return {
        x: prevPosition.y < position.y ? position.x - 1 : position.x,
        y: prevPosition.x < position.x ? position.y - 1 : position.y,
      };
    case "7":
      return {
        x: prevPosition.y > position.y ? position.x - 1 : position.x,
        y: prevPosition.x < position.x ? position.y + 1 : position.y,
      };
    case "F":
      return {
        x: prevPosition.y > position.y ? position.x + 1 : position.x,
        y: prevPosition.x > position.x ? position.y + 1 : position.y,
      };
  }
};

const findFurthestPosition = (grid, start, firstPipes) => {
  let steps = 0;
  let previousPositions = start;
  let pipes = firstPipes;
  let isMidPoint = false;
  while (!isMidPoint) {
    steps++;
    const nextPositions = [
      getNextPosition(grid, previousPositions[0], pipes[0]),
      getNextPosition(grid, previousPositions[1], pipes[1]),
    ];
    isMidPoint =
      (pipes[0].x === pipes[1].x && pipes[0].y === pipes[1].y) ||
      (nextPositions[0].x === pipes[1].x &&
        nextPositions[0].y === pipes[1].y &&
        nextPositions[1].x === pipes[0].x &&
        nextPositions[1].y === pipes[0].y);
    previousPositions = pipes;
    pipes = nextPositions;
  }

  return steps;
};

const findFirstPipePositions = (grid, start) => {
  let pipes = [];

  const up = grid[Math.max(0, start.y - 1)][start.x];
  if (["|", "7", "F"].includes(up)) {
    pipes = [...pipes, { x: start.x, y: start.y - 1 }];
  }
  const down = grid[Math.min(grid.length - 1, start.y + 1)][start.x];
  if (["|", "L", "J"].includes(down)) {
    pipes = [...pipes, { x: start.x, y: start.y + 1 }];
  }
  const left = grid[start.y][Math.max(0, start.x - 1)];
  if (["-", "L", "F"].includes(down)) {
    pipes = [...pipes, { x: start.x - 1, y: start.y }];
  }
  const right = grid[start.y][Math.min(grid[start.y].length - 1, start.x + 1)];
  if (["-", "J", "7"].includes(right)) {
    pipes = [...pipes, { x: start.x + 1, y: start.y }];
  }

  return pipes;
};

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const grid = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  const start = findStart(grid);
  const pipes = findFirstPipePositions(grid, start);
  const furthestSteps = findFurthestPosition(grid, [start, start], pipes);

  return furthestSteps;
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const grid = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => line.split(""));

  console.log(grid);

  return 0;
};

const expectedTestAnswer = [8, 2];

const answerOneTest = one("test.txt");
console.log("Answer one test", answerOneTest);
console.log(
  `Answer is ${answerOneTest === expectedTestAnswer[0] ? "correct" : "wrong"}`
);
const answerOne = one("input.txt");
console.log("Answer one", answerOne);

/* const answerTwoTest = two("test.txt");
console.log("Answer two test", answerTwoTest);
console.log(
  `Answer is ${answerTwoTest === expectedTestAnswer[1] ? "correct" : "wrong"}`
);

const answerTwo = two("input.txt");
console.log("Answer two", answerTwo);
 */
