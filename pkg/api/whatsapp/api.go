package whatsapp

import (
	"context"
	"fmt"

	businessProfileAPI "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/business_profile"
	customClient "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/client"
	mediaAPI "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/media"
	messageAPI "github.com/jkandasa/whatsapp-cloud-api/pkg/api/whatsapp/message"
	types "github.com/jkandasa/whatsapp-cloud-api/pkg/types"
	whatsappTY "github.com/jkandasa/whatsapp-cloud-api/pkg/types/whatsapp"
	loggerUtils "github.com/jkandasa/whatsapp-cloud-api/pkg/utils/logger"
	"go.uber.org/zap"
)

type WhatsAppClient struct {
	ctx    context.Context
	logger *zap.Logger
	client *customClient.Client
	cfg    types.WhatsAppConfig
}

func New(ctx context.Context, cfg types.WhatsAppConfig) (*WhatsAppClient, error) {
	logger, err := loggerUtils.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	logger = logger.Named("whatsapp_client")

	version := cfg.Version
	// if the version not defined in the config, use the default version
	if version == "" {
		version = whatsappTY.DEFAULT_API_VERSION
	}
	baseURL := fmt.Sprintf("%s/%s", whatsappTY.BASE_URL, version)
	logger.Debug("base url formed", zap.String("baseUrl", baseURL))

	// get custom http client
	// inject required authentication parameters
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", cfg.AccessToken),
	}
	client, err := customClient.New(ctx, baseURL, headers)
	if err != nil {
		logger.Error("error on getting custom http client", zap.Error(err))
		return nil, err
	}

	whatsAppClient := &WhatsAppClient{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
		client: client,
	}

	return whatsAppClient, nil
}

func (wc *WhatsAppClient) BusinessProfile() *businessProfileAPI.BusinessProfileAPI {
	return businessProfileAPI.New(wc.ctx, wc.client, wc.cfg.PhoneNumberID)
}

func (wc *WhatsAppClient) Media() *mediaAPI.MediaAPI {
	return mediaAPI.New(wc.ctx, wc.client, wc.cfg.PhoneNumberID)
}

func (wc *WhatsAppClient) Message() *messageAPI.MessageAPI {
	return messageAPI.New(wc.ctx, wc.client, wc.cfg.PhoneNumberID)
}
