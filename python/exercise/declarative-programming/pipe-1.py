def even_filter(nums):
    for n in nums:
        if n % 2 == 0:
            yield n


def multiply_by_three(nums):
    for n in nums:
        yield n * 3


def output(nums):
    for n in nums:
        yield "The number is: %s" % n

nums = [1, 2, 3, 4, 5, 6]
nums = output(multiply_by_three(even_filter(nums)))
for n in nums:
    print n
