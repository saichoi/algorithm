// 코드에서 괄호와 따옴표를 추출하는 함수
function bracketPicker(code) {
    let codeString = code.split("");
    let brackets = [];
    for (let i = 0; i < codeString.length; i++) {
        if (codeString[i] === '{' || codeString[i] === '}' ||
            codeString[i] === '(' || codeString[i] === ')' ||
            codeString[i] === '[' || codeString[i] === ']' ) {
            brackets.push(codeString[i])
        }
    }
    return brackets;
}

// 중괄호{}를 체크하는 함수
function braceChecker(codeString, codeMap) {
    // value에 들어갈 값을 선언한다.
    let braceCount = 0;
    codeMap.set('{}', 0);

    for (let i = 0; i < codeString.length; i++) {
        // 문자열에 '{'가 있을 경우 key가 '{}'인 value에 +1을 하고
        if (codeString[i] === '{') {
            braceCount++;
            codeMap.set('{}', braceCount);
        }
        // 문자열에 '}'가 있을 경우 key가 '{}'인 value에 -1을 한다.
        else if (codeString[i] === '}') {
            braceCount--;
            codeMap.set('{}', braceCount);
        } 
    }

    // 문자열을 한번씩 다 체크하고 나면 key: value로 담아둔 값을 확인하고 true인지 false인지 정한다.
    if (codeMap.get('{}') === 0) {
        return true;
    } else {
        return false;
    }
}

// 소괄호()를 체크하는 함수
function parenthesesChecker(codeString, codeMap) {
    // value에 들어갈 값을 선언한다.
    let parenthesesCount = 0;
    codeMap.set('()', 0);

    for (let i = 0; i < codeString.length; i++) {
        // 문자열에 '('가 있을 경우 key가 '()'인 value에 +1을 하고
        if (codeString[i] === '(') {
            parenthesesCount++;
            codeMap.set('()', parenthesesCount);
        }
        // 문자열에 ')'가 있을 경우 key가 '())'인 value에 -1을 한다.
        else if (codeString[i] === ')') {
            parenthesesCount--;
            codeMap.set('()', parenthesesCount);
        }
    }

    // 문자열을 한번씩 다 체크하고 나면 key: value로 담아둔 값을 확인하고 true인지 false인지 정한다.
    if (codeMap.get('()') === 0) {
        return true;
    } else {
        return false;
    }    
}

// 대괄호[]를 체크하는 함수
function bracketsChecker(codeString, codeMap) {
    // value에 들어갈 값을 선언한다.
    let bracketsCount = 0;
    codeMap.set('[]', 0);

    for (let i = 0; i < codeString.length; i++) {
        // 문자열에 '['가 있을 경우 key가 '[]'인 value에 +1을 하고
        if (codeString[i] === '[') {
            bracketsCount++;
            codeMap.set('[]', bracketsCount);
        }
        // 문자열에 ')'가 있을 경우 key가 '[]'인 value에 -1을 한다.
        else if (codeString[i] === ']') {
            bracketsCount--;
            codeMap.set('[]', bracketsCount);
        }
    }

    // 문자열을 한번씩 다 체크하고 나면 key: value로 담아둔 값을 확인하고 true인지 false인지 정한다.
    if (codeMap.get('[]') === 0) {
        return true;
    } else {
        return false;
    }    
}

function codeChecker(code) {
    // 1. 문자열에서 {}, ()을 담는다.
    let codeString = bracketPicker(code);

    // 2. key: value 형태의 자료구조를 만들어준다.
    const codeMap = new Map();

    // 3. 문자열에 {}, ()를 하나씩 짝이 맞는지 검사하여 true 혹은 false를 반환한다.
    let isBraceOk = braceChecker(codeString, codeMap);
    let isParenthesesOk = parenthesesChecker(codeString, codeMap);
    let isBracketsOk = bracketsChecker(codeString, codeMap);
    if (isBraceOk && isParenthesesOk && isBracketsOk) {
        console.log('true');
        return true;
    } else {
        console.log('false');
        return false;
    }
}

// 코드 검사 함수 테스트
let code1 = 'function(){{}}'; // true
let code2 = 'function(){{'; // false
let code3 = 'function()'; // true
let code4 = 'function(()'; // false
let code5 = 'let test = [];' // true
let code6 = 'let test = []];' // false
codeChecker(code1);
codeChecker(code2);
codeChecker(code3);
codeChecker(code4);
codeChecker(code5);
codeChecker(code6);