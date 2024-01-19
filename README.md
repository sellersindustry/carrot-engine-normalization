# normalization-tts
- Step 1 - Detect
    - words
    - numbers
    - currency
    - punctuation
    - symbols
    - hashtags
    - special
        - email
        - dates
        - hashtags
        - urls



    - acronym
    - initialism
    - abbreviation (punctuation)
- Step 2 - Classification
    - word
        - unit
            - overlapping units
            - m, k, b (after number)
            - cm, km (meters or millions???)
        - proper nouns
    - dates ??
    - numbers
        - Roman Numeral `John XI` -> `John the nineth`
            - make posesive roman numeral
        - phone numbers (Phone number 911 or 311)
        - years
            - remove years from the other classification
    - symbols
        - punctuation
        - math operator (numbers both sides)
            - `-`, `+`, `-`, `*`, `^`, `/`
            - `2/3` -> two thirds
            - `21/32` -> twenty one divied by thirty two
        - math prefix (`+` or `-` before number, no number behind)
            - `+` - positive
            - `-` - negative
            - `=` - equals
        - quotes
            - Quote Long Start (quote)
            - Quote Long End (end quote)
            - Quote Word (quote on quote)
        - per
            - 8/year (8 per year)
        - non-silent
            - `@` - at
            - `#` - pound
            - `$` - dollars
            - `%` - precent
            - `&` - and
            - `-` - to
            - `=` - is
            - letters
        - `%` - word unit
        - everything else is silent
- Step 3 - Division
    - emails
    - hashtags
    - phone
    - url
    - phone
    - dates
    - numbers
- Step 4 - Expansion
- Step 5 - Grouping
    - math equations
    - unit
    - nouns and posessive roman numerals


## Issues
- Roman Numeral or `I`
- Roman Numeral `John XI` -> `John the nineth`
- Phone number 911 or 311
- Dates desifer
- "m" or "k" or "b" after a number
- quote, end quote
- colloquial terms - `9/11` -> `9 11`


## Notes
- https://github.com/neurosnap/sentences?tab=readme-ov-file
- https://tomaarsen.github.io/SpanMarkerNER/index.html