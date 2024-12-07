from sys import stdin
from time import sleep


# eight directions 
directions = [
    (0, 1),    # right
    (1, 1),    # down-right
    (1, 0),    # down
    (1, -1),   # down-left
    (0, -1),   # left
    (-1, -1),  # up-left
    (-1, 0),   # up
    (-1, 1)    # up-right
]
puzzel = []

def find_word(word, start, direction):
    if len(word) == 0:
        # print("found")
        return 1
    if start[0] < 0 or start[0] >= len(puzzle) or start[1] < 0 or start[1] >= len(puzzle[0]):
        # print("out of bounds", start)
        return 0
    if puzzle[start[0]][start[1]] != word[0]:
        # print("not matching" , word[0], puzzle[start[0]][start[1]])
        return 0
    sum = 0
    # print("one letter matched", word[0], puzzle[start[0]][start[1]])
    # print("sum", sum)
    # sleep(3)
    return find_word(word[1:], (start[0] + direction[0], start[1] + direction[1]), direction)




puzzle = [ list(line)[:-1] for line in stdin][:-1]
sum = 0
for i in range(len(puzzle)):
    for j in range(len(puzzle[i])):
        for direction in directions:
            sum += find_word("XMAS", (i, j), direction)


print(sum)


