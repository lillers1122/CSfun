def rotate(mat)
  i = 0
  while i < mat.length
    j = i
    while j < mat[0].length #could be using mat[0].length; same size
      mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
      j += 1
    end
    i += 1
  end
  return mat
end

matrix = [[1,1,1,1],
          [2,2,2,2],
          [3,3,3,3],
          [4,4,4,4]]

p rotate(matrix)