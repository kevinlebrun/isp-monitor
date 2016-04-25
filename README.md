# isp-monitor

Some script I run on my Raspberry Pi to track my ISP (Internet Service
Provider) performance.

## Usage

On OS X:

```
$ go build
$ ./isp-monitor --with-headers > logfile
$ ./isp-monitor >> logfile
```

You can cross-compile on ARM (Rapsberry Pi):

```
$ GOARCH=arm GOOS=linux go build
```

I usually run this script every 10 minutes with `cron`.

## License

The MIT License
