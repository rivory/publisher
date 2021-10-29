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

	"github.com/spf13/cobra"
)

// createTopicCmd represents the createTopic command
var createTopicCmd = &cobra.Command{
	Use:   "createTopic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		createTopic(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(createTopicCmd)

	createTopicCmd.Flags().String("topic", "", "Topic ID you want to create")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTopicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTopicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createTopic(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	client, err := InitPubsubClient(ctx, cmd)
	if err != nil {
		panic(err)
	}

	name, _ := cmd.Flags().GetString("topic")
	if name == "" {
		fmt.Fprintln(os.Stderr, "Missing topic name for pubsub")

		return
	}

	t, err := client.CreateTopic(ctx, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, " CreateTopic: %v\n", err)

		return
	}
	fmt.Fprintf(os.Stdout, "Topic created: %v\n", t)

	return
}
