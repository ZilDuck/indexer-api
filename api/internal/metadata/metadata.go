package metadata

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/hashicorp/go-retryablehttp"
)

type Service interface {
	GetZrc6Metadata(nft entity.Nft) (interface{}, error)
}

type service struct {
	client *retryablehttp.Client
}

func NewMetadataService(client *retryablehttp.Client) Service {
	return service{client}
}

func (s service) GetZrc6Metadata(nft entity.Nft) (interface{}, error) {
	resp, err := retryablehttp.Get(nft.MetadataUri())
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


