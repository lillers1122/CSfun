#arrToBST(sortedArr) -> balanced BST
# [1,2,3,4,5,6] => 4

class Node
  attr_accessor :left, :right
  attr_reader :data

  def initialize(input)
    @data = input
    @left = nil
    @right = nil
  end
end

def arrToBST(ar)
  return if !ar || ar.length == 0

  midIndex = ar.length/2
  treeNode = Node.new(ar[midIndex])

  treeNode.left = arrToBST(ar[0...midIndex])
  treeNode.right = arrToBST(ar[(midIndex+1)..-1])
  return treeNode
end

myArray = [1,2,3,4,5,6,7,8]
p arrToBST(myArray)