package webhook

import (
	"fmt"
	"log"

	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/gtuk/discordwebhook"
)

func Sendreport(report model.Report) {

	var username = "BotUser"
	var content = "This is a test message"
	var url = "https://discord.com/api/webhooks/1077711746413899946/xqFikJU4fJTw89S7rLi9qVnQjFNkg7eDeJMETS2lnQotZ55wo4CKPbTV_cJvmhrs5OQZ"
	reporturl := fmt.Sprintf("[View report](http://%s/report/%s)", config.LoadedConfig.PublicURL, report.UUID)

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
		Embeds: &[]discordwebhook.Embed{
			{
				Title:       &report.Title,
				Description: &reporturl,
				Fields: &[]discordwebhook.Field{
					{
						Value: &reporturl,
					},
				},
			},
		},
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

}
