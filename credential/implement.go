package credential

import (
	"encoding/json"
	"fmt"
	"github.com/cariapo/cservice/credential/model/request"
	"github.com/cariapo/cservice/credential/model/response"
	"github.com/cuwand/pondasi/helper/headerHelper"
	"github.com/cuwand/pondasi/helper/httpHelper"
	"github.com/cuwand/pondasi/logger"
	"net/http"
)

type implementOutbound struct {
	config            httpHelper.HttpConfig
	Logger            logger.Logger
	oauthUrl          string
	authorizeUrl      string
	registerUrl       string
	changePasswordUrl string
	resetPasswordUrl  string
}

func (i implementOutbound) Register(req request.Register) (*response.User, error) {
	resp := &response.User{}

	marshaledObj, _ := json.Marshal(req)

	fmt.Println("string(marshaledObj)")
	fmt.Println(string(marshaledObj))

	coreUA := headerHelper.GenerateUserAudit(req.User)

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.registerUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CredentialConverter,
		Header: &http.Header{
			"Content-Type": []string{httpHelper.ContentTypeApplicationJson},
			"Accept":       []string{httpHelper.ContentTypeApplicationJson},
			//"Authorization": []string{*req.AccessToken},
			"X-CORE-UA": []string{coreUA},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) ChangePassword(req request.ChangePassword) (*response.User, error) {
	resp := &response.User{}

	marshaledObj, _ := json.Marshal(req)

	fmt.Println("string(marshaledObj)")
	fmt.Println(string(marshaledObj))

	coreUA := headerHelper.GenerateUserAudit(req.User)

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.changePasswordUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CredentialConverter,
		Header: &http.Header{
			"Content-Type": []string{httpHelper.ContentTypeApplicationJson},
			"Accept":       []string{httpHelper.ContentTypeApplicationJson},
			//"Authorization": []string{*req.AccessToken},
			"X-CORE-UA": []string{coreUA},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) ResetPassword(req request.ResetPassword) (*response.User, error) {
	resp := &response.User{}

	marshaledObj, _ := json.Marshal(req)

	fmt.Println("string(marshaledObj)")
	fmt.Println(string(marshaledObj))

	coreUA := headerHelper.GenerateUserAudit(req.User)

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.resetPasswordUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CredentialConverter,
		Header: &http.Header{
			"Content-Type": []string{httpHelper.ContentTypeApplicationJson},
			"Accept":       []string{httpHelper.ContentTypeApplicationJson},
			//"Authorization": []string{*req.AccessToken},
			"X-CORE-UA": []string{coreUA},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func ImplementOutbound(httpConfig httpHelper.HttpConfig, logger logger.LogConfig) CredentialOutbound {
	return implementOutbound{
		config:            httpConfig,
		Logger:            logger,
		oauthUrl:          httpConfig.BuildBaseUrl() + "/v1/oauth/token",
		authorizeUrl:      httpConfig.BuildBaseUrl() + "/v1/oauth/authorize",
		registerUrl:       httpConfig.BuildBaseUrl() + "/v1/users/register",
		changePasswordUrl: httpConfig.BuildBaseUrl() + "/v1/users/change-password",
		resetPasswordUrl:  httpConfig.BuildBaseUrl() + "/v1/users/reset-password",
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
		Converter:  CredentialConverter,
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
		Converter:  CredentialConverter,
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
