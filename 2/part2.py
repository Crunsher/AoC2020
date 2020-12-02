f = open("input.txt", "r")

correct = 0

for line in f:
    k = line.split(' ')
    ns = k[0].split('-')
    c = k[1][0]

    if (k[2][int(ns[0])-1] == c) != (k[2][int(ns[1])-1] == c):
        correct = correct + 1 


print("Correct: ", correct)
