package field

import "go.mongodb.org/mongo-driver/bson"

var (
	// EmptyDoc is empty document.
	EmptyDoc = bson.M{}
)

// $lookup fields
const (
	From         = "from"
	LocalField   = "localField"
	ForeignField = "foreignField"
	As           = "as"
	Let          = "let"
	Pipeline     = "pipeline"
)

// $merge
const (
	Into           = "into"
	On             = "on"
	WhenMatched    = "whenMatched"
	WhenNotMatched = "whenNotMatched"
)

// $replaceRoot
const (
	NewRoot = "newRoot"
)

// $sample
const (
	Size = "size"
)

// $unwind
const (
	Path                       = "path"
	IncludeArrayIndex          = "includeArrayIndex"
	PreserveNullAndEmptyArrays = "preserveNullAndEmptyArrays"
)

// Cond
const (
	If   = "if"
	Then = "then"
	Else = "else"
	Null = "null"
)
