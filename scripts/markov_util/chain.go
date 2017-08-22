package markov_util

import (
  "bufio"
  "io"
  "math/rand"
  "strings"
)

type Chain struct {
  chain     map[string][]string
  prefixLen int
}

func NewChain(prefixLen int) *Chain {
  return &Chain{ make(map[string][]string), prefixLen }
}

func (c *Chain) Build(r io.Reader) {
  scan := bufio.NewScanner(r)
  p    := make(Prefix, c.prefixLen)
  scan.Split(bufio.ScanWords)
  for scan.Scan() {
    key         := p.String()
    c.chain[key] = append(c.chain[key], scan.Text())
    p.Shift(scan.Text())
  }
}

func (c *Chain) Generate(n int, seed io.Reader) string {
  p := PrefixFromSeed(c.prefixLen, seed)

  var words []string
  for i := 0; i <n; i++ {
    choices := c.chain[p.String()]
    if choices == nil { choices = c.RandomChain() }
    if len(choices) == 0   { break }
    next := choices[rand.Intn(len(choices))]
    words = append(words, next)
    p.Shift(next)
  }
  return strings.Join(words, " ")
}

func (c *Chain) RandomChain() (chain []string) {
  stop := rand.Intn(len(c.chain))
  i    := 0
  for _, rand_chain := range c.chain {
    if i == stop {
      chain = rand_chain
      break
    }
    i++
  }
  return
}
