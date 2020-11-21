// Fetch prints the content found at a URL
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	// parameters
	// 1: member github id
	// 2: repo
	// 3: tree/blob
	// 4: 1, 2, 3, 4 (i-th week)
	// 5: mon/tue/wed...
	dailyURL       = "https://github.com/%s/%s/%s/master/week%d/%s/README.md"
	memberFilepath = "config/members.csv"
)

var course = *flag.String("c", "mit-test-repo", "course name given")
var week = *flag.Int("w", 1, "the i-th week of the course")
var day = *flag.Int("d", 1, "the i-th day of the week")

var days = []string{1: "mon", 2: "tue", 3: "wed", 4: "thur", 5: "fri", 6: "sat", 7: "sun"}

func main() {

	flag.Parse()

	members := loadMembers()
	dailyAssess(members)
}

func loadMembers() []string {
	memberFile, err := os.Open(memberFilepath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "getMembers open file: %v", err)
		os.Exit(1)
	}

	var rets []string

	reader := csv.NewReader(memberFile)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "getMember read csv: %v", err)
		}

		rets = append(rets, record[0])
	}

	return rets
}

func dailyAssess(members []string) {
	res := assessAll(members)

	// folder structure: data/course_name/week/day/assess.csv
	path := fmt.Sprintf("data/%s/week%d/%s/", course, week, days[day])
	os.MkdirAll(path, 0777)

	csvFile := path + "assess.csv"

	// write to a csv file
	writeMapCSV(res, csvFile)
}

func writeMapCSV(m map[string]int, filepath string) {
	var records [][]string
	for k, v := range m {
		record := []string{k, strconv.Itoa(v)}
		records = append(records, record)
	}

	out, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0777)

	if err != nil {
		fmt.Fprintf(os.Stderr, "writeMapCSV: %v", err)
		os.Exit(1)
	}

	writer := csv.NewWriter(out)

	writer.WriteAll(records)
}

func assessAll(members []string) map[string]int {
	// read members from csv file, and add them to a map

	var rets = make(map[string]int)

	for _, member := range members {
		var url = fmt.Sprintf(dailyURL, member, course, "tree", week, days[day])
		rets[member] = assess(url)
		// sleep before crawling the next one
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
	return rets
}

// return values:
// 0: if not checked in
// 1: if checked in
func assess(url string) int {
	resp, err := http.Get(url)
	fmt.Printf("fetch: %s\n", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	// if the page does not exist, not checked in yet
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("%s not found", url)
		return 0
	}

	defer resp.Body.Close()

	// TODO: tokenize the html body and check more rigorously
	// 1. check checkin date(README.md creation date)

	return 1
}
