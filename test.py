# import nltk

# # sentence = """Hello world dec. This"""
# # tokens = nltk.word_tokenize(sentence)
# # print(tokens)

# sentence = "Reilly"
# tagged_sent = nltk.pos_tag(sentence.split())
# propernouns = [word for word,pos in tagged_sent if pos == 'NNP']
# print(propernouns)

from span_marker import SpanMarkerModel

model    = SpanMarkerModel.from_pretrained("tomaarsen/span-marker-bert-base-fewnerd-fine-super")
print("Starting...")
entities = model.predict("Amelia Earhart flew her single engine Lockheed Vega 5B across the Atlantic to Paris.")
print(entities)