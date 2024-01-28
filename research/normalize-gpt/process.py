from index import *


def process_file(input, output):
    inputFile  = open(input, "r")
    outputFile = open(output, "w")
    inputs     = inputFile.read().split("------INPUT------\n")

    buffer = ""
    for input in inputs:
        if input == "":
            continue
        input = input.strip()
        buffer += "------INPUT------\n" + input + "\n------OUTPUT------\n" + getResponse(input) + "\n"
    outputFile.write(buffer.strip())

    outputFile.close()
    inputFile.close()


# process_file("../../tests/flint-8.txt", "../../tests/flint-8-output.txt")