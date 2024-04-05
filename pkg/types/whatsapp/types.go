package whatsapp

type StatusResponse struct {
	Success bool `json:"success,omitempty"`
}

type BusinessProfile struct {
	About             string   `json:"about,omitempty"`
	Address           string   `json:"address,omitempty"`
	Description       string   `json:"description,omitempty"`
	Email             string   `json:"email,omitempty"`
	MessagingProduct  string   `json:"messaging_product,omitempty"`
	ProfilePictureURL string   `json:"profile_picture_url,omitempty"`
	Vertical          string   `json:"vertical,omitempty"`
	Websites          []string `json:"websites,omitempty"`
}

type Media struct {
	ID               string `json:"id,omitempty"`
	File             string `json:"file,omitempty"`
	MediaType        string `json:"type,omitempty"`
	MessagingProduct string `json:"messaging_product,omitempty"`
}

// https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages
type Message struct {
	BizOpaqueCallbackData string          `json:"biz_opaque_callback_data,omitempty"`
	Context               *MessageContext `json:"context,omitempty"`

	// object types
	Audio       *MessageMediaObject    `json:"audio,omitempty"`
	Contacts    interface{}            `json:"contacts,omitempty"`
	Document    *MessageMediaObject    `json:"document,omitempty"`
	Image       *MessageMediaObject    `json:"image,omitempty"`
	Interactive *InteractiveObject     `json:"interactive,omitempty"`
	Location    interface{}            `json:"location,omitempty"`
	Sticker     *MessageMediaObject    `json:"sticker,omitempty"`
	Template    *MessageTemplateObject `json:"template,omitempty"`
	Text        *MessageTextObject     `json:"text,omitempty"`

	MessagingProduct string `json:"messaging_product,omitempty"`
	PreviewURL       bool   `json:"preview_url,omitempty"`
	RecipientType    string `json:"recipient_type,omitempty"`
	Status           string `json:"status,omitempty"`
	To               string `json:"to,omitempty"`
	Type             string `json:"type,omitempty"` // optional, default: text
}

type MessageResponse struct {
	MessagingProduct string        `json:"messaging_product,omitempty"`
	Contacts         interface{}   `json:"contacts,omitempty"`
	WaID             string        `json:"wa_id,omitempty"`
	Messages         []interface{} `json:"messages,omitempty"`
}

type MessageContext struct {
	MessageID string `json:"message_id,omitempty"`
}

// object definitions

// media object
// used in: audio, document, image, sticker
type MessageMediaObject struct {
	ID       string `json:"id,omitempty"`
	Link     string `json:"link,omitempty"`
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// text object
type MessageTextObject struct {
	Body       string `json:"body,omitempty"`
	PreviewURL bool   `json:"preview_url,omitempty"`
}

// template object
type MessageTemplateObject struct {
	Name       string            `json:"name,omitempty"`
	Language   Language          `json:"language,omitempty"`
	Components []ComponentObject `json:"components,omitempty"`
}

type ComponentObject struct {
	Type       string            `json:"type,omitempty"`
	SubType    string            `json:"sub_type,omitempty"`
	Parameters []ParameterObject `json:"parameters,omitempty"`
	Index      uint64            `json:"index,omitempty"`
}

type ParameterObject struct {
	Type     string          `json:"type,omitempty"` // options: currency, date_time, document, image, text, video
	Text     string          `json:"text,omitempty"`
	Currency *CurrencyObject `json:"currency,omitempty"`
	DateTime *DateTimeObject `json:"date_time,omitempty"`
	Image    *Media          `json:"image,omitempty"`
	Document *Media          `json:"document,omitempty"`
	Video    *Media          `json:"video,omitempty"`
}

type CurrencyObject struct {
	FallbackValue string `json:"fallback_value,omitempty"`
	Code          string `json:"code,omitempty"`
	Amount1000    string `json:"amount_1000,omitempty"`
}

type DateTimeObject struct {
	FallbackValue string `json:"fallback_value,omitempty"`
}

type Language struct {
	Code string `json:"code,omitempty"`
}

// reaction object
type MessageReactionObject struct {
	MessageID string `json:"message_id,omitempty"`
	Emoji     string `json:"emoji,omitempty"`
}

type InteractiveObject struct {
	Type   string        `json:"type,omitempty"` // options: button, catalog_message, list, product, product_list, flow
	Action *ActionObject `json:"action,omitempty"`
	Body   *TextObject   `json:"body,omitempty"`
	Footer *TextObject   `json:"footer,omitempty"`
	Header *HeaderObject `json:"header,omitempty"`
}

type HeaderObject struct {
	Type     string `json:"type,omitempty"` // options: text, video, image, document
	Document *Media `json:"document,omitempty"`
	Image    *Media `json:"image,omitempty"`
	Text     string `json:"text,omitempty"`
	Video    *Media `json:"video,omitempty"`
}

type TextObject struct {
	Text string `json:"text"`
}

type ActionObject struct {
	Button             string              `json:"button,omitempty"`
	Buttons            []InteractiveButton `json:"buttons,omitempty"`
	CatalogID          string              `json:"catalog_id,omitempty"`
	ProductRetailerID  string              `json:"product_retainer_id,omitempty"`
	Sections           []any               `json:"sections,omitempty"`
	Mode               string              `json:"mode,omitempty"`
	FlowMessageVersion string              `json:"flow_message_version,omitempty"` // must be 3
	FlowToken          string              `json:"flow_token,omitempty"`
	FlowID             string              `json:"flow_id,omitempty"`
	FlowAction         string              `json:"flow_action,omitempty"`
	FlowActionPayload  any                 `json:"flow_action_payload,omitempty"`
}

type InteractiveButton struct {
	Type  string                  `json:"type,omitempty"` // only supported type "reply"
	Title string                  `json:"title,omitempty"`
	ID    string                  `json:"id,omitempty"`
	Reply *InteractiveReplyButton `json:"reply,omitempty"`
}

type InteractiveReplyButton struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
}
