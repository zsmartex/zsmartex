package domain

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
