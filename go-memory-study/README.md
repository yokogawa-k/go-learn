[Goならわかる Linuxのメモリ管理](https://speakerdeck.com/juntaki/gonarawakaru-linuxfalsememoriguan-li) の勉強

### slide03

4 の stack なのかな？

```console
[yokogawa-k@arch:~/.ghq/github.com/yokogawa-k/go-learn/go-memory-study]$ go build -o slide03 ./slide03.go
[yokogawa-k@arch:~/.ghq/github.com/yokogawa-k/go-learn/go-memory-study]$ gdb slide03
GNU gdb (GDB) 8.1.1
Copyright (C) 2018 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.  Type "show copying"
and "show warranty" for details.
This GDB was configured as "x86_64-pc-linux-gnu".
Type "show configuration" for configuration details.
For bug reporting instructions, please see:
<http://www.gnu.org/software/gdb/bugs/>.
Find the GDB manual and other documentation resources online at:
<http://www.gnu.org/software/gdb/documentation/>.
For help, type "help".
Type "apropos word" to search for commands related to "word"...
Reading symbols from slide03...done.
warning: File "/usr/lib/go/src/runtime/runtime-gdb.py" auto-loading has been declined by your `auto-load safe-path' set to "$debugdir:$datadir/auto-load".
To enable execution of this file add
        add-auto-load-safe-path /usr/lib/go/src/runtime/runtime-gdb.py
line to your configuration file "/home/yokogawa-k/.gdbinit".
To completely disable this security protection add
        set auto-load safe-path /
line to your configuration file "/home/yokogawa-k/.gdbinit".
For more information about this security protection see the
"Auto-loading safe path" section in the GDB manual.  E.g., run from the shell:
        info "(gdb)Auto-loading safe path"
(gdb) b main.main
Breakpoint 1 at 0x482070: file /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03.go, line 5.
(gdb) r
Starting program: /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03
[New LWP 18234]
[New LWP 18235]
[New LWP 18236]
[New LWP 18237]

Thread 1 "slide03" hit Breakpoint 1, main.main () at /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03.go:5
warning: Source file is more recent than executable.
5       func main() {
(gdb) i proc mappings
process 18230
Mapped address spaces:

          Start Addr           End Addr       Size     Offset objfile
            0x400000           0x483000    0x83000        0x0 /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03
            0x483000           0x514000    0x91000    0x83000 /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03
            0x514000           0x528000    0x14000   0x114000 /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03
            0x528000           0x547000    0x1f000        0x0 [heap]
        0xc000000000       0xc000001000     0x1000        0x0
        0xc41fff8000       0xc420100000   0x108000        0x0
      0x7ffff7f5a000     0x7ffff7ffa000    0xa0000        0x0
      0x7ffff7ffa000     0x7ffff7ffd000     0x3000        0x0 [vvar]
      0x7ffff7ffd000     0x7ffff7fff000     0x2000        0x0 [vdso]
      0x7ffffffde000     0x7ffffffff000    0x21000        0x0 [stack]
(gdb) n

Thread 1 "slide03" hit Breakpoint 1, main.main () at /home/yokogawa-k/.ghq/github.com/yokogawa-k/go-learn/go-memory-study/slide03.go:5
5       func main() {
(gdb)
6               a := "hoge"
(gdb) c
Continuing.
0xc42007c1c0
Couldn't get registers: No such process.
Couldn't get registers: No such process.
(gdb) [LWP 18236 exited]
[LWP 18234 exited]
[LWP 18230 exited]
[LWP 18237 exited]
[Inferior 1 (process 18230) exited normally]

The program is not being run.
```

答えは 5 なのか...

### slide06

```console
[yokogawa-k@arch:~/.ghq/github.com/yokogawa-k/go-learn/go-memory-study]$ go build -gcflags '-m -N -l' slide06.go
# command-line-arguments
./slide06.go:7:13: a escapes to heap
./slide06.go:6:10: new(int) escapes to heap
./slide06.go:7:13: main ... argument does not escape
./slide06.go:9:10: main new(int) does not escape
```

