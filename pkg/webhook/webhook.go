package webhook

import (
	"fmt"

	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

func SendReport(report model.Report) (string, error) {
	if config.LoadedConfig.DiscordWebhook == "" {
		return "", nil
	}

	client, err := webhook.NewWithURL(config.LoadedConfig.DiscordWebhook)
	if err != nil {
		return "", err
	}
	evidenceAvailable := "Evidence is not available"

	var priority string

	switch report.Priority {
	case 1:
		priority = "Low"
	case 2:
		priority = "Medium"
	case 3:
		priority = "High"
	default:
		priority = "Not Set"
	}

	if report.Evidence {
		evidenceAvailable = "Evidence is available"
	}

	message, err := client.CreateEmbeds([]discord.Embed{
		discord.NewEmbedBuilder().
			SetTitle(report.Title).
			SetDescriptionf("[View report](%s/report/%s)", config.LoadedConfig.PublicURL, report.UUID).
			AddField("Issue Type", report.Type, false).
			AddField("Issue Priority", priority, false).
			AddField("Platform", report.Platform, false).
			AddField("Report Details", fmt.Sprintf("```\n%s\n```", report.Description), false).
			AddField("Is Evidence Available?", evidenceAvailable, false).
			SetTimestamp(report.SubmitTime).
			SetFooterTextf("This report was submitted by %s", report.Username).Build(),
	})
	if err != nil {
		return "", err
	}

	return message.ID.String(), nil
}
