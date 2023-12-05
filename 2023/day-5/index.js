const fs = require("fs");

const map = {
  category: "",
  to: "",
  ranges: [{ destinationRangeStart: 0, sourceRangeStart: 0, rangeLength: 1 }],
};

const findCategoryNumberInRanges = (sourceCategoryNr, ranges) => {
  const range = ranges.find(
    (range) =>
      range.sourceRangeStart <= sourceCategoryNr &&
      range.sourceRangeStart + range.rangeLength >= sourceCategoryNr
  );
  if (range) {
    const diff = range.sourceRangeStart - range.destinationRangeStart;
    return sourceCategoryNr - diff;
  }
  // is same nr
  return sourceCategoryNr;
};

const findCategoryNr = (maps, category, goalCategory, sourceCategoryNr) => {
  if (category === goalCategory) return sourceCategoryNr;

  const map = maps.find((m) => m.category === category);
  const categoryNumber = findCategoryNumberInRanges(
    sourceCategoryNr,
    map.ranges
  );
  return findCategoryNr(maps, map.to, goalCategory, categoryNumber);
};

const findSeedLocation = (maps, seedNr) => {
  return findCategoryNr(maps, "seed", "location", seedNr);
};

const parseMaps = (mapLines) => {
  const [categoryLine, ...rangeLines] = mapLines.split(/\r?\n/);
  const [category, to] = categoryLine.split("map")[0].trim().split("-to-");

  const ranges = rangeLines.map((line) => {
    const [destinationRangeStart, sourceRangeStart, rangeLength] =
      line.split(" ");
    return {
      destinationRangeStart: parseInt(destinationRangeStart),
      sourceRangeStart: parseInt(sourceRangeStart),
      rangeLength: parseInt(rangeLength),
    };
  });

  return {
    category,
    to,
    ranges,
  };
};

const one = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  const [seedLine, ...mapLines] = allFileContents.trim().split(/\r?\n\n/);

  const seeds = seedLine.split("seeds: ")[1].split(" ");

  const maps = mapLines.map(parseMaps);

  const seedLocations = seeds.map((s) => findSeedLocation(maps, s));
  seedLocations.sort((a, b) => a - b);

  return seedLocations[0];
};

const two = (inputFile) => {
  const allFileContents = fs.readFileSync(inputFile, "utf-8");
  const [seedLine, ...mapLines] = allFileContents.trim().split(/\r?\n\n/);

  const seeds = seedLine.split("seeds: ")[1].split(" ");

  const maps = mapLines.map(parseMaps);

  let smallest = Number.MAX_SAFE_INTEGER;

  //Naive solution, takes 10 mins
  const seedPairs = seeds.reduce(function (result, value, index, array) {
    if (index % 2 === 0)
      result.push(array.slice(index, index + 2).map((s) => parseInt(s)));
    return result;
  }, []);

  seedPairs.forEach(([start, rangeLength]) => {
    for (let i = 0; i <= rangeLength - 1; i++) {
      const location = findSeedLocation(maps, start + i);
      if (location < smallest) {
        smallest = location;
      }
    }
  });

  return smallest;
};

const expectedTestAnswer = [35, 46];

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
