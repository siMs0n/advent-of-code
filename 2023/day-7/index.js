const fs = require("fs");

const TYPES = {
  FiveOfAKind: 7,
  FourOfAKind: 6,
  FullHouse: 5,
  ThreeOfAKind: 4,
  TwoPair: 3,
  OnePair: 2,
  HighCard: 1,
};

const hand = {
  hand: "AAQQQ",
  type: "",
  bid: 123,
};

const one = (inputFile) => {
  const faceCardValue = {
    A: 14,
    K: 13,
    Q: 12,
    J: 11,
    T: 10,
  };

  const getType = (hand) => {
    const groupsMap = {};
    hand
      .split("")
      .forEach((c) => (groupsMap[c] = groupsMap[c] ? groupsMap[c] + 1 : 1));

    const groups = Object.values(groupsMap);
    if (groups.length === 1) return TYPES.FiveOfAKind;
    if (groups.length === 5) return TYPES.HighCard;
    if (groups.length === 4) return TYPES.OnePair;

    if (groups.length === 3) {
      // Either two pair or three of a kind
      return groups.some((g) => g === 3) ? TYPES.ThreeOfAKind : TYPES.TwoPair;
    }

    if (groups.length === 2) {
      // Either four of a kind or full house
      return groups.some((g) => g === 4) ? TYPES.FourOfAKind : TYPES.FullHouse;
    }
  };

  const sortHands = (a, b) => {
    if (a.type === b.type) {
      // Go card by card, when one is higher return
      for (let i = 0; i < a.hand.length; i++) {
        if (a.hand[i] !== b.hand[i]) {
          return (
            (faceCardValue[a.hand[i]] || a.hand[i]) -
            (faceCardValue[b.hand[i]] || b.hand[i])
          );
        }
      }
      return 0;
    }
    return a.type - b.type;
  };

  const allFileContents = fs.readFileSync(inputFile, "utf-8");

  // Parse to hands
  // Sort hands, first on type then on first highest card
  // Loop through whole array and multiply bid by index + 1 (rank)

  const hands = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [hand, bidAsStr] = line.split(" ");
      return {
        hand,
        type: getType(hand),
        bid: parseInt(bidAsStr),
      };
    });

  hands.sort(sortHands);

  const totalWinnings = hands.reduce(
    (sum, current, index) => sum + current.bid * (index + 1),
    0
  );

  return totalWinnings;
};

const two = (inputFile) => {
  const faceCardValue = {
    A: 14,
    K: 13,
    Q: 12,
    J: 1,
    T: 10,
  };

  const getTypeWithJoker = (groups, nrOfJokers) => {
    if (nrOfJokers === 5) return TYPES.FiveOfAKind;
    if (nrOfJokers === 4) return TYPES.FiveOfAKind;
    if (nrOfJokers === 3) {
      return groups.length === 2 ? TYPES.FiveOfAKind : TYPES.FourOfAKind;
    }
    if (nrOfJokers === 2) {
      return groups.length === 2
        ? TYPES.FiveOfAKind
        : groups.length === 3
        ? TYPES.FourOfAKind
        : TYPES.ThreeOfAKind;
    }
    //Four others are same
    if (groups.length === 2) return TYPES.FiveOfAKind;

    if (groups.length === 3) {
      // Three others are same or two pairs
      return groups.some((g) => g === 3) ? TYPES.FourOfAKind : TYPES.FullHouse;
    }
    // Two others are same
    if (groups.length === 4) return TYPES.ThreeOfAKind;

    return TYPES.OnePair;
  };

  const getType = (hand) => {
    const groupsMap = {};
    hand
      .split("")
      .forEach((c) => (groupsMap[c] = groupsMap[c] ? groupsMap[c] + 1 : 1));

    const groups = Object.values(groupsMap);

    if (Object.keys(groupsMap).includes("J"))
      return getTypeWithJoker(groups, groupsMap["J"]);

    if (groups.length === 1) return TYPES.FiveOfAKind;
    if (groups.length === 5) return TYPES.HighCard;
    if (groups.length === 4) return TYPES.OnePair;

    if (groups.length === 3) {
      // Either two pair or three of a kind
      return groups.some((g) => g === 3) ? TYPES.ThreeOfAKind : TYPES.TwoPair;
    }

    if (groups.length === 2) {
      // Either four of a kind or full house
      return groups.some((g) => g === 4) ? TYPES.FourOfAKind : TYPES.FullHouse;
    }
  };

  const sortHands = (a, b) => {
    if (a.type === b.type) {
      // Go card by card, when one is higher return
      for (let i = 0; i < a.hand.length; i++) {
        if (a.hand[i] !== b.hand[i]) {
          return (
            (faceCardValue[a.hand[i]] || a.hand[i]) -
            (faceCardValue[b.hand[i]] || b.hand[i])
          );
        }
      }
      return 0;
    }
    return a.type - b.type;
  };

  const allFileContents = fs.readFileSync(inputFile, "utf-8");
  const hands = allFileContents
    .trim()
    .split(/\r?\n/)
    .map((line) => {
      const [hand, bidAsStr] = line.split(" ");
      return {
        hand,
        type: getType(hand),
        bid: parseInt(bidAsStr),
      };
    });

  hands.sort(sortHands);

  const totalWinnings = hands.reduce(
    (sum, current, index) => sum + current.bid * (index + 1),
    0
  );

  return totalWinnings;
};

const expectedTestAnswer = [6440, 5905];

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
