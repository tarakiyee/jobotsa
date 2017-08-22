package main

import (
  "flag"
  "fmt"
  "math/rand"
  "os"
  "time"
  "log"
  "dapplebeforedawn/markovmotron/markov_util"
)

var numWords    = flag.Int("words", 100, "maximum words to print")
var prefixLen   = flag.Int("prefix", 2, "number of prefix words")
var corpus_path = flag.String("corpus", "./corpus.txt", "corpus text")
var input_path  = flag.String("react-to", "", "something for the machine to respond to")

func main() {

  flag.Parse()

  corpus, err := os.Open(*corpus_path)
  if err != nil { log.Fatal(err) }

  var input *os.File
  switch name := flag.Arg(0); {
  case name == "":
    input = os.Stdin
  default:
    input, err = os.Open(*input_path)
    if err != nil { log.Fatal(err) }
  }

  rand.Seed(time.Now().UnixNano())

  c := markov_util.NewChain(*prefixLen)
  c.Build(corpus)

  text := c.Generate(*numWords, input)
  fmt.Println(text)
}
