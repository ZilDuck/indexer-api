package metadata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/hashicorp/go-retryablehttp"
)

type Service interface {
	GetZrc6Metadata(nft entity.Nft) (interface{}, error)
	GetZrc6Media(nft entity.Nft) ([]byte, string, error)
}

type service struct {
	client *retryablehttp.Client
}

func NewMetadataService(client *retryablehttp.Client) Service {
	return service{client}
}

func (s service) GetZrc6Metadata(nft entity.Nft) (interface{}, error) {
	resp, err := retryablehttp.Get(fmt.Sprintf("%s%d", nft.TokenUri, nft.TokenId))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var md interface{}

	if err := json.Unmarshal(buf.Bytes(), &md); err != nil {
		return nil, err
	}

	return md, nil
}

func (s service) GetZrc6Media(nft entity.Nft) ([]byte, string, error) {
	if nft.MediaUri == "" {
		return nil, "", errors.New("media not found")
	}

	resp, err := retryablehttp.Get(nft.MediaUri)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != 200 {
		return nil, "", errors.New(resp.Status)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), resp.Header.Get("Content-Type"), nil
}

