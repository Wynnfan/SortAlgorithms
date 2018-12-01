
var array = [5, 9, 3, 6, 11];

function quickSort(left, right) {

    if (left > right) {
        return;
    }

    var i = left, j = right, temp = array[left];
    var t;

    while (i < j) {

        while (i < j && array[j] >= temp) {
            j--;
        }

        while (i < j && array[i] <= temp) {
            i++;
        }

        if (i < j) {
            t = array[i];
            array[i] = array[j];
            array[j] = t;
        }
    }



    t = array[left];
    array[left] = array[i];
    array[i] = t;

    quickSort(left, i-1);
    quickSort(i+1, right);
}

quickSort(0, array.length-1);
console.log(array);