from sys import stdin


map = [ list(line.strip()) for line in stdin.readlines() ][:-1]

directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

def out_of_bounds(position):
    return position[0] < 0 or position[0] >= len(map) or position[1] < 0 or position[1] >= len(map[0])

def travel(position):
    positions = {position}
    direction = directions[0]
    while True:
        next_position = (position[0]+direction[0], position[1]+direction[1])
        
        # moving out of bounds
        if out_of_bounds(next_position):
            break

        # change directions
        if map[position[0]+direction[0]][position[1]+direction[1]] == '#':
            direction = directions[(directions.index(direction)+1)%4]
            continue

        # move forward
        position = next_position
        positions.add(position)
    return len(positions)

start =  next((row_idx, row.index('^')) for row_idx, row in enumerate(map) if '^' in row)

print(travel(start))



