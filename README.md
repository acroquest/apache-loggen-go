# apache-loggen
Apache-loggen is a Golang script generating dummy apache-formatted log data.
It creates long-term log data very easily.

## Usage
```bash
$ git clone github.com/acroquest/apache-loggen
$ go run example/main.go -d 3
```

- by setting `-d 3` option, you can generate log data from 3 day's before the present to now.

## Output

```
100.174.112.61 - - [06/Feb/2017:14:41:08 +0900] "POST /category/games HTTP/1.1" 403 1286 "-" "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36" 24478
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

## Contact
Kohei Suzuki [@skjune12](http://github.com/skjune12)
