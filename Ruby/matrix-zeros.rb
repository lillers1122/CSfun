def matrixZeros(mat)
    return if mat.nil?

    rows, columns = toChangeHelper(mat)
    rows.each do |i|
        j = 0
        while j < mat.length
            mat[i][j] = 0
            j += 1
        end
    end

    columns.each do |j|
        i = 0
        while i < mat.length
            mat[i][j] = 0
            i += 1
        end
    end

    return mat
end

def toChangeHelper(mat)
    rows = []
    columns = []
    i = 0
    while i < mat.length
        j = 0
        while j < mat[0].length
            if mat[i][j] == 0
                rows.push(i)
                columns.push(j)
            end
            j += 1
        end
        i += 1
    end
    return rows, columns
end

matrix = [[0,1,1,0],
          [1,1,0,1],
          [1,1,1,1]]

matrix2 = [[1,1,0,1],
          [2,2,2,2],
          [3,3,3,3],
          [4,0,4,4]]

puts ""
p matrixZeros(matrix)
puts ""
p matrixZeros(matrix2)
puts ""