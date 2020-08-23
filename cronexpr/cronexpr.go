/*!
 *
 * Project: github.com/kingstar-go/commons/cronexpr
 * File: cronexpr.go
 * Version: 1.0
 * License: pick the one which suits you :
 *   GPL v3 see <https://www.gnu.org/licenses/gpl.html>
 *   APL v2 see <http://www.apache.org/licenses/LICENSE-2.0>
 *
 */

// Package cronexpr parses cron time expressions.
package cronexpr

import (
	"errors"
	"kingstar-go/commons/timelib"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/******************************************************************************/

// A Expression represents a specific cron time expression as defined at
type Expression struct {
	expression string   // 表达式
	first      string   // 第1个：单位 （日、周、月）
	second     string   // 第2个：周期 （每）
	third      string   // 第3个：每 X 个（日、周、月）\ 工作日
	fourth     string   // 第4个：几 (日：1-5、ANY)（周：ANY、0：周日；1：周一；2：周二；3：周三；4：周四；5：周五；6：周六）（月：第X天、ANY）
	weekList   []string // 第4个列表 星期列表 0代表周日 1-6
	fifth      string   // 第5个：第几个  1-4 L，
	sixth      string   // 第6个：星期几，只有在 fourth = 第
}

/******************************************************************************/

// Parse 解析表达式
func Parse(cronLine string) (*Expression, error) {
	var expr = &Expression{
		expression: cronLine,
	}
	cron := strings.Split(cronLine, " ")

	// 格式判断开始了
	if cronLen := len(cron); cronLen < 3 || cronLen > 6 {
		return expr, errors.New("没有传入值，没法检查！")
	}

	expr.first = cron[0] // 第一个
	if expr.first != "日" && expr.first != "周" && expr.first != "月" {
		return expr, errors.New("参数第一项不符合")
	}

	expr.second = cron[1] // 第二个
	if expr.second != "每" {
		return expr, errors.New("参数第二项不符合")
	}

	expr.third = cron[2] // 第三个
	// 判断 日 的格式
	if expr.first == "日" {
		if expr.third == "工作日" {
			return expr, nil
		}
		reg := regexp.MustCompile(`^[1-5]$`)
		if reg.MatchString(expr.third) {
			return expr, nil
		}

		return expr, errors.New("判断日：参数第三项不符合")
	}

	// 判断 周 的格式
	if expr.first == "周" {

		reg := regexp.MustCompile(`^[1-5]$`)
		if !reg.MatchString(expr.third) {
			return expr, errors.New("判断周：参数第三项不符合")
		}

		expr.fourth = cron[3] // 第四个
		// 每周任意一天
		if expr.fourth == "ANY" {
			return expr, nil
		}

		// 枚举值的判断
		expr.weekList = strings.Split(expr.fourth, ",") // 第四个列表
		if len(expr.weekList) == 0 {
			return expr, errors.New("判断周：参数第四项不符合")
		}

		WeekReg := regexp.MustCompile(`^[0-6]$`)
		for _, i2 := range expr.weekList {
			if !WeekReg.MatchString(i2) {
				return expr, errors.New("判断周：参数第四项不符合")
			}
		}

		return expr, nil
	}

	// 判断 月 的格式
	if expr.first == "月" {
		reg := regexp.MustCompile(`^[1-5]$`)
		if !reg.MatchString(expr.third) {
			return expr, errors.New("判断月：参数第三项不符合")
		}

		expr.fourth = cron[3] // 第四个
		// 每月任意一天
		if expr.fourth == "ANY" {
			return expr, nil
		}

		// 每月指定的一天
		reg = regexp.MustCompile(`^(([1-9])|((1|2)[0-9])|30|31)$`)
		if reg.MatchString(expr.fourth) {
			return expr, nil
		}

		// 每月 第X周的周几
		if expr.fourth != "第" {
			return expr, errors.New("判断月：参数第四项不符合")
		}

		expr.fifth = cron[4]
		expr.sixth = cron[5]

		reg = regexp.MustCompile(`^(["1","2","3","4","L"])$`)
		if !reg.MatchString(expr.fifth) {
			return expr, errors.New("判断月：参数第五项不符合")
		}

		reg = regexp.MustCompile(`^[0-6]$`)
		if !reg.MatchString(expr.sixth) {
			return expr, errors.New("判断月：参数第六项不符合")
		}

		return expr, nil
	}

	return nil, nil
}

/******************************************************************************/

// Validate 验证
func (expr *Expression) Validate(curTime, startTime time.Time) (*Expression, error) {
	// 一天的秒数
	var dateVal int64 = 24 * 60 * 60

	//开始时间
	startTimeVal := startTime.Unix()

	// 当前时间
	curTimeVal := curTime.Unix()

	if expr.first == "日" {
		if expr.third == "工作日" {
			reg := regexp.MustCompile(`^[1-5]$`)
			if reg.MatchString(expr.third) {
				return expr, nil
			}

			curWeekDay := curTime.Weekday() // 周几
			if curWeekDay > 0 && curWeekDay < 6 {
				return expr, nil
			}
			return expr, errors.New("不在计划内")
		}

		third, _ := strconv.ParseInt(expr.third, 10, 64)
		if (curTimeVal-startTimeVal)%(third*dateVal) == 0 { // 秒为单位
			return expr, nil
		}

		return expr, errors.New("不在计划内")
	}

	if expr.first == "周" {
		// 根据开始时间获取一周的时间
		startWeeks := timelib.GetWeeksByTime(startTime) // 0 代表周日 1-6

		// 每 X 周 任意一天
		if expr.fourth == "ANY" {
			third, _ := strconv.ParseInt(expr.third, 10, 64)
			if (curTimeVal-startWeeks[1].Unix())%(third*7*dateVal) == 0 {
				return expr, nil
			}

			return expr, errors.New("不在计划内")
		}

		// 每 X 周 指定日期
		for i := 0; i < len(expr.weekList); i++ {

			weekVal, _ := strconv.ParseInt(expr.weekList[i], 10, 64)
			third, _ := strconv.ParseInt(expr.third, 10, 64)

			if (curTimeVal-startWeeks[weekVal].Unix())%(third*7*dateVal) == 0 {
				return expr, nil
			}
		}

		return expr, errors.New("不在计划内")
	}

	if expr.first == "月" {
		curYear, _ := strconv.Atoi(curTime.Format("2006"))
		curMonth, _ := strconv.Atoi(curTime.Format("01"))
		curWeekDay := int(curTime.Weekday()) // 周几
		curDay, _ := strconv.Atoi(curTime.Format("02"))

		startYear, _ := strconv.Atoi(startTime.Format("2006"))
		startMonth, _ := strconv.Atoi(startTime.Format("01"))

		third, _ := strconv.Atoi(expr.third) // 第3个
		if expr.fourth == "ANY" {
			// 每 X 月 任意一天

			// 如果 (当前月份 - 计划开始日期月份) 整除 X  并且 是当月1号， 则可以生成
			if ((curYear-startYear)*12+curMonth-startMonth)%third == 0 && curDay == 1 {
				return expr, nil
			}

			return expr, errors.New("不在计划内")
		}

		if expr.fourth == "第" {
			// 每 X 月 的 第 Y 个周 Z

			// 如果 (当前月份 - 计划开始日期月份) 整除 X  不通过，直接否了
			if ((curYear-startYear)*12+curMonth-startMonth)%third != 0 {
				return expr, errors.New("不在计划内")
			}

			// 如果 周几不符合，则直接过
			sixth, _ := strconv.Atoi(expr.sixth) // 第6个
			if curWeekDay != sixth {
				return expr, errors.New("不在计划内")
			}

			var findTimes int64 = 0

			curFirstOfMonth := timelib.GetFirstDayOfMonth(curTime) // 当前月的第一天
			curLastOfMonth := curFirstOfMonth.AddDate(0, 1, -1)    // 当前月的最后一天
			if expr.fifth == "L" {                                 // 第5个 L

				for i := curLastOfMonth.Unix(); i > curFirstOfMonth.Unix(); i -= dateVal {

					//时间戳 to 时间
					iWeekday := int(time.Unix(i, 0).Weekday())
					if iWeekday != sixth { // 周几
						continue
					}

					findTimes++
					if findTimes == 1 && curTimeVal == i {
						return expr, nil
					}
				}
				return expr, errors.New("不在计划内")
			}

			for i := curFirstOfMonth.Unix(); i < curLastOfMonth.Unix()+dateVal; i = (i + dateVal) {
				if int(time.Unix(i, 0).Weekday()) != sixth { // 周几
					continue
				}

				findTimes++
				fifth, _ := strconv.ParseInt(expr.fifth, 10, 64) // 第5个
				if findTimes == fifth && curTimeVal == i {
					return expr, nil
				}

			}
			return expr, errors.New("不在计划内")
		}

		// ------------------------ 每 X 月 的 第 Y 天 ------------------------

		// 如果 (当前月份 - 计划开始日期月份) 整除 X  并且 日期， 则可以生成
		if ((curYear-startYear)*12+curMonth-startMonth)%third != 0 {
			return expr, errors.New("不在计划内")
		}
		curFirstOfMonth := timelib.GetFirstDayOfMonth(curTime)       // 当前月的第一天
		curLastOfMonth := curFirstOfMonth.AddDate(0, 1, -1)          // 当前月的最后一天的 日期
		curLastOfDay, _ := strconv.Atoi(curLastOfMonth.Format("02")) // 当前月的最后一天的 天

		// 当月最后一天 <= 如果指定日期 ， 则指定日期 = 当月最后一天
		fourth, _ := strconv.Atoi(expr.fourth)
		if curLastOfDay <= fourth {
			fourth = curLastOfDay
		}

		if curDay == fourth {
			return expr, nil
		}

		return expr, errors.New("不在计划内")
	}

	return expr, nil
}

/******************************************************************************/

/******************************************************************************/
