#!/usr/bin/python
from typing import List


class Solution:
    def reverseString(self, s: List[str]) -> None:
        # s.reverse()
        # s[:] = s[::-1]
        i = 0
        len_list = len(s)
        while i < len_list/2:
            s[i], s[len_list-i-1] = s[len_list-i-1], s[i]
            i += 1
        """
        Do not return anything, modify s in-place instead.
        """


solution = Solution()
l = ['Google', 'Runoob']
solution.reverseString(l)
print(l)
