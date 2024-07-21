Go Heap demo
----

# Sample 1 - Methed Return Pointer

```go
go run -gcflags "-m -m" sample1/main.go

# command-line-arguments
./escape_analysis.go:9:6: cannot inline createSuperMan: marked go:noinline
./escape_analysis.go:19:6: cannot inline createTheFlash: marked go:noinline
./escape_analysis.go:28:6: cannot inline main: function too complex: cost 133 exceeds budget 80
./escape_analysis.go:20:2: h escapes to heap:
./escape_analysis.go:20:2:   flow: ~r0 = &h:
./escape_analysis.go:20:2:     from &h (address-of) at ./escape_analysis.go:25:9
./escape_analysis.go:20:2:     from return &h (return) at ./escape_analysis.go:25:2
./escape_analysis.go:20:2: moved to heap: h
Superman  0xc0000566d8
The Flash  0xc00009e000
Superman  0xc000056720 The Flash 0xc000056718
```

We can see varaible `h` is moved to heap
```
./escape_analysis.go:20:2: h escapes to heap:
./escape_analysis.go:20:2:   flow: ~r0 = &h:
./escape_analysis.go:20:2:     from &h (address-of) at ./escape_analysis.go:25:9
./escape_analysis.go:20:2:     from return &h (return) at ./escape_analysis.go:25:2
./escape_analysis.go:20:2: moved to heap: h
```


# Sample 2 - Unknown Variable Size


```go
go run -gcflags "-m -m" sample2/main.go

# command-line-arguments
sample2/main.go:13:6: cannot inline createSuperMan: marked go:noinline
sample2/main.go:23:6: cannot inline createTheFlash: marked go:noinline
sample2/main.go:33:6: can inline createSomeHeros with cost 57 as: func(*int) []hero { heros := make([]hero, *n, *n); for loop; println(&heros); return heros }
sample2/main.go:54:6: cannot inline main: function too complex: cost 339 exceeds budget 80
sample2/main.go:63:15: inlining call to flag.Int
sample2/main.go:64:12: inlining call to flag.Parse
sample2/main.go:65:23: inlining call to createSomeHeros
sample2/main.go:24:2: h escapes to heap:
sample2/main.go:24:2:   flow: ~r0 = &h:
sample2/main.go:24:2:     from &h (address-of) at sample2/main.go:29:9
sample2/main.go:24:2:     from return &h (return) at sample2/main.go:29:2
sample2/main.go:24:2: moved to heap: h
sample2/main.go:34:15: make([]hero, *n, *n) escapes to heap:
sample2/main.go:34:15:   flow: {heap} = &{storage for make([]hero, *n, *n)}:
sample2/main.go:34:15:     from make([]hero, *n, *n) (non-constant size) at sample2/main.go:34:15
sample2/main.go:33:22: n does not escape
sample2/main.go:34:15: make([]hero, *n, *n) escapes to heap
sample2/main.go:65:23: make([]hero, *n, *n) escapes to heap:
sample2/main.go:65:23:   flow: {heap} = &{storage for make([]hero, *n, *n)}:
sample2/main.go:65:23:     from make([]hero, *n, *n) (non-constant size) at sample2/main.go:65:23
sample2/main.go:65:23: make([]hero, *n, *n) escapes to heap
Superman  0xc0000a4e40
The Flash  0xc0000b6020
Superman  0xc0000a4f20 The Flash 0xc0000a4f18
0xc0000a4f00
Heros  9 0xc0000a4ee0
```

escape here
```
sample2/main.go:34:15: make([]hero, *n, *n) escapes to heap
sample2/main.go:65:23: make([]hero, *n, *n) escapes to heap:
sample2/main.go:65:23:   flow: {heap} = &{storage for make([]hero, *n, *n)}:
sample2/main.go:65:23:     from make([]hero, *n, *n) (non-constant size) at sample2/main.go:65:23
```