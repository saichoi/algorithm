// 버블 정렬 (오름차순)
function AscBubbleSort(array) {
    for (let i = 0; i < array.length - 1; i++) {
        for (let j = 0; j < array.length - 1; j++) { 
            if (array[j] > array[j + 1]) {
                // 서로 값을 교환하기 위해 일단 다른 변수에 값을 담아둔다.
                let temp = array[j];
                array[j] = array[j + 1];
                array[j + 1] = temp;
            }
        }
    }
    console.log('array', array);
    return array;
}

// 버블 정렬 (내림차순)
function DescBubbleSort(array) {
    for (let i = 0; i < array.length; i++) {
        for (let j = 0; j < array.length - 1; j++) { 
            if (array[j] < array[j + 1]) {
                let temp = array[j];
                array[j] = array[j + 1];
                array[j + 1] = temp;
            }
        }
    }
    console.log('array', array);
    return array;
}

const array = [5, 4, 3, 2, 1];
AscBubbleSort(array);
DescBubbleSort(array);