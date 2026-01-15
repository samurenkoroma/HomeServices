// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
//
// This example program shows how to create a simple OPC UA server with data backed by a map.
// This allows you to easily create a server with a simple data model that can be updated from
// other parts of your application.  This example also shows how to monitor the data for changes
// and how to trigger change notifications to clients when the data changes.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/simonvetter/modbus"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

var (
	endpoint = flag.String("endpoint", "0.0.0.0", "OPC UA Endpoint URL")
	port     = flag.Int("port", 48400, "OPC UA Endpoint port")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	var opts []server.Option

	// Set your security options.
	opts = append(opts,
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
	)
	//opts = append(opts,
	//	server.EnableAuthMode(ua.UserTokenTypeAnonymous),
	//)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting host name %v", err)
	}

	opts = append(opts,
		server.EndPoint(*endpoint, *port),
		server.EndPoint("127.0.0.1", *port),
		server.EndPoint(hostname, *port),
	)
	s := server.New(opts...)

	greenhouseNS := server.NewMapNamespace(s, "greenhouse")
	greenhouseNS.Data["i0"] = .0
	greenhouseNS.Data["i1"] = .0
	greenhouseNS.Data["c0"] = false
	greenhouseNS.Data["c1"] = false
	greenhouseNS.Data["c2"] = false
	greenhouseNS.Data["c3"] = false

	client, _ := modbus.NewClient(&modbus.ClientConfiguration{URL: "tcp://192.168.22.129:502"})

	if err = client.Open(); err != nil {
		log.Fatal(err)
	}

	var inputRegs []uint16
	var coils []bool

	go func() {
		for {
			inputRegs, err = client.ReadRegisters(0, 6, modbus.INPUT_REGISTER)
			for i, reg := range inputRegs {
				greenhouseNS.SetValue(fmt.Sprintf("i%d", i), int16(reg))
			}

			coils, err = client.ReadCoils(0, 4)
			for i, reg := range coils {
				greenhouseNS.SetValue(fmt.Sprintf("c%d", i), reg)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			changed_key := <-greenhouseNS.ExternalNotification
			client.WriteCoil(0, greenhouseNS.GetValue(changed_key).(bool))
			log.Printf("%s changed to %v", changed_key, greenhouseNS.GetValue(changed_key))
		}
	}()

	root_ns, _ := s.Namespace(0)
	root_obj_node := root_ns.Objects()
	root_obj_node.AddRef(greenhouseNS.Objects(), id.HasComponent, true)

	// Start the server
	// Note that you can add namespaces before or after starting the server.
	if err := s.Start(context.Background()); err != nil {
		log.Fatalf("Error starting server, exiting: %s", err)
	}
	defer s.Close()

	// catch ctrl-c and gracefully shutdown the server.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	defer signal.Stop(sigch)
	log.Printf("Press CTRL-C to exit")

	<-sigch
	log.Printf("Shutting down the server...")
}
