// 코드에서 괄호를 추출하는 함수
function bracketPicker(code, open, close) {
    let codeString = code.split("");
    let brackets = [];
    for (let i = 0; i < codeString.length; i++) {
        if (codeString[i] === open || codeString[i] === close) {
            brackets.push(codeString[i])
        }
    }
    return brackets;
}

// 중괄호{}를 체크하는 함수
function braceChecker(codeString, open, close) {
    // 1. stack으로 사용할 배열을 만든다.
    let stack = [];

    for (let i = 0; i < codeString.length; i++) {
        // 1. 문자열에 '{'가 있을 경우 stack에 넣어준다.
        if (codeString[i] === open) {
            stack.push(open)
        }
        // 2. 문자열에 '}'가 있을 경우
        else if (codeString[i] === close) {
            // 2-1. 짝이되는 괄호가 없으면 false를 반환한다.
            if (stack.length === 0) {
                return false;    
            // 2-2. 짝이되는 괄호가 있으면 stack에서 꺼낸다.
            } else {
                stack.pop();
            }
        } 
    }

    // 3. 배열을 모두 체크하고 나서 stack에 값이 없으면 true, 그 외에는 false를 반환한다.
    if (stack.length === 0) {
        return true;
    } else {
        return false;
    }
}

function codeChecker(code) {
    // 1. 문자열에서 괄호만 추출한 배열을 변수에 담는다.
    let codeString1 = bracketPicker(code, '{', '}');
    let codeString2 = bracketPicker(code, '(', ')');
    let codeString3 = bracketPicker(code, '[', ']');

    // 3. 문자열에 {}를 하나씩 짝이 맞는지 검사하여 true 혹은 false를 반환한다.
    let isBraceOk = braceChecker(codeString1, '{', '}');
    let isParenthesOk = braceChecker(codeString2, '(', ')');
    let isBracketOk = braceChecker(codeString3, '[', ']');
    if (isBraceOk && isParenthesOk && isBracketOk) {
        console.log('true');
        return true;
    } else {
        console.log('false');
        return false;
    }
}

// 코드 검사 함수 테스트
let code1 = '{{}}'; // true
let code2 = '{{{{}}'; // false
let code3 = '{}{}'; // true
let code4 = '{}}{{}'; // false

let code5 = '(())'; // true
let code6 = '((())'; // false
let code7 = '()()'; // true
let code8 = '())(()' // false

let code9 = '[[]]'; // true
let code10 = '[[[]]' // false
let code11 = '[][]'; // true
let code12 = '[]][[]'; // false

codeChecker(code1);
codeChecker(code2);
codeChecker(code3);
codeChecker(code4);
console.log('==========================');
codeChecker(code5);
codeChecker(code6);
codeChecker(code7);
codeChecker(code8);
console.log('==========================');
codeChecker(code9);
codeChecker(code10);
codeChecker(code11);
codeChecker(code12);
