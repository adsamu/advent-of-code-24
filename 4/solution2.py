from sys import stdin
from time import sleep
import numpy as np


x_masks = [
    np.array([[1, 0, 1],
              [0, 2, 0],
              [3, 0, 3]]),

    np.array([[3, 0, 1],
              [0, 2, 0],
              [3, 0, 1]]),

    np.array([[3, 0, 3],
              [0, 2, 0],
              [1, 0, 1]]),

    np.array([[1, 0, 3],
              [0, 2, 0],
              [1, 0, 3]])
]


puzzle = [list(line)[:-1] for line in stdin][:-1]
mapping = {'M': 1, 'A': 2, 'S': 3, 'X': 0}
puzzle = [[mapping[puzzle[i][j]]
           for j in range(len(puzzle[i]))] for i in range(len(puzzle))]
puzzle = np.array(puzzle)

puzzle_rows, puzzle_cols = puzzle.shape
mask_rows, mask_cols = 3, 3

mask = [[1, 0, 1],
        [0, 1, 0],
        [1, 0, 1]]
mask = np.array(mask)


total_matches = 0
# Slide the mask over the larger matrix
for i in range(puzzle_rows - mask_rows + 1):
    for j in range(puzzle_cols - mask_cols + 1):
        # print(i, i + mask_rows) 
        # print(j, j + mask_cols)
        # Extract the submatrix of the same size as the mask
        test = puzzle[i:i + mask_rows, j:j + mask_cols]
        # print(test)
        sub_matrix = test * mask

        # Count matches between the mask and the current sub-matrix
        # print("*******************")
        for x_mask in x_masks:
            # print(x_masks[i,:,:])
            # print(sub_matrix)
            # print(x_masks[i,:,:])
            # print(sub_matrix)
            # print(x_mask)
            if np.all(sub_matrix == x_mask):
                # print(sub_matrix)
                # print(x_mask)
                total_matches += 1
                break
            # total_matches += 1 if np.sum(sub_matrix - x_mask) == 0 else 0
            # print("-------------------")
        # print(total_matches)
        # sleep(3)
        # Accumulate the count of matches
        # total_matches += matches


# print(mask.shape)
# print(puzzle.shape)
print(total_matches)
