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

	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		publish(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().String("topic", "", "Topic ID you want to create")

	publishCmd.Flags().String("message", "", "Message content to publish")

	publishCmd.Flags().StringToString("attribute", nil, "Attribute to publish with the message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func publish(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	client, err := InitPubsubClient(ctx, cmd)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	name, _ := cmd.Flags().GetString("topic")
	if name == "" {
		fmt.Fprintln(os.Stderr, "Missing topic name for pubsub")

		return
	}

	message, _ := cmd.Flags().GetString("message")
	if message == "" {
		fmt.Fprintln(os.Stderr, "Missing message content")

		return
	}

	attribute, _ := cmd.Flags().GetStringToString("attribute")

	topic := client.Topic(name)
	if topic == nil {
		fmt.Fprintln(os.Stderr, "Cannot find topic")

		return
	}

	event := &pubsub.Message{
		Data: []byte(message),
	}
	if attribute != nil {
		event.Attributes = attribute
	}

	res := topic.Publish(ctx, event)
	fmt.Fprintf(os.Stdout, "Message published: %v\n", res)

	topic.Stop()
}
