import math

START_POSITION = 50
sum = 0
input = "./input_one.txt"
# input = "./start.txt"


class Node:
    def __init__(self, number: int):
        self.value = number
        self.next = None


class CircularLinkedList:
    def __init__(self):
        self.head = None

    def append(self, number: int):
        new_node = Node(number)

        if not self.head:
            new_node.next = new_node
            self.head = new_node
        else:
            current = self.head
            while current.next != self.head:
                current = current.next
            current.next = new_node
            new_node.next = self.head

    def search(self, number):
        current = self.head

        while True:
            if current.value == number:
                return current

            current = current.next

    def count(self, number, steps, node):
        count = 0
        sum = 0
        current = node
        while count < steps:
            # print(f"current={current.value} steps={steps} count={count}")
            if current.value == number:
                sum += 1
            current = current.next
            count += 1
        return sum

    def print(self):
        count = 0
        current = self.head
        while count < 100:
            print(f"{current.value}", end=" -> ")
            current = current.next
            count += 1


with open(input, "r") as f:
    lines = [l.strip() for l in f.readlines()]


def part_one():
    sum = 0
    last_position = START_POSITION
    for l in lines:
        # print(l[0], l[1:])
        direction = l[0]
        movement = int(l[1:])
        if direction == "L":
            last_position -= movement
        else:
            last_position += movement

        if last_position % 100 == 0:
            sum += 1
    return sum


def part_two():
    sum = 0
    new_position = START_POSITION
    last_position = START_POSITION

    pos_nodes = CircularLinkedList()
    neg_nodes = CircularLinkedList()

    [pos_nodes.append(i) for i in range(100)]
    [neg_nodes.append(i) for i in range(99, -1, -1)]

    for l in lines:
        direction = l[0]
        movement = int(l[1:])
        if direction == "L":
            node = neg_nodes.search(new_position)
            new_position -= movement
            new_position = new_position % 100
            sum += neg_nodes.count(0, movement, node)
        else:
            node = pos_nodes.search(new_position)
            new_position += movement
            new_position = new_position % 100
            sum += pos_nodes.count(0, movement, node)

    return sum


def count_crossings(start, movement, size=100, threshold=0, direction="R"):
    """
    Counts how many times a threshold (default 0) is crossed on a circular array of given size.

    start: current position (0..size-1)
    movement: number of steps to move
    size: size of circular array
    threshold: value to count crossings of
    direction: "R" for forward, "L" for backward
    """

    leftover = movement % size  # leftover steps after full loops
    total_crossings = 0

    # Compute new position modulo size
    if direction == "R":
        if (start + movement) // size > 0:
            total_crossings += 1
        new_pos = (start + movement) % size
    else:
        if (start - movement) < threshold:
            total_crossings += 1
        new_pos = (start - movement) % size
    return total_crossings, new_pos


def part_two_testing():
    total_sum = 0
    position = START_POSITION
    for line in lines:
        direction = line[0]
        movement = int(line[1:])
        crossings, position = count_crossings(
            position, movement, size=100, direction=direction
        )
        total_sum += crossings

    return total_sum


sum_one = part_one()
print(f"SUM_ONE={sum_one}")
sum_two = part_two()
print(f"SUM_TWO={sum_two}")
sum_two_b = part_two_testing()
print(f"SUM_TWO_B={sum_two_b}")
