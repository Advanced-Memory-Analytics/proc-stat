# proc-stat
Lightweight Go utility that utilizes `ps -ef` command to get process status information.

### Structure of data

```go
type Proc struct {
	Uid   string
	Pid   int
	Ppid  int
	C     int
	Stime string
	Tty   string
	Time  string
	Cmd   string
	Args  map[string]string
}
```
Each field within the `Proc` struct corresponds to a field returned by the `ps -ef` command except 
the `Args` field. This field is parsed from the `CMD` field of `ps -ef` and corresponds to the arguments
supplied to that command.  This field is a map that can be accessed by `Key`.  For example:

```
501  1891  1886   0 10:48AM ttys002    0:00.00 /tester.sh -name test2 -runAsUser
```

The `Args` map associated with this process has the key "name" associated with the value "test2".
For boolean flags, if the flag is present, the key for that flag is associated with the value "true".
The key "runAsUser" in the example above, the value will be "true" since it is present.  Using Go maps, 
you can also easily determine if an argument was supplied by simply checking the second return type of 
the map, which is a boolean value.  See here: https://go.dev/blog/maps.