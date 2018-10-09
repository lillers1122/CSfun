#input a,b,c arrays
# output 1,2,3,6,7
a = [1,2,3,4,5,6,7]
b = [1,2,3,6,7,8,9,10]
c = [1,2,3,4,5,6,7,10,11,12,13,14]

def inThree(arr1,arr2,arr3)
  dict = {}
  arr1.each { |item| dict[item] = true }
  p dict
  dict2 = inArr?(arr2,dict)
  dict = inArr?(arr3,dict2)
  dict
end

def inArr?(arr,dict)
  newDict = {}
  arr.each do |item|
    if dict[item]
      newDict[item] = true
    end
  end
  newDict
end

p inThree(a,b,c)