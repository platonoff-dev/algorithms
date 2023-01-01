class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        left_pos = 0
        right_pos = len(nums) - 1

        while right_pos >= left_pos:
            idx = (left_pos + right_pos) // 2
            if target < nums[idx]:
                right_pos = idx -1
            elif target > nums[idx]:
                left_pos = idx + 1
            else:
                return idx

        return left_pos

