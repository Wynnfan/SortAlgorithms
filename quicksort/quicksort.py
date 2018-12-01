
array = [5, 9, 3, 6, 1]

def quickSort(left, right):
	
	if left > right:
		return
	
	i, j = left, right
	temp = array[left]
	
	while i < j:
		
		while i < j and array[j] >= temp:
			j -= 1
		while i < j and array[i] <= temp:
			i += 1
		if i < j:
			array[i], array[j] = array[j], array[i]
			
	array[left], array[i] = array[i], array[left]
	
	quickSort(left, i-1)
	quickSort(i+1, right)
	
quickSort(0, len(array)-1)
print(array)