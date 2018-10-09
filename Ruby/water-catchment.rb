topo = [1,0,2,1,1,3,0,3] #=> 3

def catchment(arr)
  total = 0
  i = 0
  j = 1
  sum = 0
  while i != arr.length && j != arr.length
    if arr[j] >= arr[i]
      total += sum
      sum = 0
      i = j
      j += 1
    elsif arr[j] < arr[i]
      sum += (arr[i] - arr[j])
      j += 1
    end
  end
  return total
end

p catchment(topo)

