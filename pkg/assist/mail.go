package assist

import (
	"xddq/pkg/game/client"
	"xddq/pkg/game/protocol/pb"
)

const (
	GatherEnergyDislodged = "MAIL_GATHER_ENERGY_DISLODGED_TITLE"
)

type MailManager struct {
	a *Assistant
}

func NewMailManager(a *Assistant) *MailManager {
	m := &MailManager{
		a: a,
	}

	a.Client().OnMailListMsg20551(m.onMailDataMsg)
	a.Client().OnMailListMsg20552(m.onMailDataMsg)

	return m
}

func (m *MailManager) onMailDataMsg(ctx client.Context, msg *pb.MailListMsg) error {
	m.a.L().Info("处理邮件数据")

	for _, mail := range msg.GetMailList() {
		if !mail.GetIsRead() {
			return ctx.Send(ctx, &pb.MailGetAllRewardReq{})
		}
	}

	return nil
}
