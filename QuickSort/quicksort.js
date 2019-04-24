// 采用ES6标准
let array = [5, 9, 3, 6, 11];

function quickSort(left, right) {

    if (left > right) {
        return;
    }

    let i = left, j = right, temp = array[left];

    while (i < j) {

        while (i < j && array[j] >= temp) {
            j--;
        }

        while (i < j && array[i] <= temp) {
            i++;
        }

        if (i < j) {
			[array[i], array[j]] = [array[j], array[i]];
        }
    }



    [array[i], array[left]] = [array[left], array[i]];

    quickSort(left, i-1);
    quickSort(i+1, right);
}

quickSort(0, array.length-1);
console.log(array);