def condense(string)
  key = {}
  string.each_char do |char|
    key[char] ? key[char] += 1 : key[char] = 1
  end

  results = ""
  key.each do |k,v|
    results += "#{k}" + "#{v}"
  end
  return puts results
end

condense("aaaabbbcccddeeeed")