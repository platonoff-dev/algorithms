class Solution:
    def romanToInt(self, s: str) -> int:
        values = {
            "I": 1,
            "IV": 4,
            "V": 5,
            "IX": 9,
            "X": 10,
            "XL": 40,
            "L": 50,
            "XC": 90,
            "C": 100,
            "CD": 400,
            "D": 500,
            "CM": 900,
            "M": 1000
        }

        n = 0
        i = 0
        while i < range(len(s)):   
            if i != len(s) - 1 and values.get(s[i:i+2], None):
                n += values[s[i:i+2]]
                i += 2
            else:
                n += values[s[i]]
                i += 1
        
        return n


if __name__ == "__main__":
    Solution().romanToInt("MCMXCIV")
