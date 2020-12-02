f = open("input.txt", "r")

nums = []

for line in f:
    nums.append(int(line))

def mksmolchk(a):
    def chk(b):
        return a+b<2020
    return chk

def mkfitchk(a):
    def chk(b):
        return a+b==2020
    return chk

for elem in nums:
    smolnums = list(filter(lambda b, a=elem : a+b<2020, nums))
    for snom in smolnums:
        med = elem+snom
        for big in nums:
            if big + med == 2020:
                print(elem, snom, big)
                print(elem*snom*big)


f.close()
