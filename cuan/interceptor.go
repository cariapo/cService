package cuan

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/logger"
	"github.com/cuwand/pondasi/response"
	"io"
	"net"
	"net/http"
	"strings"
)

func CuanErrorInterceptor(resp *http.Response, logger logger.Logger, err error) error {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.InternalServerError("Internal Service Request Timeout - Credential Service")
	}

	if err, ok := err.(net.Error); ok && strings.Contains(err.Error(), "connection refused") {
		return errors.InternalServerError("Internal Service Down - Credential Service")
	}

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			redisClient.Del(accessTokenKey)
		}

		defer resp.Body.Close()

		readResp, err := io.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		var errResp response.ModelError

		if err := json.Unmarshal(readResp, &errResp); err != nil {
			return err
		}

		return errors.Error(errResp.ErrorString())
	}

	return nil
}
