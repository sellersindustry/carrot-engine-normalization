from openai import OpenAI
import os
from dotenv import load_dotenv
load_dotenv()

OPENAI_TOKEN = os.getenv("OPENAI_TOKEN")


task = """
	Transform the text into how it would be spoken word for word by a human or
	text-to-speech system. Ensure only letters, commas, and puncuations is left,
	no symbols or numbers. Ensure the following steps are followed.
    Step 1 - Keep only letters, commas and puncuations. Convert all symbols to their words forms based on the context.
    Step 2 - Convert special symbols for example "$" to "dollar", "%" to "percent"
	Step 3 - Convert quotes short quotes under 2 words to say "quote on quote" and then the phrase.
	Step 4 - Convert longer quotes to say "quote" at the begining and "end quote" at the end.
	Step 5 - Convert all numbers and roman numerals to words. For example "4" to "four", "10" to "ten"
	Step 6 - Convert all time to the appropiate word. For example "12:00" to "Twelve o'clock", "12:50" to "Twelve fifty".
	Step 7 - Convert all years to the appropiate word. For example "1984" to "Nineteen Eighty Four".
	Step 8 - Convert all dates to their spoken form. For example "1984/4/12" to "April twelveth, Nineteen Eighty Four"
    Step 9 - Convert all intialisms to seperation for each letter. For example "FBI" should become "F B I", "U.S.A" should be "U S A" 
    Step 10 - Convert unit to full word. For example "cm" to "centimeter", "ft" to "feet" 
	Step 11 - Convert abbreviation to full word. For example "Mr." to "mister".
    Step 12 - Convert math symbol. For example "+" to plus, "*" to "multiply".
"""


examples =[[
	"The new race-car that I bought is really fast. It only cost about $123.43 and it a 2024 model. Mr. Bob III does all the maintenance work on the car for only about $20-30/year. You can reach out to him at (123) 456-7890 or bob3rd@gmail.com.",
	"The new race-car that I bought is really fast. It only cost about one hundred and twenty three dollars and four three cents and it a twenty twenty four model. Mister Bob the third does all the maintenance work on the car for only about twenty to thirty dollars per year. You can reach out to him at one two three, four five six, seven eight nine zero or bob three r d at g mail dot com."
], [
	"-1 + 32 ^ 123",
	"negative one plus three to the power of one hundred and twenty three"
]]


def getResponse(sentence):
	response = OpenAI(api_key=OPENAI_TOKEN).chat.completions.create(
		model="gpt-3.5-turbo",
		messages=[{
			"role": "user",
			"content": task
        }, {
			"role": "user",
			"content": examples[0][0]
		}, {
			"role": "system",
			"content": examples[0][1]
		}, {
			"role": "user",
			"content": examples[1][0]
		}, {
			"role": "system",
			"content": examples[1][1]
		}, {
			"role": "user",
			"content": sentence
		}],
		temperature = 0
    )
	return response.choices[0].message.content


if __name__ == "__main__":
	sentence = input("Enter sentence: ")
	print(getResponse(sentence))

