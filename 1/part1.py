f = open("input.txt", "r")

nums = []

for l in f:
    nums.append(int(l))

def chk2020(x, nums):
    for y in nums:
        if x+y == 2020:
            return x,y

def mkchk(a):
    def chk(b):
        return a+b==2020
    return chk

def all(l):
    result = 1
    for elem in l:
        chk = mkchk(elem)
        num = list(filter(chk, l))
        if len(num) > 0:
            result = result * num[0] 
    print(result)


all(nums)

f.close()
