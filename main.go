import (
	"bufio",
	"net",
	"time",
	"io",
	"strconv",
	"log"
);

func handle(net.Conn conn) {
	conn.SetReadDeadline(time.Now().Add(time.ParseDuration("10s")));
	reader := bufio.Reader(conn);
	requestline, err := reader.ReadString('\r\n');
	if (err!=nil) {
		log.Fatal(err);
		return;
	}
	parts = strings.Split(requestline, " ");
	inputlen, err = strconv.ParseInt(parts[2]);
	if (err!=nil) {
		log.Fatal(err);
		return;
	}
	data = make([]byte, inputlen);
	_, err = io.ReadFull(reader, data);
	if (err!=nil) {
		log.Fatal(err);
		return;
	}
	// this is where the request handling goes when i get to it
}

func startServer() {
	// this is where i put listening logic
	// go handle(conn); when it gets a connection
};

func main() {
	startServer();
}
