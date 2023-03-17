package cdds

/*
#cgo LDFLAGS: -L ${SRCDIR}/library/lib -lddsc
#cgo CFLAGS: -I ${SRCDIR}/library/include
#include "dds/dds.h"
*/
import "C"

type Subscriber Participant

func (p *Subscriber) CreateReader(topic interface{}, elmSize uint32, qos *QoS, listener *Listener) (*Reader, error) {
	return (*Participant)(p).CreateReader(topic, elmSize, qos, listener)
}

func (p *Subscriber) Delete() error {
	return (*Participant)(p).Delete()
}
