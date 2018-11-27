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
# a.left.left = Node.new(4)
# a.left.right = Node.new(15)
# a.left.right.left = Node.new(9)
a.right = Node.new(51)
# a.right.right = Node.new(80)

def balancedBST?(node)
  return false if node.nil?
  return balanced(node) != -1
end

def balanced(node)
  return 0 if node.nil?

  leftHeight = balanced(node.left)
  rightHeight = balanced(node.right)

  if leftHeight == -1 || rightHeight == -1 || (leftHeight - rightHeight).abs > 1
    return -1
  end
  if rightHeight > leftHeight
    return (1 + rightHeight)
  else
    return (1 + leftHeight)
  end
end

p balancedBST?(a)