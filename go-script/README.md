[Using Go as a scripting language in Linux](https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/)

```console
$ go get github.com/erning/gorun
$ sudo cp $(which gorun) /usr/local/bin
$ echo ':golang:E::go::/usr/local/bin/gorun:OC' | sudo tee /proc/sys/fs/binfmt_misc/register
```
