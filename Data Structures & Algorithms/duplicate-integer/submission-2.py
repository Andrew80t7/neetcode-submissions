class Solution:
    def hasDuplicate(self, arr: List[int]) -> bool:
        if (len(arr)) == len(set(arr)):
            return False
        else:
            return True