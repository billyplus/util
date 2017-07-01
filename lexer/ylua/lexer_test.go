package ylua

import (
	"testing"

	"git.erigame.com/billyplus/util/test"
)

func Test_newLexer(x *testing.T) {
	//t:=test.T
	t := (*test.T)(x)
	t.Log("hello")
	l := lex("test", `   first123="googd", name="my name"
		
		'testsible'
			"test"`)

	expected := []TokenType{
		TokIdent,
		TokAssign,
		TokDoubleQuot,
		TokIdent,
		TokAssign,
		TokDoubleQuot,
		TokSingleQuot,
		TokDoubleQuot,
	}
	valExpected := []string{
		"first123",
		"=",
		"\"googd\"",
		"name",
		"=",
		"\"my name\"",
		"'testsible'",
		"\"test\"",
	}

	i := 0
	for tok := l.nextToken(); tok.typ != TokEOF; tok = l.nextToken() {
		if tok.typ == TokError {
			t.Logf("tok is %v", tok.typ.String())
		}

		t.Logf("%v---%v---%v---%v\n", i, tok.typ.String(), expected[i].String(), tok.val)
		// if tok.typ != expected[i] {
		// 	t.Errorf("got %v expected %v at %v\n", tok.typ.String(), expected[i].String(), tok.val)
		// }
		t.Assert(tok.typ == expected[i], "got %s expected %s at %s\n", tok.typ.String(), expected[i].String(), tok.val)
		t.Assert(tok.String() == valExpected[i], "got %s expected %s\n", tok.String(), valExpected[i])

		i++
	}
}
