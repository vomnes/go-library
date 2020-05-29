package tests

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/kylelemons/godebug/pretty"
	mgo "gopkg.in/mgo.v2"
)

var (
	// MongoDB corresponds to the test database
	MongoDB *mgo.Database
	// TimeTest allows to round about time for tests
	TimeTest = time.Now()
)

// InitTimeTest allows to round about time for tests
func InitTime() {
	cfg := pretty.CompareConfig
	cfg.Formatter[reflect.TypeOf(time.Time{})] = func(t time.Time) string {
		if t.Nanosecond() == 0 {
			return fmt.Sprint(t)
		}
		diff := t.Sub(TimeTest)
		if diff.Nanoseconds() < 0 {
			diff = -diff
		}
		if diff.Nanoseconds() < 50000 {
			return "Now rounded to 0.5 secondes"
		}
		return fmt.Sprintf("%d-%d-%d %d:%d:%d.%s\n", TimeTest.Year(), TimeTest.Month(), TimeTest.Day(),
			TimeTest.Hour(), TimeTest.Minute(), TimeTest.Second(), string(strconv.Itoa(TimeTest.Nanosecond())[0]))
	}
}
