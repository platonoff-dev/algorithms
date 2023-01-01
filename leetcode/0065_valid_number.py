class Solution:
    def is_digit(self, s: str) -> bool:
        return s >= '0' and s <= '9'
    
    def is_digit_or_dot(self, s: str) -> bool:
        return self.is_digit(s) or s == '.'
    
    def is_natural(self, s: str) -> bool:
        if len(s) == 0:
            return False
        
        for ch in s:
            if not self.is_digit(ch):
                return False
        
        return True
    
    def is_integer(self, s: str) -> bool:
        if s[0] == '-' or s[0] == '+':
            return self.is_natural(s[1:])
        else:
            return self.is_natural(s)
    
    def is_decimal(self, s: str) -> bool:
        if s[0] == '-' or s[0] == '+':
            s = s[1:]
        
        comma_index = -1
        for i, ch in enumerate(s):
            if ch == '.':
                comma_index = i
                break
            
        if comma_index == '-1':
            return False
        
        if comma_index == 0:
            return self.is_natural(s[1:])
        elif comma_index == len(s) - 1:
            return self.is_natural(s[:len(s) - 1])
        else:
            return self.is_natural(s[:comma_index]) and self.is_natural(s[comma_index + 1:])
    
    def isNumber(self, s: str) -> bool:
        if len(s) == 0:
            return False
        
        e_idx = -1
        for i, ch in enumerate(s):
            if ch == 'e' or ch == 'E':
                e_idx = i
                break
        
        if e_idx == 0 or e_idx == len(s) - 1:
            return False
            
        if e_idx == -1:
            return self.is_integer(s) or self.is_decimal(s)
        else:
            return (self.is_integer(s[:e_idx]) or self.is_decimal(s[:e_idx])) and self.is_integer(s[e_idx + 1:])


