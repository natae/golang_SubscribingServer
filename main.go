package main

import (
	"fmt"
	subscription "subscribing_server/subscription_origin"
	"time"
)

func main() {
	// Subscribe to some feeds, and create a merged update stream.
	merged := subscription.Merge(
		subscription.Subscribe(subscription.Fetch("blog.golang.org")),
		subscription.Subscribe(subscription.Fetch("googleblog.blogspot.com")),
		subscription.Subscribe(subscription.Fetch("googledevelopers.blogspot.com")),
	)

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	panic("show me the stacks")
}
