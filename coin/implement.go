package coin

import (
	"encoding/json"
	"fmt"
	"github.com/cariapo/cservice/coin/model/request"
	"github.com/cariapo/cservice/coin/model/response"
	"github.com/cariapo/cservice/credential"
	request2 "github.com/cariapo/cservice/credential/model/request"
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
	config                         httpHelper.HttpConfig
	Logger                         logger.Logger
	storeAccountUrl                string
	findAccountByIdUrl             string
	findAccountByNumberUrl         string
	storeCustomerUrl               string
	findCustomerByIdUrl            string
	findCustomerByCifUrl           string
	existsCustomerByPhoneNumberUrl string
	statementsUrl                  string
	overbookingUrl                 string
	overbookingMultiBeneficiaryUrl string
	reversalUrl                    string
	redisClient                    redis.Redis
	credentialOutbound             credential.CredentialOutbound
	credentialUsername             string
	credentialPassword             string
}

var redisClient redis.Redis
var accessTokenKey string

type CoinConfig struct {
	HttpConfig         httpHelper.HttpConfig
	Logger             logger.LogConfig
	RedisClient        redis.Redis
	CredentialUsername string
	CredentialPassword string
	CredentialOutbound credential.CredentialOutbound
}

func ImplementOutbound(config CoinConfig) CoinOutbound {
	redisClient = config.RedisClient

	return implementOutbound{
		config:                         config.HttpConfig,
		Logger:                         config.Logger,
		storeAccountUrl:                config.HttpConfig.BuildBaseUrl() + "/v1/accounts",
		findAccountByIdUrl:             config.HttpConfig.BuildBaseUrl() + "/v1/accounts/:id",
		findAccountByNumberUrl:         config.HttpConfig.BuildBaseUrl() + "/v1/accounts/by-number/:accountNumber",
		storeCustomerUrl:               config.HttpConfig.BuildBaseUrl() + "/v1/customers",
		findCustomerByIdUrl:            config.HttpConfig.BuildBaseUrl() + "/v1/customers/:id",
		findCustomerByCifUrl:           config.HttpConfig.BuildBaseUrl() + "/v1/customers/by-cif/:cif",
		existsCustomerByPhoneNumberUrl: config.HttpConfig.BuildBaseUrl() + "/v1/customers/exists-by-phone-number/:phoneNumber",
		statementsUrl:                  config.HttpConfig.BuildBaseUrl() + "/v1/statements/:accountNumber",
		overbookingUrl:                 config.HttpConfig.BuildBaseUrl() + "/v1/transactions/overbooking",
		overbookingMultiBeneficiaryUrl: config.HttpConfig.BuildBaseUrl() + "/v1/transactions/overbooking-multi-beneficiary",
		reversalUrl:                    config.HttpConfig.BuildBaseUrl() + "/v1/transactions/reversal",
		redisClient:                    config.RedisClient,
		credentialOutbound:             config.CredentialOutbound,
		credentialUsername:             config.CredentialUsername,
		credentialPassword:             config.CredentialPassword,
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

func (i implementOutbound) StoreAccount(req request.StoreAccount) (*response.Account, error) {
	resp := &response.Account{}

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
		Url:        i.storeAccountUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
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

func (i implementOutbound) FindAccountByNumber(req request.FindAccountByNumber) (*response.Account, error) {
	resp := &response.Account{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.findAccountByNumberUrl, ":accountNumber", req.AccountNumber),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindAccountById(req request.FindAccountById) (*response.Account, error) {
	resp := &response.Account{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.findAccountByIdUrl, ":id", req.Id),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) StoreCustomer(req request.StoreCustomer) (*response.Customer, error) {
	resp := &response.Customer{}

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
		Url:        i.storeCustomerUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
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

func (i implementOutbound) FindCustomerByCIF(req request.FindCustomerByCif) (*response.Customer, error) {
	resp := &response.Customer{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.findCustomerByCifUrl, ":cif", req.CIF),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindCustomerById(req request.FindCustomerById) (*response.Customer, error) {
	resp := &response.Customer{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.findCustomerByIdUrl, ":id", req.Id),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) ExistsCustomerByPhoneNumber(req request.ExistsCustomerByPhoneNumber) (*response.ExistsCustomer, error) {
	resp := &response.ExistsCustomer{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.existsCustomerByPhoneNumberUrl, ":phoneNumber", req.PhoneNumber),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{*req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) Overbooking(req request.Overbooking) (*response.Overbooking, error) {
	resp := &response.Overbooking{}

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

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.overbookingUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
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

func (i implementOutbound) OverbookingMultiBeneficiary(req request.OverbookingMultiBeneficiary) (*response.OverbookingMultiBeneficiary, error) {
	resp := &response.OverbookingMultiBeneficiary{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.overbookingMultiBeneficiaryUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
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

func (i implementOutbound) Reversal(req request.Reversal) (*response.Reversal, error) {
	resp := &response.Reversal{}

	if req.AccessToken == nil {
		accessToken, err := i.getAccessToken()

		if err != nil {
			return nil, err
		}

		req.AccessToken = accessToken
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        i.reversalUrl,
		Method:     http.MethodPost,
		Body:       req,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
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

func (i implementOutbound) Statements(req request.Statement) (*response.ListStatement, error) {
	resp := &response.ListStatement{}

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

	if req.TransactionFlow != nil {
		queryParams["transaction_flow"] = req.TransactionFlow.String()
	}

	if req.TransactionStatus != nil {
		queryParams["transaction_status"] = req.TransactionStatus.String()
	}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:         strings.ReplaceAll(i.statementsUrl, ":accountNumber", req.AccountNumber),
		Method:      http.MethodGet,
		Result:      resp,
		QueryParams: queryParams,
		Client:      i.config.HttpClient,
		TimeoutReq:  i.config.Timeout,
		Logger:      i.Logger,
		Converter:   CoinConverter,
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
