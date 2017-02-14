package loggen

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	useragentList []string
	categoryList  []string
	endTime       = time.Now()
	currentTime   = endTime
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

func intToString(input int) string {
	return strconv.Itoa(input)
}

func outputMultipleRecord(days int, errRate float64) {
	for i := 0; i < randInt(1, 3); i++ {
		fmt.Println(GetRecord(days, errRate))
	}
}

func outputToFile(days int, errRate float64, filename string) {
	// 1. check whether file is exist or not, and if file is exist, delete original file.
	_, err := os.Stat(filename)
	if err == nil {
		// file is exist
		if err := os.Remove(filename); err != nil {
			panic(err)
		}
	}

	// 2. create a new file and write the record to the file.
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < randInt(1, 3); i++ {
		file.Write(([]byte)(GetRecord(days, errRate)))
		file.Write(([]byte)("\n"))
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

// TODO: exclude private IP address
func Ipv4Address() string {
	var ipStr string
	ipStr = intToString(randInt(1, 223)) + "."
	ipStr += intToString(randInt(0, 255)) + "."
	ipStr += intToString(randInt(0, 255)) + "."
	ipStr += intToString(randInt(0, 255))
	return ipStr
}

func RequestTime(i int) string {
	returnTime := currentTime.Add(time.Second * time.Duration(i))
	return returnTime.Format("02/Jan/2006:15:04:05 -0700")
}

func RequestType() string {
	s := []string{"GET", "POST", "PUT", "DELETE"}
	return s[rand.Intn(len(s))]
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

func UserAgent() string {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(useragentList) == 0 {
		useragentList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen-go/resources/useragents.txt")
	}
	useragent := useragentList[rand.Intn(len(useragentList))]
	return useragent
}

func Request() string {
	if len(categoryList) == 0 {
		categoryList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen-go/resources/categories.txt")
	}
	category := categoryList[rand.Intn(len(categoryList))]

	i := rand.Intn(10)
	if i < 7 {
		return "\"" + RequestType() + " /category/" + category + " HTTP/1.1\" "
	} else {
		return "\"" + RequestType() + " /" + category + "/" + intToString(randInt(1, 999)) + " HTTP/1.1\" "
	}
}

func Referer() string {
	referer := "-"
	return "\"" + referer + "\""
}

func SizeofBytes(bytes int) string {
	return floatToIntString(randLogNormal(0.0, 0.5) * float64(bytes))
}

func ResponseTime(millisecond int) string {
	return floatToIntString(randLogNormal(0.0, 0.5) * float64(millisecond))
}

func GetRecord(i int, errRate float64) string {
	return Ipv4Address() + " - - [" + RequestTime(i) + "] " + Request() + HttpStatusCode(errRate) + " " + SizeofBytes(2000) + " " + Referer() + " \"" + UserAgent() + "\" " + ResponseTime(20000)
}

// TODO change the amount of log data every day.
func GenerateLog(days int, errRate float64) {
	var weight int
	currentTime = endTime.Add(-24 * time.Hour * time.Duration(days))
	beforeHour := endTime.Hour()

	// generating log data every 1 second
	for i := 0; endTime.Sub(currentTime.Add(time.Second*time.Duration(i))) >= 0; i += 1 {
		hour := currentTime.Add(time.Second * time.Duration(i)).Hour()
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(10)

		// add weight for adding variation to the amount of data
		if hour != beforeHour {
			weight = zitter(randInt(1, 3))
			beforeHour = hour
		}

		switch {
		case hour >= 1 && hour <= 5:
			if j <= 2+weight {
				outputMultipleRecord(i, errRate)
			}
		case hour >= 6 && hour <= 9:
			if j <= 4+weight {
				outputMultipleRecord(i, errRate)
			}
		case hour >= 10 && hour <= 17:
			if j <= 6+weight {
				outputMultipleRecord(i, errRate)
			}
		case hour >= 18 && hour <= 23:
			if j <= 6+weight {
				outputMultipleRecord(i, errRate)
			}
		default:
			if j <= 4+weight {
				outputMultipleRecord(i, errRate)
			}
		}
	}
}

func GenerateLogToFile(days int, errRate float64, filename string) {
	var weight int
	currentTime = endTime.Add(-24 * time.Hour * time.Duration(days))
	beforeHour := endTime.Hour()

	// generating log data every 1 second
	for i := 0; endTime.Sub(currentTime.Add(time.Second*time.Duration(i))) >= 0; i += 1 {
		hour := currentTime.Add(time.Second * time.Duration(i)).Hour()
		rand.Seed(time.Now().UnixNano())
		j := rand.Intn(10)

		// add weight for adding variation to the amount of data
		if hour != beforeHour {
			weight = zitter(randInt(1, 3))
			beforeHour = hour
		}

		switch {
		case hour >= 1 && hour <= 5:
			if j <= 2+weight {
				outputToFile(i, errRate, filename)
			}
		case hour >= 6 && hour <= 9:
			if j <= 4+weight {
				outputToFile(i, errRate, filename)
			}
		case hour >= 10 && hour <= 17:
			if j <= 6+weight {
				outputToFile(i, errRate, filename)
			}
		case hour >= 18 && hour <= 23:
			if j <= 6+weight {
				outputToFile(i, errRate, filename)
			}
		default:
			if j <= 4+weight {
				outputToFile(i, errRate, filename)
			}
		}
	}
}
