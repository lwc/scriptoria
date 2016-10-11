package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

const sampleMD = `
	code block on the first line

_Regular_ text.

    code block indented by spaces

Regular *text.*

	the lines in this block
	all contain trailing spaces

Regular Text.
-------------

	code block on the last line
`

const expectedHTML = `
<pre><code>code block on the first line
</code></pre>

<p><em>Regular</em> text.</p>

<pre><code>code block indented by spaces
</code></pre>

<p>Regular <em>text.</em></p>

<pre><code>the lines in this block
all contain trailing spaces
</code></pre>

<h2>Regular Text.</h2>

<pre><code>code block on the last line
</code></pre>
`

func TestScriptoria(t *testing.T) {

	app := scriptoria{
		secret: "Wow this is my test secret",
	}
	r := httptest.NewRequest("POST", "/", strings.NewReader(sampleMD))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, r)

	if strings.TrimSpace(w.Body.String()) != strings.TrimSpace(expectedHTML) {
		t.Error("The response body did not match expectation")
	}
}

func BenchmarkScriptoria(b *testing.B) {

	app := scriptoria{
		secret: "Wow this is my test secret",
	}
	r := httptest.NewRequest("POST", "/", strings.NewReader(sampleMD))
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		app.ServeHTTP(w, r)
	}
}
