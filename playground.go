//go:build ignore
// +build ignore

// Package declaration
package main

// Import necessary packages
import (
	"fmt"
	"strings"
)

// StringBuilder defines an interface for string building operations
type StringBuilder interface {
	Append(string) StringBuilder
	Prepend(string) StringBuilder
	ToString() string
}

// stringBuilder is the concrete implementation of StringBuilder
type stringBuilder struct {
	builder strings.Builder
}

// NewStringBuilder creates a new StringBuilder
func NewStringBuilder() StringBuilder {
	return &stringBuilder{}
}

// Append adds a string to the end of the builder
func (sb *stringBuilder) Append(s string) StringBuilder {
	sb.builder.WriteString(s)
	return sb
}

// Prepend adds a string to the beginning of the builder
func (sb *stringBuilder) Prepend(s string) StringBuilder {
	temp := sb.builder.String()
	sb.builder.Reset()
	sb.builder.WriteString(s)
	sb.builder.WriteString(temp)
	return sb
}

// ToString returns the built string
func (sb *stringBuilder) ToString() string {
	return sb.builder.String()
}

func main() {
	// Create a new StringBuilder and chain methods
	result := NewStringBuilder().
		Append("Hello").
		Append(" ").
		Append("World").
		Prepend("Prepend1 ").
		Append(" Append1 ").
		Prepend("Prepend2 ").
		Prepend("Greeting: ").
		ToString()

	fmt.Println(result)
}
