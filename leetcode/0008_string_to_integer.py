class Solution:
    def myAtoi(self, s: str) -> int:
        digits = {
            '0': 0,
            '1': 1,
            '2': 2,
            '3': 3,
            '4': 4,
            '5': 5,
            '6': 6,
            '7': 7,
            '8': 8,
            '9': 9,
            
        }
        
        number = 0
        sign = 1
        read_index = 0
        
        while read_index < len(s) and s[read_index] == ' ':
            read_index += 1
        
        if read_index < len(s):
            if s[read_index] == '-':
                sign = -1
                read_index += 1
            elif s[read_index] == '+':
                read_index += 1

        while read_index < len(s) and s[read_index] in digits:
            number = number * 10 + digits[s[read_index]]
            read_index += 1
        
        number *= sign
        if number < -(2 ** 31):
            number = -(2 ** 31)
        
        if number > 2 ** 31 - 1:
            number = 2 ** 31 - 1
        
        return number

