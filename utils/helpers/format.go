package helpers

import (
	"fmt"
	"strings"
	"time"

	"todo/utils/errors"

	"github.com/jinzhu/now"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FormatFloatNumber(value interface{}) string {
	p := message.NewPrinter(language.Indonesian)
	return strings.Replace(p.Sprintf("%.2f", value), ",00", "", 1)
}

func MaskString(str string) string {
	if len(str) <= 3 {
		return strings.Repeat("*", len(str))
	} else {
		return str[:3] + strings.Repeat("*", len(str[3:]))
	}
}

func FormatDayIndonesia(value string) string {

	switch value {
	case "Monday":
		return "Senin"
	case "Tuesday":
		return "Selasa"
	case "Wednesday":
		return "Rabu"
	case "Thursday":
		return "Kamis"
	case "Friday":
		return "Jumat"
	case "Saturday":
		return "Sabtu"
	case "Sunday":
		return "Minggu"
	}

	return ""
}

func FormatMonthIndonesia(value string) string {

	switch value {
	case "January":
		return "Januari"
	case "February":
		return "Februari"
	case "March":
		return "Maret"
	case "April":
		return "April"
	case "May":
		return "Mei"
	case "June":
		return "Juni"
	case "July":
		return "Juli"
	case "August":
		return "Agustus"
	case "September":
		return "September"
	case "October":
		return "Oktober"
	case "November":
		return "November"
	case "December":
		return "Desember"
	}

	return ""
}

func FormatTimeStringToUTC(strTimeISO string) (string, error) {
	date := time.Now().Format("2006-01-02")
	ifTodayTime, err := now.Parse(fmt.Sprint(date, "T", strTimeISO))
	if err != nil {
		err := errors.ErrBadRequest
		err.Message = "time not valid"
		return "", err
	}
	ifTodayTimeISO := ifTodayTime.UTC().Format("15:04:05Z07:00") // 15:04:05Z

	return ifTodayTimeISO, nil
}
