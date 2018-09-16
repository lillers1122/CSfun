# Excel column conversion
def columnNum(str)
  guide = alphabetHelper
  counter = 0
  str.each_char do |char|
    val = guide[char.upcase]
    counter *= 26
    counter += val
  end
  counter
end

def alphabetHelper
  alphabet = {}
  ("A".."Z").to_a.each_with_index do |let,i|
    alphabet[let] = i+1
  end
  alphabet
end

def columnLet(num)
  guide = numberHelper
  result = ""
  while num >  26
    temp = num % 26
    result = guide[temp] + result
    num /= 26
  end
  guide[num%26] + result
end
def numberHelper
  alphabet = {}
  ("A".."Z").to_a.each_with_index do |let,i|
    alphabet[i+1] = let
  end
  alphabet
end

p columnNum("AAA")
p columnLet(45)

# O(n) Time
# O(1) Space