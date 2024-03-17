package main

import (
	"testing"
)

func TestTokenParse(t *testing.T) {
	for key, val := range tokens {
		parser := parser{input: key}
		parsed := parser.parse()
		if parsed != val {
			t.Errorf("Expected %s, got %s", val, parsed)
		}
	}
}

func TestBasicSelect(t *testing.T) {
	parser := parser{input: "yoink * skibity table"}
	parsed := parser.parse()
	expected := "select * from table"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicWhere(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5"}
	parsed := parser.parse()
	expected := "select * from table where column = 5"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicAnd(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5 goon column2 = 6"}
	parsed := parser.parse()
	expected := "select * from table where column = 5 and column2 = 6"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicOr(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god column = 5 edge column2 = 6"}
	parsed := parser.parse()
	expected := "select * from table where column = 5 or column2 = 6"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicNot(t *testing.T) {
	parser := parser{input: "yoink * skibity table on god bruh column = 5"}
	parsed := parser.parse()
	expected := "select * from table where not column = 5"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicAsc(t *testing.T) {
	parser := parser{input: "yoink * skibity table short king"}
	parsed := parser.parse()
	expected := "select * from table asc"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicDesc(t *testing.T) {
	parser := parser{input: "yoink * skibity table tall king"}
	parsed := parser.parse()
	expected := "select * from table desc"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicDelete(t *testing.T) {
	parser := parser{input: "yeet skibity table on god column = 5"}
	parsed := parser.parse()
	expected := "delete from table where column = 5"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}

func TestBasicInsert(t *testing.T) {
	parser := parser{input: "slide into table (column1, column2) values (5, 6)"}
	parsed := parser.parse()
	expected := "insert into table (column1, column2) values (5, 6)"

	if parsed != expected {
		t.Errorf("Expected %s, got %s", expected, parsed)
	}
}
