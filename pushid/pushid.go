/**
* This code comes from https://gist.github.com/themartorana/8c8b704432c8be1fed9a
* Only modification is change in main function to output just one pushid.
*
* Fancy ID generator that creates 20-character string identifiers with the following properties:
*
* 1. They're based on timestamp so that they sort *after* any existing ids.
* 2. They contain 72-bits of random data after the timestamp so that IDs won't collide with other clients' IDs.
* 3. They sort *lexicographically* (so the timestamp is converted to characters that will sort properly).
* 4. They're monotonically increasing. Even if you generate more than one in the same timestamp, the
* latter ones will sort after the former ones. We do this by using the previous random bits
* but "incrementing" them by 1 (only in the case of a timestamp collision).
 */
package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Timestamp of last push, used to prevent local collisions if you push twice in one ms.
var lastPushTime int64

// We generate 72-bits of randomness which get turned into 12 characters and appended to the
// timestamp to prevent collisions with other clients. We store the last characters we
// generated because in the event of a collision, we'll use those same characters except
// "incremented" by one.
var lastRandChars []int8

// Modeled after base64 web-safe chars, but ordered by ASCII.
const PUSH_CHARS string = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

func generatePushID() string {
	now := time.Now().UTC().UnixNano() / 1000000
	duplicateTime := now == lastPushTime
	lastPushTime = now

	timeStampChars := make([]string, 8, 8)
	for i := 7; i >= 0; i-- {
		pcIndex := int64(math.Mod(float64(now), 64.0))
		timeStampChars[i] = string(PUSH_CHARS[pcIndex])
		now = int64(math.Floor(float64(now) / 64.0))
	}

	if now != 0 {
		panic("We should have converted the entire timestamp.")
	}

	id := strings.Join(timeStampChars, "")

	if !duplicateTime {
		for i := 0; i < 12; i++ {
			lastRandChars[i] = int8(math.Floor(rand.Float64() * 64.0))
		}
	} else {
		var i int
		for i = 11; i >= 0 && lastRandChars[i] == 63; i-- {
			lastRandChars[i] = 0
		}

		lastRandChars[i]++
	}

	for i := 0; i < 12; i++ {
		id = fmt.Sprintf("%s%s", id, string(PUSH_CHARS[lastRandChars[i]]))
	}

	if len(id) != 20 {
		panic("Length should be 20")
	}

	return id
}

func main() {
	lastRandChars = make([]int8, 12, 12)
	fmt.Println(generatePushID())
}
