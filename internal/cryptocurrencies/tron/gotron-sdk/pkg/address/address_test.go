package address

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAddress_Scan(t *testing.T) {
	validAddress, err := Base58ToAddress("TSvT6Bg3siokv3dbdtt9o4oM1CTXmymGn1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// correct case
	want := validAddress
	a := &Address{}
	src := validAddress.Bytes()
	err = a.Scan(src)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !bytes.Equal(a.Bytes(), want.Bytes()) {
		t.Errorf("got %v, want %v", *a, want)
	}

	// invalid type of src
	a = &Address{}
	err = a.Scan("not a byte slice")
	if err == nil {
		t.Errorf("expected an error, but got none")
	}

	// invalid length of src
	a = &Address{}
	src = make([]byte, 4)
	err = a.Scan(src)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}
	src = make([]byte, 22) // Створюємо байтовий масив з неправильною довжиною
	err = a.Scan(src)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}
}

func TestAddress_DecodeString(t *testing.T) {
	// Base64 字符串
	encodedStr := "QaYU+AO2/XgJhqQseOycf3fm3tE8"
	// 尝试解码
	//decodedData, err := base64.StdEncoding.DecodeString(encodedStr)
	decodedData, err := Base64ToAddress(encodedStr)
	if err != nil {
		fmt.Println("解码错误:", err)
		return
	}

	fmt.Println("解码结果:", decodedData.String())

}
