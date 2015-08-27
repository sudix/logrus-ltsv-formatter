// LTSV formatter for logrus
package logrusltsv

import (
	"bytes"
	"fmt"
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
}

// New create LogrusLTSVFormatter.
func NewDefaultFormatter() *LogrusLTSVFormatter {
	c := LogrusLTSVConfig{
		TimestampFormat: logrus.DefaultTimestampFormat,
		FieldPrefix:     "field.",
	}
	return NewFormatter(c)
}

// NewWithTimestampFormat create LogrusLTSVFormatter
// with timestamp format.
func NewFormatter(config LogrusLTSVConfig) *LogrusLTSVFormatter {
	return &LogrusLTSVFormatter{
		conf: config,
	}
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

	for _, k := range keys {
		switch v := entry.Data[k].(type) {
		case time.Time:
			fmt.Fprintf(buf, "%s%s:%s\t", f.conf.FieldPrefix, k, v.Format(timestampFormat))
		default:
			fmt.Fprintf(buf, "%s%s:%v\t", f.conf.FieldPrefix, k, v)
		}
	}

	fmt.Fprintf(buf, "msg:%s\n", entry.Message)

	return buf.Bytes(), nil
}
