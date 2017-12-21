package iso8601

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	result, err := PeriodFromString("DT1S")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "invalid period representation, must start with P", err.Error())

	result, err = PeriodFromString("P1S")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "found time component without time enabler 'T', 'S'", err.Error())

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
