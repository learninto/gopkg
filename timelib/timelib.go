package timelib

import "time"

/******************************************************************************/

// StrToTime 字符串转时间对象
func StrToTime(layout, value string) (t time.Time) {
	t, _ = time.ParseInLocation(layout, value, time.Local)
	return
}

/******************************************************************************/

// SubMonth 计算日期相差多少月
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

/******************************************************************************/

// SubDays 计算日期相差多少天
func SubDays(t1, t2 time.Time) (day int) {
	day = int(t1.Sub(t2).Hours() / 24)
	return
}

/******************************************************************************/

// GetWeeksByTime 根据时间获取周的日期列表
func GetWeeksByTime(t time.Time) (startWeeks []time.Time) {
	weekStartDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, -int(t.Weekday()))
	startWeeks = append(startWeeks, weekStartDate)
	for i := 0; i < 6; i++ {
		startWeeks = append(startWeeks, startWeeks[i].AddDate(0, 0, 1))
	}
	// 这条是根据中国人的习惯，将周日放在最后一天
	startWeeks[0] = startWeeks[6].AddDate(0, 0, 1)
	return startWeeks
}

/******************************************************************************/

// GetFirstDayOfMonth 获取月的第一天
func GetFirstDayOfMonth(t time.Time) time.Time {
	currentYear, currentMonth, _ := t.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, t.Location())
	return firstOfMonth
}

/******************************************************************************/

// GetCurDayTime 获取当前天的时间戳
func GetCurDayTime() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

/******************************************************************************/

// GetLastMonthStartEnd 获取上个月开始结束时间
func GetLastMonthStartEnd() (lastMonthStart, lastMonthEnd time.Time) {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart = time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location())

	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd = time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, now.Location())
	return
}

/******************************************************************************/
