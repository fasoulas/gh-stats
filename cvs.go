package main

import (
	"strings"
)

func getCVSHeader(items ...string) string {
	var sb strings.Builder
	for i := range items {
		sb.WriteString(items[i])
		sb.WriteString(",")
	}
	header := sb.String()
	return header[:len(header)-1]
}

func getCVSLine(items []string) string {
	var sb strings.Builder
	for i := range items {
		sb.WriteString(items[i])
		sb.WriteString(",")
	}
	header := sb.String()
	return header[:len(header)-1]
}
