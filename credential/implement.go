package credential

import (
	"github.com/cuwand/cservice/credential/model/request"
	"github.com/cuwand/cservice/credential/model/response"
	"github.com/cuwand/pondasi/helper/httpHelper"
	"github.com/cuwand/pondasi/logger"
	"net/http"
)

type implementOutbound struct {
	config       httpHelper.HttpConfig
	Logger       logger.Logger
	oauthUrl     string
	authorizeUrl string
}

func ImplementOutbound(httpConfig httpHelper.HttpConfig, logger logger.LogConfig) CredentialOutbound {
	return implementOutbound{
		config:       httpConfig,
		Logger:       logger,
		oauthUrl:     httpConfig.BuildBaseUrl() + "/v1/oauth/token",
		authorizeUrl: httpConfig.BuildBaseUrl() + "/v1/oauth/authorize",
	}
}

func (i implementOutbound) Authorize(accessToken string) (*response.Authorize, error) {
	resp := &response.Authorize{}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.authorizeUrl,
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{accessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) Oauth(req request.StoreToken) (*response.StoreToken, error) {
	resp := &response.StoreToken{}

	if err := httpHelper.HttpRequestFormUrlEncoded(httpHelper.HttpRequestUrlEncodedPayload{
		Url:    i.oauthUrl,
		Method: http.MethodPost,
		FormData: map[string]string{
			"grant_type":    req.GrantType.String(),
			"refresh_token": req.RefreshToken,
		},
		Result:     resp,
		Logger:     i.Logger,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationFormUrlEncoded},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{"Basic " + httpHelper.GenerateBasicAuth(req.Username, req.Password)},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}
