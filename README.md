# etod

Turn epochs to dates:

```
$ ./etod -h
Usage of ./etod:
  -fmt string
    	time format to display dates in (default "2006-01-02T15:04:05")
  -now int
    	alternative 'current' epoch (default 1500472320)
  -span int
    	span in days around 'now' to consider a date (default 730)
$ ./etod -h |& etod -fmt '2006-01-02'
Usage of ./etod:
  -fmt string
    	time format to display dates in (default "2006-01-02T15:04:05")
  -now int
    	alternative 'current' epoch (default 2017-07-19)
  -span int
    	span in days around 'now' to consider a date (default 730)
```

It is quite simplistic and just matches on `\d{10}`, meaning it only works for dates between Sep 8 2001 and Nov 20
2286.
