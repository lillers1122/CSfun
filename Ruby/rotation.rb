def isSubstring?(string1,string2)
  false if string1.length != string2.length
  i = 0
  while i < string1.length
    string1 = string1[-1] + string1[0..-2]
    if string1 == string2
      return true
    end
    i += 1
  end
  false
end

p isSubstring?("cat", "tac") #false
p isSubstring?("cat", "atc") #true
p isSubstring?("waterbottle", "erbottlewat") #true

