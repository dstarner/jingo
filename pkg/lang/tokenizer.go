package lang

type Tokenizer interface {
	Tokenize(tokenDefinitions string, input string) ([]lexer.Token, error)
}

type DefaultTokenizer struct {

}

func (tokenizer DefaultTokenizer) Tokenize(tokenDefinitions string, input string) ([]lexer.Token, error) {
	d, err := regex.New(language)
	if err != nil {
		return nil, errors.New("could not generate regex tokenizer")
	}
	l, err := d.Lex(strings.NewReader(input))
}