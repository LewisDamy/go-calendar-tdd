package go_calendar

import "github.com/lewisdamy/go-calendar-tdd/pkg/date"

type HolidayCalendar struct {
	weekdayHolidays  []string
	specificHolidays []date.Date
}

func (c *HolidayCalendar) SetWeekdayHoliday(DayOfTheWeek string) {
	c.weekdayHolidays = append(c.weekdayHolidays, DayOfTheWeek)
}

func (c *HolidayCalendar) SetDateHoliday(year, month, day int) {
	c.specificHolidays = append(c.specificHolidays, date.NewDate(year, month, day))
}

func (c *HolidayCalendar) SetDayOfMonthHoliday(month, day int) {

}

func (c *HolidayCalendar) isWeekdayHoliday(inputDate date.Date) bool {
	for _, v := range c.weekdayHolidays {
		if inputDate.Weekday() == v {
			return true
		}
	}
	return false
}

func (c *HolidayCalendar) isSpecificHoliday(inputDate date.Date) bool {
	for _, v := range c.specificHolidays {
		if v == inputDate {
			return true
		}
	}
	return false
}

func (c HolidayCalendar) IsHoliday(date date.Date) bool {
	return c.isWeekdayHoliday(date) || c.isSpecificHoliday(date)
}
