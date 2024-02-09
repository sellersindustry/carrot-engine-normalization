# Text Normalization for Text-to-Speech 
Text normalization is a fundamental aspect of TTS, and is highly dependent on natural language processing (NLP). But text normalization for TTS is very different from text normalization for applications. Unlike typical NLP, TTS normalization focuses on transforming non-standard words (numbers, symbols, abbreviations, quotations, etc.) into an understandable format for speech synthesis, addressing issues with numerical representations and abbreviations.

In TTS every symbol and entity that our brains normally process without a second thought must be transformed. For instance, normalization would take the text "\$20.43" and convert it into its spoken word form, "twenty dollars and forty-three cents." Additionally, different types of acronyms must be considered like "FBI" to "F B I," but "NASA" to "nasa". These examples just cover a few of the syntax-based problems that the normalization system will encounter, there are plenty more. Because of these complications, normalization of text requires a deep understanding of written language. 

The goal of this project is build a fast and customizable text normalization system based off the [Normalise Library](https://github.com/EFord36/normalise). We took two different approaches. First, a pattern recongization approach which was difficult to build and can be finicky, but was very fast. Second, a GPT-2 based model, which was much slower, and required alot of training to get it to be accurate.

## Benchmarks
| System | Pass | Fail | WER | Time (secs) |
|---|---|---|---|---|
| [Nvidia NeMo](https://github.com/NVIDIA/NeMo-text-processing)             | 347 | 559 | 0.605448141189473  | 130.87 |
| [Normalise](https://github.com/EFord36/normalise)               | 12  | 894 | 0.6641477767423963 | 63.81  |
| Pattern Recognized | 879 |  27 | 0.007661769201242707 | 12.433961 |
| GPT2 - 10 Epoch    | 285 | 621 | 1.4240638465147457   | 470.03    |
| GPT2 - 20 Epoch    | 237 | 669 | 1.2405701855542082   | 475.08    |


## Process
Expanded on from research from [Flint's Normalise Library](https://github.com/EFord36/normalise)

**Step 1 - Detection** \
Identify class of tokens such as controls, spaces, dates, times, phone numbers, numbers, hashtags, words, and symbols.

**Step 2 - Classification** \
Classify each token subclass based on it's context to other words.

**Step 3 - Transform** \
Based on each tokens class and subclass classification wordify each word.


## Known Issues and Futures Notes
- Recognizing Unicode Characters like (×, ÷, ², ³, ½, ‐ (dashes), “, Ⅻ)
- Recognizing Unicode Characters for euro and pounds
- Differation between "3 - 4" and if it's "three minus four" or "three to four"
- Differation between "3 x 4" and if it's "three time four" or "three by four"
- Differation between initialism and roman numerals
- Differation between initialism and acryonym
- Differation between 1 and I
- Differation between in and inches
- Differation between (m) million and meters
- Chunking Domains


## Notes for Future
- **Step 1 - Remapping** - Unicode remapping, remap odd unicode characters to their regular form. For example `“` to `"`, `÷` to `/`, `×` to `*`, `½` to `1 / 2`, all the dashes to a regular dash.
- **Step 2 - Splitting** - Detecting if new line is used to wrap text or as a new pharagraph. Do this by splitting at each new line, getting the length of each line, and calulating standard deviation. Split the text into each paragraph. Split text into each paragraph by sentence.
- **Step 3 - Tokenization** - Split each token at space and punctuation. Ensure some elements are grouped togeather properly, like `1/2 + 23`, `$23-34 million`, `1.23x10^3`, and phone numbers. Keeping elements togeather will help will contextual understanding and classification. This also prevent's items from being degroupped, like `IXs`, `$23.43`, and `4's` from being split odd. Tokenization can classify both class and subclass. Detect and remap coloclio terms like `WW1`, `WW2`, and `9/11`.
- **Step 4 - Classification** - Classify tokens which don't have a subclass yet, based on context.
- **Step 5 - Transform** - Transform each token into it's word form.


## Notes
- https://huggingface.co/datasets/sellersew/carrot-engine-normalization-translation-v2
- https://huggingface.co/sellersew/carrot-engine-normalization
- https://www.kaggle.com/datasets/google-nlu/text-normalization
- https://huggingface.co/google/flan-t5-base
- https://github.com/neurosnap/sentences?tab=readme-ov-file
- https://tomaarsen.github.io/SpanMarkerNER/index.html
- https://platform.openai.com/docs/guides/fine-tuning

