// Code generated by "stringer -type=Type"; DO NOT EDIT.

package data

import "strconv"

const _Type_name = "TxDataStringBytesAddressUintIntBool"

var _Type_index = [...]uint8{0, 6, 12, 17, 24, 28, 31, 35}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}