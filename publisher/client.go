package publisher

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func ProvidePubSubClient(ctx context.Context, host string, projectID string) (*pubsub.Client, error) {
	os.Setenv("PUBSUB_EMULATOR_HOST", host)

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile("{}"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot start pubsub client :", err)

		return nil, err
	}

	return client, nil
}
