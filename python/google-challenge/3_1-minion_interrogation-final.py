def answer(minions):
    # your code here
    mns = []
    for m in minions:
        mns.append([float(i) for i in m])

    ordered = []
    end = -1
    for i in range(len(mns)):
        if end == -1:
            end = i
        else:
            res = better_end(mns[i], mns[end])
            if res == 0:
                end = i
            elif res == 1:
                if i > end:
                    end = i

        sit = False
        for j in range(len(ordered)):
            res = better(mns[i], mns[ordered[j]])
            if res == 0:
                ordered = ordered[:j] + [i] + ordered[j:]
                sit = True
                break
            elif res == 1:
                for k in range(j, len(ordered)):
                    if better(mns[i], mns[ordered[k]]) == 0:
                        ordered = ordered[:k] + [i] + ordered[k:]
                        sit = True
                        break
                    if i <= ordered[k]:
                        print "fuck here"
                        ordered = ordered[:k] + [i] + ordered[k:]
                        sit = True
                        break
                break
        if not sit:
            ordered.append(i)
    ordered.pop(ordered.index(end))
    return ordered + [end]


def better(a, b):
    av = a[0] * a[1] / a[2] + (1 - a[1] / a[2]) * (b[1] / b[2]) * (a[0] + b[0])
    bv = b[0] * b[1] / b[2] + (1 - b[1] / b[2]) * (a[1] / a[2]) * (a[0] + b[0])
    if av < bv:
        return 0
    elif av == bv:
        return 1
    else:
        return 2


def better_end(a, b):
    av = a[0] * a[1] / a[2] + (1 - a[1] / a[2]) * (a[0] + b[0])
    bv = b[0] * b[1] / b[2] + (1 - b[1] / b[2]) * (a[0] + b[0])
    if av > bv:
        return 0
    elif av == bv:
        return 1
    else:
        return 2


print answer([
    [5, 1, 5], [10, 1, 2],
    [5, 1, 5], [10, 1, 2]
])

print answer([
    [30, 185, 624], [66, 351, 947], [26, 1023, 1024], [19, 148, 250],
    [5, 1, 5], [10, 1, 2]
])
