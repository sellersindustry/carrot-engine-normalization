import os
import csv

INPUT_DIRECTORY = "../../tests"
OUTPUT_FILE     = "native-tests.csv"

data_array = [["input", "output"]]
for filename in os.listdir(INPUT_DIRECTORY):
    if filename.endswith(".txt"):
        with open(os.path.join(INPUT_DIRECTORY, filename), "r", encoding="utf-8") as file:
            buffer = file.read()
            for entry in buffer.split("------INPUT------\n"):
                entry = entry.split("------OUTPUT------\n")
                if len(entry) == 2:
                    data_array.append([entry[0].strip(), entry[1].strip()])


with open(OUTPUT_FILE, "w", newline='', encoding="utf-8") as csvfile:
    writer = csv.writer(csvfile, quoting=csv.QUOTE_NONNUMERIC)
    writer.writerows(data_array)
