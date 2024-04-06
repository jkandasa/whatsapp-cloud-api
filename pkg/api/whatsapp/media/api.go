package media

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"

	customClient "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/client"
	whatsappTY "github.com/jkandasa/whatsapp-cloud-api/pkg/types/whatsapp"
	loggerUtils "github.com/jkandasa/whatsapp-cloud-api/pkg/utils/logger"
	"go.uber.org/zap"
)

type MediaAPI struct {
	logger        *zap.Logger
	phoneNumberID string
	client        *customClient.Client
}

func New(ctx context.Context, client *customClient.Client, phoneNumberID string) *MediaAPI {
	logger, err := loggerUtils.FromContext(ctx)
	if err != nil {
		logger = zap.NewNop()
	}
	return &MediaAPI{
		phoneNumberID: phoneNumberID,
		client:        client,
		logger:        logger.Named("media_api"),
	}
}

func (m *MediaAPI) getMediaPayloadBody(media *whatsappTY.Media) ([]byte, string, error) {
	// verify the file path or fileBytes should present
	if len(media.FileBytes) == 0 && len(media.File) == 0 {
		return nil, "", errors.New("either file (path to file) or fileBytes should be present")
	}

	// verify the file name
	if len(media.Filename) == 0 {
		return nil, "", errors.New("filename can not be empty")
	}

	// create a body buffer
	var body bytes.Buffer

	// create multi part writer
	writer := multipart.NewWriter(&body)

	// set filename header
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name=file; filename="%s"`, media.Filename))

	// set content type header
	contentType := mime.TypeByExtension(filepath.Ext(media.Filename))
	header.Set("Content-Type", contentType)

	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, "", fmt.Errorf("error on creating part: %w", err)
	}

	// source data reader
	var sourceReader io.Reader

	// if filebytes present, use that source data
	if len(media.FileBytes) > 0 {
		sourceReader = bytes.NewReader(media.FileBytes)
	} else { // otherwise, open the file and convert it to buffered reader
		file, err := os.Open(media.File)
		if err != nil {
			return nil, "", fmt.Errorf("error on opening a file[%s]: %w", media.File, err)
		}
		sourceReader = bufio.NewReader(file)
		defer func() {
			err := file.Close()
			if err != nil {
				m.logger.Error("error on closing a file", zap.String("file", media.Filename), zap.Error(err))
			}
		}()
	}

	// copy the source bytes
	_, err = io.Copy(part, sourceReader)
	if err != nil {
		return nil, "", fmt.Errorf("error on copying source data: %w", err)
	}

	// update media type
	err = writer.WriteField("type", string(media.MediaType))
	if err != nil {
		return nil, "", fmt.Errorf("error on setting type: %w", err)
	}

	// update messaging_product
	err = writer.WriteField("messaging_product", media.MessagingProduct)
	if err != nil {
		return nil, "", fmt.Errorf("error on setting messaging_product: %w", err)
	}

	// update filename
	//  see: https://stackoverflow.com/questions/58024665/how-to-set-filename-parameter-in-whatsapp-business-api-while-sending-document-at
	// err = writer.WriteField("filename", media.Filename)
	// if err != nil {
	// 	return nil, "", fmt.Errorf("error on setting filename: %w", err)
	// }

	err = writer.Close()
	if err != nil {
		return nil, "", fmt.Errorf("error on closing writer: %w", err)
	}

	return body.Bytes(), writer.FormDataContentType(), nil
}

func (m *MediaAPI) Upload(media *whatsappTY.Media) (*whatsappTY.Media, error) {
	// /{{Phone-Number-ID}}/media
	api := fmt.Sprintf("/%s/media", m.phoneNumberID)
	out := &whatsappTY.Media{}

	body, contentType, err := m.getMediaPayloadBody(media)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{"Content-Type": contentType}

	err = m.client.Post(api, headers, nil, body, out)
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
