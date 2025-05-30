package go_calendar

import "github.com/lewisdamy/go-calendar-tdd/pkg/date"

type HolidayCalendar struct {
	holidays []Holiday
}

type Holiday interface {
	Equals(date.Date) bool
}

type DayOfWeekHoliday struct {
	dayOfWeek string
}

type SpecificDateHoliday struct {
	specificDate date.Date
}

type DayOfMonthHoliday struct {
	day   int
	month int
}

func (c *HolidayCalendar) SetWeekdayHoliday(DayOfTheWeek string) {
	c.holidays = append(c.holidays, DayOfWeekHoliday{dayOfWeek: DayOfTheWeek})
}

func (h DayOfWeekHoliday) Equals(inputDate date.Date) bool {
	if inputDate.Weekday() == h.dayOfWeek {
		return true
	}
	return false
}

func (c *HolidayCalendar) SetDateHoliday(year, month, day int) {
	c.holidays = append(c.holidays, SpecificDateHoliday{
		specificDate: date.NewDate(year, month, day),
	})
}

func (h SpecificDateHoliday) Equals(inputDate date.Date) bool {
	if inputDate == h.specificDate {
		return true
	}
	return false
}

func (c *HolidayCalendar) SetDayOfMonthHoliday(month, day int) {
	c.holidays = append(c.holidays, DayOfMonthHoliday{
		day:   day,
		month: month,
	})
}

func (h DayOfMonthHoliday) Equals(inputDate date.Date) bool {
	if h.day == inputDate.Day() && h.month == inputDate.Month() {
		return true
	}
	return false
}

func (c HolidayCalendar) IsHoliday(date date.Date) bool {
	for _, v := range c.holidays {
		if v.Equals(date) {
			return true
		}
	}
	return false
}
