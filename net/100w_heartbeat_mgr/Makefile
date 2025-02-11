ALL:
	go build

test:
	./main &
	while :; do echo "1.1.1.1" | nc 127.0.0.1 9999 -u -w0;sleep 1;done

clean:
	rm -f main
