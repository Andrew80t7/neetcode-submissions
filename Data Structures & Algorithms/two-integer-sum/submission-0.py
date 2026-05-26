class Solution:
    def twoSum(self, array: List[int], target: int) -> List[int]:
            for i in range(len(array)):
                for j in range(i+1, len(array)):
                    if array[i]+array[j] == target:
                        return [i, j]