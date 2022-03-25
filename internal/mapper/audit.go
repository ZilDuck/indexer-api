package mapper

import (
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/entity"
)

func AuditsToDtos(e []entity.Audit) []dto.Audit {
	audits := make([]dto.Audit, 0)
	for idx := range e {
		audits = append(audits, AuditToDto(e[idx]))
	}

	return audits
}

func AuditToDto(e entity.Audit) dto.Audit {
	return dto.Audit{
		Time: e.Time,
		Request: e.Request,
		Network: e.Network,
		RemoteAddr: e.RemoteAddr,
	}
}
