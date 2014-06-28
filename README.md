# defme - Find definition and translation to a word.

Command line word definer.

It gets the word definition from the `Wordnik API`, and uses `google translate` to translate it.

### Prerequisites:

This program relies heavily on `wording.com` API, so you need a Wordnik API key set in the environment variables. (sign in and get it on [developer.wordnik.com](http://developer.wordnik.com/))

```sh
  export WORDNIK_API_KEY=''
```

### How to use it:

Simple CLI usage:

```sh
./defme d --to pl come along

Definition:
   To progress to the next level of player character stats and abilities, often by acquiring experience points in role-playing games.
Related:
Translation:
   poziom w górę
   podnieść do właściwego poziomu
```

It can returns JSON:

```sh
./defme d --to pl --json 1 level up
```

It has built-in HTTP server:

```sh
./defme server --port 9292

curl "localhost:9292?lang=pl&phrase=level%20up"

> {"translation":"poziom w górę","extraTranslations":["podnieść do właściwego poziomu"],"definitions":["To progress to the next level of player character stats and abilities, often by acquiring experience points in role-playing games."],"synonyms":null}
```

### TODOS:

1. Better error reporting
  1. server doesn't die if the google is not reachable
  1. server doesn't die when wordink is ...
  1. ... when there are connection problems
  1. ... it timeouts afte 5s

