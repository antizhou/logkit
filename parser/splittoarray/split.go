package parser

import (
	"github.com/qiniu/logkit/conf"
	"github.com/qiniu/logkit/parser"
	. "github.com/qiniu/logkit/parser/config"
	. "github.com/qiniu/logkit/utils/models"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	parser.RegisterConstructor(TypeSplit, NewParser)
}

type Parser struct {
	name               string
	lineSeparator      string
	fieldOrderMappings map[string]string
	characterReplace   map[string]string
}

func NewParser(c conf.MapConf) (parser.Parser, error) {
	name, _ := c.GetStringOr(KeyParserName, "")
	lineSeparator, _ := c.GetStringOr("line_separator", "")
	fieldOrderMappings, _ := c.GetAliasMap("field_order_mappings")

	return &Parser{
		name:               name,
		lineSeparator:      lineSeparator,
		fieldOrderMappings: fieldOrderMappings,
	}, nil
}

func (p *Parser) Name() string {
	return p.name
}

func (p *Parser) Parse(lines []string) ([]Data, error) {
	datas := make([]Data, len(lines))

	mappings := p.fieldOrderMappings
	for i := 0; i < len(lines); i++ {
		datas[i] = Data{}

		line := lines[i]

		reg, err := regexp.Compile(p.lineSeparator)
		if err != nil {
			continue
		}
		fields := reg.Split(line, -1)
		for field, order := range mappings {
			o, err := strconv.Atoi(order)
			if err != nil {
				continue
			}
			datas[i][field] = p.replace(fields[o])
		}
	}
	return datas, nil
}
func (p *Parser) replace(field string) string {
	for pre, next := range p.characterReplace {
		field = strings.ReplaceAll(field, pre, next)
	}
	return field
}
