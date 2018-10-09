symbols = ["Ac", "Ag", "Al", "Am", "Ar", "As", "At", "Au", "B", "Ba", "Be", "Bh", "Bi", "Bk", "Br", "C", "Ca", "Cd", "Ce", "Cf", "Cl", "Cm", "Co", "Cr", "Cs", "Cu", "Db", "Ds", "Dy", "Er", "Es", "Eu", "F", "Fe", "Fm", "Fr", "Ga", "Gd", "Ge", "H", "He", "Hf", "Hg", "Ho", "Hs", "I", "In", "Ir", "K", "Kr", "La", "Li", "Lr", "Lu", "Md", "Mg", "Mn", "Mo", "Mt", "N", "Na", "Nb", "Nd", "Ne", "Ni", "No", "Np", "O", "Os", "P", "Pa", "Pb", "Pd", "Pm", "Po", "Pr", "Pt", "Pu", "Ra", "Rb", "Re", "Rf", "Rg", "Rh", "Rn", "Ru", "S", "Sb", "Sc", "Se", "Sg", "Si", "Sm", "Sn", "Sr", "Ta", "Tb", "Tc", "Te", "Th", "Ti", "Tl", "Tm", "U", "Uub", "Uuh", "Uuo", "Uup", "Uuq", "Uus", "Uut", "V", "W", "Xe", "Y", "Yb", "Zn", "Zr"]

def make_dictionary(arr)
  dictionary = {}
  arr.each {|elem| dictionary[elem.downcase] = elem.downcase }
  return dictionary
end

def bold_elements(my_str, guide)
  str = my_str.split("")
  results = ""
  i = 0
  while i < str.length
    if guide[str[i]]
      results += "<strong>" + guide[str[i]] + "</strong>"
    elsif str[i+1] && guide[(str[i] + str[i + 1])]
      results += "<strong>" + guide[(str[i] + str[i + 1])] + "</strong>"
      i += 1
    else
      results += str[i]
    end
    i += 1
  end
  results
end

dict = make_dictionary(symbols)
phrase = "cats are great!"
p bold_elements(phrase,dict)
