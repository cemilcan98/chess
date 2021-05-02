package operator

// Comparison
const (
	Eq  = "$eq"
	Gt  = "$gt"
	Gte = "$gte"
	In  = "$in"
	Lt  = "$lt"
	Lte = "$lte"
	Ne  = "$ne"
	Nin = "$nin"
)

// Logical
const (
	And = "$and"
	Not = "$not"
	Nor = "$nor"
	Or  = "$or"
)

// Element
const (
	Exists = "$exists"
	Type   = "$type"
)

// Evaluation
const (
	Expr       = "$expr"
	JSONSchema = "$jsonSchema"
	Mod        = "$mod"
	Regex      = "$regex"
	Text       = "$text"
	Where      = "$where"
)

// Array
const (
	All       = "$all"
	ElemMatch = "$elemMatch"
	Size      = "$size"
)
