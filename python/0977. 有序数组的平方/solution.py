from typing import List


# class Solution:
#     def sortedSquares(self, nums: List[int]) -> List[int]:
#         return sorted([num * num for num in nums])

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n

        i, j, pos = 0, n - 1, n-1
        while i <= j:
            if nums[i] * nums[i] > nums[j] * nums[j]:
                ans[pos] = nums[i] * nums[i]
                i += 1
            else:
                ans[pos] = nums[j] * nums[j]
                j -= 1
            pos -= 1

        return ans


s = Solution()
print(s.sortedSquares([-5, 2, 7, 9]))



