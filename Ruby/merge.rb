#merge sort takes in 2 sorted arrays

def merge(arr1,arr2)
  arr3 = []
  i = 0
  j = 0
  while i < arr1.length && j < arr2.length
      if arr1[i] < arr2[j]
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
  p arr3
end

array1 = [1,2,4]
array2 = [3,5,6,7]

merge(array1,array2)