package util

import (
  "testing"
)

func TestNewSymbolIterator(t *testing.T) {
  _, err := NewSymbolIterator([]string{"", "", ""})
  if err != nil {
    t.Fail()
    t.Log("Constructor function failed")
  } 

  _, err = NewSymbolIterator([]string{"", ""})
  if err == nil {
    t.Fail()
    t.Log("Constructor function accepts wrong input")
  } 
  
  _, err = NewSymbolIterator([]string{""})
  if err == nil {
    t.Fail()
    t.Log("Constructor function accepts wrong input")
  } 
}

func TestSymbolIteratorNextWithCorrectLineLength(t *testing.T) {
  teststring := ""
  for i := 0; i < 6; i++ {
    teststring = teststring + "   "
  }
  si, _ := NewSymbolIterator([]string{teststring, teststring, teststring})

  for i := 0; i < 6; i++ {
    next := si.Next()
    if(!next) {
      t.Fail()
      t.Logf("Existing symbol %d not iterated", i)
    }
  }

  if si.Next() != false {
    t.Fail()
    t.Log("Iteration does not stop correctly")
  }
}

func TestSymbolIteratorNextWithWrongLineLength(t *testing.T) {
  teststring := ""
  for i := 0; i < 6; i++ {
    teststring = teststring + "   "
  }
  teststring = teststring + "  "
  si, _ := NewSymbolIterator([]string{teststring, teststring, teststring})

  for i := 0; i < 6; i++ {
    next := si.Next()
    if(!next) {
      t.Fail()
      t.Logf("Existing symbol %d not iterated", i)
    }
  }

  if si.Next() != false {
    t.Fail()
    t.Log("Less than 3 symbols gets parsed")
  }
}

func TestSymbolIteratorValue(t *testing.T) {
  si, _ := NewSymbolIterator([]string{"abc", "def", "ghi"})

  si.Next()
  exp := OCRSymbol{"abc", "def", "ghi"}
  if si.Value() != exp {
    t.Fail()
    t.Log("Wrong value returned")
  }
}
