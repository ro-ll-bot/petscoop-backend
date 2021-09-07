package seller

import (
	"github.com/ro-ll-bot/petscoop/entity"

type Seller struct {
	entity.User
	Role string
	Permissions []string
}
