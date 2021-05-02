package builder

import (
	f "chess-api/pkg/mongoextentions/aggregateBuilder/field"
	o "chess-api/pkg/mongoextentions/aggregateBuilder/operator"
	"go.mongodb.org/mongo-driver/bson"
)

func Project(params bson.M) Operator {
	m := bson.M{}

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Project, m)
}

func Out(outCollection string) Operator {

	return New(o.Out, outCollection)
}

func Match(params bson.M) Operator {
	m := bson.M{}

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Match, m)
}

func Group(ID interface{}, params bson.M) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.ID, ID)

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Group, m)
}
func Sort(params bson.M) Operator {
	m := bson.M{}

	for key, val := range params {
		appendIfHasVal(m, key, val)
	}

	return New(o.Sort, m)
}
func LookupExpert(from string, let bson.M, pipeline interface{}, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.Pipeline, pipeline)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}
func Lookup(from, localField, foreignField, as interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.From, from)
	appendIfHasVal(m, f.LocalField, localField)
	appendIfHasVal(m, f.ForeignField, foreignField)
	appendIfHasVal(m, f.As, as)

	return New(o.Lookup, m)
}

func Merge(into, on, let, whenMatched, whenNotMatched interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Into, into)
	appendIfHasVal(m, f.On, on)
	appendIfHasVal(m, f.Let, let)
	appendIfHasVal(m, f.WhenMatched, whenMatched)
	appendIfHasVal(m, f.WhenNotMatched, whenNotMatched)

	return New(o.Merge, m)
}
func Limit(limit interface{}) Operator {

	return New(o.Limit, limit)
}

func Sample(size interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Size, size)

	return New(o.Sample, m)
}

func Unwind(path, includeArrayIndex, preserveNullAndEmptyArrays interface{}) Operator {
	m := bson.M{}

	appendIfHasVal(m, f.Path, path)
	appendIfHasVal(m, f.IncludeArrayIndex, includeArrayIndex)
	appendIfHasVal(m, f.PreserveNullAndEmptyArrays, preserveNullAndEmptyArrays)

	return New(o.Unwind, m)
}
