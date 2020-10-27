// time.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package time

import (
	"fmt"
	"time"
)

// Time Format.
const (
	FormatDayTimeString = "2006-01-02"
)

var emptyTime time.Time

func AddHours(
	timeStart time.Time,
	timeAddedDeltaHours float64, // Must be > 0.
) time.Time {

	return timeStart.Add(HoursToMicroseconds(timeAddedDeltaHours))
}

func GetLocationOffsetSec(
	location *time.Location,
) (offsetSec int, err error) {

	const TimeFormat = "2006-01-02 15:04:05"

	var sampleTimeString string
	var timeInLocation time.Time
	var timeInUtcTimezone time.Time
	var timeOffset time.Duration

	sampleTimeString = "2019-09-01 00:00:00"

	// Time in Location.
	timeInLocation, err = time.ParseInLocation(
		TimeFormat,
		sampleTimeString,
		location,
	)
	if err != nil {
		return
	}

	// Time in UTC Time Zone.
	timeInUtcTimezone, err = time.Parse(
		TimeFormat,
		sampleTimeString,
	)
	if err != nil {
		return
	}

	// Delta.
	timeOffset = timeInUtcTimezone.Sub(timeInLocation)
	offsetSec = int(timeOffset.Seconds())

	return
}

func GetLocationOffsetHours(
	location *time.Location,
) (offsetHrs int, err error) {

	const (
		TimeFormat = "2006-01-02 15:04:05"
	)

	var sampleTimeString string
	var timeInLocation time.Time
	var timeInUtcTimezone time.Time
	var timeOffset time.Duration

	sampleTimeString = "2019-09-01 00:00:00"

	// Time in Location.
	timeInLocation, err = time.ParseInLocation(
		TimeFormat,
		sampleTimeString,
		location,
	)
	if err != nil {
		return
	}

	// Time in UTC Time Zone.
	timeInUtcTimezone, err = time.Parse(
		TimeFormat,
		sampleTimeString,
	)
	if err != nil {
		return
	}

	// Delta.
	timeOffset = timeInUtcTimezone.Sub(timeInLocation)
	offsetHrs = int(timeOffset.Hours())

	return
}

func HoursToMicroseconds(
	hours float64,
) time.Duration {

	return time.Duration(hours*3600*1000*1000) * time.Microsecond
}

func IntervalDurationHours(
	timeStart time.Time,
	timeEnd time.Time,
) float64 {

	return timeEnd.Sub(timeStart).Hours()
}

func IsEmpty(
	t time.Time,
) bool {

	if t == emptyTime {
		return true
	}
	return false
}

func Maximum(
	a time.Time,
	b time.Time,
) time.Time {

	if a.After(b) {
		return a
	}
	return b
}

func Minimum(
	a time.Time,
	b time.Time,
) time.Time {

	if a.After(b) {
		return b
	}
	return a
}

func NewTimeStringRFC3339(
	year uint,
	month uint,
	day uint,
	hour uint,
	minute uint,
	second uint,
) string {

	var result string

	yearStr := fmt.Sprintf("%04d", year)
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	hourStr := fmt.Sprintf("%02d", hour)
	minuteStr := fmt.Sprintf("%02d", minute)
	secondStr := fmt.Sprintf("%02d", second)

	result = fmt.Sprintf(
		"%s-%s-%sT%s:%s:%sZ",
		yearStr,
		monthStr,
		dayStr,
		hourStr,
		minuteStr,
		secondStr,
	)

	return result
}

func ParseDayTimeStringInLocation(
	dayTimeString string,
	location *time.Location,
) (dayStartTime time.Time, err error) {

	// Unfortunately, the built-in 'ParseInLocation' Function works not as it
	// could be understood from its Name. So, we are implementing a true
	// 'ParseInLocation' Method here...

	var locationOffsetSec int

	// Get Location's Time Zone Offset.
	locationOffsetSec, err = GetLocationOffsetSec(location)
	if err != nil {
		return
	}

	// Parse the Time and correct it.
	dayStartTime, err = time.Parse(
		FormatDayTimeString,
		dayTimeString,
	)
	if err != nil {
		return
	}
	dayStartTime = dayStartTime.In(location)
	dayStartTime = dayStartTime.Add(time.Second * time.Duration(-locationOffsetSec))

	return
}

func SubHours(
	timeStart time.Time,
	timeSubtractedDeltaHours float64, // Must be > 0.
) time.Time {

	return timeStart.Add(-HoursToMicroseconds(timeSubtractedDeltaHours))
}

func ToDayStart(
	timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	delta = time.Nanosecond*time.Duration(timeStart.Nanosecond()) +
		time.Second*time.Duration(timeStart.Second()) +
		time.Minute*time.Duration(timeStart.Minute()) +
		time.Hour*time.Duration(timeStart.Hour())

	result = timeStart.Add(-delta)

	return result
}

func ToHourStart(
	timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	delta = time.Nanosecond*time.Duration(timeStart.Nanosecond()) +
		time.Second*time.Duration(timeStart.Second()) +
		time.Minute*time.Duration(timeStart.Minute())

	result = timeStart.Add(-delta)

	return result
}

func ToMinuteStart(
	timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	delta = time.Nanosecond*time.Duration(timeStart.Nanosecond()) +
		time.Second*time.Duration(timeStart.Second())

	result = timeStart.Add(-delta)

	return result
}

func ToMonthStart(
	timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	delta = time.Nanosecond*time.Duration(timeStart.Nanosecond()) +
		time.Second*time.Duration(timeStart.Second()) +
		time.Minute*time.Duration(timeStart.Minute()) +
		time.Hour*time.Duration(timeStart.Hour()) +
		time.Duration(timeStart.Day()-1)*(time.Hour*24)

	result = timeStart.Add(-delta)

	return result
}

func ToNextMonthStart(
	timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	timeMonthStart := ToMonthStart(timeStart)

	delta = (time.Hour * 24) * 33
	timeNextMonthForSure := timeMonthStart.Add(delta)

	result = ToMonthStart(timeNextMonthForSure)

	return result
}

func ToPreviousMonthStart(timeStart time.Time,
) time.Time {

	var delta time.Duration
	var result time.Time

	timeMonthStart := ToMonthStart(timeStart)

	delta = (time.Hour * 24) * (-1)
	timeNextMonthForSure := timeMonthStart.Add(delta)

	result = ToMonthStart(timeNextMonthForSure)

	return result
}
