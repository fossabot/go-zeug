# go-zeug

[![GoDoc](https://godoc.org/github.com/demaggus83/go-zeug?status.svg)](https://godoc.org/github.com/demaggus83/go-zeug) [![Go Report Card](https://goreportcard.com/badge/github.com/demaggus83/go-zeug)](https://goreportcard.com/report/github.com/demaggus83/go-zeug)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdemaggus83%2Fgo-zeug.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdemaggus83%2Fgo-zeug?ref=badge_shield)

It's a small stdlib only collection of tools, libs and helpers for go >=1.13.   

"Zeug" is the german slang word for "stuff"

## Packages

### env

Package env provides a type for the current runtime environment (not env vars!).

The idea is to initialize the Env very early in your code to make decisions based on the current runtime env. 
Because Env is implementing the Unmarshaler interface it's compatible with 
[envconfig](https://github.com/vrischmann/envconfig) and that's the idea how to use it.

### slice

Package slice provides some helper methods for working with slices..

At the moment there are just 
```
HasType(t Type) bool 
```
methods. When generics arriving this package is going to shrink ;)

### sql

Package sql have some helpers and types for sql.  
Currently there is just the "Scanner" interface defined.
```
type Scanner interface {
	Scan(args ...interface{}) error
}
```

### sqlite

Package sqlite have some helpers and types for working with sqlite.
  
It provides a SQLTime with nanoseconds support which can be marshalled and unmarshalled in sqlite.

#### sqlite trigger example

A typical sqlite.SQLTime trigger looks like:

```sql
CREATE TRIGGER t_mytable_insert_updated
    AFTER INSERT ON "mytable"
    FOR EACH ROW
    WHEN NEW.updated IS NULL
BEGIN
    UPDATE "mytable" SET updated = (CAST(ROUND((julianday("now") - 2440587.5)*86400000) As INTEGER)) WHERE id = NEW.id;
END;

CREATE TRIGGER t_mytable_update_updated
    AFTER UPDATE ON "mytable"
    FOR EACH ROW
BEGIN
    UPDATE "mytable" SET updated = (CAST(ROUND((julianday("now") - 2440587.5)*86400000) As INTEGER)) WHERE id = OLD.id;
END;
```

### test

Package test provides some helpers and vars usefully for testing.


#### Golden file handling

See this [article](https://medium.com/soon-london/testing-with-golden-files-in-go-7fccc71c43d3) for more info about
golden files.

```
func TestSomething(t *testing.T) {
    // ...
    b := test.HandleGoldenFile(t.Name(), t, dataToTest);
    // ...
}
```

To update your golden files, run your tests with 

```
go test ./... -v -update
```

#### Variables

```
var (
	// Now represents the current time.
	Now = time.Now()
	// FixedTime represents a fixed time "2019-03-03T13:51:06.7369193+00:00".
	FixedTime time.Time
	// Error is a default error for mocks.
	Error = errors.New("unittest")
)
```

#### Methods

See godoc ;)

## Changelog

### 0.1.1
+ fixes some refactoring problems

### 0.1.0
+ initial release


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdemaggus83%2Fgo-zeug.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdemaggus83%2Fgo-zeug?ref=badge_large)