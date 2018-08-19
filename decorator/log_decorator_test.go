package decorator_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"gopkg.in/jwowillo/cache.v2/decorator"
)

// TestLogDecoratorDeleteLogs tests that LogDecorator's Delete logs.
func TestLogDecoratorDeleteLogs(t *testing.T) {
	b := &bytes.Buffer{}
	c := decorator.NewLogDecorator(&MockCache{}, b, "test")
	c.Delete("k")
	bs, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
	}
	out := string(bs)
	prefix := "cache test: "
	suffix := "delete k\n"
	if !strings.HasPrefix(out, prefix) {
		t.Errorf("strings.HasPrefix(%s, %s) = false, want true",
			out, prefix)
	}
	if !strings.HasSuffix(out, suffix) {
		t.Errorf("strings.HasSuffix(%s, %s) = false, want true",
			out, suffix)
	}
}

// TestLogDecoratorClearLogs test that LogDecorator's Clear logs.
func TestLogDecoratorClearLogs(t *testing.T) {
	b := &bytes.Buffer{}
	c := decorator.NewLogDecorator(&MockCache{}, b, "test")
	c.Clear()
	bs, err := ioutil.ReadAll(b)
	if err != nil {
		t.Error(err)
	}
	out := string(bs)
	prefix := "cache test: "
	suffix := "clear\n"
	if !strings.HasPrefix(out, prefix) {
		t.Errorf("strings.HasPrefix(%s, %s) = false, want true",
			out, prefix)
	}
	if !strings.HasSuffix(out, suffix) {
		t.Errorf("strings.HasSuffix(%s, %s) = false, want true",
			out, suffix)
	}
}

// TestLogDecorator tests that LogDecorator decorates properly.
func TestLogDecorator(t *testing.T) {
	f := decorator.NewLogDecoratorFactory(&bytes.Buffer{}, "test")
	DecoratorTest(t, f)
}
