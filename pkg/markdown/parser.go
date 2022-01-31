package markdown

import (
	"bufio"
	"fmt"
	"io"
)

type Parser struct {
	source io.Reader
}

func NewParser(source io.Reader) *Parser {
	return &Parser{
		source: source,
	}
}

func (p Parser) Parse() {

	scanner := bufio.NewScanner(p.source)
	buf := make([]byte, 1024)
	scanner.Buffer(buf, 1024)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Printf("============ %d\n", len(scanner.Text()))
	}
}

type Token interface {
	Key() string
	Value() string
	Child() Token
}

type headline struct {
	key, value string
}

func (h headline) Key() string   { return h.key }
func (h headline) Value() string { return h.value }
func (h headline) Child() Token  { return nil }

func tokenize(line string) Token {

}
