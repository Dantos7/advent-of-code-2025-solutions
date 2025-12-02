from pathlib import Path
import time

def main():
    with Path("input-user.txt").open() as file:
        intervals = file.readline().strip()

    start = time.time()
    value = part1(intervals)
    end = time.time()
    print("Part 1 elapsed time:", format_elapsed_time(end - start))
    print("Part 1 result:", value)

    start = time.time()
    value = part2(intervals)
    end = time.time()
    print("Part 2 elapsed time:", format_elapsed_time(end - start))
    print("Part 2 result:", value)


def format_elapsed_time(seconds):
    if seconds >= 1:
        return f"{round(seconds, 2)} s"
    if seconds >= 0.001:
        return f"{round(seconds * 1000, 2)} ms"
    if seconds >= 0.000001:
        return f"{round(seconds * 1000000, 2)} µs"


def part1(intervals: str):
    total_sum = 0
    for interval in intervals.split(","):
        total_sum += get_repeated_numbers_sum_part1(interval)
    return total_sum


def part2(intervals: str):
    total_sum = 0
    for interval in intervals.split(","):
        total_sum += get_repeated_numbers_sum_part2(interval)
    return total_sum


def get_repeated_numbers_sum_part1(interval: str):
    start_end = interval.split("-")
    start, end = int(start_end[0]), int(start_end[1])
    return sum(i for i in range(start, end + 1) if is_repeated_part1(str(i)))

def is_repeated_part1(input):
    if len(input) % 2 == 0:
        return input[:len(input) // 2] == input[len(input) // 2:]
    return False

def get_repeated_numbers_sum_part2(interval: str):
    start_end = interval.split("-")
    start, end = int(start_end[0]), int(start_end[1])
    return sum(i for i in range(start, end + 1) if is_repeated_part2(str(i)))

def is_repeated_part2(input: str):
    for span_size in range(1, len(input) // 2 + 1):
        if len(input) % span_size == 0:
            repeat_count = len(input) // span_size
            repeated = input[:span_size] * repeat_count
            if repeated == input:
                return True
    return False

main()
