def even_filter(nums):
    return filter(lambda x: x % 2 == 0, nums)


def multiply_by_three(nums):
    return map(lambda x: x * 3, nums)


def output(nums):
    return map(lambda x: "The numbers is: %s" % x, nums)


def pipe(nums, function):
    return reduce(lambda x, f: f(x), function, nums)

nums = [1, 2, 3, 4, 5, 6]
nums = pipe(nums, [
    even_filter,
    multiply_by_three,
    output
])

for n in nums:
    print n
