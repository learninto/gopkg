package cronexpr

import (
	"errors"
	"kingstar-go/commons/timelib"
	"reflect"
	"testing"
	"time"
)

// TestExpression_Next TODO 待完善
func TestExpression_Next(t *testing.T) {
	startDate := timelib.StrToTime("2006-01-02", "2020-05-09")
	endDate := timelib.StrToTime("2006-01-02", "2021-05-09")
	fromTime := timelib.StrToTime("2006-01-02", "2020-05-16")

	type args struct {
		cronLine  string
		fromTime  time.Time
		startDate time.Time
		endDate   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"表达式： 每 X 天",
			args{
				"日 每 1",
				fromTime,
				startDate,
				endDate,
			},
			timelib.StrToTime("2006-01-02  15:04:05", "2020-05-16 00:00:00"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, _ := Parse(tt.args.cronLine)
			got, err := expr.Next(tt.args.fromTime, tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpression_NextN(t *testing.T) {
	startDate := timelib.StrToTime("2006-01-02", "2020-05-08")
	endDate := timelib.StrToTime("2006-01-02", "2021-05-09")
	fromTime := timelib.StrToTime("2006-01-02", "2020-05-18")

	type args struct {
		cronLine  string
		fromTime  time.Time
		startDate time.Time
		endDate   time.Time
		n         int
	}
	tests := []struct {
		name         string
		args         args
		wantListCron []time.Time
		wantErr      bool
	}{
		// --------------- 日 ---------------
		{
			"表达式：日 每 1",
			args{
				"日 每 1",
				fromTime,
				startDate,
				endDate,
				10,
			},
			[]time.Time{
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-18 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-19 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-20 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-21 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-22 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-23 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-24 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-25 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-26 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-27 00:00:00"),
			},
			false,
		},
		{
			"表达式： 日 每 2",
			args{
				"日 每 2",
				fromTime,
				startDate,
				endDate,
				20,
			},
			[]time.Time{
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-18 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-20 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-22 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-24 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-26 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-28 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-30 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-01 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-03 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-05 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-07 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-09 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-11 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-13 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-15 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-17 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-19 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-21 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-23 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-25 00:00:00"),
			},
			false,
		},
		{
			"表达式： 日 每 工作日",
			args{
				"日 每 工作日",
				fromTime,
				startDate,
				endDate,
				20,
			},
			[]time.Time{
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-18 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-19 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-20 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-21 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-22 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-25 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-26 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-27 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-28 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-05-29 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-01 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-02 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-03 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-04 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-05 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-08 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-09 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-10 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-11 00:00:00"),
				timelib.StrToTime("2006-01-02  15:04:05", "2020-06-12 00:00:00"),
			},
			false,
		},
		// --------------- 周 ---------------
		//{
		//	"表达式： 周 每 1 1",
		//	args{
		//		"周 每 1 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 1 1,3",
		//	args{
		//		"周 每 1 1,3",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 1 ANY",
		//	args{
		//		"周 每 1 ANY",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 2 ANY",
		//	args{
		//		"周 每 2 ANY",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 10 1,3",
		//	args{
		//		"周 每 10 1,3",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 4 ANY",
		//	args{
		//		"周 每 4 ANY",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 周 每 1 1,0",
		//	args{
		//		"周 每 1 1,0",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 1 第 2 1",
		//	args{
		//		"月 每 1 第 2 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 1 第 1 1",
		//	args{
		//		"月 每 1 第 1 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 1 1",
		//	args{
		//		"月 每 1 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 1 第 4 4",
		//	args{
		//		"月 每 1 第 4 4",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 1 第 L 1,
		//	args{
		//		"月 每 1 第 L 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式：月 每 1 ANY",
		//	args{
		//		"月 每 1 ANY",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式： 月 每 2 ANY",
		//	args{
		//		"月 每 2 ANY",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式：月 每 2 1",
		//	args{
		//		"月 每 2 1",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
		//{
		//	"表达式：月 每 3 31",
		//	args{
		//		"月 每 3 31",
		//		fromTime,
		//		startDate,
		//		endDate,
		//		20,
		//	},
		//	[]time.Time{},
		//	false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, _ := Parse(tt.args.cronLine)
			gotListCron, err := expr.NextN(tt.args.fromTime, tt.args.startDate, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotListCron, tt.wantListCron) {
				t.Log(gotListCron)
				t.Errorf("NextN() gotListCron = %v, want %v", gotListCron, tt.wantListCron)
			}
		})
	}
}

func TestExpression_Validate(t *testing.T) {
	startDate := timelib.StrToTime("2006-01-02", "2020-05-09")
	fromTime := timelib.StrToTime("2006-01-02", "2020-05-16")

	type args struct {
		cronLine  string
		fromTime  time.Time
		startTime time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"表达式： 每 X 天",
			args{
				"日 每 2",
				fromTime,
				startDate,
			},
			errors.New("不在计划内"),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, _ := Parse(tt.args.cronLine)
			expr, got := expr.Validate(tt.args.fromTime, tt.args.startTime)
			if (got != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", got, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		cronLine string
	}
	tests := []struct {
		name    string
		args    args
		want    *Expression
		wantErr bool
	}{
		{
			"表达式： 每 X 天",
			args{"日 每 2"},
			&Expression{expression: "日 每 2", first: "日", second: "每", third: "2"},
			false,
		}, {
			"表达式：每个工作日",
			args{"日 每 工作日"},
			&Expression{expression: "日 每 工作日", first: "日", second: "每", third: "工作日"},
			false,
		}, {
			"表达式：每 X 周 的 周Y",
			args{"周 每 2 1,2"},
			&Expression{expression: "周 每 2 1,2", first: "周", second: "每", third: "2", fourth: "1,2", weekList: []string{"1", "2"}},
			false,
		},
		{
			"表达式： 每 X 周 的 任意一天",
			args{"周 每 2 ANY"},
			&Expression{expression: "周 每 2 ANY", first: "周", second: "每", third: "2", fourth: "ANY"},
			false,
		},
		{
			"表达式： 每 X 月 的 第 Y 天",
			args{"月 每 2 16"},
			&Expression{expression: "月 每 2 16", first: "月", second: "每", third: "2", fourth: "16"},
			false,
		},
		{
			"表达式： 每 X 月 的 第 Y 个周 Z",
			args{"月 每 2 第 2 1"},
			&Expression{expression: "月 每 2 第 2 1", first: "月", second: "每", third: "2", fourth: "第", fifth: "2", sixth: "1"},
			false,
		},
		{
			"表达式： 每 X 月 任意一天",
			args{"月 每 2 ANY"},
			&Expression{expression: "月 每 2 ANY", first: "月", second: "每", third: "2", fourth: "ANY"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.cronLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
