package events

import (
	"github.com/modernice/goes/codec"
	"github.com/zsmartex/zsmartex/internal/code/domain"
	"github.com/zsmartex/zsmartex/pkg/setup"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	CodeGenerated = "code.generated"
	CodeValidated = "code.validated"
)

var ListEvents = []string{
	CodeGenerated,
	CodeValidated,
}

func RegisterEvents(r setup.EventRegistry) {
	codec.Register[CodeGeneratedData](r, string(CodeGenerated))
	codec.Register[CodeValidatedData](r, string(CodeValidated))
}

type CodeGeneratedData struct {
	UserID   primitive.ObjectID
	Type     domain.CodeType
	Category domain.CodeCategory
}

type CodeValidatedData struct {
}
