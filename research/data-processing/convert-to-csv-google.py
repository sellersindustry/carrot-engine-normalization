import csv

"""
    For processing google challange data
    https://www.kaggle.com/competitions/text-normalization-challenge-english-language
"""


INPUT_FILE   = "../google-data/google-raw-6.csv"
OUTPUT_FILE  = "../google-data/google-6.csv"


sentences    = [["input", "output"]]
bufferOutput = ""
bufferInput  = ""
isInQuote    = False
with open(INPUT_FILE, newline="", encoding="utf-8") as csvfile:
    reader = csv.reader(csvfile)
    for row in reader:

        if row[0] == "<eos>":
            bufferOutput = bufferOutput.strip()
            print(bufferOutput)
            if "_letter" not in bufferOutput and "_number" not in bufferOutput and "sil" not in bufferOutput:
                sentences.append([bufferInput, bufferOutput])
            bufferOutput = ""
            bufferInput  = ""
            continue

        if row[0] == "PUNCT":
            if row[1] in [".", "!", "?", ":", ";", ")", "]"]:
                bufferInput  = bufferInput.strip() + row[1] + " "
                bufferOutput = bufferOutput.strip() + row[1] + " "
                continue
            if row[1] in ["(", "["]:
                bufferInput  = row[1]
                bufferOutput = row[1]
                continue
            if row[1] == "\"":
                if isInQuote:
                    bufferInput  = bufferInput.strip() + row[1] + " "
                    bufferOutput = bufferOutput.strip() + row[1] + " "
                else:
                    bufferInput  = row[1]
                    bufferOutput = row[1]
                isInQuote = not isInQuote;
                continue


        if row[2] == "<self>":
            bufferInput  += row[1] + " "
            bufferOutput += row[1] + " "
        else:
            bufferInput  += row[1] + " "
            bufferOutput += row[2] + " "


with open(OUTPUT_FILE, "w", newline='', encoding="utf-8") as csvfile:
    writer = csv.writer(csvfile, quoting=csv.QUOTE_NONNUMERIC)
    writer.writerows(sentences)

