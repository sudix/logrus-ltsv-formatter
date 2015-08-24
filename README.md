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

func ExampleNew() {
	formatter := logrusltsv.New()
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

You can specify timestamp format.

```golang
func ExampleNewWithTimestampFormat() {
	formatter := logrusltsv.NewWithTimestampFormat("2006/01/02 15:04:05 JST")
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
