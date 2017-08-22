package markov_util

import (
  "bufio"
  "io"
  "strings"
)

type Prefix []string

func (p Prefix) String() string {
  return strings.Join(p, "")
}

func (p Prefix) Shift(word string) {
  copy(p, p[1:])
  p[len(p)-1] = word
}

func PrefixFromSeed(prefixLen int, seed io.Reader) *Prefix {
  p    := make(Prefix, prefixLen)
  scan := bufio.NewScanner(seed)
  scan.Split(bufio.ScanWords)

  for scan.Scan() {
    copy(p[:prefixLen-1], p[1:prefixLen])
    p[prefixLen-1] = scan.Text()
  }
  return &p
}
