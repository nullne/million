def answer(minions):
    # your code here
    mns = []

    for i in range(len(minions)):
        v = float(minions[i][0]) / (float(minions[i][1]) / float(minions[i][2]))
        sit = False
        for j in range(len(mns)):
            if v < mns[j][0]:
                mns = mns[:j] + [[v, i]] + mns[j:]
                sit = True
                break
            elif v > mns[j][0]:
                continue
            else:
                for k in range(j, len(mns)):
                    if v < mns[k][0]:
                        mns = mns[:k] + [[v, i]] + mns[k:]
                        sit = True
                        break
                    if minions[i] <= minions[mns[k][1]]:
                        mns = mns[:k] + [[v, i]] + mns[k:]
                        sit = True
                        break
                    else:
                        continue
                break
        if not sit:
            mns.append([v, i])
    return [i[1] for i in mns]
