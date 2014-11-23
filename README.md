# defme

`Defme` returns a word definition, synonyms, examples from `Wordnik API` and translation from https://github.com/sparrovv/gotr

### Prerequisites:

This program relies heavily on `wordnik.com` API. To use it, you need a wordnik key set in the environment.
(Sign in and get it on [developer.wordnik.com](http://developer.wordnik.com/))

```sh
export WORDNIK_API_KEY='xxxxxxxxxx'
```

### CLI

```sh
› ./defme d --to pl scry
Translation:
   wróżyć z kuli szklanej
Definition:
   To see or predict the future by means of a crystal ball.
Related:
   descry, crystal-gaze
Examples:
   Dee was haunted by his shortcomings: "You know I cannot see, nor scry" he lamented.
   The Savant (aka nobody knows his identity) - cybernetic mathematical supergenius who can scry into the futures of many possible timelines.
   Annika had already warped past aneurismal straight into action, dispatching search parties and hiring witches to scry.
   Rydstrom told Sabine, “Cwena, the witch will scry for Lanthe—”
   Perhaps she was using the surface of the water to scry into her future.
```

### HTTP server

```sh
./defme server --port 9292

curl "localhost:9292?to=pl&word=scry"

{"translation":"wróżyć z kuli szklanej","extraTranslations":["wróżyć z kuli szklanej"],"definitions":["To see or predict the future by means of a crystal ball."],"synonyms":["descry","crystal-gaze"],"examples":["   Dee was haunted by his shortcomings: \"You know I cannot see, nor scry\" he lamented.","   The Savant (aka nobody knows his identity) - cybernetic mathematical supergenius who can scry into the futures of many possible timelines.","   Annika had already warped past aneurismal straight into action, dispatching search parties and hiring witches to scry.","   Rydstrom told Sabine, “Cwena, the witch will scry for Lanthe—”","   Perhaps she was using the surface of the water to scry into her future."]}
```

### TODOS:

1. Better exception handling and error reporting
1. Add performance tests
1. Add caching
