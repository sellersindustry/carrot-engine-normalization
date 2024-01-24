from openai import OpenAI
import os
from dotenv import load_dotenv

load_dotenv()

OPENAI_TOKEN = os.getenv('OPENAI_TOKEN')
# print(OPENAI_TOKEN)

client = OpenAI(api_key=OPENAI_TOKEN)

def get_response(prompt):
  response = client.chat.completions.create(
    model="gpt-3.5-turbo",
    messages=[{"role": "user", "content": prompt}], 
    temperature = 0)
  return response.choices[0].message.content

def normalize_sentence_with_gpt(sentence):
    prompt = """
    Help me to normalize the sentence with the following steps
    Step 1 - Keep commas and puncuations
    Step 2 - Remove dashes in words for example "4-year-old boy" to "4 year old boy"
    Step 3 - Convert special symbols for example "$" to "dollar", "%" to "percent" 
    Step 4 - Convert date to words. For example "1984" to "Nineteen Eighty Four", "1984/4/12" to "Nineteen Eighty Four April twelveth"
    Step 5 - Convert all numbers. For example "4" to "four", "10" to "ten"
    Step 6 - Intialisms. For example "FBI" should become "F B I", "USA" should be "U S A" 
    Step 7 - Convert unit to full word. For example "cm" to "centimeter", "ft" to "feet" 
    Step 8 - Convert math symbol. For example "+" to plus, "*" to "multiply"
    """ + sentence
    response = get_response(prompt)
    print(response)

# sentence = "100 + 200 = 300"
# normalize_sentence_with_gpt(sentence)