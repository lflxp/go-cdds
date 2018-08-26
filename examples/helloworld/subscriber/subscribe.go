package main

/*
#cgo CFLAGS: -I/usr/local/include/ddsc
#cgo LDFLAGS: -lddsc ${SRCDIR}/../HelloWorldData.o
#include "ddsc/dds.h"
#include "../HelloWorldData.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	cdds "github.com/ami-GS/go-cdds"
)

const MAX_SAMPLES = 1

func main() {
	var msg *C.HelloWorldData_Msg
	participant, err := cdds.CreateParticipant(cdds.DomainDefault, nil, nil)
	defer participant.Delete()
	if err != nil {
		panic(err)
	}

	_, err = participant.CreateTopic(unsafe.Pointer(&C.HelloWorldData_Msg_desc), "HelloWorldData_Msg", nil, nil)
	if err != nil {
		panic(err)
	}
	qos := cdds.CreateQoS()
	qos.SetReliability(cdds.Reliable, time.Second*10)
	reader, err := participant.CreateReader("HelloWorldData_Msg", uint32(unsafe.Sizeof(*msg)), qos, nil)
	if err != nil {
		panic(err)
	}
	qos.Delete()
	fmt.Println("=== [Subscriber] Waiting for sample ...")

	// finCh := make(chan error)
	// go reader.ReadWithCallback(MAX_SAMPLES, MAX_SAMPLES, &finCh, func(samples *cdds.Array) {
	// 	msg = (*C.HelloWorldData_Msg)(samples.At(0))
	// 	fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))
	// })
	// err <-finCh
	// if err != nil {
	// panic(err)
	// }

	sample, err := reader.BlockAllocRead(MAX_SAMPLES, MAX_SAMPLES)
	if err != nil {
		panic(err)
	}
	msg = (*C.HelloWorldData_Msg)(sample.At(0))
	fmt.Print("=== [Subscriber] Received : ")
	fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))

	// for {
	// 	samples := reader.AllocRead(MAX_SAMPLES, MAX_SAMPLES)
	// 	if samples.IsValidAt(0) {
	// 		/* Print Message. */
	// 		msg = (*C.HelloWorldData_Msg)(samples.At(0))
	// 		fmt.Print("=== [Subscriber] Received : ")
	// 		fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))
	// 		break
	// 	}
	// 	cdds.SleepFor(time.Millisecond * 20)
	// }
}
