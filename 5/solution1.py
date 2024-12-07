from sys import stdin
from collections import defaultdict


rules = defaultdict(list)

for line in stdin:
    line = line.strip()
    if not line:
        break
    rule = line.split('|')
    rules[int(rule[1])].append(int(rule[0]))

updates = [line.strip().split(',') for line in stdin][:-1]
updates = [list(map(int, sublist)) for sublist in updates]


def check_update(update):
    curr_rules = set()
    for i, page in enumerate(update):
        curr_rules.update(rules[page])
        if page in curr_rules:
            incorrect_idx = i
            return incorrect_idx
    return -1


def correct_update(update, incorrect_idx):
    for i, page in enumerate(update):
        if update[incorrect_idx] in rules[page]:
            temp = page
            update[i] = update[incorrect_idx]
            update[incorrect_idx] = temp

            return update


result = 0
for update in updates:
    save = False
    while (incorrect_idx := check_update(update)) != -1:
        save = True
        update = correct_update(update, incorrect_idx)
    if save:
        result += update[len(update) // 2]

print(result)

