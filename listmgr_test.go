package yfgo

import (
    "testing"
)

func TestSymbolList(t *testing.T) {
    symlist := SymbolList()
    if len(symlist) == 0 {
        t.Errorf("The list of symbols is empy")
    }
}

func TestSymbolAddRemove(t *testing.T) {
    sym := "__TEST__"
    SymbolAdd(sym)
    if !SymbolExists(sym) {
       t.Errorf("Symbol was not added") 
    }
    SymbolRemove(sym)
    if SymbolExists(sym) {
       t.Errorf("Symbol was not removed") 
    }
}
