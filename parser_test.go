package iso8601

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	result, err := Parse("R5/2008-03-01T13:00:00Z/P1Y2M10DT2H30M/2009-03-01T13:00:00Z")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 5, result.Repeats)
	s, _ := time.Parse(time.RFC3339, "2008-03-01T13:00:00Z")
	e, _ := time.Parse(time.RFC3339, "2009-03-01T13:00:00Z")
	assert.Equal(t, s, result.Start)
	assert.Equal(t, e, result.End)
	assert.Equal(t, &Period{Years: 1, Months: 2, Days: 10, Hours: 2, Minutes: 30}, result.Period)
}

func TestParseJustDate(t *testing.T) {
	result, err := Parse("2008-03-01T13:00:00Z")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	s, _ := time.Parse(time.RFC3339, "2008-03-01T13:00:00Z")
	assert.Equal(t, s, result.Start)
	assert.Equal(t, time.Time{}, result.End)
	assert.Equal(t, 0, result.Repeats)
}

func TestParseRepetitionsNotAtTheBeginning(t *testing.T) {
	result, err := Parse("2008-03-01T13:00:00Z/R5")
	assert.Nil(t, result)
	assert.Equal(t, "repetitions component must be at the beginning of the string", err.Error())
}

func TestParseInvalidRepetitions(t *testing.T) {
	result, err := Parse("Rw/2008-03-01T13:00:00Z")
	assert.Nil(t, result)
	assert.Equal(t, "unable to parse repetitions, strconv.Atoi: parsing \"w\": invalid syntax", err.Error())
}

func TestParseZeroRepetitions(t *testing.T) {
	result, err := Parse("R-1/2008-03-01T13:00:00Z")
	assert.Nil(t, result)
	assert.Equal(t, "repeat value must be greater than zero", err.Error())
}

func TestParseUnboundRepetitions(t *testing.T) {
	result, _ := Parse("R/2008-03-01T13:00:00Z")
	assert.Equal(t, -1, result.Repeats)
}

func TestParseMultiplePeriods(t *testing.T) {
	result, err := Parse("R1/2008-03-01T13:00:00Z/P1Y2M10DT2H30M/P1Y2M10DT2H30M")
	assert.Nil(t, result)
	assert.Equal(t, "invalid iso8601, more than one period component detected", err.Error())
}

func TestParseInvalidPeriod(t *testing.T) {
	result, err := Parse("R1/2008-03-01T13:00:00Z/P1Y2X10DT2H30M")
	assert.Nil(t, result)
	assert.Equal(t, "invalid period, unable to parse, invalid value found, 'X'", err.Error())
}

func TestParseInvalidTimeComponent(t *testing.T) {
	result, err := Parse("R1/2008-03-01X13:00:00Z/P1Y2D10DT2H30M")
	assert.Nil(t, result)
	assert.Equal(t, "unable to parse time component, parsing time \"2008-03-01X13:00:00Z\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"X13:00:00Z\" as \"T\"", err.Error())
}

func TestParseMultipleEndDates(t *testing.T) {
	result, err := Parse("R1/P1Y2D10DT2H30M/2008-03-01T13:00:00Z/2008-03-01T13:00:00Z/2008-03-01T13:00:00Z")
	assert.Nil(t, result)
	assert.Equal(t, "invalid iso8601, more than one end date detected", err.Error())
}

func TestToString(t *testing.T) {
	isoStr := "R5/2008-03-01T13:00:00Z/P1Y2M10DT2H30M/2009-03-01T13:00:00Z"
	result, _ := Parse(isoStr)
	str := result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "2008-03-01T13:00:00Z/P1Y2M10DT2H30M/2009-03-01T13:00:00Z"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "P1Y2M10DT2H30M/2009-03-01T13:00:00Z"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "R5/P1Y2M10DT2H30M/2009-03-01T13:00:00Z"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "2008-03-01T13:00:00Z/P1Y2M10DT2H30M"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "R5/2008-03-01T13:00:00Z/P1Y2M10DT2H30M"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "2009-03-01T13:00:00Z"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)

	isoStr = "P1Y2M10DT2H30M"
	result, _ = Parse(isoStr)
	str = result.ToString()
	assert.Equal(t, isoStr, str)
}
