package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)
var (
	ServiceMq *Mq
	ch *amqp.Channel
	conn *amqp.Connection
)
type Mq struct {
	action string
	option *Options
}
type Options struct {
	Host     string
	Port     string
	Username string
	Pwd      string
	Vh       string
	// Queue    string
}
type newOption func(options *Options)
func init() {
	ServiceMq = &Mq{}
}

func (mq *Mq) New (action string, opts ...newOption) *Mq {
	option := &Options{
		Host:     "39.99.174.39",
		Port:     "5672",
		Username: "admin",
		Pwd:      "pwd",
		Vh:       "todsp",
	}
	for _, o := range opts {
		o(option)
	}
	mq.action = action
	mq.option = option
	return mq
}

func (mq *Mq) Go ()  {
	var err error
	// defer conn.Close()
	// defer ch.Close()
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", mq.option.Username, mq.option.Pwd, mq.option.Host, mq.option.Port, mq.option.Vh)
	conn, err = amqp.Dial(RabbitUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("-----------------[ RabbitMq Start ]-----------------------")
	fmt.Println("-----------------["+mq.action+":"+time.Now().String()+"]------------------------------------")
	fmt.Println(ServiceMq.option)

}


func SendMsg(queue string, msg []byte) {
	if conn == nil || conn.IsClosed() || ch == nil {
		ServiceMq.Go()
	}
	err := ch.Publish(
		"",    //exchange
		queue, //routing key
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //Msg set as persistent
			ContentType:  "text/plain",
			Body:         msg,
		})
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%v :Error: %s: %s\n", ServiceMq.action, msg, err)
		panic(err)
	}
}
