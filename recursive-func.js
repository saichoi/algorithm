// 1. 재귀함수를 사용하여 n! 구현
function factorial(number) {
    if(number <= 1){ // 0과 1일때는 재귀함수를 돌릴 필요가 없다.
        return number;
    }
    return number * factorial(number - 1);
    // 5가 들어오면 5 * factorial(4)로 함수를 한번 더 호출하고 4를 리턴한다.
    // 4가 들어오면 4 * factorial(3)로 함수를 한번 더 호출하고 3를 리턴한다.
    // 3가 들어오면 3 * factorial(2)로 함수를 한번 더 호출하고 2를 리턴한다.
    // 2가 들어오면 2 * factorial(1)로 함수를 한번 더 호출하고 if문에서 걸려 1을 리턴한다.
    // 이 과정을 다 연결하면 5 * 4 * 3 * 2 * 1이 된다.
}

console.log(factorial(5))

// 2. 재귀함수를 사용하여 피보나치 수열을 구현
// 1, 1, 2, 3, 5, 8, 13 ...
function fibonacci(index) {
    if (index < 3) { // 6, 5, 4, 3
        return 1
    }
    return fibonacci(index - 1) + fibonacci(index - 2)
    // fibonacci(6) = fibonacci(5) + fibonacci(4)
                  // fibonacci(5) = fibonacci(4) + fibonacci(3)                                   = 5 + 3 = 8
                                // fibonacci(4) = fibonacci(3) + fibonacci(2)                     = 3 + 2 = 5
                                              // fibonacci(3) = fibonacci(2) + fibonacci(1)       = 2 + 1 = 3
                                                            // fibonacci(2)                     = 1 + 1 = 2
                                                                          // fibonacci(1)       = 1
}

console.log(fibonacci(6))

// 3. 재귀함수를 사용하여 하노이의 탑 구현
let count = 0;
// from -> to ( A -> C )
function hanoi(num, from, other, to){
    if(num === 0) {
        return;
    }
    // from -> other ( A -> B ) 로 이동
    hanoi(num - 1, from, to, other);
    count++;
    // other -> to ( B -> C ) 로 이동
    hanoi(num - 1, other, from, to);
}

hanoi(5, 'A', 'B', 'C');
console.log(count);

//                                                      h(3, 'A', 'B', 'C');
//                       h(2, 'A', 'C', 'B');                                         h(2, 'B', 'A', 'C');
//         h(1, 'A', 'C', 'B');         h(1, 'B', 'A', 'C');         h(1, 'A', 'C', 'B');               h(1, 'B', 'A', 'C');
