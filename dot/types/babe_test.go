package types

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/ChainSafe/gossamer/lib/crypto/sr25519"
	"github.com/ChainSafe/gossamer/lib/keystore"
)

func TestAuthorityDataRaw(t *testing.T) {
	ad := new(AuthorityDataRaw)
	buf := &bytes.Buffer{}
	data := []byte{0, 91, 50, 25, 214, 94, 119, 36, 71, 216, 33, 152, 85, 184, 34, 120, 61, 161, 164, 223, 76, 53, 40, 246, 76, 38, 235, 204, 43, 31, 179, 28, 1, 0, 0, 0, 0, 0, 0, 0}
	buf.Write(data)

	_, err := ad.Decode(buf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuthorityData(t *testing.T) {
	kr, err := keystore.NewKeyring()
	if err != nil {
		t.Fatal(err)
	}

	ad := NewAuthorityData(kr.Alice.Public().(*sr25519.PublicKey), 77)
	enc := ad.Encode()

	buf := &bytes.Buffer{}
	buf.Write(enc)

	res := new(AuthorityData)
	err = res.Decode(buf)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res.ID.Encode(), ad.ID.Encode()) {
		t.Fatalf("Fail: got %v expected %v", res.ID.Encode(), ad.ID.Encode())
	}

	if res.Weight != ad.Weight {
		t.Fatalf("Fail: got %d expected %d", res.Weight, ad.Weight)
	}
}

func TestAuthorityData_ToRaw(t *testing.T) {
	kr, err := keystore.NewKeyring()
	if err != nil {
		t.Fatal(err)
	}

	ad := NewAuthorityData(kr.Alice.Public().(*sr25519.PublicKey), 77)
	raw := ad.ToRaw()

	res := new(AuthorityData)
	err = res.FromRaw(raw)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(res.ID.Encode(), ad.ID.Encode()) {
		t.Fatalf("Fail: got %v expected %v", res.ID.Encode(), ad.ID.Encode())
	}

	if res.Weight != ad.Weight {
		t.Fatalf("Fail: got %d expected %d", res.Weight, ad.Weight)
	}
}