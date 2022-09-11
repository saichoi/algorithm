// 중복 검사 알고리즘 1
// 빅 오 표기법 : O(N²)
function hasDuplicatedValue1(array) {
    for (let i = 0; i < array.length; i++) {
        for (let j = 0; j < array.length; j++) {
            if (i != j && array[i] === array[j]) {
                console.log('true');
                return true;
            }
        }
    }
    console.log('false');
    return false;    
}

const array = [1, 2, 3, 4, 5, 6, 6];
hasDuplicatedValue1(array);

// 중복 검사 알고리즘 2 (책깔피 기법)
// 빅 오 표기법 : O(N)
function hasDuplicatedValue2(array) {
    const existNumber = new Array(100); // 100개 배열 생성
    existNumber.fill(0); // 배열을 0으로 초기화
    for (let i = 0; i < array.length; i++) {
        // 처음 들어오는 인덱스에는 초기화해둔 값 0이 들어가 있으므로 
        // 이를 1로 바꿔 같은 값이 있다는 것을 체크해 둔다.
        if (existNumber[array[i]] === 0) { 
            existNumber[array[i]] = 1;
        } else {
            console.log('true');
            return true;
        }
    }
    console.log('false');
    return false;    
}

const array2 = [1, 2, 3, 4, 5, 6, 6];
hasDuplicatedValue2(array2);