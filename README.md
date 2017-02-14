# apache-loggen-go
Apache-loggen-go is a Golang script to generate dummy apache-formatted log data.
It creates long-term log data very easily.

## Usage

```bash
$ go get github.com/acroquest/apache-loggen-go
$ go run example/main.go -d 3 -e 0.3 -f apache-3days.log
```

- by setting `-d` option, you can generate log data from 3 days before the present to now. (default parameter is `1`)
- by setting `-e` option, you can generate log data including bad http status (like 403, 404, 500, etc). (default parameter is `0.1`%)
- by setting `-f` option, you can output log data to the specified file. When `-f` option does not set, data is output to stdout.

## Output

```
100.174.112.61 - - [06/Feb/2017:14:41:08 +0900] "POST /category/games HTTP/1.1" 403 1286 "-" "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36" 24478
...
158.208.79.152 - - [07/Feb/2017:14:41:08 +0900] "DELETE /category/health HTTP/1.1" 403 1977 "-" "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36" 23184
```

|No.|Value|Description|
|:--|:--|:--|
|1| 100.174.112.61 | Client IP Address |
|2| - | RemoteLogname |
|3| - | RemoteUser |
|4| [06/Feb/2017:14:41:08 +0900] | RequestTime |
|5| POST /category/games HTTP/1.1 | Request |
|6| 403 | HttpStatusCode |
|7| 1286 | Size of Bytes |
|8| -   |Referer |
|9| Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36 |  User-Agent |
|10| 24478 | ResponseTime (ms) |

## TODO
- Write Test
- Enable to configure parameter

## References
Some functions and ideas are used as a reference from...
- [tamtam180/apache_log_gen](https://github.com/tamtam180/apache_log_gen)
- [Art-Wolf/ApacheLogGenerator](https://github.com/Art-Wolf/ApacheLogGenerator)

## Contact
Kohei Suzuki [@skjune12](http://github.com/skjune12)
