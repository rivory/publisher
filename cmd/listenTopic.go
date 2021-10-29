/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

const SUBNAME = "test"

// listenTopicCmd represents the listenTopic command
var listenTopicCmd = &cobra.Command{
	Use:   "listenTopic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listener(cmd)
	},
}

func init() {
	rootCmd.AddCommand(listenTopicCmd)

	listenTopicCmd.Flags().String("topic", "", "Topic ID you want to create")
}

func subscriber(topic *pubsub.Topic, projectID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("pubsub.NewClient: %v", err)
		return err
	}
	defer client.Close()

	exist, err := client.Subscription(SUBNAME).Exists(ctx)
	if err != nil {
		fmt.Printf("SubscriptionExist: %v", err)
		return err
	}
	if exist {
		return nil
	}

	sub, err := client.CreateSubscription(ctx, SUBNAME, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		fmt.Printf("CreateSubscription: %v", err)
		return err
	}
	fmt.Printf("Created subscription: %v\n", sub)

	return nil

}

func listener(cmd *cobra.Command) {
	ctx := context.Background()

	name, _ := cmd.Flags().GetString("topic")
	if name == "" {
		fmt.Fprintln(os.Stderr, "Missing topic name for pubsub")

		return
	}
	project, _ := cmd.Flags().GetString("project")
	if project == "" {
		fmt.Fprintln(os.Stderr, "Missing project for pubsub")

		return
	}

	client, err := InitPubsubClient(ctx, cmd)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	topic := client.Topic(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, " GetTopic: %v\n", err)

		return
	}

	err = subscriber(topic, project)
	if err != nil {
		spew.Dump(err)
		return
	}

	sub := client.Subscription(SUBNAME)

	var mu sync.Mutex
	received := 0
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		// fmt.Printf("Got message: %q\n", string(msg.Data))
		spew.Dump("event name :", msg.Attributes)
		spew.Dump(string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		fmt.Printf("Receive: %v", err)
	}
}
