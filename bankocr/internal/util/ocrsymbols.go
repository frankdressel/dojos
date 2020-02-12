package util

import (
  "errors"
  "fmt"
)

type OCRSymbol struct {
  Line0, Line1, Line2 string
}

type SymbolIterator struct {
  threelines []string
  index int
  linelength int
}

func NewSymbolIterator(threelines []string) (*SymbolIterator, error) {
  if len(threelines) != 3 {
    return nil, errors.New(fmt.Sprintf("Invalid argument: Three lines needed but received %d", len(threelines)))
  }
  for i := 1; i < len(threelines); i++ {
    if len(threelines[i]) != len(threelines[i-1]) {
      return nil, errors.New(fmt.Sprintf("Invalid argument: Lines have different lengths %d and %d", len(threelines[i]), len(threelines[i-1])))
    }
  }

  linelength := len(threelines[0])
  s := SymbolIterator{threelines, -3, linelength}

  return &s, nil  
}

func (s *SymbolIterator) Next() bool {
  s.index = s.index + 3
  return s.index + 3 <= s.linelength
}

func (s *SymbolIterator) Value() OCRSymbol {
  return OCRSymbol{
    s.threelines[0][s.index:s.index + 3],
    s.threelines[1][s.index:s.index + 3],
    s.threelines[2][s.index:s.index + 3]}
}
