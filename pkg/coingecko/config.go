package coingecko

import (
	"time"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/valueobject"
)

type CoingeckoPlatform string

type Config struct {
	Endpoint  string `mapstructure:"endpoint" json:"endpoint"`
	ProAPIKey string `mapstructure:"proAPIKey" json:"proAPIKey"`
}

const (
	DefaultAPIEndpoint = "https://api.coingecko.com/api/v3"
	DefaultTimeout     = 10 * time.Second
	ProAPIHeaderKey    = "x-cg-pro-api-key"
)

var (
	CoingeckoPlatformEthereum     CoingeckoPlatform = "ethereum"
	CoingeckoPlatformPolygon      CoingeckoPlatform = "polygon-pos"
	CoingeckoPlatformBSC          CoingeckoPlatform = "binance-smart-chain"
	CoingeckoPlatformAvalanche    CoingeckoPlatform = "avalanche"
	CoingeckoPlatformFantom       CoingeckoPlatform = "fantom"
	CoingeckoPlatformCronos       CoingeckoPlatform = "cronos"
	CoingeckoPlatformArbitrum     CoingeckoPlatform = "arbitrum-one"
	CoingeckoPlatformBittorrent   CoingeckoPlatform = "bittorrent"
	CoingeckoPlatformVelas        CoingeckoPlatform = "velas"
	CoingeckoPlatformAurora       CoingeckoPlatform = "aurora"
	CoingeckoPlatformOasis        CoingeckoPlatform = "oasis"
	CoingeckoPlatformOptimism     CoingeckoPlatform = "optimistic-ethereum"
	CoingeckoPlatformZKSync       CoingeckoPlatform = "zksync"
	CoingeckoPlatformSolana       CoingeckoPlatform = "solana"
	CoingeckoPlatformPolygonZkEVM CoingeckoPlatform = "polygon-zkevm"
	CoingeckoPlatformLinea        CoingeckoPlatform = "linea"
	CoingeckoPlatformBase         CoingeckoPlatform = "base"
	CoingeckoPlatformScroll       CoingeckoPlatform = "scroll"
)

// ChainIdCoingeckoPlatformMap can be fetched by using https://api.coingecko.com/api/v3/asset_platforms
// Remember to update this map when integrating new chain
var ChainIdCoingeckoPlatformMap = map[valueobject.ChainID]CoingeckoPlatform{
	valueobject.ChainIDEthereum:        CoingeckoPlatformEthereum,
	valueobject.ChainIDPolygon:         CoingeckoPlatformPolygon,
	valueobject.ChainIDBSC:             CoingeckoPlatformBSC,
	valueobject.ChainIDAvalancheCChain: CoingeckoPlatformAvalanche,
	valueobject.ChainIDFantom:          CoingeckoPlatformFantom,
	valueobject.ChainIDCronos:          CoingeckoPlatformCronos,
	valueobject.ChainIDArbitrumOne:     CoingeckoPlatformArbitrum,
	valueobject.ChainIDBitTorrent:      CoingeckoPlatformBittorrent, // Coingecko supports BTTC now, but only a few tokens
	valueobject.ChainIDVelasEVM:        CoingeckoPlatformVelas,
	valueobject.ChainIDAurora:          CoingeckoPlatformAurora,
	valueobject.ChainIDOasisEmerald:    CoingeckoPlatformOasis,
	valueobject.ChainIDOptimism:        CoingeckoPlatformOptimism,
	valueobject.ChainIDZKSync:          CoingeckoPlatformZKSync,
	valueobject.ChainIDSolana:          CoingeckoPlatformSolana,
	valueobject.ChainIDPolygonZkEVM:    CoingeckoPlatformPolygonZkEVM,
	valueobject.ChainIDLinea:           CoingeckoPlatformLinea,
	valueobject.ChainIDBase:            CoingeckoPlatformBase,
	valueobject.ChainIDScroll:          CoingeckoPlatformScroll,
}
