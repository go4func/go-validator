package main

import (
	"errors"
	"fmt"
)

type NonF2FRequest struct {
	Occupation       string   `json:"occupation" validate:"max=50,required"`
	SubOccupation    string   `json:"sub_occupation,omitempty"`
	WorkplaceName    string   `json:"workplace_name,omitempty" validate:"max=50"`
	WorkplaceAddress *Address `json:"workplace_address,omitempty"`
}

type Address struct {
	Address1 string `bson:"address_1,omitempty" json:"address_1" validate:"required,max=100"`
}

func main() {
	// n := NonF2FRequest{
	// 	Occupation:    "s",
	// 	SubOccupation: "a",
	// }
	// v := validator.New()
	// if err := v.Struct(n); err != nil {
	// 	panic(err)
	// }

	nonF2FRequests := []NonF2FRequest{
		NonF2FRequest{ // 0 pass
			Occupation:       "Government Officer",
			SubOccupation:    "Police Officer/Military Officer",
			WorkplaceName:    "KBTG",
			WorkplaceAddress: &Address{Address1: "address of working place"},
		},
		NonF2FRequest{ // 1 failed
			Occupation: "State Enterprise Employee",
			// SubOccupation:    "Police Officer/Military Officer",
			WorkplaceName:    "KBTG",
			WorkplaceAddress: &Address{Address1: "address of working place"},
		},
		NonF2FRequest{ // 2 pass
			Occupation: "Stay-at-Home Spouse",
			// SubOccupation:    "Police Officer/Military Officer",
			// WorkplaceName:    "KBTG",
			// WorkplaceAddress: Address{Address1: "address of working place"},
		},
		NonF2FRequest{ // 3 failed
			Occupation: "อาชีพอิสระ",
			// SubOccupation:    "Police Officer/Military Officer",
			WorkplaceName:    "KBTG",
			WorkplaceAddress: &Address{Address1: "address of working place"},
		},
		NonF2FRequest{ // 4 failed
			Occupation:    "Hired Worker/Temporary Worker",
			SubOccupation: "Police Officer/Military Officer",
			// WorkplaceName:    "KBTG",
			// WorkplaceAddress: Address{Address1: "address of working place"},
		},
	}
	for i, nonF2FRequest := range nonF2FRequests {
		fmt.Println(i, nonF2FRequest.Occupation)
		fmt.Println(i, validateOccupation(&nonF2FRequest))
	}

}

var THoccEN = map[string]string{
	"รับราชการ":                        "Government Officer",
	"พนักงานรัฐวิสาหกิจ":               "State Enterprise Employee",
	"พนักงานบริษัทเอกชน":               "Private Enterprise Employee",
	"เจ้าของกิจการที่จดทะเบียนพาณิชย์": "Business Proprietor",
	"พ่อบ้าน/แม่บ้าน":                  "Stay-at-Home Spouse",
	"นักเรียน/นักศึกษา":                "Student",
	"อาชีพอิสระ":                       "Freelance",
	"เกษตรกร":                          "Farmer",
	"รับจ้าง/พนักงานรายวัน/พนักงานชั่วคราว": "Hired Worker/Temporary Worker",
	"พระภิกษุ/นักบวช":                       "Monk/Priest",
	"เกษียณ":                                "Retiree",
}

var occs = map[string][]bool{
	"Government Officer":            []bool{true, true, true},
	"State Enterprise Employee":     []bool{true, true, true},
	"Private Enterprise Employee":   []bool{true, true, true},
	"Business Proprietor":           []bool{true, true, true},
	"Stay-at-Home Spouse":           []bool{false, false, false},
	"Student":                       []bool{false, false, false},
	"Freelance":                     []bool{true, false, false},
	"Farmer":                        []bool{false, false, false},
	"Hired Worker/Temporary Worker": []bool{true, false, true},
	"Monk/Priest":                   []bool{false, false, false},
	"Retiree":                       []bool{false, false, false},
}

func validateOccupation(req *NonF2FRequest) error {
	en, isTH := THoccEN[req.Occupation]
	if isTH {
		req.Occupation = en
	}

	occ, ok := occs[req.Occupation]
	if !ok {
		return errors.New("invalid occupation")
	}
	if occ[0] && (req.SubOccupation == "") {

		return errors.New("sub occupation is required")
	}
	if occ[1] && (req.WorkplaceName == "") {

		return errors.New("workplace name is required")
	}
	if occ[2] && (req.WorkplaceAddress == nil) {

		return errors.New("workplace address is required")
	}

	return nil
}

// Name of the struct tag used in examples
// const tagName = "validate"

// type User struct {
// 	Id    int    `validate:"-"`
// 	Name  string `validate:"presence,min=2,max=32"`
// 	Email string `validate:"email,required"`
// }

// func main() {
// 	user := User{
// 		Id:    1,
// 		Name:  "John Doe",
// 		Email: "john@example",
// 	}

// 	// TypeOf returns the reflection Type that represents the dynamic type of variable.
// 	// If variable is a nil interface value, TypeOf returns nil.
// 	t := reflect.TypeOf(user)

// 	// Get the type and kind of our user variable
// 	fmt.Println("Type:", t.Name())
// 	fmt.Println("Kind:", t.Kind())

// 	// Iterate over all available fields and read the tag value
// 	for i := 0; i < t.NumField(); i++ {
// 		// Get the field, returns https://golang.org/pkg/reflect/#StructField
// 		field := t.Field(i)

// 		// Get the field tag value
// 		tag := field.Tag.Get(tagName)

// 		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
// 	}
// }
