package queries

// EsQuery struct
type EsQuery struct{
	Equals 		[]FieldValue `json:"equals"`
	NotEquals 	[]FieldValue `json:"not_equals"`
}

// FieldValue struct
type FieldValue struct{
	Field string 		`json:"field"`
	Value interface{} 	`json:"value"`
}
/*
{
	
		equal: [
			{
				field: "",
				value: ""
			},{
				field: "",
				value: ""
			},
		]
	
}
*/
