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

class LLNode
  attr_accessor :next
  attr_reader :value

  def initialize(input)
    @value = input
    @next = nil
  end
end

class LinkedList
  attr_accessor :head
  def initialize
    @head = nil
  end

  # 1. Add New Node
  def add_node(input)
    if !@head
      @head = LLNode.new(input)
    else
      temp = @head
      while temp.next
        temp = temp.next
      end
      temp.next = LLNode.new(input)
    end
  end
end

def makeLinkedList(node)
  queue = [[node,0]]
  results = []
  while queue.length > 0
    current = queue.shift
    if !results[current[1]]
      results[current[1]] = LinkedList.new
    end
    results[current[1]].add_node(current[0].data)
    if current[0].left
      queue.push([current[0].left,current[1]+1])
    end
    if current[0].right
      queue.push([current[0].right,current[1]+1])
    end
  end
  return results
end

p makeLinkedList(a)
# [50->nil,6->1,4->2->8,9->nil]