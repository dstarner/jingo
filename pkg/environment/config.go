package environment

// Configuration for an Environment instance
type Configuration struct {
	blockStartString string
	blockEndString   string

	variableStartString string
	variableEndString   string
}

// NewConfiguration returns a new, default configuration for the environment
func NewConfiguration() *Configuration {
	return &Configuration{
		blockStartString: "{%",
		blockEndString:   "%}",

		variableStartString: "{{",
		variableEndString:   "}}",
	}
}

// Option allows easily setting Configuration instances
type Option func(*Configuration)

// ---------------
// List of Options
// ---------------

// BlockStartString is the string marking the beginning of a block
func BlockStartString(blockStartString string) Option {
	return func(c *Configuration) {
		c.blockStartString = blockStartString
	}
}

// BlockEndString is the string marking the ending of a block
func BlockEndString(blockEndString string) Option {
	return func(c *Configuration) {
		c.blockEndString = blockEndString
	}
}

// VariableStartString is the string marking the beginning of a block
func VariableStartString(variableStartString string) Option {
	return func(c *Configuration) {
		c.variableStartString = variableStartString
	}
}

// VariableEndString is the string marking the ending of a block
func VariableEndString(variableEndString string) Option {
	return func(c *Configuration) {
		c.variableEndString = variableEndString
	}
}
