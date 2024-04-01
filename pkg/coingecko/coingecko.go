package coingecko

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
	"github.com/KyberNetwork/toolkit-go/pkg/logger"
	"github.com/go-resty/resty/v2"
)

type coingeckoPriceSource struct {
	client *resty.Client
}

func NewCoingeckoPriceSource(
	baseURL string,
	proAPIKey string,
	timeout time.Duration,
) *coingeckoPriceSource {
	// Override MaxConnsPerHost, MaxIdleConnsPerHost and TLSHandshakeTimeout
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	transport.TLSHandshakeTimeout = timeout

	client := resty.New()
	client.SetBaseURL(baseURL)

	// Set the Pro API key if present
	if proAPIKey != "" {
		client.SetHeader(ProAPIHeaderKey, proAPIKey)
	}

	client.SetTimeout(timeout)
	client.SetTransport(transport)

	return &coingeckoPriceSource{
		client: client,
	}
}

func (s *coingeckoPriceSource) ListPrices(_ context.Context, chainId valueobject.ChainID, keys []string) (map[string]float64, error) {
	chain, ok := ChainIdCoingeckoPlatformMap[valueobject.ChainID(chainId)]
	if !ok {
		return nil, errors.New("coingecko does not support this chain")
	}

	url := fmt.Sprintf(
		"/simple/token_price/%s?contract_addresses=%s&vs_currencies=usd",
		chain,
		strings.Join(keys, ","),
	)

	var response map[string]PriceResponse

	resp, err := s.client.R().SetResult(&response).Get(url)
	if err != nil {
		logger.Errorf("failed to call Coingecko - simple price api, err: %v", err)
		return nil, err
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 400 {
		logger.Errorf("fail to get Coingecko request url: %v, header: %v", resp.Request.URL, resp.Request.Header)
		return nil, fmt.Errorf("fail to fetch Coingecko prices, response: %+v", resp)
	}

	result := make(map[string]float64)

	for address, tokenPriceResp := range response {
		if tokenPriceResp.Usd > 0 {
			result[address] = tokenPriceResp.Usd
		}
	}

	return result, nil
}

func (s *coingeckoPriceSource) ListCoins(_ context.Context) ([]CoingeckoCoin, error) {
	url := "/coins/list?include_platform=true"

	var response []CoingeckoCoin

	_, err := s.client.R().SetResult(&response).Get(url)
	if err != nil {
		logger.Errorf("failed to call Coingecko - list coins api, err: %v", err)
		return nil, err
	}

	return response, nil
}
