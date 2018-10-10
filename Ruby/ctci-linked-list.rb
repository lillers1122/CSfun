class Node
  attr_accessor :next, :value

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
  def add_node_front(input)
    if !@head
      @head = Node.new(input)
    else
      temp = @head
      @head = Node.new(input)
      @head.next = temp
    end
  end

  def add_node_back(input)
    if !@head
      @head = Node.new(input)
    else
      current = @head
      while current.next
        current = current.next
      end
      current.next = Node.new(input)
    end
  end

  def visit
    current = @head
    while current
      print current.value.to_s + " "
      current = current.next
    end
    puts ""
  end

  # 2.1 Remove Dupes
  def remove_dupes
    return puts "insufficient content!" if !@head || !@head.next
    current = @head
    previous = nil
    library = {}

    while current
      if !library[current.value]
        library[current.value] = true
        previous = current
      else
        previous.next = current.next
      end
      current = current.next
    end
  end

  # 2.2 Kth to last // returns head if insufficient nodes to get Kth
  def kth_from_last(k)
    i = 0
    l = @head
    t = @head

    while l
      l = l.next
      i += 1
      t = t.next if i > k
    end
    t
  end

  # 2.3 Delete Middle Node
  def delete_mid_node(myNode)
    current = myNode
    current.value = current.next.value
    current.next = current.next.next
  end

  # 2.4 Partition

  # 2.5 Sum List - see below

  # 2.6 Palindrome // boolean
  def palindrome?
    l = @head
    t = @head
    previous = {}

    while l.next
      previous[l.next] = l
      l = l.next
    end

    i = (previous.length+1)/2
    i.times do
      return false if l.value != t.value
      l = previous[l]
      t = t.next
    end
    true
  end

  # 2.7 Intersection // return node if both pointers -> same node?


  # 2.8 Loop Detection
  def loop?
    l = @head
    t = @head

    while l
      l = l.next.next
      t = t.next
      return true if l == t
    end
    false
  end

  def loop_start?

  end

end

#2.5 Sum List: 1 -> 3 -> 4   +   2 -> 9  =>  3 -> 2 -> 5// currently only works for non-nil values
def sumLists(list1,list2)
  first = list1.head
  second = list2.head
  remainder = 0
  answer = LinkedList.new()

  while first && second
    temp = first.value += second.value += remainder
    if temp > 9
      remainder = 1
      temp -= 10
    else
      remainder = 0
    end
    answer.add_node_back(temp)
    first = first.next
    second = second.next
  end
  answer.add_node_back(remainder) if remainder > 0
  return answer
end

# --------------------------------------------------
# INITIALIZE
myList = LinkedList.new
myList.add_node_front(9)
myList.add_node_front(6)
myList.add_node_front(7)
myList.add_node_front(6)
myList.add_node_front(0)
# myList.add_node(8)
# myList.add_node(1)
# myList.add_node(9)
# myList.add_node(3)
# myList.add_node(2)
# myList.add_node(10)
myNextList = LinkedList.new
myNextList.add_node_front(2)
myNextList.add_node_front(9)
myNextList.add_node_front(5)
myList.visit
myNextList.visit
puts "--------------"

#Remove dupes
# myList.remove_dupes

j = myList.kth_from_last(4)

# myList.delete_mid_node(j)
#
# n =  sumLists(myList,myNextList)
# n.visit
#
# p myList.palindrome?
