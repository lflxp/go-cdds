pub: install
	cd publisher && go run publish.go

install: clean
	gcc -E HelloWorldData.c -o h.i -I ../../library/include
	gcc -S h.i -o h.s
	gcc -c h.s -o HelloWorldData.o
	cp *.o publisher/
	cp *.o subscriber/

clean:
	rm -f *.i *.s *.o
	rm -f publisher/*.o
	rm -f subscriber/*.o