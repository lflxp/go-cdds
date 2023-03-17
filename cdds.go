package cdds

/*
#cgo LDFLAGS: -L ${SRCDIR}/library/lib -lddsc
#cgo CFLAGS: -I ${SRCDIR}/library/include
#include "dds/dds.h"
*/
import "C"
import (
	"unsafe"
)

type DomainID C.dds_domainid_t
type Listener C.dds_listener_t

// type TopicDescriptor C.dds_topic_descriptor_t
type Attach C.dds_attach_t

// originally argument is void* arg
func CreateListener(arg unsafe.Pointer) *Listener {
	return (*Listener)(C.dds_create_listener(arg))
}
