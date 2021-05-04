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
	"publisher/publisher"

	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

// listTopicCmd represents the listTopic command
var listTopicCmd = &cobra.Command{
	Use:   "listTopic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTopic(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listTopicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listTopic(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	host, _ := cmd.Flags().GetString("host")
	if host == "" {
		fmt.Fprintln(os.Stderr, "Missing host for pubsub")

		return
	}

	client, err := publisher.ProvidePubSubClient(ctx, host)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stdout, "List of topics :")

	it := client.Topics(ctx)

	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Issue iterating over topics :", err)

			return
		}
		fmt.Println(topic.String())
	}
}