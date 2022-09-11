// 선택 정렬 (오름차순)
function AscSelectedSort(array) {
    for (let i = 0; i < array.length-1; i++) {
        let minIndex = i;
        for (let j = i; j < array.length; j++) {
            if (array[j] < array[minIndex]) {
                let temp = array[j];
                array[j] = array[minIndex];
                array[minIndex] = temp;
            }
        }
    }
    console.log('array', array);
    return array
}

const array1 = [5, 4, 3, 2, 1, 0];
AscSelectedSort(array1);

// 선택 정렬 (내림차순)
function DescSelectedSort(array) {
    for (let i = 0; i < array.length-1; i++) {
        let minIndex = i;
        for (let j = i; j < array.length; j++) {
            if (array[j] > array[minIndex]) {
                let temp = array[j];
                array[j] = array[minIndex];
                array[minIndex] = temp;
            }
        }
    }
    console.log('array', array);
    return array
}

const array2 = [0, 1, 2, 3, 4, 5];
DescSelectedSort(array2);