class Solution:   
    def reverseString(self, s: List[str]) -> None:
        for i in range(len(s) // 2):
            left = i
            right = len(s) - 1 - i
            tmp = s[left]
            s[left] = s[right]
            s[right] = tmp

