def answer(minions):
    # your code here
    mns = []
    for minion in minions:
        mns.append([float(i) for i in minion])
    bests = []
    min_time = 0
    for arr in arrangement([i for i in range(len(mns))]):
        expected = 0.0
        for i in range(len(arr)):
            pi = arr[i]
            p = mns[pi][1] / mns[pi][2]
            if i + 1 == len(arr):
                p = 1.0
            base = mns[pi][0]
            for j in range(i):
                pj = arr[j]
                p *= (1 - (mns[pj][1] / mns[pj][2]))
                base += mns[pj][0]
            expected += p * base

        if len(bests) == 0:
            bests.append(arr)
            min_time = expected
        else:
            if expected == min_time:
                bests.append(arr)
            elif expected < min_time:
                bests = [arr]
                min_time = expected
    best = bests[0]
    for b in bests:
        if b < best:
            best = b



    # m = 0.0
    # for b in best:
    #     v = mns[b][0] / (mns[b][1] / mns[b][2])
    #     if m == 0:
    #         m = v
    #     elif m < v:
    #         print "fuck"



    return best


def arrangement(n):
    if len(n) == 1:
        yield n
    for i in range(len(n)):
        ne = n[:i] + n[i+1:]
        r = arrangement(ne)
        for t in r:
            yield [n[i]] + t

res =  answer([
    [30, 185, 624], [66, 351, 947], [26, 1023, 1024], [19, 148, 250],
    [5, 1, 5], [10, 1, 2]
])

print res

