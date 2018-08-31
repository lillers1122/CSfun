# RUBY LINKED LIST IMPLEMENTATION
class Node
  attr_accessor :next
  attr_reader :value

  def initialize(input)
    @value = input
    @next = nil
  end
end

class LinkedList
  def initialize
    @head = nil
  end

  # 1. Add New Node
  def add_node(input)
    if !@head
      @head = Node.new(input)
    else
      temp = @head
      @head = Node.new(input)
      @head.next = temp
    end
  end

  # 2. Contains Value? Boolean
  def has_value?(input)
    current = @head
    while current
      return true if current.value == input
      current = current.next
    end
    return false
  end

  # 3. Max Value
  def get_max
    if !@head
      return "no content!"
    else
      max = @head.value
      current = @head
      while current
        if max < current.value
          max = current.value
        end
        current = current.next
      end
      return max
    end
  end

  # 4. Min Value
  def get_min
    if !@head
      return "no content!"
    else
      min = @head.value
      current = @head
      while current
        if min > current.value
          min = current.value
        end
        current = current.next
      end
      return min
    end
  end

  # 5. Count Nodes
  def get_size
    count = 0
    current = @head
    while current
      count += 1
      current = current.next
    end
    return count
  end

  # 6. Nth Node from the Beginning; Index Starts at 0
  def nth_node(input)
    current = @head
    (input).times do
      return "not enough content!" if !current
      current = current.next
    end
    return current.value
  end

  # 7. Insert Node Asc Order
  def asc_insert(input)
    if !@head
      #list is empty
      @head = Node.new(input)
    elsif @head.value > input
      #input smaller than head value
      add_node(input)
    else
      #input larger than head value
      current = @head
      if !current.next
        #if list has 1 item
        current.next = Node.new(input)
      else
        #if list has more than 1 item iterate until current's next node is bigger than input
        while current.next && current.next.value < input
          current = current.next
        end
        ourNode = Node.new(input)
        ourNode.next = current.next
        current.next = ourNode
      end
    end
  end

  # 8. Print Values Separated by a Space
  def visit
    return puts "no content!" if !@head

    current = @head
    while current
      print current.value.to_s + " "
      current = current.next
    end
    puts ""
  end

  # 9. Delete First Found of Value
  def delete(input)
    #TODO what if the input cannot be found?
    #if empty
    return puts "no content!" if !@head
    current = @head
    #if 1st node has value
    if current.value == input
      if current.next
        #if there is 2nd node in list
        @head = current.next
        return
      else
        #if only 1 node
        @head = nil
        return
      end
    end
      #keep looking for the value
    while current
      if current.next.value == input && current.next.next
        #if value is found
        current.next = current.next.next
        return
      end
      current = current.next
    end
  end

  # 10. Reverse List
  def reverse
    #no content or only 1 node
    return puts "insufficient content!" if !@head || !@head.next

    current = @head
    previous = nil
    while current
      nextNode = current.next
      current.next = previous
      previous = current
      current = nextNode
    end
    @head = previous
  end

end

# --------------------------------------------------
# INITIALIZE
# myList = LinkedList.new()
# myList.visit
# puts "--------------"

# CHECK METHODS
# 1 Add
# myList.add_node(7)
# myList.add_node(6)
# myList.add_node(1)

#2 Contains?
# puts myList.has_value?(1)
# puts myList.has_value?(5)

#3 Max, #4 Min
# puts myList.get_max
# puts myList.get_min

#5 Size
# puts myList.get_size

#6 Nth Node
# myList.visit
# puts myList.nth_node(0)
# puts myList.nth_node(1)
# puts myList.nth_node(5)

#7 Insert Asc
# puts myList.get_size
# myList.asc_insert(6)
# puts myList.get_size

#8. Visit
# myList.visit

#9 Delete
# myList.visit
# myList.delete(6)
# myList.visit
# puts myList.get_size
# myList.delete(1)
# myList.visit
# puts myList.get_size
# myList.delete(7)
# myList.visit
# puts myList.get_size

#10 Reverse
# myList.visit
# myList.reverse
# myList.visit
# puts "--------------"
# myList.delete(6)
# myList.delete(7)
# myList.reverse
# myList.visit
