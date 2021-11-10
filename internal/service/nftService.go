package service

import (
	"github.com/dantudor/zilkroad-txapi/internal/dto"
	"github.com/dantudor/zilkroad-txapi/internal/factory"
	"github.com/dantudor/zilkroad-txapi/internal/repository"
	"go.uber.org/zap"
)

type NFTService interface {
	GetForAddress(ownerAddr string, from, size int) (dto.NFTs, int64, error)
}

type nftService struct {
	nftRepo repository.NftRepository
}

func NewNFTService(nftRepo repository.NftRepository) NFTService {
	return nftService{nftRepo}
}

func (s nftService) GetForAddress(ownerAddr string, from, size int) (dto.NFTs, int64, error) {
	zap.S().Infof("NFTService:GetNftsForAddress(%s, %d, %d)", ownerAddr, from, size)

	nfts, total, err := s.nftRepo.GetForAddress(ownerAddr, from, size)
	if err != nil {
		return nil, 0, err
	}
	zap.S().Infof("Found %d NFT For %s", total, ownerAddr)

	return factory.NftsIndexToDto(nfts), total, err
}
