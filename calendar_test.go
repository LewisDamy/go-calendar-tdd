package go_calendar

import (
	"github.com/lewisdamy/go-calendar/pkg/date"
	"testing"
)

func TestWeekdayHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 6)

	c.SetWeekdayHoliday("Saturday")

	if c.IsHoliday(d) != true {
		t.Errorf("Saturday should be an holiday")
	}
}

func TestWeekdayNotHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 1)

	if c.IsHoliday(d) != false {
		t.Errorf("Monday should not be an holiday")
	}
}

func TestMultipleWeekdayNotHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 6)
	d2 := date.NewDate(2020, 6, 7)

	c.SetWeekdayHoliday("Saturday")
	c.SetWeekdayHoliday("Sunday")

	if c.IsHoliday(d) != true {
		t.Errorf("Saturday should not be an holiday")
	}
	if c.IsHoliday(d2) != true {
		t.Errorf("Sunday should be an holiday")
	}
}

func TestSpecificHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 1, 1)

	c.SetDayHoliday(2020, 1, 1)

	if c.IsHoliday(d) != true {
		t.Errorf("New years Date should be a holiday")
	}
}

func TestDayOfMonthHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 7, 4)
	d2 := date.NewDate(2021, 7, 4)

	c.SetDayOfMonthHoliday(7, 4)

	if c.IsHoliday(d) != true {
		t.Errorf("July 4th 2020 should be a holiday")
	}
	if c.IsHoliday(d2) != true {
		t.Errorf("July 4th 2021 should be a holiday")
	}
}
