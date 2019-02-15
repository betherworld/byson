package tokenservice

import (
	"fin4-core/server/datatype"
)

func (db *Service) FindTokenByID(id datatype.ID) (*datatype.Token, error) {
	return db.findTokenBy("id", string(id))
}
