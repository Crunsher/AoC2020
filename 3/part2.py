f = open("input.txt", "r")


def goSlope(speed, fall):
    f.seek(0)
    trees = 0
    p = 0-speed
    ln = -1
    print("Speed: {} Fall: {}".format(speed, fall))
    for line in f:
        ln = ln+1
        if ln % fall != 0:
            print("continue")
            continue
        if p + speed >= len(line)-1:
            p = (p+speed)%(len(line)-1)
        else:
            p = p + speed

        print("LN: {} P: {} Char: {}".format(ln, p, line[p]))
        if line[p] == "#":
            trees = trees + 1

    print(trees)
    return trees

trees = goSlope(1,1) * goSlope(3,1) * goSlope(5,1) * goSlope(7,1) * goSlope(1,2)
print("All trees multi: ", trees)
