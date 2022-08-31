// 선형 검색
function linearSearch(array, value) {
  for (let i = 0; i < array.length; i++) {
    if (array[i] === value) return i;
  }
  return '찾는 값이 없습니다.';
}

// 이진 검색
function binarySearch(array, target) {
    // 1. 이진 검색의 전제 조건인 정렬된 배열을 만든다.
    array = array.sort(function(a,b){
        return a - b;
    });

    // 2. 소거해 나갈 배열의 길이를 정의하기 위해 시작값과 끝값을 기억한다.
    let start = 0;
    let end = array.length - 1;

    // 3. 중간 값을 정의한다.
    let middleIndex = Math.floor((end - start) / 2);

    // 4. 서치가 완료될 때까지 반복문으로 돌린다.
    while(true){
        // 5. 중간 값과 타겟 값이 같다면 몇번째 배열에 타겟값이 있는지 로그에 찍는다.
        if(array[middleIndex] === target) {
            return middleIndex;
        }
        // 7-1. 모든 비교가 완료되고 마지막으로 시작값과 끝값이 같아지면, 즉 값이 하나 남으면 마지막으로 타겟값과 중간값을 비교한다.
        if(start === end) {
            // 7-2. 타겟값과 중간 값이 같으면 해당 값이 몇번째 배열에 있는지 로그에 찍고, 같지 않다면 해당 값이 없다는 로그를 찍는다.
            return array[middleIndex] === target ? start : '찾는 값이 없습니다.';
        }
        // 6-1. 타겟값과 중간값을 비교하여 중간값이 작다면 중간값보다 작은 값을 버린다.
        if(target>array[middleIndex]) {
            // 중간값보다 작은 값을 버리기 때문에 시작값의 인덱스가 중간값에서 +1을 한 값이 된다.
            start = middleIndex + 1;
        }
        // 6-2. 타겟값과 중간값을 비교하여 중간값이 크다면 중간값보다 큰 값을 버린다.
        if(target<array[middleIndex]) {
            // 중간값보다 큰 값을 버리기 때문에 끝 값의 인덱스가 중간 값에서 -1을 한 값이 된다.
            end = middleIndex - 1;
        }
        // 한 바퀴를 돌 때마다 중간값이 바뀌므로 중간 값을 변경해준다. start를 더하는 이유는 최초에 정의된 인덱스 값을 기억하기 위해서이다.
        middleIndex = Math.floor((end-start)/2) + start;
    }
}

const example = [1, 3, 5, 7, 9, 11, 13];
const example2 = [3, 6, 1, 2, 4, 7, 0, 5, 9, 4];

// 선형 검색 테스트
console.log(linearSearch(example, 1)); // 0
console.log(linearSearch(example, 13)); // 6
console.log(linearSearch(example, 2)); // 찾는 값이  없습니다.
console.log(linearSearch(example2, 0)); // 0
console.log(linearSearch(example2, 9)); // 9
console.log(linearSearch(example2, 13)); // 찾는 값이 없습니다.

// 이진 검색 테스트
console.log(binarySearch(example, 1)) // 0
console.log(binarySearch(example, 13)) // 6
console.log(binarySearch(example, 2)) // 찾는 값이 없습니다.
console.log(binarySearch(example2, 0)) // 0
console.log(binarySearch(example2, 9)) // 9
console.log(binarySearch(example2, 13)) // 찾는 값이 없습니다.
