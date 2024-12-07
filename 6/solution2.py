from collections import defaultdict
from sys import stdin


map = [ list(line.strip()) for line in stdin.readlines() ][:-1]

directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

def out_of_bounds(position):
    return position[0] < 0 or position[0] >= len(map) or position[1] < 0 or position[1] >= len(map[0])

def travel(position, new_obstacle):
    direction = directions[0]
    positions = {(position, direction)}
    while True:
        next_position = (position[0]+direction[0], position[1]+direction[1])
        
        # moving out of bounds
        if out_of_bounds(next_position):
            return False

        # change directions
        if map[next_position[0]][next_position[1]] == '#' or next_position == new_obstacle:
            direction = directions[(directions.index(direction)+1)%4]
            continue

        # move forward
        position = next_position

        if (position, direction) in positions:
            return True
        positions.add((position, direction))
    return False


start =  next((row_idx, row.index('^')) for row_idx, row in enumerate(map) if '^' in row)
paths = set()
nr_loops = 0
for i in range(len(map)):
    for j in range(len(map[i])):
        if map[i][j] != '^' and map[i][j] != '#':
            temp = map[i][j]
            nr_loops += 1 if travel(start, (i, j)) else 0
            

print(nr_loops)



