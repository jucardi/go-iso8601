package iso8601

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeriod_FromString(t *testing.T) {
	result, err := PeriodFromString("DT1S")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid period representation, must start with P", err.Error())

	result, err = PeriodFromString("P1S")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "found time component without time enabler 'T', 'S'", err.Error())

	result, err = PeriodFromString("P1H")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "found time component without time enabler 'T', 'H'", err.Error())

	result, err = PeriodFromString("PS")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "attempting to assign 'S' but no value found", err.Error())

	result, err = PeriodFromString("PT1S4")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "the last character cannot be a number", err.Error())

	result, err = PeriodFromString("P1Q")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid value found, 'Q'", err.Error())

	result, err = PeriodFromString("PT1S")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Seconds: 1}, result)

	result, err = PeriodFromString("PT2M")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Minutes: 2}, result)

	result, err = PeriodFromString("PT3H")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Hours: 3}, result)

	result, err = PeriodFromString("P4D")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Days: 4}, result)

	result, err = PeriodFromString("P5M")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Months: 5}, result)

	result, err = PeriodFromString("P6Y")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Years: 6}, result)

	result, err = PeriodFromString("P6Y5M4DT3H2M1S")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Years: 6, Months: 5, Days: 4, Hours: 3, Minutes: 2, Seconds: 1}, result)

	result, err = PeriodFromString("P7W")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Weeks: 7}, result)

	result, err = PeriodFromString("P1Y2M10DT2H30M")
	assert.Nil(t, err)
	assert.Equal(t, &Period{Years: 1, Months: 2, Days: 10, Hours: 2, Minutes: 30}, result)
}

func TestPeriod_Normalize(t *testing.T) {
	period := Period{Years: 1, Months: 15, Weeks: 2, Days: 31, Hours: 27, Minutes: 73, Seconds: 91}
	result := period.Normalize()

	assert.Equal(t, 31, result.Seconds)
	assert.Equal(t, 14, result.Minutes)
	assert.Equal(t, 4, result.Hours)
	assert.Equal(t, 15, result.Days)
	assert.Equal(t, 0, result.Weeks)
	assert.Equal(t, 4, result.Months)
	assert.Equal(t, 2, result.Years)
}

func TestPeriod_ToDuration(t *testing.T) {
	p := &Period{
		Hours:   12,
		Minutes: 30,
		Seconds: 20,
	}

	d := time.Duration(12)*time.Hour + time.Duration(30)*time.Minute + time.Duration(20)*time.Second
	assert.Equal(t, d, p.ToDuration())
}

func TestPeriod_ToString(t *testing.T) {
	p := &Period{Years: 1, Months: 2, Days: 3, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, "P1Y2M3DT4H5M6S", p.ToString())

	p = &Period{Years: 0, Months: 2, Days: 3, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, "P2M3DT4H5M6S", p.ToString())

	p = &Period{Years: 1, Months: 0, Days: 3, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, "P1Y3DT4H5M6S", p.ToString())

	p = &Period{Years: 1, Months: 2, Days: 0, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, "P1Y2MT4H5M6S", p.ToString())

	p = &Period{Years: 1, Months: 2, Days: 3, Hours: 0, Minutes: 5, Seconds: 6}
	assert.Equal(t, "P1Y2M3DT5M6S", p.ToString())

	p = &Period{Years: 1, Months: 2, Days: 3, Hours: 4, Minutes: 0, Seconds: 6}
	assert.Equal(t, "P1Y2M3DT4H6S", p.ToString())

	p = &Period{Years: 1, Months: 2, Days: 3, Hours: 4, Minutes: 5, Seconds: 0}
	assert.Equal(t, "P1Y2M3DT4H5M", p.ToString())

	p = &Period{Years: 1, Months: 2, Days: 3, Hours: 0, Minutes: 0, Seconds: 0}
	assert.Equal(t, "P1Y2M3D", p.ToString())

	p = &Period{Years: 0, Months: 0, Days: 0, Hours: 4, Minutes: 5, Seconds: 6}
	assert.Equal(t, "PT4H5M6S", p.ToString())
}
