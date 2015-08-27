package logrusltsv_test

import (
	"bytes"
	"regexp"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/sudix/logrus-ltsv-formatter"
)

func TestNewDefaultFormatter(t *testing.T) {
	out := &bytes.Buffer{}

	logrus.SetFormatter(logrusltsv.NewDefaultFormatter())
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.DebugLevel)

	now := time.Now()

	logrus.WithFields(logrus.Fields{
		"stringKey":  "foo",
		"booleanKey": true,
		"numberKey":  122,
		"msg":        "msg 1",
		"timeKey":    now,
	}).Debug("test message 1")

	expectedPattern := "time:[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\\+[0-9]{2}:[0-9]{2}\tlevel:debug\tfield.booleanKey:true\tfield.msg:msg 1\tfield.numberKey:122\tfield.stringKey:foo\tfield.timeKey:[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\\+[0-9]{2}:[0-9]{2}\tmsg:test message 1\n"
	actual := out.String()
	expected := regexp.MustCompile(expectedPattern)
	if !expected.MatchString(actual) {
		t.Errorf("\nexpectedPattern:%s\n but got:%s\n", expectedPattern, actual)
	}
}

func TestNewFormatter(t *testing.T) {
	out := &bytes.Buffer{}

	timestampFormat := "2006/01/02 15:04:05 JST"

	c := logrusltsv.LogrusLTSVConfig{
		TimestampFormat: timestampFormat,
		FieldPrefix:     "prefix_",
	}

	logrus.SetFormatter(logrusltsv.NewFormatter(c))
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.DebugLevel)

	now := time.Now()

	logrus.WithFields(logrus.Fields{
		"stringKey":  "foo",
		"booleanKey": true,
		"numberKey":  122,
		"msg":        "msg 1",
		"timeKey":    now,
	}).Debug("test message 1")

	expectedPattern := "time:[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} JST\tlevel:debug\tprefix_booleanKey:true\tprefix_msg:msg 1\tprefix_numberKey:122\tprefix_stringKey:foo\tprefix_timeKey:" + now.Format(timestampFormat) + "\tmsg:test message 1\n"
	actual := out.String()
	expected := regexp.MustCompile(expectedPattern)
	if !expected.MatchString(actual) {
		t.Errorf("\nexpectedPattern:%s\n but got:%s\n", expectedPattern, actual)
	}
}

type dummyWriter struct{}

func (w *dummyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func BenchmarkFormat(b *testing.B) {
	out := &dummyWriter{}

	logrus.SetFormatter(logrusltsv.NewDefaultFormatter())
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.DebugLevel)

	now := time.Now()

	for i := 0; i < b.N; i++ {
		logrus.WithFields(logrus.Fields{
			"stringKey":  "foo",
			"booleanKey": true,
			"numberKey":  122,
			"msg":        "msg 1",
			"timeKey":    now,
		}).Debug("test message 1")
	}
}
