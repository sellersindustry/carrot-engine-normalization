from normalise import normalise, rejoin
from datasets import load_dataset
import nltk
import time

for dependency in ("brown", "names", "wordnet", "averaged_perceptron_tagger", "universal_tagset", "stopwords", "punkt"):
    nltk.download(dependency)

def evalulate(text):
    try:
        return rejoin(normalise(text))
    except:
        return text


def calculate_wer(reference, hypothesis):
	ref_words = reference.split()
	hyp_words = hypothesis.split()
	substitutions = sum(1 for ref, hyp in zip(ref_words, hyp_words) if ref != hyp)
	deletions = len(ref_words) - len(hyp_words)
	insertions = len(hyp_words) - len(ref_words)
	total_words = len(ref_words)
	wer = (substitutions + deletions + insertions) / total_words
	return wer

testing_data = load_dataset("sellersew/carrot-engine-normalization-translation-v2", data_files="native-tests.csv")

import time
from tqdm import tqdm

start_time = time.time()
passing    = 0
failing    = 0
WER        = []

for entry in tqdm(testing_data["train"], desc='Processing', unit='iteration'):
    input  = entry["input"]
    output = evalulate(input)
    WER.append(calculate_wer(input, output))
    if input == output:
        passing += 1
    else:
        failing += 1
    
time_taken = time.time() - start_time


average = sum(WER) / len(WER)
print("Passing: " + str(passing))
print("Failing: " + str(failing))
print("WER    : " + str(average))
print(f'Time : {time_taken:.2f} seconds')

