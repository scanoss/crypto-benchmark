// Package check holds the problem report shared by the corpus, ground-truth
// and prediction parsers, and the strict YAML and JSON decoding they use.
package check

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type Severity int

const (
	Error Severity = iota
	Warning
)

func (s Severity) String() string {
	if s == Warning {
		return "warning"
	}
	return "error"
}

// Problem names one defect by file and by its path inside the YAML document.
type Problem struct {
	Severity Severity
	File     string
	Path     string
	Msg      string
}

func (p Problem) String() string {
	if p.Path == "" {
		return fmt.Sprintf("%s: %s: %s", p.Severity, p.File, p.Msg)
	}
	return fmt.Sprintf("%s: %s: %s: %s", p.Severity, p.File, p.Path, p.Msg)
}

type Report struct {
	Problems []Problem
}

func (r *Report) Errorf(file, path, format string, args ...any) {
	r.Problems = append(r.Problems, Problem{Error, file, path, fmt.Sprintf(format, args...)})
}

func (r *Report) Warnf(file, path, format string, args ...any) {
	r.Problems = append(r.Problems, Problem{Warning, file, path, fmt.Sprintf(format, args...)})
}

func (r *Report) HasErrors() bool {
	for _, p := range r.Problems {
		if p.Severity == Error {
			return true
		}
	}
	return false
}

// DecodeStrict decodes one YAML document, rejecting fields the target does not declare.
func DecodeStrict(data []byte, v any) error {
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(v); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("empty document")
		}
		return err
	}
	var extra yaml.Node
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		return errors.New("more than one YAML document")
	}
	return nil
}

// OneOf parses s into a closed enum.
func OneOf[T ~string](s string, allowed ...T) (T, error) {
	names := make([]string, len(allowed))
	for i, a := range allowed {
		if string(a) == s {
			return a, nil
		}
		names[i] = string(a)
	}
	return "", fmt.Errorf("%q is not one of %s", s, strings.Join(names, "|"))
}

// DecodeStrictJSON decodes one JSON document, rejecting fields the target does not declare.
func DecodeStrictJSON(data []byte, v any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("empty document")
		}
		return err
	}
	if _, err := dec.Token(); !errors.Is(err, io.EOF) {
		return errors.New("trailing data after the JSON document")
	}
	return nil
}
