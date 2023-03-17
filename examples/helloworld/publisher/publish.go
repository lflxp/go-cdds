package main

/*
#cgo LDFLAGS: -L ../../../library/lib -lddsc ${SRCDIR}/HelloWorldData.o
#cgo CFLAGS: -I ../../../library/include
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

func main() {
	var msg C.HelloWorldData_Msg

	participant, err := cdds.CreateParticipant(cdds.DomainDefault, nil, nil)
	defer participant.Delete()
	if err != nil {
		panic(err)
	}

	_, err = participant.CreateTopic(unsafe.Pointer(&C.HelloWorldData_Msg_desc), "HelloWorldData_Msg", nil, nil)
	if err != nil {
		panic(err)
	}
	writer, err := participant.CreateWriter("HelloWorldData_Msg", nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("=== [Publisher] Waiting for a reader to be discovered ...")
	C.fflush(C.stdout)

	// err = writer.SearchTopic(time.Millisecond * 20)
	// if err != nil {
	// 	panic(err)
	// }

	err = writer.SetEnabledStatus(cdds.PublicationMatched)
	if err != nil {
		panic(err)
	}
	var status cdds.CommunicationStatus
	for status != cdds.PublicationMatched {
		// for status > 0 && cdds.PublicationMatched > 0 {
		status, err = writer.GetStatusChanges()
		if err != nil {
			panic(err)
		}
		// fmt.Println("status: rc", status)
		cdds.SleepFor(time.Millisecond * 20)
	}

	for x := 0; x < 10000; x++ {
		msg.userID = (C.int)(x)

		jsonStr := "{\"Name\":\"cyclone\", \"Age\":22}"

		msg.message = C.CString(jsonStr)
		// msg.message = (*C.char)(unsafe.Pointer(ms))

		fmt.Println("=== [Publisher] Writing : ")
		fmt.Printf("Message (%d, %s)\n", msg.userID, C.GoString(msg.message))
		cdds.SleepFor(time.Second)
		writer.Write(unsafe.Pointer(&msg))
	}

}
