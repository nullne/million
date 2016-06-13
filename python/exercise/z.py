#!/usr/bin/python2.7
import os
import time

def getLastModified(path):
    dirname = os.path.dirname(path)
    src = ""
    last = None
    for d in os.listdir(path):
        tmp = os.path.getmtime(os.path.join(dirname, d))
        if last is None:
            last = tmp
            src = d
        else:
            if last < tmp:
                last = tmp
                src = d
    return src

def main(src, dest):
    d = os.path.dirname(src)
    last = getLastModified(src)
    print last
    print d
    if src == "":
        return
    os.system("xcopy /e %s %s" % (os.path.join(d, last), dest))

if __name__ == "__main__":
    main("D:\\backup\\", "D:\\testbk\\")
