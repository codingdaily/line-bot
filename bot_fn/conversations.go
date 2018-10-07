package bot_fn

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewJobOffer(jobListing string) *linebot.TemplateMessage {
	apply := linebot.NewMessageAction("Ya", "apply")
	skip := linebot.NewMessageAction("Tidak", "skip")
	tpl := linebot.NewConfirmTemplate("Apa Kakak Tertarik?", skip, apply)
	return linebot.NewTemplateMessage(jobListing, tpl)
}
