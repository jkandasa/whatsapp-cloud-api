package media

import (
	"context"
	"fmt"

	customClient "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/client"
	whatsappTY "github.com/jkandasa/whatsapp-cloud-api/pkg/types/whatsapp"
)

type MediaAPI struct {
	phoneNumberID string
	client        *customClient.Client
}

func New(ctx context.Context, client *customClient.Client, phoneNumberID string) *MediaAPI {
	return &MediaAPI{
		phoneNumberID: phoneNumberID,
		client:        client,
	}
}

func (m *MediaAPI) Upload(media whatsappTY.Media) (*whatsappTY.Media, error) {
	// /{{Phone-Number-ID}}/media
	api := fmt.Sprintf("/%s/media", m.phoneNumberID)
	out := &whatsappTY.Media{}
	err := m.client.Post(api, nil, nil, &media, out)
	return out, err
}

func (m *MediaAPI) Retrieve(mediaID string) (*whatsappTY.Media, error) {
	// /{{Media-ID}}?phone_number_id=<PHONE_NUMBER_ID>
	api := fmt.Sprintf("/%s", mediaID)
	out := &whatsappTY.Media{}
	err := m.client.Get(api, nil, nil, out)
	return out, err
}

func (m *MediaAPI) Delete(mediaID string) error {
	// /{{Media-ID}}/?phone_number_id=<PHONE_NUMBER_ID>
	api := fmt.Sprintf("/%s", mediaID)
	out := &whatsappTY.StatusResponse{}
	err := m.client.Delete(api, nil, nil, out)
	if err != nil {
		return err
	}

	if !out.Success {
		return fmt.Errorf("error on deleting media:%s", mediaID)
	}

	return nil
}

func (m *MediaAPI) Download(mediaURL string) ([]byte, error) {
	// /{{Media-URL}}
	api := fmt.Sprintf("/%s", mediaURL)
	out := []byte{}
	err := m.client.Get(api, nil, nil, &out)
	return out, err
}
