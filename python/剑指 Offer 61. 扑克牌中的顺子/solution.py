from typing import List


class Solution:
    def isStraight(self, nums: List[int]) -> bool:
        repeat = set()
        mi, ma = 14, 0

        for num in nums:
            if num in repeat:
                return False
            if num != 0:
                mi = min(num, mi)
                ma = max(num, ma)
                repeat.add(num)
        return ma - mi < 5

# class Solution:
#     def isStraight(self, nums: List[int]) -> bool:
#         nums.sort()
#         joker = 0
#         for i in range(4):
#             if nums[i] == 0:
#                 joker += 1
#             elif nums[i] == nums[i+1]:
#                 return False
#
#         return nums[4] - nums[joker] < 5

# class Solution:
#     def isStraight(self, nums: List[int]) -> bool:
#         nums.sort()
#         count, zero_count = 0, 0
#         for i, num in enumerate(nums):
#             if num == 0:
#                 zero_count += 1
#                 continue
#             if i != 0 and nums[i-1] != 0:
#                 if nums[i-1] == num:
#                     return False
#                 count += num - nums[i-1] - 1
#             i += 1
#         return count <= zero_count


s = Solution()
print(s.isStraight([9,10,4,0,9]))
