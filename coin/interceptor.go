package coin

import (
	"fmt"
	"github.com/cuwand/pondasi/errors"
	"github.com/cuwand/pondasi/logger"
	"io"
	"net"
	"net/http"
)

func CoinErrorInterceptor(resp *http.Response, logger logger.Logger, err error) error {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return errors.BadRequest("Request Timeout")
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		readResp, err := io.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		logger.Error(string(readResp))
		logger.Error(fmt.Sprintf("%v", resp.StatusCode))

		return errors.BadRequest("request error.")
	}

	return nil
}
