package strutil

import (
	"bytes"
	"html/template"
)

// StringFromTemplate renders a given template with variables to a string
func StringFromTemplate(templateText string, vars interface{}) (string, error) {
	// parse template
	tmpl, err := template.New("").Parse(templateText)
	if err != nil {
		return "", err
	}
	// render template
	var templateOutput bytes.Buffer
	err = tmpl.Execute(&templateOutput, vars)
	if err != nil {
		return "", err
	}
	return templateOutput.String(), nil
}
