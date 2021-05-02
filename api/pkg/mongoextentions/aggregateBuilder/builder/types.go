package builder

import "go.mongodb.org/mongo-driver/bson"

type SMap struct {
	Operators []Operator
}

func (s *SMap) ToMap() bson.M {
	m := bson.M{}

	for _, o := range s.Operators {
		m[o.GetKey()] = o.GetVal()
	}

	return m
}

func S(operators ...Operator) bson.M {
	s := &SMap{Operators: operators}

	return s.ToMap()
}

// bson.E STUFF!!!!!
func (s *SMap) ToMapD() []bson.E {
	dArray := []bson.E{}

	for _, o := range s.Operators {
		//m[o.GetKey()] = o.GetVal()
		dArray = append(dArray, bson.E{
			Key:   o.GetKey(),
			Value: o.GetVal(),
		})
	}

	return dArray
}
