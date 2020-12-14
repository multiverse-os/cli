package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

func LoadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("[cli/template] failed to load template file:", err.Error)
	}
	return string(content), nil
}

func StdOut(content string, data interface{}) error {
	uiTemplate := template.Must(template.New("ui").Parse(content))
	err := uiTemplate.Execute(os.Stdout, data)
	if err != nil {
		return fmt.Errorf("[cli/template] failed to render template:", err)
	}
	return nil
}

func OutputStream(content string, data interface{}) (output io.Writer, err error) {
	uiTemplate := template.Must(template.New("ui").Parse(content))
	err = uiTemplate.Execute(output, data)
	if err != nil {
		return nil, fmt.Errorf("[cli/template] failed to render template:", err)
	}
	return output, nil
}

func OutputString(content string, data interface{}) (string, error) {
	uiTemplate := template.Must(template.New("ui").Parse(content))
	buffer := new(bytes.Buffer)
	err := uiTemplate.Execute(buffer, data)
	if err != nil {
		return "", fmt.Errorf("[cli/template] failed to convert template stream to string:", err.Error)
	}
	return buffer.String(), nil
}
