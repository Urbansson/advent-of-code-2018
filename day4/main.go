package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/urbansson/advent-of-code/util"
)

var dateLayout = "2006-01-02 15:04"

type log struct {
	Time    time.Time
	Message string
}

type sleep struct {
	from time.Time
	to   time.Time
}

type guard struct {
	id    string
	sleep []sleep
}

func (g guard) totalSleepDuration() time.Duration {
	var total time.Duration
	for _, s := range g.sleep {
		total += s.to.Sub(s.from)
	}
	return total
}

func (g guard) mostCommonSleepMinute() (int, int) {
	minutes := make(map[int]int, 0)
	for _, s := range g.sleep {
		for i := s.from.Minute(); i < s.to.Minute(); i++ {
			minutes[i] = minutes[i] + 1
		}
	}

	var maxKey int
	var maxMinute int
	for k, v := range minutes {
		if v > maxMinute {
			maxKey = k
			maxMinute = v
		}
	}

	return maxKey, maxMinute
}

func main() {
	args := os.Args[1:]
	file := args[0]
	fmt.Println("Using input:", file)
	fc := util.ReadFile(file)

	sortedLogs := parse(fc)

	guards := make(map[string]*guard, 0)

	var from time.Time
	var guardID string
	for _, sl := range sortedLogs {

		if strings.HasPrefix(sl.Message, "Guard") {
			guardID = getGuardID(sl.Message)
			if _, ok := guards[guardID]; !ok {
				guards[guardID] = &guard{guardID, make([]sleep, 0)}
			}

		} else if strings.HasPrefix(sl.Message, "falls asleep") {
			from = sl.Time
		} else if strings.HasPrefix(sl.Message, "wakes up") {
			guard := guards[guardID]
			guard.sleep = append(guard.sleep, sleep{from, sl.Time})
		}
	}

	var max *guard
	var max2 *guard

	for _, g := range guards {
		if max == nil {
			max = g
		} else if max.totalSleepDuration() < g.totalSleepDuration() {
			max = g
		}

		if max2 == nil {
			max2 = g
		}
		_, mmcsm := max.mostCommonSleepMinute()
		_, omcsm1 := g.mostCommonSleepMinute()

		if mmcsm < omcsm1 {
			max2 = g
		}
	}

	min, _ := max.mostCommonSleepMinute()
	maxMin, _ := max2.mostCommonSleepMinute()

	id, _ := strconv.Atoi(max.id)
	id2, _ := strconv.Atoi(max2.id)

	fmt.Println("part1:", id*min)
	fmt.Println("part2:", id2*maxMin)
}

var replacer = strings.NewReplacer("Guard #", "", " begins shift", "")

func getGuardID(message string) string {
	//Guard #349 begins shift
	return replacer.Replace(message)
}

func parse(input []string) []*log {
	logs := make([]*log, 0, len(input))

	for _, f := range input {
		p := strings.Split(f, "] ")
		t, err := time.Parse(dateLayout, p[0][1:])
		if err != nil {
			panic(err)
		}
		logs = append(logs, &log{t, p[1]})
	}
	sort.Slice(logs, func(i, j int) bool { return logs[i].Time.Before(logs[j].Time) })

	return logs
}
