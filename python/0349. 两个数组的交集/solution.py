from typing import List


class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        s1, s2 = set(nums1), set(nums2)
        return list({x for x in s1 if x in s2})


s = Solution()
print(s.intersection([1, 2, 3], [2, 3, 4]))
