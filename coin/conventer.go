package coin

import (
	"encoding/json"
	"fmt"
	"github.com/cuwand/pondasi/logger"
	"io"
	"net/http"
)

func CoinConverter(body interface{}, logger logger.Logger, httpResponse *http.Response, err error) error {
	if err := CoinErrorInterceptor(httpResponse, logger, err); err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	readResp, err := io.ReadAll(httpResponse.Body)

	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("[Response] Body | %v", string(readResp)))

	if err := json.Unmarshal(readResp, body); err != nil {
		return err
	}

	return nil
}
