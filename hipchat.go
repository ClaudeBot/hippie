package main

import (
	"fmt"
	"github.com/tbruyelle/hipchat-go/hipchat"
	"net/url"
	"strings"
)

type Message struct {
	Date    string `json:"date,omitempty"`
	Message string `json:"message"`
}

type Item struct {
	Message Message `json:"message"`
}

type RoomMessage struct {
	Event string `json:"event,omitempty"`
	Item  Item   `json:"item"`
}

func getRoomID(rn string) (int, error) {
	rooms, _, err := client.Room.List()
	if err != nil {
		return 0, err
	}

	for _, room := range rooms.Items {
		if strings.EqualFold(room.Name, rn) {
			return room.ID, nil
		}
	}

	return 0, fmt.Errorf("room '%s' could not be found; you may not have the necessary permissions to view it", rn)
}

func webhookGC(rid int, k string) error {
	webhooks, _, err := client.Room.ListWebhooks(rid, new(hipchat.ListWebhooksOptions))
	if err != nil {
		return err
	}

	for _, webhook := range webhooks.Webhooks {
		// TODO: replace `Name` with `Key` in the future
		if strings.EqualFold(webhook.Name, k) {
			if _, err := client.Room.DeleteWebhook(rid, webhook.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

func webhookRegister(rn string, k string, u *url.URL) error {
	if client == nil {
		return fmt.Errorf("HipChat client is not initialised")
	}

	rid, err := getRoomID(rn)
	if err != nil {
		return err
	}
	roomID = rid

	err = webhookGC(roomID, k)
	if err != nil {
		return err
	}

	req := hipchat.CreateWebhookRequest{
		Name:  k,
		Event: "room_message",
		URL:   u.String(),
	}

	wh, _, err := client.Room.CreateWebhook(roomID, &req)
	if err != nil {
		return err
	}
	webhookID = wh.ID

	return nil
}
