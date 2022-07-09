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

Here is a sample run:

```
flashcards ▶ flashcards
Pick a topic to practice:
0: Chapter 1
1: Chapter 2
Type the number of the topic to practice: 1
Q: Who is your cat?
A? Shortcut
✅ 1/1 CORRECT
Q: Who is your favorite person in the 0 to 3 age catgory?
A? Diana
✅ 2/2 CORRECT
👏👏👏 GOOD JOB PRACTICING. Hope to see you again soon!
```
