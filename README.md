Flashcards
==========

A simple flashcard program. Do `go install` to install then type
`flashcards` to run it in the current directory. When it runs, it
finds all topic files in the current directory and loads each one
as a separate topic. It then asks you which topic to practice and
walks you through the cards with an updating score.

A properly formatted topic file ends with the `.topic` extension
and is a JSON file that follows this format:

```
{
  "title": "Chapter 1",
  "cards": [
    {"q": "Who are you?",
     "a": "Rick"},
    {"q": "Who is your favorite person in the 4 to 33 age catgory?",
     "a": "Celeste"}
  ]
}
```

...add as many cards as you want.
