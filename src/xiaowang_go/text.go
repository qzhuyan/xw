package xiaowang_go

import (
	"math/rand"
	"strings"
)

type text string //text can have any chars including blankspace
type word string // word has no blank space

func (s text) Transform(seed int64) string {
	ws := strings.Split(string(s), " ")
	for i := 0; i < len(ws); i++ {
		ws[i] = word(ws[i]).Transform(seed)
	}
	return strings.Join(ws, " ")
}

func (w word) Transform(seed int64) string {
	return RandStringBytes(len(w))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytes(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
