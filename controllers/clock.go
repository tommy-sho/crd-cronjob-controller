package controllers

import "time"

type realClock struct{}

func (_ realClock) Now() time.Time { return time.Now() }

// testように時間を入れ替えるためのinterface.Testの時はフェイクの時間を入れることで時刻を固定する
type Clock interface {
	Now() time.Time
}
