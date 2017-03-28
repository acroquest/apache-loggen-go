package loggen

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	useragentList []string
	categoryList  []string
	ipList        []string
	endTime       = time.Now()
	marker        = endTime
)

func causeErr(errRate float64) bool {
	rand.Seed(time.Now().UnixNano())
	for i := 1.0; ; i *= 10 {
		if errRate*i >= 1.0 {
			n := rand.Intn(100 * int(i))
			if int(errRate*i) > n {
				return true
			} else {
				return false
			}
		}
	}
}

func floatToIntString(input float64) string {
	return strconv.Itoa(int(input))
}

func increnemt(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func intToString(input int) string {
	return strconv.Itoa(input)
}

func outputRecord(tick int, config Config) {
	filename := config.Filename
	if filename == "" {
		// ファイル名が指定されているときの処理
		for i := 0; i < randInt(1, 3); i++ {
			fmt.Println(GetRecord(tick, config))
		}
	} else {
		// ファイル名が指定されていないときの処理
		var f string
		splitted := strings.Split(config.Filename, ".")
		if len(splitted) == 1 {

			// TODO: 関数化
			for i := 1; i <= config.NumOfFiles; i++ {
				f = filename + "-" + strconv.Itoa(i) + ".log"

				// BEGIN
				file, _ := os.OpenFile(f, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
				defer file.Close()

				writer := bufio.NewWriter(file)
				for j := 0; j < randInt(1, 3); j++ {
					record := ([]byte)(GetRecord(tick, config) + "\n")
					writer.Write(record)
					writer.Flush()
				}
				// END
			}
		} else {

			// TODO: 関数化
			head := strings.Join(splitted[:len(splitted)-1], "")
			for i := 1; i <= config.NumOfFiles; i++ {
				f = head + "-" + strconv.Itoa(i) + "." + splitted[len(splitted)-1]

				// BEGIN
				file, _ := os.OpenFile(f, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
				defer file.Close()

				writer := bufio.NewWriter(file)
				for j := 0; j < randInt(1, 3); j++ {
					record := ([]byte)(GetRecord(tick, config) + "\n")
					writer.Write(record)
					writer.Flush()
				}
				// END
			}
		}
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// Generate random number based on log-normal distribution
func randLogNormal(mu, sigma float64) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Float64()
	r2 := rand.New(s1).Float64()
	z := mu + sigma*math.Sqrt(-2.0*math.Log(r1))*math.Sin(2.0*math.Pi*r2)
	return math.Exp(z)
}

func returnNewList(path string) []string {
	var newList []string

	absPath, _ := filepath.Abs(path)
	fp, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		newList = append(newList, scanner.Text())
	}

	return newList
}

func zitter(i int) int {
	rand.Seed(time.Now().UnixNano())
	min := i - randInt(1, 2)
	max := i + randInt(1, 2)
	if min < 0 {
		min = 1
	}
	return min + rand.Intn(max-min)
}

func GenerateNewRecord(config Config) {
	var weight int
	var days = config.Days

	marker = endTime.Add(-24 * time.Hour * time.Duration(days))
	beforeHour := endTime.Hour()

	for tick := 0; endTime.Sub(marker.Add(time.Second*time.Duration(tick))) >= 0; tick++ {
		hour := marker.Add(time.Second * time.Duration(tick)).Hour()
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(10)

		// weight is used for adding variation to the amount of data
		if hour != beforeHour {
			weight = zitter(randInt(1, 3))
			beforeHour = hour
		}

		switch {
		case hour >= 1 && hour <= 5:
			if j <= 2+weight {
				outputRecord(tick, config)
			}
		case hour >= 6 && hour <= 9:
			if j <= 4+weight {
				outputRecord(tick, config)
			}
		case hour >= 10 && hour <= 17:
			if j <= 6+weight {
				outputRecord(tick, config)
			}
		case hour >= 18 && hour <= 23:
			if j <= 6+weight {
				outputRecord(tick, config)
			}
		default:
			if j <= 4+weight {
				outputRecord(tick, config)
			}
		}
	}
}

func GetRecord(i int, config Config) string {
	if config.Format == "" {
		return Ipv4Address(config.Prefix) + " - - [" + RequestTime(i) + "] \"" + Request() + "\" " + HttpStatusCode(config.ErrRate) + " " + SizeofBytes(config.Bytes) + " " + "\"" + Referer() + "\" " + " \"" + UserAgent() + "\" " + ResponseTime(config.ResponseTime)
	} else {
		return parseFormat(i, config)
	}
}

func parseFormat(i int, config Config) string {
	base := config.Format

	formatted := strings.Replace(base, "%h", Ipv4Address(config.Prefix), 1)
	formatted = strings.Replace(formatted, "%l", "-", 1)
	formatted = strings.Replace(formatted, "%u", "-", 1)
	formatted = strings.Replace(formatted, "%t", "["+RequestTime(i)+"]", 1)
	formatted = strings.Replace(formatted, "%r", Request(), 1)
	formatted = strings.Replace(formatted, "%>s", HttpStatusCode(config.ErrRate), 1)
	formatted = strings.Replace(formatted, "%b", SizeofBytes(config.Bytes), 1)
	formatted = strings.Replace(formatted, "%{Referer}i", Referer(), 1)
	formatted = strings.Replace(formatted, "%{User-Agent}i", UserAgent(), 1)
	formatted = strings.Replace(formatted, "%D", ResponseTime(config.ResponseTime), 1)

	return formatted
}

func Ipv4Address(cidr string) string {
	if len(ipList) == 0 {
		v4addr, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}

		for v4addr := v4addr.Mask(ipnet.Mask); ipnet.Contains(v4addr); increnemt(v4addr) {
			ipList = append(ipList, v4addr.String())
		}
	}

	ip := ipList[rand.Intn(len(ipList))]
	return ip
}

func HttpStatusCode(errRate float64) string {
	rand.NewSource(time.Now().UnixNano())
	if causeErr(errRate) == false {
		return "200"
	} else {
		s := []string{"301", "403", "404", "500"}
		return s[rand.Intn(len(s))]
	}
}

func Request() string {
	if len(categoryList) == 0 {
		categoryList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen-go/resources/categories.txt")
	}
	category := categoryList[rand.Intn(len(categoryList))]

	i := rand.Intn(10)
	if i < 7 {
		return RequestType() + " /category/" + category + " HTTP/1.1"
	} else {
		return RequestType() + " /" + category + "/" + intToString(randInt(1, 999)) + " HTTP/1.1"
	}
}

func Referer() string {
	referer := "-"
	return referer
}

func RequestTime(i int) string {
	returnTime := marker.Add(time.Second * time.Duration(i))
	return returnTime.Format("02/Jan/2006:15:04:05 -0700")
}

func RequestType() string {
	s := []string{"GET", "POST", "PUT", "DELETE"}
	return s[rand.Intn(len(s))]
}

func ResponseTime(millisecond int) string {
	return floatToIntString(randLogNormal(0.0, 0.5) * float64(millisecond))
}

func UserAgent() string {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(useragentList) == 0 {
		useragentList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen-go/resources/useragents.txt")
	}
	useragent := useragentList[rand.Intn(len(useragentList))]
	return useragent
}

func SizeofBytes(bytes int) string {
	return floatToIntString(randLogNormal(0.0, 0.5) * float64(bytes))
}
