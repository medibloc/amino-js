package types

import (
	"encoding/json"
)

type DID string

type NetworkID string

type DIDDocument struct {
	Contexts        Contexts         `json:"@context"`
	ID              DID              `json:"id"`
	VeriMethods     []VeriMethod     `json:"verificationMethod"`
	Authentications []Authentication `json:"authentication"`
}

type Contexts []Context

func (ctxs Contexts) MarshalJSON() ([]byte, error) {
	if len(ctxs) == 1 { // if only one, treat it as a single string
		return json.Marshal(ctxs[0])
	}
	return json.Marshal([]Context(ctxs)) // if not, as a list
}

func (ctxs *Contexts) UnmarshalJSON(bz []byte) error {
	var single Context
	err := json.Unmarshal(bz, &single)
	if err == nil {
		*ctxs = Contexts{single}
		return nil
	}

	var multiple []Context
	if err := json.Unmarshal(bz, &multiple); err != nil {
		return err
	}
	*ctxs = multiple
	return nil
}

type Context string

type VeriMethodID string

type KeyType string

type VeriMethod struct {
	ID           VeriMethodID `json:"id"`
	Type         KeyType      `json:"type"`
	Controller   DID          `json:"controller"`
	PubKeyBase58 string       `json:"publicKeyBase58"`
}

type Authentication struct {
	VeriMethodID    VeriMethodID
	DedicatedMethod *VeriMethod
}

func (a Authentication) MarshalJSON() ([]byte, error) {
	// if dedicated
	if a.DedicatedMethod != nil {
		return json.Marshal(a.DedicatedMethod)
	}
	// if not dedicated
	return json.Marshal(a.VeriMethodID)
}

func (a *Authentication) UnmarshalJSON(bz []byte) error {
	// if not dedicated
	var veriMethodID VeriMethodID
	err := json.Unmarshal(bz, &veriMethodID)
	if err == nil {
		*a = Authentication{VeriMethodID: veriMethodID, DedicatedMethod: nil}
		return nil
	}

	// if dedicated
	var veriMethod VeriMethod
	if err := json.Unmarshal(bz, &veriMethod); err != nil {
		return err
	}
	*a = Authentication{VeriMethodID: veriMethod.ID, DedicatedMethod: &veriMethod}
	return nil
}

type DIDDocumentWithSeq struct {
	Document DIDDocument `json:"document"`
	Seq      Sequence    `json:"sequence"`
}
