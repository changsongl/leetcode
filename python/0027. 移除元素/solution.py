from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        pos = 0
        while i < len(nums):
            if nums[i] != val:
                nums[pos] = nums[i]
                pos += 1
            i += 1
        return pos

