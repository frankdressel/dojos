package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "strings"

  "github.com/frankdressel/dojos/bankocr/internal/util"
)

var (
  One = util.OCRSymbol{
    Line0: "   ",
    Line1: "  |",
    Line2: "  |"}
  Two = util.OCRSymbol{
    Line0: " _ ",
    Line1: " _|",
    Line2: "|_ "}
  Three = util.OCRSymbol{
    Line0: " _ ",
    Line1: " _|",
    Line2: " _|"}
  Four = util.OCRSymbol{
		Line0: "   ",
		Line1: "|_|",
		Line2: "  |"}
  Five = util.OCRSymbol{
		Line0: " _ ",
		Line1: "|_ ",
		Line2: " _|"}
  Six = util.OCRSymbol{
		Line0: " _ ",
		Line1: "|_ ",
		Line2: "|_|"}
  Seven = util.OCRSymbol{
		Line0: " _ ",
		Line1: "  |",
		Line2: "  |"}
  Eight = util.OCRSymbol{
		Line0: " _ ",
		Line1: "|_|",
		Line2: "|_|"}
  Nine = util.OCRSymbol{
		Line0: " _ ",
		Line1: "|_|",
		Line2: " _|"}
)

func number(o util.OCRSymbol) (uint8, error) {
  switch {
    case o == One:
      return 1, nil
    case o == Two:
      return 2, nil
    case o == Three:
      return 3, nil
    case o == Four:
      return 4, nil
    case o == Five:
      return 5, nil
    case o == Six:
      return 6, nil
    case o == Seven:
      return 7, nil
    case o == Eight:
      return 8, nil
    case o == Nine:
      return 9, nil
    default:
      return 0, errors.New(fmt.Sprintf("Unable to parse\n%s\n to number", o))
  } 
}

func main() {
  bytes, err := ioutil.ReadFile("testfile_large.txt")
  if err !=nil {
    log.Fatal(err)
  }
  lines := strings.Split(string(bytes), "\n")

  for i := 0; i <= len(lines) - 4; i = i + 4 {
    it, error := util.NewSymbolIterator(lines[i: i+3])
    if error != nil {
      continue
    }
    var num uint32 = 0
    for it.Next() {
      ocrnumber, err := number(it.Value())
      if err != nil {
        goto EndOfSymbolLine
      }
      num = num * 10 + uint32(ocrnumber)
    }

    fmt.Println(num)

    EndOfSymbolLine:
  }
}
