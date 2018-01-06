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
		l := []byte(`{"109":{"childId":109,"type":3,"reward":[[{"limit":1,"item":50087,"cost":[3,88,188]},{"limit":1,"item":50088,"cost":[3,288,588]}],[{"limit":1,"item":50089,"cost":[3,88,188]},{"limit":1,"item":50090,"cost":[3,288,588]}],[{"limit":1,"item":50091,"cost":[3,88,188]},{"limit":1,"item":50092,"cost":[3,288,588]}],[{"limit":1,"item":50093,"cost":[3,88,188]},{"limit":1,"item":50094,"cost":[3,288,588]}],[{"limit":1,"item":50095,"cost":[3,88,188]},{"limit":1,"item":50096,"cost":[3,288,588]}],[{"limit":1,"item":50097,"cost":[3,88,188]},{"limit":1,"item":50098,"cost":[3,288,588]}],[{"limit":1,"item":50099,"cost":[3,88,188]},{"limit":1,"item":50100,"cost":[3,288,588]}]],"skintype":"7day","iconskin":"temp_json.temp_limitbuy"}}`)
		var v map[string]interface{}
		if err := json.Unmarshal(l, &v); err != nil {
			b.Error(err.Error())
		}
	}
}
