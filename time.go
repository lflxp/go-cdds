package cdds

/*
#cgo LDFLAGS: -L ${SRCDIR}/library/lib -lddsc
#cgo CFLAGS: -I ${SRCDIR}/library/include
#include "dds/dds.h"
*/
import "C"
import "time"

type Time C.dds_time_t
type Duration C.dds_duration_t

func SleepFor(n time.Duration) {
	C.dds_sleepfor(C.int64_t(int64(n)))
}

/*
func SleepUntil(n Time) {

}
*/

func DdsTime() Time {
	return Time(C.dds_time_t(C.dds_time()))
}
