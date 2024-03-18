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
				if strings.Contains(currToken, "short") ||
					strings.Contains(currToken, "tall") ||
					strings.Contains(currToken, "on") ||
					strings.Contains(currToken, "them") ||
					strings.Contains(currToken, "w") ||
					strings.Contains(currToken, "fanum") {
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
