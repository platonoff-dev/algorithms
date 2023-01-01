class Solution:
    def countPrimes(self, n: int) -> int:
        primes = [True for _ in range(n)]
        i = 2
        while i * i < n:
            if primes[i]:
                for p in range(i * i, n, i):
                    primes[p] = False
            i += 1
        
        counter = 0
        for i in range(2, len(primes)):
            if primes[i]:
                counter += 1
        
        return counter

