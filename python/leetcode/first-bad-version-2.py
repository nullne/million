# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        lf, rt = 1, n
        while True:
            if lf == rt:
                if isBadVersion(lf):
                    return lf
                else:
                    break
            ct = lf + (rt - lf) / 2
            if isBadVersion(ct):
                if ct > 1:
                    if not isBadVersion(ct - 1):
                        return ct
                    else:
                        rt = ct - 1
                else:
                    return ct
            else:
                lf = ct + 1


def isBadVersion(n):
    if n % 3 == 0:
        return True
    else:
        return False

solution = Solution()
print solution.firstBadVersion(2126753390)
