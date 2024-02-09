from normalise import normalise
import os
import time
import nltk
for dependency in ("brown", "names", "wordnet", "averaged_perceptron_tagger", "universal_tagset"):
    nltk.download(dependency)    
class Stat:
  passed = 0
  total = 0
  wer = 0
def process(file_name):
    inputFile = open(file_name, "r")
    lines = inputFile.read().replace("------OUTPUT------", "////").replace("------INPUT------", "////").split("////")
    result = Stat()
    for i in range(1,len(lines),2):
        input = lines[i]
        output = lines[i+1]
        normalize_output = normalise(input, verbose=True)
        if(normalize_output == output):
            result.passed = result.passed + 1
        result.total = result.total + 1
        result.wer = result.wer + calculate_wer(input, normalize_output)
    return result

def calculate_wer(reference, hypothesis):
	ref_words = reference.split()
	hyp_words = hypothesis.split()
	substitutions = sum(1 for ref, hyp in zip(ref_words, hyp_words) if ref != hyp)
	deletions = len(ref_words) - len(hyp_words)
	insertions = len(hyp_words) - len(ref_words)
	total_words = len(ref_words)
	wer = (substitutions + deletions + insertions) / total_words
	return wer

files = os.listdir('tests/')
tests_passed = 0
tests_total = 0
wer_total = 0
start_time = time.time()
for file in files:
    processed = process("tests/" + file)
    tests_passed = tests_passed + processed.passed
    tests_total = tests_total + processed.total
    wer_total = wer_total + processed.wer
print("TOTAL: " + str(tests_total) + "\n"+
      "PASSING: " + str(tests_passed) + "\n" +
      "Total time: " + str(time.time() - start_time) + " seconds\n" +
      "Average WER: " + str(wer_total/tests_total) + " \n")