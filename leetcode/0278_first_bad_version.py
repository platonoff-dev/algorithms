# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        left = 1
        right = n + 1
        
        while right >= left:
            idx = (left + right) // 2
            print(left, right, idx)
            if isBadVersion(idx):
                if not isBadVersion(idx - 1):
                    return idx
                if isBadVersion(idx + 1):
                    right = idx
                    continue
            else:
                if isBadVersion(idx + 1):
                    return idx + 1
                if not isBadVersion(idx - 1):
                    left = idx
                    continue

        return -1

