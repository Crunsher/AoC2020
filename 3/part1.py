f = open("input.txt", "r")

trees = 0
p = -3
for line in f:
    if p + 3 >= len(line)-1:
        p = (p+3)%(len(line)-1)
    else:
        p = p + 3

    print("P: {} Char: {}".format(p, line[p]))
    if line[p] == "#":
        trees = trees + 1

print(trees)


