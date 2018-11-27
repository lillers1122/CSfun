class Node
  attr_accessor :left, :right
  attr_reader :data

  def initialize(input)
    @data = input
    @left = nil
    @right = nil
  end
end

a = Node.new(50)
a.left = Node.new(6)
a.left.left = Node.new(4)
a.left.right = Node.new(2)
a.left.right.left = Node.new(9)
a.right = Node.new(1)
a.right.right = Node.new(8)

def validBST?(node)
  return [] if node.nil?

  nodeValues = inOrderArray(node)
  i = 0
  j = 1
  while j < nodeValues.length
    if nodeValues[i] > nodeValues[j]
      return false
    end
    i += 1
    j += 1
  end
  return true
end

def inOrderArray(node,results=[])
  return results if node.nil?
  inOrderArray(node.left,results)
  results << node.data
  inOrderArray(node.right,results)
end

p validBST?(a)