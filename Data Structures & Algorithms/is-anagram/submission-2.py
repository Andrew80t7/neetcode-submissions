class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        return True if sorted([i for i in s]) == sorted([i for i in t]) else False
