package sqlite_test

import (
	"testing"
	"time"

	"github.com/demaggus83/go-zeug/sqlite"
	"github.com/demaggus83/go-zeug/test"
)

func TestSQLTimeFromTime(t *testing.T) {
	t.Parallel()

	st := sqlite.SQLTimeFromTime(test.Now)
	test.Equals(t, st.Time, test.Now.UTC(), `Expected st to be "%s" but got "%s"`)
	test.True(t, st.Valid)
}

var fixedTimeMilli = test.FixedTime.UnixNano() / 1000000

var sqliteScanTestTable = []struct {
	Input         interface{}
	ExpectedError error
	ExpectedTime  time.Time
	Valid         bool
	Equals        bool
}{
	{fixedTimeMilli, nil, test.FixedTime, true, true},
	{int(fixedTimeMilli), nil, test.FixedTime, true, true},
	{test.FixedTime, nil, test.FixedTime, true, true},
	{time.Now(), nil, test.FixedTime, true, false},
	{"2019-03-03T13:51:06.7369193+01:00", sqlite.ErrSQLTimeUnmarshallType, time.Time{}, false, false},
}

func TestScan(t *testing.T) {
	for counter, tt := range sqliteScanTestTable {
		sqltime := &sqlite.SQLTime{}

		err := sqltime.Scan(tt.Input)
		if tt.ExpectedError != nil {
			test.IsError(t, tt.ExpectedError, err)
			return
		}

		if tt.ExpectedError == nil {
			test.NoError(t, err)
		}

		equal := sqltime.Time.UnixNano()/1000000 == tt.ExpectedTime.UnixNano()/1000000
		test.Equals(t, tt.Equals, equal, "Expected time comparison to be '%s' but it is '%s' (counter: %d)", counter)
		test.Equals(t, sqltime.Valid, tt.Valid, "Expected valid == '%s' but it is '%s' (counter: %d)", counter)
	}
}
