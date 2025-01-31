package main

/*
#cgo CFLAGS: -I ../../../library/include
#cgo LDFLAGS: -L ../../../library/lib -lddsc ${SRCDIR}/HelloWorldData.o
#include "dds/dds.h"
#include "../HelloWorldData.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"

	cdds "github.com/ami-GS/go-cdds"
)

const MAX_SAMPLES = 3

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
	fmt.Println("=== [Subscriber] Waiting for sample ...")

	// 1. callback
	// finCh := make(chan error)
	// go reader.ReadWithCallback(MAX_SAMPLES, MAX_SAMPLES, false, &finCh, func(samples *cdds.Array) {
	// 	msg = (*C.HelloWorldData_Msg)(samples.At(0))
	// 	fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))
	// })
	// err = <-finCh
	// if err != nil {
	// 	panic(err)
	// }

	// 2. blocking read
	// sample, err := reader.BlockAllocRead(MAX_SAMPLES, MAX_SAMPLES, false)
	// if err != nil {
	// 	panic(err)
	// }
	// msg = (*C.HelloWorldData_Msg)(sample.At(0))
	// fmt.Print("=== [Subscriber] Received : ")
	// fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))

	// 3. basic loop
	//samples := reader.Alloc(MAX_SAMPLES)
	var num int
	var samples *cdds.Array
	for {
		// WARN: Just using AllocRead() use much heap space
		if samples == nil {
			samples, num, err = reader.AllocRead(MAX_SAMPLES, MAX_SAMPLES, false)
			fmt.Printf("nil sample num is %d\n", num)
		} else {
			num, err = reader.ReadWithBuff(samples, true)
			fmt.Printf("sample num is %d\n", num)
		}

		if err != nil {
			panic(err)
		}
		for i := 0; ; {
			if samples.IsValidAt(i) {
				msg = (*C.HelloWorldData_Msg)(samples.At(i))
				fmt.Print("=== [Subscriber] Received : ")
				fmt.Printf("Message %d:(%d, %s)\n", i, msg.userID, C.GoString(msg.message))
				i++
				if i >= num {
					fmt.Printf("i %d num %d\n", i, num)
					goto END
				}
			} else {
				break
			}

		}
	END:
		cdds.SleepFor(time.Millisecond * 500)
		// cdds.SleepFor(time.Second)
	}
	// END:
}
