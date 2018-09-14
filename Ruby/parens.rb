#RUBY - Generate Valid Combinations of n Parentheses pairs
# Given an integer, generate an array containing all pairs of valid parentheses.
# Parentheses are valid if (()()(()))() each pair opens and closes properly. For example, )( would NOT be valid! (Neither would ()))
# generateParens(1) => ["()"]
# generateParens(2) => ["()()","(())"]
# generateParens(3) => ["()()()", "()(())", "(())()", "(()())", "((()))"]


## STRING Solution
def generateParensString(n)
  if n > 0
    result =  _generateParensString(n, "", [], 0, 0)
    print "input of #{n}: " + result.to_s + "\n"
  end
end

def _generateParensString(n, item, valid, open, close)
  if close == n  # count of closed gets to n
    valid.push(item)
    return
  else
    if open > close
      updatedItem = item + ")"
      _generateParensString(n, updatedItem, valid, open, close+1)
    end
    if open < n
      updatedItem = item + "("
      _generateParensString(n, updatedItem, valid, open+1, close)
    end
  end
  return valid
end

## ARRAY solution
def generateParensArray(n)
  if n > 0
    result = _generateParensArray(n, [], [], 0, 0, 0)
    print "input of #{n}: " + result.to_s + "\n\n"
  end
end

def _generateParensArray(n, array, valid, i, open, close)
  if close == n # count of closed gets to n
    current = array.join("").to_s
    valid.push(current)
  else
    if open > close
      array[i] = ")"
      _generateParensArray(n, array, valid, i+1, open, close + 1)
    end
    if open < n
      array[i] = "("
      _generateParensArray(n, array, valid, i+1, open + 1, close)
    end
  end
  return valid
end

puts "\n- - - - - - String - - - - -\n"
generateParensString(2)
puts "\n- - - - - - Array - - - - - -\n"
generateParensArray(3)
