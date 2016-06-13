def answer(names):
    # your code here
    sorts = []
    for name in names:
        score = sum([ord(l) - ord('a') + 1 for l in name])
        sit = False
        for i in range(len(sorts)):
            if score > sorts[i][0]:
                sorts = sorts[:i] + [[score, name]] + sorts[i:]
                sit = True
                break
            elif score < sorts[i][0]:
                continue
            else:
                for n in range(i, len(sorts)):
                    if score != sorts[n][0]:
                        sorts = sorts[:n] + [[score, name]] + sorts[n:]
                        sit = True
                        break
                    elif name >= sorts[n][1]:
                        sorts = sorts[:n] + [[score, name]] + sorts[n:]
                        sit = True
                        break
                    else:
                        continue
                break
        if sit is not True:
            sorts = sorts + [[score, name]]
    return [s[1] for s in sorts]
