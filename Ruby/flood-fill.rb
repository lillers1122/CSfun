myMatrix = [['z','z','z','y'],
            ['z','y','y','z'],
            ['z','y','z','z'],
            ['z','y','y','y']]

def floodFill(matrix,c1,c2,change)
  toChange = matrix[c1][c2]
  stack = [[c1,c2]]

  while !stack.empty?
    coords = stack.pop
    coor1,coor2 = coords[0], coords[1]

    #change current
    matrix[coor1][coor2] = change

    #check below
    if coor1+1 < matrix.length && matrix[coor1+1][coor2] == toChange
      stack.push([coor1+1,coor2])
    end

    #check above
    if coor1-1 > -1 && matrix[coor1-1][coor2] == toChange
      stack.push([coor1-1,coor2])
    end

    #check right
    if coor2+1 < matrix[0].length && matrix[coor1][coor2+1] == toChange
      stack.push([coor1,coor2+1])
    end

    #check left
    if coor2-1 > -1 && matrix[coor1][coor2-1] == toChange
      stack.push([coor1,coor2-1])
    end
  end
  return matrix
end

p floodFill(myMatrix,2,2,"a")
