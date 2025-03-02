# Add Binaries

### Problem
Consider the problem of adding two n-bit binary integers, stored in two n-element
arrays A and B. The sum of the two integers should be stored in binary form in
(n+1)-element array C. State the problem formally and write pseudocode for
adding the two integers.

### Formal problem
**Input**: A, B sequences of booleans $A = (a_1, a_2, ..., a_n)$ and $B = (b_1, b_2, ..., b_n)$, where $a_n, b_n \in \{true, false\}$, each array representing a number stored in binary format

**Output**: C - sequence of booleans with length $n + 1$ representing binary form of the result of addition.

## Functions

### `AddBinaries(a []bool, b []bool) []bool`

This function takes two slices of booleans representing binary numbers and returns their sum as a slice of booleans.

### `AddBinariesAsInts(a []int, b []int) []int`

This function takes two slices of integers (0 or 1) representing binary numbers and returns their sum as a slice of integers.

## Explanation

- The `AddBinaries` function uses boolean logic to compute the sum of two binary numbers.
- The `AddBinariesAsInts` function performs the same operation but uses integer arithmetic to compare performance.

Both functions handle the carry-over during the addition process and return the correct binary sum.
