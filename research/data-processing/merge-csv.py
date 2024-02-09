import pandas as pd
import csv


INPUT_FILES = [
    "output.csv",
    "output2.csv",
]
OUTPUT_FILE = "output-final.csv"


def merge_csv_files(file_list):
    all_data = pd.DataFrame()
    for file_path in file_list:
        df = pd.read_csv(file_path)
        all_data = all_data.append(df, ignore_index=True)
    return all_data


merge_csv_files(INPUT_FILES).to_csv(OUTPUT_FILE, index=False, quoting=csv.QUOTE_NONNUMERIC)

