# Markovmotron
> A toy markovchain program

## Usage
```bash
echo "two words" | markovmotron --corpus="sample/corpus.txt"

#  option:
#    --corpus     a single line file of ordered sentances
#    --words      how many words to output
#    --prefix     how many history words the Markov should use to find the next word
#    --react-to   something for the Markov to respond to, defaults to standard in
```

## Installion
```
go build markovmotron.go
```

## Scripts
`markovmotron.coffee` is a hubot script to include markov fun in your chat room.  For deployment on heroku the compiled markovmotron will need to compiled for linux, and inside the hubot's `bin/` directory.

## Bonus
Binaries are included for OSX and Linux in the `bin/` directory.  Go makes cross-compiling _really_ easy.  =)
