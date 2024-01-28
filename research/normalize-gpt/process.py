from index import *

def process_file(input, output):
    inputFile = open(input, "r")
    outputFile = open(output, "w")

    content = inputFile.read()
    sentences = content.split("------INPUT------")
    result = ""
    for x in range(1,len(sentences),2):
        # print(getResponse(sentences[x]))
        result = result + "------INPUT------" + sentences[x] +"------OUTPUT------\n" + getResponse(sentences[x])+"\n\t"
        # print(result)
    outputFile.write(result)

    outputFile.close()
    inputFile.close()

# process_file("tests/flint-1.txt", "tests/flint-1-output.txt")
process_file("tests/flint-2.txt", "tests/flint-2-output.txt")
# process_file("tests/flint-3.txt", "tests/flint-3-output.txt")
# process_file("tests/flint-4.txt", "tests/flint-4-output.txt")
# process_file("tests/flint-5.txt", "tests/flint-5-output.txt")
# process_file("tests/flint-6.txt", "tests/flint-6-output.txt")
# process_file("tests/flint-7.txt", "tests/flint-7-output.txt")
# process_file("tests/flint-8.txt", "tests/flint-8-output.txt")