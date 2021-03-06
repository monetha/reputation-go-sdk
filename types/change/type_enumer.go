// Code generated by "enumer -type=Type"; DO NOT EDIT.

package change

import (
	"fmt"
)

const _TypeName = "UpdatedDeleted"

var _TypeIndex = [...]uint8{0, 7, 14}

func (i Type) String() string {
	if i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

var _TypeValues = []Type{0, 1}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:7]:  0,
	_TypeName[7:14]: 1,
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}
