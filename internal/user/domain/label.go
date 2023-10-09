package domain

import commonv1 "github.com/zsmartex/zsmartex/proto/common/v1"

type LabelScope string

const (
	LabelScopePrivate LabelScope = "private"
	LabelScopePublic  LabelScope = "public"
)

type Label struct {
	Key   string     `bson:"key"`
	Value string     `bson:"value"`
	Scope LabelScope `bson:"scope"`
}

func (label Label) ProtobufValue() *commonv1.UserLabel {
	return &commonv1.UserLabel{
		Key:   label.Key,
		Value: label.Value,
		Scope: string(label.Scope),
	}
}
