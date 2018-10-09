#given n steps, return all combinations of 1,2,3 steps
def recursiveSteps(x,n,current,valid)
  if n  == 0 && !valid[current]
    valid[current]  = true
  elsif n > 0
    current2 = ""
    if x > 0
      current2 = current + x.to_s
      n -= x
    end

    recursiveSteps(1,n,current2,valid)
    recursiveSteps(2,n,current2,valid)
    recursiveSteps(3,n,current2,valid)
  end
  valid.keys
end

p recursiveSteps(0,4,"",{})