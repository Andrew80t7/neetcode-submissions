class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:

        d = {i: nums.count(i) for i in nums}
        return sorted([i for i in sorted(d, key=d.get, reverse=True)[:k]])
