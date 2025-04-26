package code

// 管理和安排课程表的 demo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	TotalWeeks  = 20 //一共有多少周
	DayPerWeek  = 7
	HoursPerDay = 24
)

type Table struct {
	table [][]int
}

func NewTable() *Table {
	table := make([][]int, DayPerWeek+1)
	for i := range table {
		table[i] = make([]int, HoursPerDay+1)
	}
	return &Table{table}
}

func (ag *Table) AddBusyTime(startHour, endHour, day int) {
	for i := startHour; i <= endHour; i++ {
		ag.table[day][i] += 1
	}
}

func (ag *Table) FindFreeTime() []string {
	var freeTimes []string
	for day := 1; day <= DayPerWeek; day++ {
		for hour := 1; hour <= HoursPerDay; hour++ {
			if ag.table[day][hour] == 0 {
				freeTimes = append(freeTimes, fmt.Sprintf("%d-%d", day, hour))
			}
		}
	}
	return freeTimes
}

// 处理调课情况
func (ag *Table) Adjust(oStartHour, oEndHour, oDay, startHour, endHour, day int) {
	for i := oStartHour; i <= oEndHour; i++ {
		ag.table[oDay][i] -= 1
	}
	for i := startHour; i <= endHour; i++ {
		ag.table[day][i] += 1
	}
}

// ExgClass 格式化输入课表
func ExgClass(time string) (int, int, int, int, int) {
	res := exg(time)
	startWeek := res[0]
	endWeek := res[1]
	day := res[2]
	startTime := res[3]
	endTime := res[4]
	return startWeek, endWeek, day, startTime, endTime
}

// ExgAdjust 格式化输入调课信息
func ExgAdjust(time string) (int, int, int, int, int, int, int, int, int, int) {
	res := exg(time)
	oStartWeek := res[0]
	oEndWeek := res[1]
	oDay := res[2]
	oStartTime := res[3]
	oEndTime := res[4]
	startWeek := res[5]
	endWeek := res[6]
	day := res[7]
	startTime := res[8]
	endTime := res[9]
	return oStartWeek, oEndWeek, oDay, oStartTime, oEndTime, startWeek, endWeek, day, startTime, endTime
}

func exg(time string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(time, -1)
	res := make([]int, 0)
	// 去除前导零并转换为整数
	for _, match := range matches {
		match = strings.TrimLeft(match, "0")
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			continue
		}
		res = append(res, num)
	}
	return res
}

// 格式化输出
func readData(week int, data []string) {

	groupedData := make(map[int][]int)

	for _, item := range data {
		parts := strings.Split(item, "-")
		if len(parts) == 2 {
			day, _ := strconv.Atoi(parts[0])
			hour, _ := strconv.Atoi(parts[1])

			groupedData[day] = append(groupedData[day], hour)
		}
	}

	for day, hours := range groupedData {
		start := hours[0]
		for k, i := range hours {
			if k > 0 {
				if hours[k]-hours[k-1] != 1 {
					fmt.Printf("第%d周：星期%d,%d-%d\n", week+1, day, start, hours[k-1])
					start = i
				}
			}
			if k == len(hours)-1 {
				fmt.Printf("第%d周：星期%d,%d-%d\n", week+1, day, start, hours[k])
			}
		}
	}
}

func RunClassArrange() {
	tables := make([]*Table, 0)

	for i := 0; i < TotalWeeks; i++ {
		table := NewTable()
		tables = append(tables, table)
	}

	//这个是课表的时间
	s := make([]string, 0)
	s = append(s, "02-03 星期4:5-6节")
	s = append(s, "10-11 星期6:1-8节")
	s = append(s, "01-16 星期1:5-6节")

	//m把节转成小时
	m := map[int]int{
		1:  8,
		2:  9,
		3:  10,
		4:  11,
		5:  14,
		6:  15,
		7:  16,
		8:  17,
		9:  19,
		10: 20,
		11: 21,
	}
	for _, time := range s {
		startWeek, endWeek, day, startTime, endTime := ExgClass(time)
		for i := startWeek; i <= endWeek; i++ {
			tables[i-1].AddBusyTime(m[startTime], m[endTime], day)
		}

	}
	for i := 0; i < TotalWeeks; i++ {
		readData(i, tables[i].FindFreeTime())
	}
}
