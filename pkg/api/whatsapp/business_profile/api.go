package businessprofile

import (
	"context"
	"fmt"

	customClient "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/client"
	whatsappTY "github.com/jkandasa/whatsapp-cloud-api/pkg/types/whatsapp"
)

type BusinessProfileAPI struct {
	phoneNumberID string
	client        *customClient.Client
}

func New(ctx context.Context, client *customClient.Client, phoneNumberID string) *BusinessProfileAPI {
	return &BusinessProfileAPI{
		phoneNumberID: phoneNumberID,
		client:        client,
	}
}

func (bp *BusinessProfileAPI) Get() (*whatsappTY.BusinessProfile, error) {
	// /{{Phone-Number-ID}}/whatsapp_business_profile
	api := fmt.Sprintf("/%s/whatsapp_business_profile", bp.phoneNumberID)
	// {"data":[{"messaging_product":"whatsapp"}]}
	out := struct {
		Data []whatsappTY.BusinessProfile `json:"data"`
	}{}
	err := bp.client.Get(api, nil, nil, &out)
	if err != nil {
		return nil, err
	}
	if len(out.Data) > 0 {
		return &out.Data[0], nil
	}
	return &whatsappTY.BusinessProfile{}, nil
}
