package mydate

import (
	"time"

	"github.com/jalaali/go-jalaali"
)

type times struct {
	year  int
	month int
	day   int
}

var N times = jalaali.ToJalaali(time.Now().Date())
