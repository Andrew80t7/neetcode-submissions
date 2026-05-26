class Solution:

    def isAnagram(self, a, b):
        res = True if sorted([i for i in a]) == sorted([j for j in b]) else False
        return res


    def groupAnagrams(self, nums: List[str]) -> List[List[str]]:
        result = []
        used_indices = set()

        for i in range(len(nums)):
            if i in used_indices:
                continue
                
            temp_array = [nums[i]]
            used_indices.add(i)
            
            for j in range(i + 1, len(nums)):
                if j not in used_indices and self.isAnagram(nums[i], nums[j]):
                    temp_array.append(nums[j])
                    used_indices.add(j)
                    
            result.append(temp_array)
        return result