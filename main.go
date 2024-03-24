package main

import (
	"strings"
)

var tokens = map[string]string{
	"yoink":      "select",
	"fr":         "=",
	"skibity":    "from",
	"on god":     "where",
	"goon":       "and",
	"edge":       "or",
	"delulu":     "not",
	"rizzler":    "update",
	"w rizz":     "set",
	"short king": "asc",
	"tall king":  "desc",
	"yeet":       "delete",
	"slide":      "insert",
	"them ones":  "order by",
	"bands":      "values",
	"dms":        "into",
	"fanum tax":  "left join",
	"ate":        "on",
	"cap":        "as",
}

type parser struct {
	curr   byte
	pos    int
	input  string
	output string
}

func UseSql(input string) string {
	p := newParser(input)
	return p.parse()
}

func newParser(input string) *parser {
	return &parser{
		input: strings.ToLower(input),
	}
}

func (p *parser) next() {
	p.pos += 1
	if p.pos >= len(p.input) {
		p.curr = 0
		return
	}
	p.curr = p.input[p.pos]
}

func (p *parser) parse() string {
	var currToken string
	p.pos = -1
	p.curr = ' '
	for p.pos <= len(p.input) && p.curr != 0 {
		p.next()
		currToken += string(p.curr)

		if val, ok := tokens[currToken]; ok {
			p.output += val
			currToken = ""
		} else {
			if p.curr == ' ' {
				if isComboToken(tokens, currToken) {
					continue
				} else if currToken != "" {
					p.output += currToken
					currToken = ""
				} else {
					currToken = ""
					p.output += " "
				}
			}
		}

	}

	if currToken != "" {
		p.output += currToken[:len(currToken)-1]
	}

	return p.output
}

func isComboToken(tokens map[string]string, currToken string) bool {
	for k := range tokens {
		if strings.Contains(k, " ") {
			splitAtSpace := strings.Split(k, " ")
			if strings.Contains(currToken, splitAtSpace[0]) {
				return true
			}
		}
	}

	return false
}
