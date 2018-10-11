myMatrix = [['z','z','z','y'],
            ['z','y','y','z'],
            ['z','y','z','z'],
            ['z','y','y','y']]

def flood_fill(matrix, coor1, coor2, v)
  previous = matrix[coor1][coor2]
  matrix[coor1][coor2] = v

  directions = [[0,1],[0,-1],[1,0],[-1,0]] #R,L,B,T

  directions.each do |coords|
    coords[0] += coor1
    coords[1] += coor2
    if matrix[coords[0]][coords[1]] == previous
      flood_fill(matrix,coords[0],coords[1],v)
    end
  end
  matrix
end
p myMatrix
p flood_fill(myMatrix,2,2,"a")