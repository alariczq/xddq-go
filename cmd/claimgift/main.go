package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: claimgift <username> <password>")
		os.Exit(1)
	}

	username := os.Args[1]
	password := os.Args[2]

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	c := client.NewClient(username, password)

	c.WithOnLogined(func(ctx client.Context) error {
		log.Println("login success")
		return nil
	})

	c.OnSendGiftSyncMsg(func(ctx client.Context, msg *pb.SendGiftSyncMsg) error {
		for _, data := range msg.GetData() {
			log.Println("data", data)
			ctx.Client().SendGetSendGiftRewardReq(
				ctx,
				&pb.GetSendGiftRewardReq{
					Id: data.Id,
				},
			)
		}
		return nil
	})

	c.Run()
}
