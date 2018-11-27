def polish_num(arr)
  stack = []
  operators = {
      "+" => true,
      "*" => true,
      "/" => true,
      "-" => true
  }

  arr.each do |item|
    if operators[item]
        j = stack.pop
        i = stack.pop
      case item
      when "+"
        stack.push(i+=j)
      when "-"
        stack.push(i-=j)
      when "*"
        stack.push(i*=j)
      when "/"
          if j == 0
            return "Divide by zero error"
          else
            stack.push(i/=j)
          end
      end
    else
      stack.push(item.to_i)
    end
  end
  return stack.pop
end

input = ["1","2","+","2","1","-","/"]
p polish_num(input)
