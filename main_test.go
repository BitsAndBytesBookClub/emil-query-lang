package main

import (
	"testing"
)

func TestTokenParse(t *testing.T) {
	for key, val := range tokens {
		p := newParser(key)
		if p.parse() != val {
			t.Errorf("Expected %s, got %s", val, p.parse())
		}
	}
}

func TestUseSql(t *testing.T) {
	input := "yoink * skibity users on god name fr 'john doe'"
	expected := "select * from users where name = 'john doe'"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleSortASC(t *testing.T) {
	input := "yoink * skibity users on god name fr 'john doe' them ones name short king"
	expected := "select * from users where name = 'john doe' order by name asc"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleSortDESC(t *testing.T) {
	input := "yoink * skibity users on god name fr 'john doe' them ones name tall king"
	expected := "select * from users where name = 'john doe' order by name desc"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleAnd(t *testing.T) {
	input := "yoink * skibity users on god name fr 'john doe' goon age fr 25"
	expected := "select * from users where name = 'john doe' and age = 25"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleOr(t *testing.T) {
	input := "yoink * skibity users on god name fr 'john doe' edge age fr 25"
	expected := "select * from users where name = 'john doe' or age = 25"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleNot(t *testing.T) {
	input := "yoink * skibity users on god bruh name fr 'john doe'"
	expected := "select * from users where not name = 'john doe'"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleUpdate(t *testing.T) {
	input := "rizzler users w rizz name fr 'john doe'"
	expected := "update users set name = 'john doe'"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleDelete(t *testing.T) {
	input := "yeet skibity users on god name fr 'john doe'"
	expected := "delete from users where name = 'john doe'"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestSimpleInsert(t *testing.T) {
	input := "slide dms users (name, age) bands ('john doe', 25)"
	expected := "insert into users (name, age) values ('john doe', 25)"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}

func TestLeftJoin(t *testing.T) {
	input := "yoink * skibity users fanum tax orders ate users.id fr orders.user_id"
	expected := "select * from users left join orders on users.id = orders.user_id"
	if UseSql(input) != expected {
		t.Errorf("Expected %s, got %s", expected, UseSql(input))
	}
}
