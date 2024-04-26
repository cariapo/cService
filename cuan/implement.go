package cuan

import (
	"encoding/json"
	"fmt"
	"github.com/cariapo/cservice/credential"
	request2 "github.com/cariapo/cservice/credential/model/request"
	"github.com/cariapo/cservice/cuan/model/request"
	"github.com/cariapo/cservice/cuan/model/response"
	"github.com/cuwand/pondasi/database/redis"
	"github.com/cuwand/pondasi/enum/grandTypeEnums"
	"github.com/cuwand/pondasi/helper/headerHelper"
	"github.com/cuwand/pondasi/helper/httpHelper"
	"github.com/cuwand/pondasi/logger"
	"net/http"
	"strings"
	"time"
)

type implementOutbound struct {
	config                           httpHelper.HttpConfig
	Logger                           logger.Logger
	redisClient                      redis.Redis
	credentialOutbound               credential.CredentialOutbound
	credentialUsername               string
	credentialPassword               string
	findProductUrl                   string
	findProductTelcoUrl              string
	getProductTelcoProviderUrl       string
	getProductByProductCodeUrl       string
	storeInquiryUrl                  string
	storePaymentUrl                  string
	getPaymentByCustomerRefNumberUrl string
}

var redisClient redis.Redis
var accessTokenKey string

type CuanConfig struct {
	HttpConfig         httpHelper.HttpConfig
	Logger             logger.LogConfig
	RedisClient        redis.Redis
	CredentialUsername string
	CredentialPassword string
	CredentialOutbound credential.CredentialOutbound
}

func ImplementOutbound(config CuanConfig) CuanOutbound {
	redisClient = config.RedisClient

	return implementOutbound{
		config:                           config.HttpConfig,
		Logger:                           config.Logger,
		redisClient:                      config.RedisClient,
		credentialOutbound:               config.CredentialOutbound,
		credentialUsername:               config.CredentialUsername,
		credentialPassword:               config.CredentialPassword,
		findProductUrl:                   config.HttpConfig.BuildBaseUrl() + "/v1/products",
		getProductByProductCodeUrl:       config.HttpConfig.BuildBaseUrl() + "/v1/products/:productCode",
		findProductTelcoUrl:              config.HttpConfig.BuildBaseUrl() + "/v1/products/telco/:phoneNumber",
		getProductTelcoProviderUrl:       config.HttpConfig.BuildBaseUrl() + "/v1/products/telco-provider/:phoneNumber",
		storeInquiryUrl:                  config.HttpConfig.BuildBaseUrl() + "/v1/inquiries",
		storePaymentUrl:                  config.HttpConfig.BuildBaseUrl() + "/v1/payments",
		getPaymentByCustomerRefNumberUrl: config.HttpConfig.BuildBaseUrl() + "/v1/payments/by-customer-ref/:customerReferenceNumber",
	}
}

func (i implementOutbound) getAccessToken() (*string, error) {
	var accessToken *string

	accessTokenKey = fmt.Sprintf("access-token:%s:%s", i.credentialUsername, i.credentialPassword)

	if err := i.redisClient.Get(accessTokenKey, &accessToken); err != nil {
		return nil, err
	}

	if accessToken != nil {
		return accessToken, nil
	}

	oauth, err := i.credentialOutbound.Oauth(request2.StoreToken{
		GrantType:    grandTypeEnums.CLIENT_CREDENTIAL,
		Username:     i.credentialUsername,
		Password:     i.credentialPassword,
		RefreshToken: "",
	})

	if err != nil {
		fmt.Println("-BACA ERROR GET ACCESS TOKEN-")
		fmt.Println(oauth)
		fmt.Println(err.Error())

		return nil, err
	}

	if err := i.redisClient.Set(accessTokenKey, oauth.GetAccessToken(), 25*time.Minute); err != nil {
		return nil, err
	}

	return i.getAccessToken()
}

func (i implementOutbound) FindProduct(req request.FindProduct) (*response.ListProduct, error) {
	resp := &response.ListProduct{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	queryParams := make(map[string]string)

	queryParams["page"] = fmt.Sprintf("%v", req.Page)

	queryParams["item_per_page"] = fmt.Sprintf("%v", req.ItemPerPage)

	queryParams["sort_by"] = req.SortBy

	if req.Provider != nil {
		queryParams["provider"] = req.Provider.String()
	}

	if req.Plan != nil {
		queryParams["plan"] = req.Plan.String()
	}

	if req.Type != nil {
		queryParams["type"] = req.Type.String()
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:         i.findProductUrl,
		Method:      http.MethodGet,
		Result:      resp,
		QueryParams: queryParams,
		Client:      i.config.HttpClient,
		TimeoutReq:  i.config.Timeout,
		Logger:      i.Logger,
		Converter:   CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindProductTelco(req request.FindProductTelco) (*response.ListProductAggregate, error) {
	resp := &response.ListProductAggregate{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	queryParams := make(map[string]string)

	queryParams["page"] = fmt.Sprintf("%v", req.Page)

	queryParams["item_per_page"] = fmt.Sprintf("%v", req.ItemPerPage)

	queryParams["sort_by"] = req.SortBy

	queryParams["plan"] = req.Plan.String()

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:         strings.ReplaceAll(i.findProductTelcoUrl, ":phoneNumber", req.PhoneNumber),
		Method:      http.MethodGet,
		Result:      resp,
		QueryParams: queryParams,
		Client:      i.config.HttpClient,
		TimeoutReq:  i.config.Timeout,
		Logger:      i.Logger,
		Converter:   CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) GetProductByProductCode(req request.GetProduct) (*response.GetProduct, error) {
	resp := &response.GetProduct{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.getProductByProductCodeUrl, ":productCode", req.ProductCode),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) GetProductProvider(req request.GetProvider) (*response.Provider, error) {
	resp := &response.Provider{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.getProductTelcoProviderUrl, ":phoneNumber", req.PhoneNumber),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) StoreInquiry(req request.StoreInquiry) (*response.Inquiry, error) {
	resp := &response.Inquiry{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	marshaledObj, _ := json.Marshal(req)

	fmt.Println("string(marshaledObj)")
	fmt.Println(string(marshaledObj))

	coreUA := headerHelper.GenerateUserAudit(req.User)

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.storeInquiryUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
			"X-CORE-UA":     []string{coreUA},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) StorePayment(req request.StorePayment) (*response.Payment, error) {
	resp := &response.Payment{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	marshaledObj, _ := json.Marshal(req)

	fmt.Println("string(marshaledObj)")
	fmt.Println(string(marshaledObj))

	coreUA := headerHelper.GenerateUserAudit(req.User)

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.storePaymentUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
			"X-CORE-UA":     []string{coreUA},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) GetPaymentByCustomerRefNumber(req request.GetPaymentByCustomerRefNumber) (*response.Payment, error) {
	resp := &response.Payment{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.getPaymentByCustomerRefNumberUrl, ":customerReferenceNumber", req.CustomerReferenceNumber),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CuanConverter,
		Header: &http.Header{
			"Content-Type":  []string{httpHelper.ContentTypeApplicationJson},
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}
