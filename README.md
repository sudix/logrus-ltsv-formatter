# LogrusLTSVFormatter

Description
====================

[LTSV](http://ltsv.org/) formatter for [Sirupsen/logrus](https://github.com/Sirupsen/logrus).

Install
====================

```
$ go get github.com/sudix/logrus-ltsv-formatter
```

Usage
====================

```golang
import (
	"bytes"
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/sudix/logrus-ltsv-formatter"
)

func ExampleNewDefaultFormatter() {
	formatter := logrusltsv.NewDefaultFormatter()
	out := &bytes.Buffer{}

	logrus.SetFormatter(formatter)
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.DebugLevel)

	now := time.Now().Format(time.RFC850)

	logrus.WithFields(logrus.Fields{
		"stringKey":  "foo",
		"booleanKey": true,
		"numberKey":  122,
		"msg":        "msg 1",
		"timeKey":    now,
	}).Debug("test message 1")

	fmt.Println(out)
}
```

You can specify configuration.

```golang
func ExampleNewFormatter() {
	formatter := logrusltsv.NewFormatter(
		logrusltsv.LogrusLTSVConfig{
			TimestampFormat: "2006/01/02 15:04:05 JST",
			FieldPrefix:     "prefix_",
			Filters:         []Filter{EscapeNewLine},
		},
	)

	out := &bytes.Buffer{}

	logrus.SetFormatter(formatter)
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.DebugLevel)

	now := time.Now().Format(time.RFC850)

	logrus.WithFields(logrus.Fields{
		"stringKey":  "foo",
		"booleanKey": true,
		"numberKey":  122,
		"msg":        "msg 1",
		"timeKey":    now,
	}).Debug("test message 1")

	fmt.Println(out)
}
```

Configuration
====================

`LogrusLTSVConfig` has three fields.

* `TimestampFormat` - Timestamp Format to apply `time.Time` type.
* `FieldPrefix` - The prefix string that attached before field key name.
* `Filters` - Filter functions that applies field values. Filter functions have to satisfy `logrusltsv.Filter` type. They are applied in declaration order. (Caution! Heavy process filter may cause performance issue.)
