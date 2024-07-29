package telegram

import "github.com/jocRednow/suggest-to-read-bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {}
