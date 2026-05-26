class Solution:
    def isAnagram(self, s1: str, s2: str) -> bool:
            if len(s1) != len(s2):
                return False
            f = 1
            for i in s1:
                if (i not in s2) or (s1.count(i) != s2.count(i)):
                    f = 0
            return bool(f)