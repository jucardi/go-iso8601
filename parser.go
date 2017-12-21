package iso8601

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Info represents the information held by an ISO8601 expression
type Info struct {
	Start   time.Time `json:"start" bson:"start"`
	End     time.Time `json:"end" bson:"end"`
	Repeats int       `json:"repeats" bson:"repeats"`
	Period  *Period   `json:"period" bson:"period"`
}

// Parse Parses an ISO8601 expression
func Parse(expression string) (*Info, error) {
	split := strings.Split(expression, "/")

	if len(split) == 1 {
		t, err := time.Parse(time.RFC3339, expression)
		return &Info{
			Start: t,
		}, err
	}

	endSet := false
	repeatsSet := false
	durationSet := false
	ret := &Info{}

	for i, v := range split {
		if strings.HasPrefix(v, "R") {
			if i != 0 {
				return nil, errors.New("repetitions component must be at the beginning of the string")
			}

			r, err := strconv.Atoi(v[1:])

			if err != nil {
				return nil, fmt.Errorf("unable to parse repetitions, %s", err.Error())
			}

			if r <= 0 {
				return nil, errors.New("repeat value must be greater than zero")
			}
			ret.Repeats = r
			repeatsSet = true
			continue
		}

		if strings.HasPrefix(v, "P") {
			if durationSet {
				return nil, errors.New("invalid iso8601, more than one period component detected")
			}

			p, err := PeriodFromString(v)
			if err != nil {
				return nil, fmt.Errorf("invalid period, unable to parse, %s", err.Error())
			}
			ret.Period = p
			durationSet = true
			continue
		}

		t, err := time.Parse(time.RFC3339, v)

		if err != nil {
			return nil, fmt.Errorf("unable to parse time component, %s", err.Error())
		}

		if i == 0 || i == 1 && repeatsSet {
			ret.Start = t
			continue
		} else if endSet {
			return nil, errors.New("invalid iso8601, more than one end date detected")
		}

		ret.End = t
		endSet = true
	}

	return ret, nil

}
