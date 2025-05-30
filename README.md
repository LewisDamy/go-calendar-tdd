# Go Calendar – TDD Practice

This is a simple Go project designed for practicing Test-Driven Development (TDD). It contains code and tests for checking whether a date is a holiday.

## Purpose

- Practice writing Go code using TDD.
- Build simple date-related functionality (e.g., checking holidays).
- Focus on clean, testable code without a main application.

## How to Run

1. Clone the repository:

```bash
git clone https://github.com/<your-username>/go-calendar.git
cd go-calendar
```

2. Run the tests:

```
go test ./...
```

## Features
	•	Check if a date is a holiday.
	•	Unit tests written first (TDD approach).
	•	Simple and extendable logic.

## Project Structure
```
go-calendar/
├── pkg/date/          # Date logic
│   └── date.go
├── calendar.go        # Example functions
├── calendar_test.go   # Tests for calendar.go
├── go.mod             # Go module
└── README.md
```

## Example Test

```
func TestWeekdayHoliday(t *testing.T) {
	c := HolidayCalendar{}
	d := date.NewDate(2020, 6, 6)

	c.SetWeekdayHoliday("Saturday")

	if c.IsHoliday(d) != true {
		t.Errorf("Saturday should be an holiday")
	}
}
```
