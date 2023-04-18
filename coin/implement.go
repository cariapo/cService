package coin

import (
	"fmt"
	"github.com/cariapo/cservice/coin/model/request"
	"github.com/cariapo/cservice/coin/model/response"
	"github.com/cuwand/pondasi/helper/httpHelper"
	"github.com/cuwand/pondasi/logger"
	"net/http"
	"strings"
)

type implementOutbound struct {
	config                         httpHelper.HttpConfig
	Logger                         logger.Logger
	findAccountByIdUrl             string
	findAccountByNumberUrl         string
	findCustomerByIdUrl            string
	findCustomerByCifUrl           string
	statementsUrl                  string
	overbookingUrl                 string
	overbookingMultiBeneficiaryUrl string
	reversalUrl                    string
}

func ImplementOutbound(httpConfig httpHelper.HttpConfig, logger logger.LogConfig) CoinOutbound {
	return implementOutbound{
		config:                         httpConfig,
		Logger:                         logger,
		findAccountByIdUrl:             httpConfig.BuildBaseUrl() + "/v1/accounts/:id",
		findAccountByNumberUrl:         httpConfig.BuildBaseUrl() + "/v1/accounts/by-number/:accountNumber",
		findCustomerByIdUrl:            httpConfig.BuildBaseUrl() + "/v1/customers/:id",
		findCustomerByCifUrl:           httpConfig.BuildBaseUrl() + "/v1/customers/by-cif/:cif",
		statementsUrl:                  httpConfig.BuildBaseUrl() + "/v1/statements/:accountNumber",
		overbookingUrl:                 httpConfig.BuildBaseUrl() + "/v1/transactions/overbooking",
		overbookingMultiBeneficiaryUrl: httpConfig.BuildBaseUrl() + "/v1/transactions/overbooking-multi-beneficiary",
		reversalUrl:                    httpConfig.BuildBaseUrl() + "/v1/transactions/reversal",
	}
}

func (i implementOutbound) FindAccountByNumber(req request.FindAccountByNumber) (*response.Account, error) {
	resp := &response.Account{}

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
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindAccountById(req request.FindAccountById) (*response.Account, error) {
	resp := &response.Account{}

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
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindCustomerByCIF(req request.FindCustomerByCif) (*response.Customer, error) {
	resp := &response.Customer{}

	if err := httpHelper.HttpRequest(httpHelper.HttpRequestPayload{
		Url:        strings.ReplaceAll(i.findCustomerByCifUrl, ":cif", req.Cif),
		Method:     http.MethodGet,
		Result:     resp,
		Client:     i.config.HttpClient,
		TimeoutReq: i.config.Timeout,
		Logger:     i.Logger,
		Converter:  CoinConverter,
		Header: &http.Header{
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) FindCustomerById(req request.FindCustomerById) (*response.Customer, error) {
	resp := &response.Customer{}

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
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) Overbooking(req request.Overbooking) (*response.Overbooking, error) {
	resp := &response.Overbooking{}

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
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) OverbookingMultiBeneficiary(req request.OverbookingMultiBeneficiary) (*response.OverbookingMultiBeneficiary, error) {
	resp := &response.OverbookingMultiBeneficiary{}

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
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) Reversal(req request.Reversal) (*response.Reversal, error) {
	resp := &response.Reversal{}

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
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}

func (i implementOutbound) Statements(req request.Statement) (*response.ListStatement, error) {
	resp := &response.ListStatement{}

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
			"Accept":        []string{httpHelper.ContentTypeApplicationJson},
			"Authorization": []string{req.AccessToken},
		},
	}); err != nil {
		return nil, err
	}

	return resp, nil
}
