// LTSV formatter for logrus
package logrusltsv

import (
	"bytes"
	"fmt"
	"sort"

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

	fmt.Fprintf(
		buf,
		"time:%s\tlevel:%s\t",
		entry.Time.Format(timestampFormat),
		entry.Level.String(),
	)

	for _, k := range keys {
		fmt.Fprintf(buf, "field.%s:%v\t", k, entry.Data[k])
	}

	fmt.Fprintf(buf, "msg:%s\n", entry.Message)

	return buf.Bytes(), nil
}
