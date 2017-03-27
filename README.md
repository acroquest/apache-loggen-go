# apache-loggen-go
Apache-loggen-go is a Golang script to generate dummy apache-formatted access log data.
It can create long-term access log very easily.

## Usage

```bash
$ git clone https://github.com/acroquest/apache-loggen-go && cd $_
$ go get
$ cd example
$ vi config.json
$ go run main.go
```

You can set some options by configuring `config.json`
There is sample configuration below.

```json
{
    "prefix": "192.168.0.0/16",
    "days": 1,
    "filename": "sample",
    "error_rate": 0.01,
    "num_of_files": 3,
    "bytes": 2000,
    "response_time": 10000
}
```

- `prefix` means client's IP address range. When you set `192.168.0.0/16`, the client IP addresses are specified in the range of this prefix.
- `days` can specify the length of log data. When you set `1` as the `days` option, you can generate log data from 3 days before the present to now.
- `error_rate` option generates a log data including bad http status (like 403, 404, 500, etc). `0.01` means the log contains error status one hundledth (that is, 1%).
- `filename` option enables the log to output the specified file. When `filename` option does not set, data is output to stdout. When it is blank, the data is written to stdout.
- `num_of_files` option creates multiple log files at once. For example, when you set `10` as a value, the files sample-1.log, sample-2.log ... sample-10.log are created. (default parameter is 0)
- `bytes` specifies the size of bytes. It follows the lognormal distribution with μ = 0, σ 0.5.
- `response_time` also specifies the time of response time. It also follows the lognormal distribution with μ = 0, σ = 0.



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
- Write Test code
- Enable to configure parameter

## References
Some functions and ideas are used as a reference from...
- [tamtam180/apache_log_gen](https://github.com/tamtam180/apache_log_gen)
- [Art-Wolf/ApacheLogGenerator](https://github.com/Art-Wolf/ApacheLogGenerator)

## Contact
- Kohei Suzuki [@skjune12](http://github.com/skjune12)
