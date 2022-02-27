package lexer

import (
	"regexp"
)

type Rules []*Rule

type Rule struct {
	Name        string     `json:"name"`
	Pattern     string     `json:"pattern"`
	DeleteToken bool       `json:"delete"`
	Replaces    []*Replace `json:"replaces"`
	reg         *regexp.Regexp
}

type Replace struct {
	Pattern string `json:"pattern"`
	Replace string `json:"replace"`
	reg     *regexp.Regexp
}

func (rules Rules) Init() error {
	var err error
	for i := range rules {
		rules[i].reg, err = regexp.Compile(rules[i].Pattern)
		if err != nil {
			return err
		}
		for j := range rules[i].Replaces {
			rules[i].Replaces[j].reg, err = regexp.Compile(rules[i].Replaces[j].Pattern)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Rule) Replace(b []byte) string {
	result := string(b)
	for _, repl := range r.Replaces {
		result = repl.reg.ReplaceAllString(result, repl.Replace)
	}
	return result
}

func (rules Rules) Find(stack []byte) (*Rule, []byte) {
	for _, re := range rules {
		result := re.reg.Find(stack)
		if result != nil {
			return re, result
		}
	}
	return nil, nil
}
