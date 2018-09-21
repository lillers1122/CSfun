def mergeSorted(arr1,arr2)
  arr3 = []
  i = 0
  j = 0
  while i < arr1.length && j < arr2.length
      if arr1[i] <= arr2[j]
        arr3 << arr1[i]
        i += 1
      elsif arr2[j] < arr1[i]
        arr3 << arr2[j]
        j += 1
      end
  end

  while i < arr1.length
    arr3 << arr1[i]
    i += 1
  end

  while j < arr2.length
    arr3 << arr2[j]
    j += 1
  end
  return arr3
end

def mergeSort(arr)
  if arr.length < 2
    return arr
  else
    mid = (arr.length)/2
    l = arr[0,mid]
    r = arr[mid,arr.length]
    return mergeSorted(mergeSort(l), mergeSort(r))
  end
end

myArray = [4,3,7,8,1]
p "Merge sort: " + myArray.to_s + " => " + mergeSort(myArray).to_s

array1 = [1,2,3]
array2 = [2,3,4]
p "Merge 2 sorted arrays: " + array1.to_s + " + " + array2.to_s + " => " + mergeSorted(array1,array2).to_s