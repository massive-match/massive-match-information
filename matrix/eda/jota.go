package eda

type Jota struct {
	Name       string      `json:"name"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Value       string   `json:"value"`
	Observation string   `json:"observation"`
	Obligatory  string   `json:"obligatory"`
	Values      []Values `json:"values"`
}

type Values struct {
	Value string `json:"value"`
}

func MatchCategories(matrices []Matrix, list []List, relations []Relation) []Jota {
	var jotas = make([]Jota, len(relations))

	for i := range relations {
		jotas[i].Name = relations[i].Jota
		jotas[i].Attributes = GetAttributes(relations[i].Matrix, matrices, list)
	}

	return jotas
}

func GetAttributes(idMatrix int, matrices []Matrix, list []List) []Attribute {
	var attributes []Attribute

	for i := range matrices {
		if matrices[i].Id == idMatrix {
			var attribute = Attribute{
				Value:       matrices[i].Attr,
				Observation: matrices[i].Obs,
				Obligatory:  matrices[i].Obligatory,
			}
			attribute.Values = GetValues(attribute.Value, list)

			attributes = append(attributes, attribute)
		}
	}

	return attributes
}

func GetValues(attr string, list []List) []Values {
	var values []Values

	for i := range list {
		if list[i].Attr == attr {
			var value = Values{
				Value: list[i].Value,
			}

			values = append(values, value)
		}
	}

	return values
}
