package model

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/caioeverest/fedits/internal/aes"
	"github.com/imroc/req/v3"
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"not null" validate:"required" json:"name" example:"Example LTDA"`
	Contact    string `json:"contact,omitempty" example:"some@email.com"`
	Slug       string `gorm:"not null;uniqueIndex" validate:"required,lowercase" json:"slug" example:"provider-slug"`
	Webhook    string `gorm:"not null" validate:"required,url" json:"webhook" example:"https://provider.com/webhook"`
	Secret     string `gorm:"not null" validate:"required" json:"secret"`
}

type ReqPayload struct {
	UserRef string `json:"user_ref" example:"user-ref"`
	Method  string `json:"method" example:"method-name"`
	Params  []any  `json:"params" example:"[\"param1\", \"param2\"]"`
}

func (p Provider) CallProviderMethod(ctx context.Context, hashSecret, userRef, methodName string, params []any) (result *req.Response, err error) {
	var (
		secret    string
		bytes     []byte
		signature string
		payload   = ReqPayload{
			UserRef: userRef,
			Method:  methodName,
			Params:  params,
		}
	)

	if bytes, err = json.Marshal(payload); err != nil {
		return
	}
	if secret, err = aes.Decrypt(hashSecret, p.Secret); err != nil {
		return
	}
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(bytes)
	signature = hex.EncodeToString(hash.Sum(nil))

	return req.R().
		SetContext(ctx).
		SetBodyJsonBytes(bytes).
		SetHeader("X-Signature", signature).
		Post(p.Webhook)
}
