const fs = require('fs');

const one = (inputFile) => {
    const allFileContents = fs.readFileSync(inputFile, 'utf-8');
    
    const calibrationValues = allFileContents.trim().split(/\r?\n/).map((line) => {
        let first, last;
        line.split("").forEach(char => {
            if(parseInt(char)){
                if(!first){
                    first = char;
                }
                last = char;
            }
        })
        return parseInt(first + last);
    });

    const sum = calibrationValues.reduce((sum, current) => sum + current, 0);

    return sum;
}

const two = (inputFile) => {
    const numbersRegex = /one|two|three|four|five|six|seven|eight|nine/g
    const numbersAsStringsToNumbers = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9
    }
    const allFileContents = fs.readFileSync(inputFile, 'utf-8');

    const calibrationValues = allFileContents.trim().split(/\r?\n/).map((line) => {
        let first, last, firstIndex = 9999, lastIndex = 0;
        line.split("").forEach((char, index) => {
            if(parseInt(char)){
                if(!first){
                    first = char;
                    firstIndex = index;
                }
                last = char;
                lastIndex = index;
            }
        })
        
        const firstAsLettersIndex = line.search(numbersRegex);
       
        if(firstAsLettersIndex > -1 && firstAsLettersIndex < firstIndex){
            const firstAsLetters = line.match(numbersRegex)[0];
            first = numbersAsStringsToNumbers[firstAsLetters];
        }

        if(firstAsLettersIndex > -1){
            const findLastAsLetters = () => {
                for (let i = line.length; i >= lastIndex; i--) {
                    const foundNumber = line.slice(i-1).search(numbersRegex) > -1;
                    if(foundNumber) {
                        const numAsLetter = line.slice(i-1).match(numbersRegex)[0];
                        return numbersAsStringsToNumbers[numAsLetter];
                    }
                }
            }
            const lastAsLetter = findLastAsLetters();
            if(lastAsLetter){
                last = lastAsLetter;
            }
            
        }

        return parseInt("" + first + last);
    });

    const sum = calibrationValues.reduce((sum, current) => sum + current, 0);
 
    return sum;
}

const expectedTestAnswer = [142, 281]

const answerOneTest = one('test.txt');
console.log("Answer one test", answerOneTest);
console.log(`Answer is ${answerOneTest === expectedTestAnswer[0] ? 'correct' : 'wrong'}`)
const answerOne = one('input.txt');
console.log("Answer one", answerOne);

const answerTwoTest = two('test2.txt');
console.log("Answer two test", answerTwoTest);
console.log(`Answer is ${answerTwoTest === expectedTestAnswer[1] ? 'correct' : 'wrong'}`)
const answerTwo = two('test2.txt');
console.log("Answer two", answerTwo);