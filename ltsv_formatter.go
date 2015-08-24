// LTSV formatter for logrus
package logrusltsv

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/Sirupsen/logrus"
)

type LogrusLTSVFormatter struct {
	TimestampFormat string
}

// New create LogrusLTSVFormatter.
func New() *LogrusLTSVFormatter {
	return new(LogrusLTSVFormatter)
}

// NewWithTimestampFormat create LogrusLTSVFormatter
// with timestamp format.
func NewWithTimestampFormat(timestampFormat string) *LogrusLTSVFormatter {
	return &LogrusLTSVFormatter{
		TimestampFormat: timestampFormat,
	}
}

func (f *LogrusLTSVFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var keys []string
	for k, _ := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := &bytes.Buffer{}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = logrus.DefaultTimestampFormat
	}

	buf.WriteString(fmt.Sprintf(
		"time:%s\tlevel:%s\t",
		entry.Time.Format(timestampFormat),
		entry.Level.String(),
	))

	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("field.%s:%v\t", k, entry.Data[string(k)]))
	}

	buf.WriteString(fmt.Sprintf("msg:%s\n", entry.Message))

	return buf.Bytes(), nil
}

func (f *LogrusLTSVFormatter) Format2(entry *logrus.Entry) ([]byte, error) {
	var buf []string
	var keys []string

	for k, _ := range entry.Data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = logrus.DefaultTimestampFormat
	}

	buf = append(buf, fmt.Sprintf("time:%s", entry.Time.Format(timestampFormat)))
	buf = append(buf, fmt.Sprintf("level:%s", entry.Level.String()))

	for _, k := range keys {
		buf = append(buf, fmt.Sprintf("field.%s:%v", k, entry.Data[string(k)]))
	}

	buf = append(buf, fmt.Sprintf("msg:%s\n", entry.Message))

	return []byte(strings.Join(buf, "\t")), nil
}
