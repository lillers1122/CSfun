# Given a map of water, 0s, and islands, 1s, as a matrix, return the number of islands
myMap = [[0,0,1,0],
         [0,1,1,0],
         [0,0,0,0],
         [1,1,0,1]]

def count_islands(matrix)
  i = 0
  j = 0
  islands = 0
  while i < matrix.length
    while j < matrix[0].length
      if matrix[i][j] == 1
        islands += 1
        q_erase_helper(i,j,matrix)
      else
        j += 1
      end
    end
    j = 0
    i += 1
  end
  p matrix
  return islands
end

def s_erase_helper(c1,c2,myMatrix) #iterative with stack
  stack = [[c1,c2]]

  while !stack.empty?
    current = stack.pop
    myMatrix[current[0]][current[1]] = 0
    directions = [[1,0],[-1,0],[0,1],[0,-1]]

    directions.each do |coords|
      coords[0] += current[0]
      coords[1] += current[1]
      if coords[0] > -1 && coords[0] < myMatrix.length && coords[1] > -1 && coords[1] < myMatrix.length && myMatrix[coords[0]][coords[1]] == 1
        stack.push([coords[0],coords[1]])
        p stack.length
      end
    end
  end
  myMatrix
end

def q_erase_helper(c1,c2,myMatrix) #iterative with queue
  queue = [[c1,c2]]
  while !queue.empty?
    current = queue.shift
    myMatrix[current[0]][current[1]] = 0
    directions = [[1,0],[-1,0],[0,1],[0,-1]]

    directions.each do |coords|
      coords[0] += current[0]
      coords[1] += current[1]
      if coords[0] > -1 && coords[0] < myMatrix.length && coords[1] > -1 && coords[1] < myMatrix.length && myMatrix[coords[0]][coords[1]] == 1
        queue.push([coords[0],coords[1]])
        p queue
      end
    end
  end
  myMatrix
end

def r_erase_helper(c1,c2,myMatrix) #recursive
  directions = [[1,0],[-1,0],[0,1],[0,-1]]
  myMatrix[c1][c2] = 0

  directions.each do |coords|
    coords[0] += c1
    coords[1] += c2
    if coords[0] > -1 && coords[0] < myMatrix.length && coords[1] > -1 && coords[1] < myMatrix.length && myMatrix[coords[0]][coords[1]] == 1
      r_erase_helper(coords[0],coords[1],myMatrix)
    end
  end
end

p count_islands(myMap)
# p erase_helper(0,2,myMap)
