import random

import pytest

from main import sort

def random_list(length: int) -> list[int]:
    return [random.randint(0, length) for _ in range(length)]

@pytest.mark.parametrize("length", [0, 10, 10000])
def test_sort(length: int):
    arr = random_list(length)
    sort(arr)

    for i in range(1, len(arr)):
        assert arr[i] >= arr[i-1]


@pytest.mark.parametrize("length", [0, 10, 10000])
def test_sort_desc(length: int):
    arr = random_list(length)
    sort(arr, False)

    for i in range(1, len(arr)):
        assert arr[i] <= arr[i-1]
