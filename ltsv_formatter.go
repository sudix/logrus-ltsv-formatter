// LTSV formatter for logrus
package logrusltsv

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/Sirupsen/logrus"
)

type LogrusLTSVFormatter struct {
	conf LogrusLTSVConfig
}

type LogrusLTSVConfig struct {
	TimestampFormat string
	FieldPrefix     string
	Filters         []Filter
}

type Filter func(string) string

var reNL = regexp.MustCompile(`\n`)

func EscapeNewLine(s string) string {
	return reNL.ReplaceAllString(s, " ")
}

// NewDefaultFormatter create LogrusLTSVFormatter with default configuration.
func NewDefaultFormatter() *LogrusLTSVFormatter {
	c := LogrusLTSVConfig{
		TimestampFormat: logrus.DefaultTimestampFormat,
		FieldPrefix:     "field.",
		Filters:         []Filter{EscapeNewLine},
	}
	return NewFormatter(c)
}

// NewFormatter create LogrusLTSVFormatter
// with given configuration.
func NewFormatter(config LogrusLTSVConfig) *LogrusLTSVFormatter {
	return &LogrusLTSVFormatter{
		conf: config,
	}
}

func (f *LogrusLTSVFormatter) filtering(val string) string {
	s := val
	for _, filter := range f.conf.Filters {
		s = filter(s)
	}
	return s
}

func (f *LogrusLTSVFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var keys []string
	for k, _ := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := &bytes.Buffer{}

	timestampFormat := f.conf.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = logrus.DefaultTimestampFormat
	}

	fmt.Fprintf(
		buf,
		"time:%s\tlevel:%s\t",
		entry.Time.Format(timestampFormat),
		entry.Level.String(),
	)

	var val string
	for _, k := range keys {
		switch v := entry.Data[k].(type) {
		case string:
			val = v
		case time.Time:
			val = v.Format(timestampFormat)
		default:
			val = fmt.Sprintf("%v", v)
		}

		val = f.filtering(val)

		fmt.Fprintf(buf, "%s%s:%s\t", f.conf.FieldPrefix, k, val)
	}

	fmt.Fprintf(buf, "msg:%s\n", f.filtering(entry.Message))

	return buf.Bytes(), nil
}
