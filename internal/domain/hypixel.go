package domain

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/Tnze/go-mc/nbt"
	"io"
)

type HypixelAuctionResponse struct {
	LastUpdated int64 `json:"lastUpdated"`
	Auctions    []struct {
		AuctionId     string `json:"auction_id"`
		Seller        string `json:"seller"`
		SellerProfile string `json:"seller_profile"`
		Buyer         string `json:"buyer"`
		BuyerProfile  string `json:"buyer_profile"`
		Timestamp     int64  `json:"timestamp"`
		Price         int    `json:"price"`
		Bin           bool   `json:"bin"`
		ItemBytes     string `json:"item_bytes"`
	} `json:"auctions"`
}

type HypixelAuction struct {
	IsBin bool
}

type HypixelBoughtAuction struct {
}

func DecodeNBT(data []byte) (map[string]any, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedData))
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	var nbtData map[string]any
	if err := nbt.Unmarshal(decompressedData, &nbtData); err != nil {
		return nil, err
	}

	return nbtData, nil
}
