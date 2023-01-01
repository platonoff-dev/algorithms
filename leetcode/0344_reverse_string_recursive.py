class Solution:
    def helper(self, s: List[str], left: int, right: int) -> None:
        if right - left < 1:
            return
        
        s[left], s[right] = s[right], s[left]
        self.helper(s, left + 1, right - 1)
    
    def reverseString(self, s: List[str]) -> None:
        self.helper(s, 0, len(s) - 1)

