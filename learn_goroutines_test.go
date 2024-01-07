package learn

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestBufferedChannel(t *testing.T) {
	// Membuat buffered channel dengan kapasitas 3
	bufferedCh := make(chan int, 3)

	// Mengirim data ke buffered channel
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	// Uncomment baris di bawah untuk melihat contoh buffer penuh dan blok
	// bufferedCh <- 4 // Blokir, karena buffer sudah penuh

	// Menunggu sebentar
	time.Sleep(2 * time.Second)

	// Menerima data dari buffered channel
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)

	// Uncomment baris di bawah untuk melihat contoh buffer kosong dan blok
	// fmt.Println(<-bufferedCh) // Blokir, karena buffer kosong
	defer close(bufferedCh)
}
func TestCreateGoroutines(t *testing.T) {
	channel := make(chan time.Time)
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Defer closing the file until the function exits.
	defer close(channel)
	// Create a bufio.Writer to efficiently write to the file.
	writer := bufio.NewWriter(file)
	go writeData(writer, channel) // async
	getWriteTime := <-channel     // blocking

	log.Println(getWriteTime)
	log.Println("File written successfully.")
}

func writeData(writer *bufio.Writer, channel chan time.Time) {
	time.Sleep(2 * time.Second)
	log.Println("processing")

	// Write content to the file.
	content := "Hello, this is some text that will be written to the file."
	_, err := writer.WriteString(content)
	if err != nil {
		fmt.Println("Error writer:", err)
	}
	channel <- time.Now() // blocking
	// Flush the bufio.Writer to ensure that all buffered data is written to the file.
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
	log.Println("finish")
}
func TestRangeChannel(t *testing.T) {
	// Membuat channel
	channel := make(chan int)
	go simulateGetAPI(channel)

	for data := range channel {
		log.Println(data)
	}
	log.Println("END OF PROGRAM")
}
func TestSelectChannel(t *testing.T) {
	chA := make(chan int)
	chB := make(chan int)

	go selectSimulateGetAPI(chA)
	go selectSimulateGetAPI(chB)

	for {
		select {
		case data, ok := <-chA:
			if ok {
				log.Println("Data from CH A : ", data)
			} else {
				close(chA)
				log.Println("Channel A is closed")
			}
		case data, ok := <-chB:
			if ok {
				log.Println("Data from CH B : ", data)
			} else {
				close(chB)
				log.Println("Channel B is closed")
			}
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout! No data received in time.")
			return
		}
	}
}
func selectSimulateGetAPI(channel chan int) {
	for i := 1; i < 10; i++ {
		channel <- i
	}
}
func simulateGetAPI(channel chan int) {
	for i := 1; i < 10; i++ {
		channel <- i
	}
	defer close(channel) // sinyal untuk memberi tahu bahwa tidak ada lagi data yang akan dikirim
}
