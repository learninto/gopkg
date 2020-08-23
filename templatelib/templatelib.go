package templatelib

import (
	"bytes"
	"html/template"
)

func process(t *template.Template, vars interface{}) (tmplBytes bytes.Buffer, err error) {
	err = t.Execute(&tmplBytes, vars)
	if err != nil {
		return tmplBytes, err
	}
	return tmplBytes, nil
}

func ProcessString(str string, vars interface{}) (string, error) {
	tmpl, err := template.New("tmpl").Parse(str)
	if err != nil {
		return "", err
	}

	tmplBytes, err := process(tmpl, vars)
	if err != nil {
		return "", err
	}

	return tmplBytes.String(), nil
}
