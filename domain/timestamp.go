package domain

type Timestamp struct {
	Interval string
	Timezone string
	Start    string
	End      string
}

func (t Timestamp) GenerateTimestampDomain(interval string, timezone string, start string, end string) Timestamp {

	return Timestamp{
		Interval: interval,
		Timezone: timezone,
		Start:    start,
		End:      end,
	}
}
