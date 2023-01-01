

def sort(arr: list[int], asc: bool = True):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1

        while j >= 0 and (arr[j] > key and asc or arr[j] < key and not asc):
            arr[j+1] = arr[j]
            j -= 1
        
        arr[j+1] = key
