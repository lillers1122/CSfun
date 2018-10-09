# in-order subsets:  cat -> [c, ca, ct, cat, at, a, t, '']

def subsets(str,item, res,i)
  if i == str.length
    res << item
    return res
  else
    subsets(str,item, res,i+1)
    subsets(str,item+=str[i], res,i+1)
  end
  return res
end

def in_order_subsets(string)
  p subsets(string, "",[],0)
end

in_order_subsets("cat")
