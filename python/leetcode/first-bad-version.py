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
        return self.helper(range(1, n+1))

    def helper(self, n):
        center = len(n)/2
        if isBadVersion(n[center]):
            if center > 0:
                if not isBadVersion(n[center] - 1):
                    return n[center]
                else:
                    return self.helper(n[:center])
            else:
                return n[center]
        else:
            return self.helper(n[center:])

def isBadVersion(n):
    if n % 3 == 0:
        return True
    else:
        return False

solution = Solution()
print solution.firstBadVersion(2126753390)
