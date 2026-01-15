package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fatih/color"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message from topic: %s\n: %s\n", msg.Topic(), msg.Payload())
}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	c := color.New(color.FgHiGreen)
	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Create new metrics and register them using the custom registry.
	m := NewMetrics(reg)
	var broker = "lab.raspi"
	var port = 1883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("polevod_client[go]")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	//opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		c.Add(color.FgHiGreen).Printf("Received message from topic: \"%s\"\n", msg.Topic())
		c.Add(color.FgHiYellow).Printf("%s\n", msg.Payload())
		var device Device
		json.Unmarshal(msg.Payload(), &device)
		for _, v := range device.Sensors {

			m.sensors.With(prometheus.Labels{"pin": fmt.Sprintf("%s-%d", device.Device, v.Pin), "type": v.Type}).Set(v.Value)
		}
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go sub(client)
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.HandleFunc("/api", apiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func apiHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("{\"status\": \"ok\"}"))
}

func sub(client mqtt.Client) {
	topic := "homeapp/meteo"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s \n", topic)
}

type Device struct {
	Device  string   `json:"device"`
	Sensors []Sensor `json:"sensors"`
}
type Sensor struct {
	Pin   int     `json:"p"`
	Type  string  `json:"t"`
	Value float64 `json:"v"`
}
type metrics struct {
	sensors *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		sensors: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "sensors",
				Help: "Текущие показания",
			},
			[]string{"type", "pin"},
		)}
	reg.MustRegister(m.sensors)
	return m
}
