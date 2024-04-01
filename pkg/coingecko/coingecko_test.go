package coingecko

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

func mockHandleSuccess(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]PriceResponse{
		"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0": {Usd: 0.670028},
		"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": {Usd: 5.73},
		"0x3845badade8e6dff049820680d1f14bd3903a5d0": {Usd: 1.32},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(jsonResponse)
}

func mockHandleFailure(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]PriceResponse{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(jsonResponse)
}

func TestFetchPrices(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				switch strings.TrimSpace(r.URL.Path) {
				case "/simple/token_price/ethereum":
					mockHandleSuccess(w, r)
				default:
					http.NotFoundHandler().ServeHTTP(w, r)
				}
			},
		),
	)
	defer server.Close()

	chainId := valueobject.ChainIDEthereum

	s := NewCoingeckoPriceSource(server.URL, "", time.Second)

	result, err := s.ListPrices(
		context.Background(),
		chainId, []string{
			"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0",
			"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
			"0x3845badade8e6dff049820680d1f14bd3903a5d0",
		},
	)

	if err != nil {
		t.Errorf("TestFetchPrice failed, err: %v", err)
		return
	}

	want := map[string]float64{
		"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0": 0.670028,
		"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": 5.73,
		"0x3845badade8e6dff049820680d1f14bd3903a5d0": 1.32,
	}

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("want %v\n, got %v\n", want, result)
	}
}

func TestFetchPricesFailed(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				switch strings.TrimSpace(r.URL.Path) {
				case "/simple/token_price/ethereum":
					mockHandleFailure(w, r)
				default:
					http.NotFoundHandler().ServeHTTP(w, r)
				}
			},
		),
	)
	defer server.Close()

	chainId := valueobject.ChainIDEthereum

	s := NewCoingeckoPriceSource(server.URL, "", time.Second)

	_, err := s.ListPrices(
		context.Background(),
		chainId, []string{
			"0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0",
			"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
			"0x3845badade8e6dff049820680d1f14bd3903a5d0",
		},
	)

	if err == nil {
		t.Error("TestFetchPriceFailed should pass")
		return
	}
}
