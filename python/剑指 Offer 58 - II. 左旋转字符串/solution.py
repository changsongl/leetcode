class Solution:
    def reverseLeftWords(self, s: str, n: int) -> str:
        return s[n:]+s[:n]

s = Solution()
print(s.reverseLeftWords("123456", 2))
