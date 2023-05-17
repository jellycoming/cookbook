# 选择排序

```py
# 每次迭代中,通过比较更新最小元素的index,每次迭代完成后进行一次交换,时间复杂度为О(n²)
def selection_sort(array):
    length = len(array)
    for i in range(length - 1):
        min_index = i
        for j in range(i + 1, length):
            if array[min_index] > array[j]:
                min_index = j
        if min_index != i:
            array[i], array[min_index] = array[min_index], array[i]
        print 'iteration #{}: {}'.format(i, array)
    return array


array = [6, 5, 7, 2, 1, 4, 3]
print 'Before selection sort: {}'.format(array)
print 'After selection sort: {}'.format(selection_sort(array))

# Output
# Before selection sort: [6, 5, 7, 2, 1, 4, 3]
# iteration #0: [1, 5, 7, 2, 6, 4, 3]
# iteration #1: [1, 2, 7, 5, 6, 4, 3]
# iteration #2: [1, 2, 3, 5, 6, 4, 7]
# iteration #3: [1, 2, 3, 4, 6, 5, 7]
# iteration #4: [1, 2, 3, 4, 5, 6, 7]
# iteration #5: [1, 2, 3, 4, 5, 6, 7]
# After selection sort: [1, 2, 3, 4, 5, 6, 7]
```

# 冒泡排序

```py
# 第一个for负责确定排序步数,例如,当i为7时执行第一次排序,参与排序的是整个序列,第二个for负责比较两个相邻元素,如果符合条件则交换;
# flag可以确保在某次迭代后,所有元素位置都没有发生改变,即已排序完成,可以减少无效的迭代次数
def bubble_sort(array):
    length = len(array)
    for j in range(length - 1, 0, -1):
        flag = True
        for i in range(0, j):
            if array[i] > array[i+1]:
                flag = False
                array[i], array[i+1] = array[i+1], array[i]
        print 'iteration #{}: {}'.format(j, array)
        if flag:
            return array
    return array

array = [6, 5, 7, 2, 1, 4, 3]
print 'Before bubble sort: {}'.format(array)
print 'After bubble sort: {}'.format(bubble_sort(array))

# Output
# Before bubble sort: [6, 5, 7, 2, 1, 4, 3]
# iteration #6: [5, 6, 2, 1, 4, 3, 7]
# iteration #5: [5, 2, 1, 4, 3, 6, 7]
# iteration #4: [2, 1, 4, 3, 5, 6, 7]
# iteration #3: [1, 2, 3, 4, 5, 6, 7]
# iteration #2: [1, 2, 3, 4, 5, 6, 7]
# After bubble sort: [1, 2, 3, 4, 5, 6, 7]
```

# 插入排序

```py
# 平均复杂度O(n²)
def insertion_sort(array):
    for i in range(1, len(array)):
        for j in range(i-1, -1, -1):
            if array[j + 1] < array[j]:
                array[j], array[j + 1] = array[j + 1], array[j]
        print 'iteration #{}: {}'.format(i, array)
    return array

array = [6, 5, 7, 2, 1, 4, 3]
print 'Before selection sort: {}'.format(array)
print 'After selection sort: {}'.format(insertion_sort(array))

# Output
# Before selection sort: [6, 5, 7, 2, 1, 4, 3]
# iteration #1: [5, 6, 7, 2, 1, 4, 3]
# iteration #2: [5, 6, 7, 2, 1, 4, 3]
# iteration #3: [2, 5, 6, 7, 1, 4, 3]
# iteration #4: [1, 2, 5, 6, 7, 4, 3]
# iteration #5: [1, 2, 4, 5, 6, 7, 3]
# iteration #6: [1, 2, 3, 4, 5, 6, 7]
# After selection sort: [1, 2, 3, 4, 5, 6, 7]
```

# 归并排序

```py
# 算法复杂度为O(nlogn),merge部分遍历了所有元素,为O(n),递归部分每次处理一半元素，为O(logn)
def merge_sort(array):
    print 'Array: {}'.format(array)
    if len(array) <= 1:
        return array

    def merge(left, right):
        merged, left, right = deque(), deque(left), deque(right)
        while left and right:
            merged.append(left.popleft() if left[0] <= right[0] else right.popleft())  # deque popleft is also O(1)
        merged.extend(right if right else left)
        print 'Merged: {}'.format(merged)
        return list(merged)

    middle = int(len(array) / 2)
    left = merge_sort(array[:middle])
    right = merge_sort(array[middle:])
    return merge(left, right)


array = [6, 5, 7, 2, 1, 4, 3]
print 'Before merge sort: {}'.format(array)
print 'After merge sort: {}'.format(merge_sort(array))

# Output
# Before merge sort: [6, 5, 7, 2, 1, 4, 3]
# Array: [6, 5, 7, 2, 1, 4, 3]
# Array: [6, 5, 7]
# Array: [6]
# Array: [5, 7]
# Array: [5]
# Array: [7]
# Merged: deque([5, 7])
# Merged: deque([5, 6, 7])
# Array: [2, 1, 4, 3]
# Array: [2, 1]
# Array: [2]
# Array: [1]
# Merged: deque([1, 2])
# Array: [4, 3]
# Array: [4]
# Array: [3]
# Merged: deque([3, 4])
# Merged: deque([1, 2, 3, 4])
# Merged: deque([1, 2, 3, 4, 5, 6, 7])
# After merge sort: [1, 2, 3, 4, 5, 6, 7]
```

# 快速排序

> 快速排序使用分治法（Divide and conquer）策略来把一个序列（list）分为两个子序列（sub-lists）。  
> 步骤为：  
> 1. 从数列中挑出一个元素，称为"基准"（pivot），  
> 2. 重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（相同的数可以到任何一边）。在这个分区结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。  
> 3. 递归地（recursively）把小于基准值元素的子数列和大于基准值元素的子数列排序

```py
# 普通实现，需要O(n)的额外存储空间
def quicksort(array):
    left = []
    pivot_list = []
    right = []
    if len(array) <= 1:
        return array
    pivot = array[0] # 将第一个元素作为基准值
    for i in array[1:]: # 遍历除基准值外的其元素
        if i < pivot:
            left.append(i)
        elif i >= pivot:
            right.append(i)
    pivot_list.append(pivot)
    return quicksort(left) + pivot_list + quicksort(right)

# 普通实现的pythonic版本
def quicksort(array):
    return quicksort([i for i in array[1:] if i < array[0]]) + array[0:1] + quicksort([i for i in array[1:] if i >= array[0]]) if len(array) > 1 else array

# Or
quicksort = lambda array: quicksort([i for i in array[1:] if i < array[0]]) + array[0:1] + quicksort([i for i in array[1:] if i >= array[0]]) if len(array) > 1 else array


# 原地分区实现，通过元素在列表中的移动实现排序，节省空间
def quicksort(array, left, right):
    def partition(array, left, right):
        pivot = array[right-1]  # 将最后一个元素作为基准值
        index = left  # 将所有小于基准值的元素移到列表的左端,索引i代表当前小于基准值的元素应该移到到的位置
        for i in range(left, right):
            if array[i] < pivot:
                array[index], array[i] = array[i], array[index]
                index += 1
        array[index], array[right-1] = pivot, array[index]
        return index

    if len(array[left:right]) > 1:
        i = partition(array, left, right)
        quicksort(array, left, i)
        quicksort(array, i, right)

if __name__ == '__main__':
    array = [4, 65, 2, -31, 0, 99, 83, 782, 1]
    quicksort(array, 0, len(array))
    print array

# Output
# [-31, 0, 1, 2, 4, 65, 83, 99, 782]
```

# 斐波那契数列

```py
# 生成式方式
def fibonacci(n):
    a, b = 0, 1
    yield a, b
    while n:
        a, b = b, a + b
        n -= 1
        yield a, b

n = 20
res = [a for a, b in fibonacci(n)]
print res

# Output
# [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765]


# 递归方式
def fibonacci_recursion(n):
    if n <= 1:
        return 0 if n == 0 else 1
    return fibonacci_recursion(n-1) + fibonacci_recursion(n-2)
n = 20
res = [fibonacci_recursion(i) for i in range(n+1)]
print res

# Output
# [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765]
```

# 动态规划-斐波那契数列

```py
# 在递归方式生成斐波那契数列的过程中，会有很多重复调用，如[fibonacci_recursion]fib(5)会调用fib(4)及fib(3),
# fib(4)又会调用fib(3)和fib(2),如Output所示,算法一直对已经知道答案的子问题重复求解，
# 对于这类“重叠的子问题”的问题，可以采用动态规划算法。
def fibonacci_recursion(n):
    global calls
    calls += 1
    print 'fibonacci called with {}'.format(n)
    if n <= 1:
        return 0 if n == 0 else 1
    return fibonacci_recursion(n-1) + fibonacci_recursion(n-2)

calls = 0
n = 20
res = [fibonacci_recursion(i) for i in range(n+1)]
print 'fibonacci n: {}\nres: {}\ncall times: {}'.format(n, res, calls)

# Output
# ......
# fibonacci called with 1
# fibonacci called with 0
# fibonacci called with 3
# fibonacci called with 2
# fibonacci called with 1
# fibonacci called with 0
# fibonacci called with 1
# fibonacci called with 4
# fibonacci called with 3
# fibonacci called with 2
# fibonacci called with 1
# fibonacci called with 0
# fibonacci called with 1
# fibonacci called with 2
# fibonacci called with 1
# fibonacci called with 0
# fibonacci called with 7
# fibonacci called with 6
# fibonacci called with 5
# fibonacci called with 4
# fibonacci called with 3
# fibonacci called with 2
# fibonacci called with 1
# ......
# fibonacci n: 20
# res: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765]
# call times: 57291

# 采用默记法的动态规划，会大大减少调用次数
def fast_fib(n, memo):
    global calls
    calls += 1
    if n not in memo:
        memo[n] = fast_fib(n-1, memo) + fast_fib(n-2, memo)
    return memo[n]

def fib(n):
    memo = {0: 0, 1: 1}
    return fast_fib(n, memo)

calls = 0
n = 20
res = [fib(i) for i in range(n+1)]
print 'fibonacci n: {}\nres: {}\ncall times: {}'.format(n, res, calls)

# Output
# fibonacci n: 20
# res: [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765]
# call times: 401
```

# 动态规划-背包问题

> 背包最大负重为5，现有物品重量及价值如下：  
> weights = \[5, 3, 2\], values = \[9, 7, 8\]  
> 求背包能选择物品的最大价值  
> 采用深度优先及左侧优先的决策树算法,每个节点的元素表示: 当前物品索引i（右侧开始），剩余背包负重w，当前背包内物品价值v  
> ![](assets/knapsack.png)

```py
def max_val(w, v, i, aw):
    print 'called with: {}, {}'.format(i, aw)
    global calls
    calls += 1
    if i == 0:
        return v[i] if w[i] <= aw else 0
    without_i = max_val(w, v, i-1, aw)
    if w[i] > aw:
        return without_i
    else:
        with_i = v[i] + max_val(w, v, i-1, aw-w[i])
    return max(with_i, without_i)

calls = 0
weights = [5, 3, 2]
values = [9, 7, 8]
res = max_val(weights, values, len(values)-1, 5)
print 'max val res: {}, call times: {}'.format(res, calls)

# Output
# called with: 2, 5
# called with: 1, 5
# called with: 0, 5
# called with: 0, 2
# called with: 1, 3
# called with: 0, 3
# called with: 0, 0
# max val res: 15, call times: 7

# 随着物品的增多及背包负重的增加，该递归算法的方法调用次数会急剧增加，
# 同样，也会以相同参数多次调用方法，
# 类似递归方式斐波那契数列的实现。
```

#### 背包问题的动态规划算法

```py
def fast_max_val(w, v, i, aw, m):
    print 'called with: {}, {}'.format(i, aw)
    global calls
    calls += 1
    try:
        return m[(i, aw)]
    except KeyError:
        if i == 0:
            if w[i] <= aw:
                m[(i, aw)] = v[i]
                return v[i]
            else:
                m[(i, aw)] = 0
                return 0
        without_i = fast_max_val(w, v, i-1, aw, m)
        if w[i] > aw:
            m[(i, aw)] = without_i
            return without_i
        else:
            with_i = v[i] + fast_max_val(w, v, i-1, aw-w[i], m)
        res = max(with_i, without_i)
        m[(i, aw)] = res
        return res


def max_val(w, v, i, aw):
    m = {}
    return fast_max_val(w, v, i, aw, m)

calls = 0
# weights = [5, 3, 2]
# values = [9, 7, 8]
weights = [1, 1, 5, 5, 3, 3, 2, 2, 6, 6, 4, 4]
values = [15, 15, 9, 9, 7, 7, 8, 8, 16, 16, 10, 10]
res = max_val(weights, values, len(values)-1, 8)
print 'max val res: {}, call times: {}'.format(res, calls)
```

> #### _总结：动态规划就是在处理重叠的子问题时，采用默记法将子问题的结果缓存下来，避免对已知结果的子问题的重复调用。它的基本思想是以空间换时间。_

# 二分查找

```python
def binary_search(array, n, imin, imax):
    imid = (imax + imin)/2
    print 'find {} in {}'.format(n, array[imin:imax])
    print 'mid index={}'.format(imid)
    if array[imid] == n:
        return imid
    if array[imid] > n:
        imax = imid
    else:
        imin = imid
    return binary_search(array, n, imin, imax)


if __name__ == '__main__':
    array = [1, 3, 5, 6, 7, 9, 10]
    print binary_search(array, 10, 0, len(array))

# Output
# find 10 in [1, 3, 5, 6, 7, 9, 10]
# mid index=3
# find 10 in [6, 7, 9, 10]
# mid index=5
# find 10 in [9, 10]
# mid index=6
# found target on index 6
```

# 最大连续子列和问题

```python
def max_subseq_sum(array):
    max_sum = this_sum = 0
    this_tmp = []
    max_sub = []
    for i in array:
        this_sum += i  # 向右累加
        this_tmp.append(i)  # 将元素添加到临时列表中
        if this_sum > max_sum:
            max_sum = this_sum  # 发现更大和则更新当前结果
            max_sub = this_tmp[:]  # 同时更新最大子列
        elif this_sum < 0:
            this_sum = 0  # 如果当前子列和为负,则不可能使后面的子列和增大,抛弃之
            this_tmp = []
    print max_sub
    return max_sum


if __name__ == '__main__':
    array = [-1, 3, -2, 4, -6, 1, 6, -1]
    print 'max subseq sum is {}'.format(max_subseq_sum(array))

# Output
# [1, 6]
# max subseq sum is 7
```



