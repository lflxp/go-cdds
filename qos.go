package cdds

/*
#cgo LDFLAGS: -L ${SRCDIR}/library/lib -lddsc
#cgo CFLAGS: -I ${SRCDIR}/library/include
#include "dds/dds.h"
*/
import "C"
import (
	"time"
	"unsafe"
)

type QoS C.dds_qos_t

func CreateQoS() *QoS {
	return (*QoS)(C.dds_create_qos())
}

func (qos *QoS) SetReliability(rel Reliability, n time.Duration) {
	C.dds_qset_reliability((*C.dds_qos_t)(qos), C.dds_reliability_kind_t(rel), C.int64_t(int64(n)))
}

func (qos *QoS) SetWriterDataLifecycle(autoDispose bool) {
	C.dds_qset_writer_data_lifecycle((*C.dds_qos_t)(qos), C.bool(autoDispose))
}

func (qos *QoS) SetPartition(num int, partitions *string) {
	C.dds_qset_partition((*C.dds_qos_t)(qos), C.uint32_t(num), (**C.char)(unsafe.Pointer(partitions)))

}

func (qos *QoS) delete() {
	C.dds_delete_qos((*C.dds_qos_t)(qos))
}
