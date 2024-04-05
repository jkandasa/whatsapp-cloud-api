package message

import (
	"context"
	"fmt"

	customClient "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/client"
	whatsappTY "github.com/jkandasa/whatsapp-cloud-api/pkg/types/whatsapp"
)

type MessageAPI struct {
	phoneNumberID string
	client        *customClient.Client
}

func New(ctx context.Context, client *customClient.Client, phoneNumberID string) *MessageAPI {
	return &MessageAPI{
		phoneNumberID: phoneNumberID,
		client:        client,
	}
}

func (ma *MessageAPI) Post(message whatsappTY.Message) (*whatsappTY.MessageResponse, error) {
	// /{{Phone-Number-ID}}/messages
	api := fmt.Sprintf("/%s/messages", ma.phoneNumberID)
	out := &whatsappTY.MessageResponse{}
	err := ma.client.Post(api, nil, nil, &message, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
