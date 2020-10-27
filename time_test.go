// time_test.go.

// +build test

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
	"testing"
	"time"

	"github.com/vault-thirteen/tester"
)

const SecondsPerHour = 60 * 60

func Test_IsEmpty(t *testing.T) {

	var aTest *tester.Test
	var err error
	var x time.Time

	aTest = tester.New(t)

	// Test #1.
	x, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), false)

	// Test #2.
	x = time.Time{}
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(IsEmpty(x), true)
}

func Test_NewTimeStringRFC3339(t *testing.T) {

	var aTest *tester.Test
	var result string
	var resultExpected string

	aTest = tester.New(t)
	resultExpected = "2019-06-11T16:44:05Z"

	// Test.
	result = NewTimeStringRFC3339(
		2019,
		6,
		11,
		16,
		44,
		5,
	)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Minimum(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	aTest = tester.New(t)

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time1
	aTest.MustBeNoError(err)
	result = Minimum(time1, time1)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_Maximum(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result time.Time
	var resultExpected time.Time
	var time1 time.Time
	var time2 time.Time

	aTest = tester.New(t)

	time1, err = time.Parse(time.RFC3339, "2019-06-24T15:02:55Z")
	aTest.MustBeNoError(err)
	time2, err = time.Parse(time.RFC3339, "2020-06-24T15:02:55Z")
	aTest.MustBeNoError(err)

	// Test #1.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time1, time2)
	aTest.MustBeEqual(result, resultExpected)

	// Test #2.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time1)
	aTest.MustBeEqual(result, resultExpected)

	// Test #3.
	resultExpected = time2
	aTest.MustBeNoError(err)
	result = Maximum(time2, time2)
	aTest.MustBeEqual(result, resultExpected)
}

func Test_IntervalDurationHours(t *testing.T) {

	var aTest *tester.Test
	var dur float64
	var t1 time.Time
	var t2 time.Time

	aTest = tester.New(t)

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 15)
	dur = IntervalDurationHours(t1, t2)
	aTest.MustBeEqual(dur, float64(0.25))
}

func Test_HoursToMicroseconds(t *testing.T) {

	var aTest *tester.Test
	var hours float64

	aTest = tester.New(t)

	hours = 0.25 // 1/4 Hour = 15 Minutes = 900 Seconds = 900 000 000 micro Seconds.
	dur := HoursToMicroseconds(hours)
	aTest.MustBeEqual(dur, time.Minute*15)
}

func Test_AddHours(t *testing.T) {

	var aTest *tester.Test
	var t1 time.Time
	var t2 time.Time

	aTest = tester.New(t)

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 15)
	t1PlusDelta := AddHours(t1, 0.25)
	aTest.MustBeEqual(t1PlusDelta, t2)
}

func Test_SubHours(t *testing.T) {

	var aTest *tester.Test
	var t1 time.Time
	var t2 time.Time

	aTest = tester.New(t)

	t1 = time.Now()
	t2 = t1.Add(time.Minute * 15)
	t2MinusT1 := SubHours(t2, 0.25)
	aTest.MustBeEqual(t2MinusT1, t1)
}

func Test_ToMinuteStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:48:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:48:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:48:00.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:48:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMinuteStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToHourStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:00:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T12:00:00.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T12:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToHourStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToDayStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T00:00:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339Nano, "2019-06-11T00:00:00.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-11T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToDayStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToMonthStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-06-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339, "2020-02-29T01:02:03Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-02-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #4.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T01:02:03.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-02-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToNextMonthStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-07-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-07-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339, "2020-02-29T23:59:59Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-03-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #4.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T23:59:59.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-03-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToNextMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

func Test_ToPreviousMonthStart(t *testing.T) {

	var aTest *tester.Test
	var err error
	var t1 time.Time
	var t1Processed time.Time
	var t2 time.Time

	aTest = tester.New(t)

	// Test #1.
	t1, err = time.Parse(time.RFC3339, "2019-06-11T12:48:51Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-05-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #2.
	t1, err = time.Parse(time.RFC3339, "2019-06-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2019-05-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #3.
	t1, err = time.Parse(time.RFC3339, "2020-02-29T23:59:59Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)

	// Test #4.
	t1, err = time.Parse(time.RFC3339Nano, "2020-02-29T23:59:59.001Z")
	aTest.MustBeNoError(err)
	t2, err = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	aTest.MustBeNoError(err)
	t1Processed = ToPreviousMonthStart(t1)
	aTest.MustBeEqual(t1Processed, t2)
}

// Attention! This Function has been created due to the Dumbness of the
// Developers of the Go Language. The Results of this Tests may change in the
// later Versions of the Go Programming Language. This Test has been confirmed
// to work for the Go Language Version '1.15.3'. The Author of this Test does
// not guarantee the Correctness of this Test in Go Language of other Versions.
func Test_GetLocationOffsetSec(t *testing.T) {

	var err error
	var location *time.Location
	var locationOffsetSec int
	var locationOffsetSecExpected int
	var tst *tester.Test

	tst = tester.New(t)

	// Test #1. Positive Time Zone Offset.
	locationOffsetSecExpected = 5 * SecondsPerHour // 5 Earth-Planet Hours.
	location = time.FixedZone("Somewhere on Earth [EAST]", locationOffsetSecExpected)
	locationOffsetSec, err = GetLocationOffsetSec(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetSec, locationOffsetSecExpected)

	// Test #2. Negative Time Zone Offset.
	locationOffsetSecExpected = -4 * SecondsPerHour // 4 Earth-Planet Hours.
	location = time.FixedZone("Somewhere on Earth [WEST]", locationOffsetSecExpected)
	locationOffsetSec, err = GetLocationOffsetSec(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetSec, locationOffsetSecExpected)

	// Test #3. Zero Time Zone Offset.
	locationOffsetSecExpected = 0
	location = time.FixedZone("Somewhere on Earth [Greenwich]", locationOffsetSecExpected)
	locationOffsetSec, err = GetLocationOffsetSec(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetSec, locationOffsetSecExpected)
}

// Attention! This Function has been created due to the Dumbness of the
// Developers of the Go Language. The Results of this Tests may change in the
// later Versions of the Go Programming Language. This Test has been confirmed
// to work for the Go Language Version '1.15.3'. The Author of this Test does
// not guarantee the Correctness of this Test in Go Language of other Versions.
func Test_GetLocationOffsetHours(t *testing.T) {

	var err error
	var location *time.Location
	var locationOffsetHrs int
	var locationOffsetHrsExpected int
	var tst *tester.Test

	tst = tester.New(t)

	// Test #1. Positive Time Zone Offset.
	locationOffsetHrsExpected = 5 // 5 Earth-Planet Hours.
	location = time.FixedZone(
		"Somewhere on Earth [EAST]",
		locationOffsetHrsExpected*SecondsPerHour,
	)
	locationOffsetHrs, err = GetLocationOffsetHours(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetHrs, locationOffsetHrsExpected)

	// Test #2. Negative Time Zone Offset.
	locationOffsetHrsExpected = -4 // 4 Earth-Planet Hours.
	location = time.FixedZone(
		"Somewhere on Earth [WEST]",
		locationOffsetHrsExpected*SecondsPerHour,
	)
	locationOffsetHrs, err = GetLocationOffsetHours(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetHrs, locationOffsetHrsExpected)

	// Test #3. Zero Time Zone Offset.
	locationOffsetHrsExpected = 0
	location = time.FixedZone(
		"Somewhere on Earth [Greenwich]",
		locationOffsetHrsExpected*SecondsPerHour,
	)
	locationOffsetHrs, err = GetLocationOffsetHours(location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(locationOffsetHrs, locationOffsetHrsExpected)
}

// Attention! This Function has been created due to the Dumbness of the
// Developers of the Go Language. The Results of this Tests may change in the
// later Versions of the Go Programming Language. This Test has been confirmed
// to work for the Go Language Version '1.15.3'. The Author of this Test does
// not guarantee the Correctness of this Test in Go Language of other Versions.
func Test_ParseDayTimeStringInLocation(t *testing.T) {

	var dayStartTime time.Time
	var dayStartTimeExpected time.Time
	var err error
	var location *time.Location
	var tst *tester.Test

	tst = tester.New(t)

	// Test #1. Positive Time Zone Offset.
	location = time.FixedZone("+0400", 4*SecondsPerHour)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00+04:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.

	// Test #2. Negative Time Zone Offset.
	location = time.FixedZone("-0800", -8*SecondsPerHour)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00-08:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.

	// Test #3. Zero Time Zone Offset.
	location = time.FixedZone("", 0)
	dayStartTimeExpected, err = time.Parse(time.RFC3339, "2019-09-01T00:00:00-00:00")
	tst.MustBeNoError(err)
	dayStartTime, err = ParseDayTimeStringInLocation("2019-09-01", location)
	tst.MustBeNoError(err)
	tst.MustBeEqual(
		dayStartTime.Format(time.RFC3339),
		dayStartTimeExpected.Format(time.RFC3339),
	)
	fmt.Println(dayStartTime.Format(time.RFC3339)) // Report.
}
