

######## IMPORT PACKAGES ########

import re

########## DEFINE FXNS ##########
def numword_to_digit(string):
    num_map = {
        "zero": 0,
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    for k, v in num_map.items():
        if k in string:
            string = string.replace(k, str(v))
    return string


def clean_num_string(string):
    return re.sub(pattern="[^0-9]", repl="", string=string)


def first_last_nums_in_string(string, parse_numwords=False):
    first, last = None, None
    while not first:
        for idx, c in enumerate(string, 1):
            if c.isdigit():
                first = c
                break
            elif (
                parse_numwords
                and clean_num_string(numword_to_digit(string[:idx])).isdigit()
            ):
                first = clean_num_string(numword_to_digit(string[:idx]))
                break
    while not last:
        for idx, c in enumerate(reversed(string), 1):
            if c.isdigit():
                last = c
                break
            elif (
                parse_numwords
                and clean_num_string(numword_to_digit(string[-idx:])).isdigit()
            ):
                last = clean_num_string(numword_to_digit(string[-idx:]))
                break
    return int(first[0] + last[-1])


########## -------- ##########

def read_data(data_filepath):
  if data_filepath.startswith('http'):
    import requests
    data = requests.get(data_filepath).text
  else:
    data = open(data_filepath).read()
  return data.strip().split("\n")


def get_answers(data, return_answers=False):
    pt1_answer = sum(first_last_nums_in_string(line) for line in data if line)
    pt2_answer = sum(first_last_nums_in_string(line, parse_numwords=True) for line in data if line)
    print(f"Part 1 answer: {pt1_answer}", f"Part 2 answer: {pt2_answer}", sep="\n")
    if return_answers:
        return pt1_answer, pt2_answer

########## RUN ##########

if __name__ == "__main__":
    aoc_input_fp = "./input_era.txt"
    data = read_data(aoc_input_fp)
    get_answers(data)

