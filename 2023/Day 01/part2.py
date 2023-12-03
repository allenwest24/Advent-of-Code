# THIS DAY WAS AWFUL. I THINK THEY MESSED THE PROMPT UP. IGNORE THIS ONE.

# import re
# import functools

# with open('./input.txt', 'r') as f:
#     data = f.readlines()

# def calibrationValue(txt: list[str], faulty: bool) -> int:
#     nums = ['zero','one','two','three','four','five','six','seven','eight','nine','ten']
#     pattern = r'\d' if faulty else r'(?=(' + "|".join(nums) + "|\\d))"
#     digits: list[str] = [str(nums.index(digit)) if digit in nums else digit for digit in re.findall(pattern, txt)]
#     return int(digits[0] + digits[-1])

# print('Part 1', functools.reduce(lambda a, b: a + b, list(map(lambda line: calibrationValue(line, True), data))))
# print('Part 2', functools.reduce(lambda a, b: a + b, list(map(lambda line: calibrationValue(line, False), data))))