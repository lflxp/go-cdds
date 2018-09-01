package cdds

/*
#cgo CFLAGS: -I/usr/local/include/ddsc
#cgo LDFLAGS: -lddsc
#include "ddsc/dds.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

type Writer struct {
	Entity
}

func (w Writer) Write(data unsafe.Pointer) error {
	ret := C.dds_write(w.GetEntity(), data)
	ErrorCheck(ret, C.DDS_CHECK_REPORT|C.DDS_CHECK_EXIT, "tmp where")
	if ret < 0 {
		return CddsErrorType(ret)
	}
	return nil
}

func (w Writer) WriteTimeStamp(data unsafe.Pointer, ts Time) error {
	ret := C.dds_write_ts(w.GetEntity(), data, C.dds_time_t(ts))
	if ret < 0 {
		return CddsErrorType(ret)
	}
	return nil
}

func (w Writer) WriteDispose(data unsafe.Pointer) error {
	ret := C.dds_writedispose(w.GetEntity(), data)
	if ret < 0 {
		return CddsErrorType(ret)
	}
	return nil
}

func (w *Writer) SearchTopic(d time.Duration) error {
	// need mutex lock?
	// WARN: this cause error
	w.SetEnabledStatus(PublicationMatched)
	for {
		status, err := w.GetStatusChanges()
		if err != nil {
			return err
		}
		if status == PublicationMatched {
			break
		}
		time.Sleep(d)
	}
	return nil
}

func (w *Writer) delete() error {
	if w.qos != nil {
		w.qos.delete()
	}
	return nil
}
