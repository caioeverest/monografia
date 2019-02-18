package mqtt

import (
	"fmt"
	"github.com/caioever/monografia/backend/infra/config"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/url"
	"time"
)

type Context struct {
	Client	mqtt.Client
}

func (context Context) Send(topic string, message interface{}) error{
	context.Client.Publish(topic, 0, false, message)
	return nil
}

func (context Context) Listen(topic string) {
	client := context.Client
	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	}); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

func (context Context) ConnectionTest() bool {
	return !context.Client.IsConnected()
}

func ContextBuilder(conf *config.Mqtt) *Context {
	context := &Context{}
	uri, err := url.Parse(buildUri(conf))
	if err != nil {
		log.Print(err)
	}
	opts := createClientOptions(conf.ClientName, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	context.Client = client
	return context
}

func buildUri(conf *config.Mqtt) string {
	return fmt.Sprintf("mqtts://%s:%s@%s:%s/", conf.User, conf.Pass, conf.Host, conf.Port)
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}
