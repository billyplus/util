package ylua

import (
	"testing"

	"encoding/json"

	"git.erigame.com/billyplus/util/test"
)

func TestNewLexer(x *testing.T) {
	//t:=test.T
	t := (*test.T)(x)
	t.Log("hello")
	l := lex("test", `   {stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999}`)

	expected := []TokenType{
		TokLBrace,
		TokIdent,
		TokAssign,
		TokDoubleQuot,
		TokIdent,
		TokAssign,
		TokLBrace,
		TokDoubleQuot,
		TokDoubleQuot,
		TokDoubleQuot,
		TokRBrace,
		TokIdent,
		TokAssign,
		TokInt,
		TokIdent,
		TokAssign,
		TokFloat,
		TokRBrace,
	}
	valExpected := []string{
		"{",
		"stringval",
		"=",
		"\"stringval\"",
		"tableval",
		"=",
		"{",
		"\"val1\"",
		"\"val2\"",
		"\"val3\"",
		"}",
		"intval",
		"=",
		"3232",
		"floatval",
		"=",
		"334.999",
		"}",
	}

	i := 0
	for tok := l.nextToken(); tok.typ != TokEOF; tok = l.nextToken() {
		if tok.typ == TokError {
			t.Logf("tok is %v", tok.typ.String())
		}

		//t.Logf("%v---%v---%v---%v\n", i, tok.typ.String(), expected[i].String(), tok.val)
		// if tok.typ != expected[i] {
		// 	t.Errorf("got %v expected %v at %v\n", tok.typ.String(), expected[i].String(), tok.val)
		// }
		// fmt.Printf("%q\n", tok.typ)
		t.Assert(tok.typ == expected[i], "got %s expected %s at %s\n", tok.typ.String(), expected[i].String(), tok.val)
		t.Assert(tok.String() == valExpected[i], "got %s expected %s\n", tok.String(), valExpected[i])

		i++
	}
}

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := lex("test", `   {stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999,
stringval="stringval", tableval={"val1","val2","val3"}, intval=3232, floatval=334.999}`)
		for tok := l.nextToken(); tok.typ != TokEOF; tok = l.nextToken() {
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := []byte(`   {"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,
"stringval"="stringval", "tableval"=["val1","val2","val3"], "intval"=3232, "floatval"=334.999,}`)
		var v map[string]interface{}
		json.Unmarshal(l, &v)
	}
}
