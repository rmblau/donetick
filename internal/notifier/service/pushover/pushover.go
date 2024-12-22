package pushover

import (
	"context"

	"donetick.com/core/config"
	nModel "donetick.com/core/internal/notifier/model"
	"donetick.com/core/logging"
	"github.com/gregdel/pushover"
)

type Pushover struct {
	pushover *pushover.Pushover
}

func NewPushover(cfg *config.Config) *Pushover {

	pushoverApp := pushover.New(cfg.Pushover.Token)

	return &Pushover{
		pushover: pushoverApp,
	}
}

func (p *Pushover) SendNotification(c context.Context, notification *nModel.Notification) error {
	log := logging.FromContext(c)
	recipient := pushover.NewRecipient(notification.TargetID)
	message := pushover.NewMessageWithTitle(notification.Text, "Donetick")

	_, err := p.pushover.SendMessage(message, recipient)
	if err != nil {
		log.Debug("Error sending pushover notification", err)
		return err
	}

	return nil
}