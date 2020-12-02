f = open("input.txt", "r")

correct = 0

for line in f:
    k = line.split(' ')
    ns = k[0].split('-')
    n = k[2].count(k[1][0])
    if n >= int(ns[0]) and n <= int(ns[1]):
        correct = correct + 1 

print("Correct: ", correct)
