package extensions

import (
	aol "github.com/cosmos/amino-js/go/extensions/panacea-core/x/aol/types"
	did "github.com/cosmos/amino-js/go/extensions/panacea-core/x/did/types"
	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	// AOL
	codec.RegisterConcrete(aol.MsgCreateTopic{}, aolMsgCreateTopic, nil)
	codec.RegisterConcrete(aol.MsgAddWriter{}, aolMsgAddWriter, nil)
	codec.RegisterConcrete(aol.MsgDeleteWriter{}, aolMsgDeleteWriter, nil)
	codec.RegisterConcrete(aol.MsgAddRecord{}, aolMsgAddRecord, nil)

	// DID
	codec.RegisterConcrete(did.MsgCreateDID{}, didMsgCreateDID, nil)
	codec.RegisterConcrete(did.MsgUpdateDID{}, didMsgUpdateDID, nil)
	codec.RegisterConcrete(did.MsgDeactivateDID{}, didMsgDeactivateDID, nil)
}
