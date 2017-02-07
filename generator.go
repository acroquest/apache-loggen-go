package loggen

import (
	"bufio"
	"fmt"
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

func randInt(min int, max int) string {
	return strconv.Itoa(min + rand.Intn(max-min))
}

// TODO: exclude private IP address
func Ipv4Address() string {
	var ipStr string

	ipStr = randInt(1, 223) + "."
	ipStr += randInt(0, 255) + "."
	ipStr += randInt(0, 255) + "."
	ipStr += randInt(0, 255)

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

func DocType() string {
	s := []string{".html", ".gif", ".jpeg", ".png"}
	return s[rand.Intn(len(s))]
}

func ReturnCode() string {
	s := []string{"200", "301", "403", "404", "500"}
	return s[rand.Intn(len(s))]
}

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)

	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
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

func ReturnUserAgent() string {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(useragentList) == 0 {
		useragentList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen/resources/useragents.txt")
	}
	useragent := useragentList[rand.Intn(len(useragentList))]
	return useragent
}

func ReturnRequest() string {
	if len(categoryList) == 0 {
		categoryList = returnNewList(os.Getenv("GOPATH") + "/src/github.com/acroquest/apache-loggen/resources/categories.txt")
	}
	category := categoryList[rand.Intn(len(categoryList))]

	i := rand.Intn(10)
	if i < 7 {
		return "\"" + RequestType() + " /category/" + category + " HTTP/1.1\" "
	} else {
		return "\"" + RequestType() + " /" + category + "/" + randInt(1, 999) + " HTTP/1.1\" "
	}
}

func ReturnReferer() string {
	referer := "-"
	return "\"" + referer + "\""
}

func ReturnRecord(i int) string {
	bytes := randInt(20, 5000)
	referer := ReturnReferer()
	responseTime := randInt(20000, 30000)
	if random := rand.Intn(1000); random == 1 {
		responseTime = randInt(60000, 80000)
	}
	return Ipv4Address() + " - - [" + RequestTime(i) + "] " + ReturnRequest() + ReturnCode() + " " + bytes + " " + referer + " \"" + ReturnUserAgent() + "\" " + responseTime

}

func GenerateLog(days int) {
	currentTime = endTime.Add(-24 * time.Hour * time.Duration(days))
	for i := 0; endTime.Sub(currentTime.Add(time.Second*time.Duration(i))) >= 0; i += 1 {
		rand.Seed(time.Now().UTC().UnixNano())
		j := rand.Intn(10)
		hour := currentTime.Add(time.Second * time.Duration(i)).Hour()

		switch {
		case hour >= 1 && hour <= 5:
			if j <= 2 {
				fmt.Println(ReturnRecord(i))
			}
		case hour >= 6 && hour <= 9:
			if j <= 4 {
				fmt.Println(ReturnRecord(i))
			}
		case hour >= 10 && hour <= 17:
			if j <= 6 {
				fmt.Println(ReturnRecord(i))
			}
		case hour >= 18 && hour <= 23:
			if j <= 8 {
				fmt.Println(ReturnRecord(i))
			}
		case hour == 0:
			if j <= 6 {
				fmt.Println(ReturnRecord(i))
			}
		}
	}

}
