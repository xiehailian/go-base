package rmqconn

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync/atomic"
	"testing"
	"time"
)

var (
	attempts int32
)

type channMock struct {
	с chan *amqp.Error
}

func (ch *channMock) Close() error {
	close(ch.с)
	return nil
}

func (ch *channMock) NotifyClose(c chan *amqp.Error) chan *amqp.Error {
	ch.с = c

	return nil
}

func (ch *channMock) GetChannel() *amqp.Channel {
	return new(amqp.Channel)
}

type connMock struct {
}

func (c *connMock) Close() error {
	return nil
}

func (c *connMock) Channel() (Channer, error) {
	if atomic.LoadInt32(&attempts) == 3 {
		return new(channMock), nil
	}
	return new(channMock), fmt.Errorf("get chan err")
}

func DialMock(url string) (Conner, error) {
	defer atomic.AddInt32(&attempts, 1)
	if atomic.LoadInt32(&attempts) <= 2 {
		return new(connMock), nil
	}
	return new(connMock), fmt.Errorf("dial err")
}

func TestConn(t *testing.T) {
	c, err := Open("amqp://search:search@192.168.56.2:5672/search_new", Dial)
	time.Sleep(time.Second * 10)

	fmt.Println(c, err)
}

func TestDial(t *testing.T) {
	c, err := Dial("amqp://search:search@192.168.56.2:5672/search_new")
	fmt.Println(c, err)
	if err != nil {
		t.Fail()
	}
}

func TestTypes(t *testing.T) {
	ch := &chann{c :new(amqp.Channel)}
	ch.GetChannel()
	c := make(chan *amqp.Error)
	ch.NotifyClose(c)
}


