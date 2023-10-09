package commands

import (
	"encoding/json"

	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/internal/code/domain"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	GenerateCodeCmd = "generate.code"
	ValidateCodeCmd = "validate.code"
)

func RegisterCommands(r setup.CommandRegistry) {
	codec.Register[GenerateCodeData](r, string(GenerateCodeCmd))
	codec.Register[ValidateCodeData](r, string(ValidateCodeCmd))
}

type GenerateCodeData struct {
	UserID   primitive.ObjectID
	Code     string
	Email    string
	Type     domain.CodeType
	Category domain.CodeCategory
	Data     json.RawMessage
}

type ValidateCodeData struct {
	UserID   primitive.ObjectID
	Code     string
	Type     domain.CodeType
	Category domain.CodeCategory
}
