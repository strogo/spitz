// Copyright (C) 2012 Numerotron Inc.
// Use of this source code is governed by an MIT-style license
// that can be found in the LICENSE file.

package spitz

import (
	"fmt"
	"html/template"
	"strings"
)

// some functions to use inside templates

var funcMap = map[string]interface{}{"plural": pluralize, "mailto": mailto, "simpleformat": simpleFormat}

func pluralize(singular string, count int32) string {
	if count == 1 {
		return singular
	}

	return fmt.Sprintf("%ss", singular)
}

func mailto(email string) template.HTML {
	pre := `<script type="text/javascript">eval(decodeURIComponent('`
	post := `'))</script>`
	inner := fmt.Sprintf("document.write('<a href=\"mailto:%s\">%s</a>');", email, email)
	pieces := []string(nil)
	for _, v := range inner {
		pieces = append(pieces, fmt.Sprintf("%%%x", v))
	}
	return template.HTML(pre + strings.Join(pieces, "") + post)
}

func simpleFormat(in string) template.HTML {
	lines := strings.Split(in, "\n")
	paragraphs := []string(nil)
	for _, v := range lines {
		line := strings.TrimSpace(v)
		if len(line) == 0 {
			continue
		}
		line = template.HTMLEscapeString(line)
		paragraphs = append(paragraphs, fmt.Sprintf("<p>%s</p>", line))
	}
	return template.HTML(strings.Join(paragraphs, ""))
}
