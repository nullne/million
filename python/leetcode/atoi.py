class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        str = str.strip()
        if len(str) < 1:
            return 0
        sign = '+'
        if str[0] in ['+', '-']:
            sign = str[0]
            str = str[1:]
        rtn = []
        for i in str:
            cur = ord(i) - ord('0')
            if cur > 9 or cur < 0:
                break
            rtn.insert(0, cur)
        p = 0
        for l in range(len(rtn)):
            p = p + rtn[l] * (10 ** l)
        if p >= 2**31 and sign == '+':
            p = 2 ** 31 - 1
        if p > 2**31 and sign == '-':
            p = 2 ** 31
        if sign == '-':
            return -p
        return p

solution = Solution()
i = solution.myAtoi("-a100a")
print i, type(i)
