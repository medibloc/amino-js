package src

import (
	rootmulti "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/store/rootmulti"
	iavl "github.com/cosmos/amino-js/go/lib/tendermint/iavl"

	crypto "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/crypto"
	keys "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/crypto/keys"
	hd "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/crypto/keys/hd"

	types "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"

	auth "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/bank"
	crisis "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/crisis"
	distribution "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/gov/types"
	ibc "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/ibc"
	params "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/params/types"
	slashing "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/slashing"
	staking "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/staking/types"

	tm_crypto "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto"
	tm_ed25519 "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto/ed25519"
	tm_multisig "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto/multisig"
	tm_secp256k1 "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto/secp256k1"

	tm_blockchain "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/blockchain"
	tm_consensus "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/consensus"
	tm_evidence "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/evidence"
	tm_mempool "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/mempool"
	tm_conn "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/p2p/conn"
	tm_pex "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/p2p/pex"
	tm_privval "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/privval"
	tm_types "github.com/cosmos/amino-js/go/lib/tendermint/tendermint/types"
)

func decodeMultiStoreProofOp(bz []byte, lengthPrefixed bool) (o rootmulti.MultiStoreProofOp, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeIAVLAbsenceOp(bz []byte, lengthPrefixed bool) (o iavl.IAVLAbsenceOp, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeIAVLValueOp(bz []byte, lengthPrefixed bool) (o iavl.IAVLValueOp, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePrivKeyLedgerSecp256k1(bz []byte, lengthPrefixed bool) (o crypto.PrivKeyLedgerSecp256k1, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeInfo(bz []byte, lengthPrefixed bool) (o keys.Info, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBIP44Params(bz []byte, lengthPrefixed bool) (o hd.BIP44Params, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeLocalInfo(bz []byte, lengthPrefixed bool) (o keys.LocalInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeLedgerInfo(bz []byte, lengthPrefixed bool) (o keys.LedgerInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeOfflineInfo(bz []byte, lengthPrefixed bool) (o keys.OfflineInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMultiInfo(bz []byte, lengthPrefixed bool) (o keys.MultiInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsg(bz []byte, lengthPrefixed bool) (o types.Msg, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeTx(bz []byte, lengthPrefixed bool) (o types.Tx, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeAccount(bz []byte, lengthPrefixed bool) (o auth.Account, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeVestingAccount(bz []byte, lengthPrefixed bool) (o auth.VestingAccount, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBaseAccount(bz []byte, lengthPrefixed bool) (o auth.BaseAccount, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBaseVestingAccount(bz []byte, lengthPrefixed bool) (o auth.BaseVestingAccount, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeContinuousVestingAccount(bz []byte, lengthPrefixed bool) (o auth.ContinuousVestingAccount, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeDelayedVestingAccount(bz []byte, lengthPrefixed bool) (o auth.DelayedVestingAccount, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeStdTx(bz []byte, lengthPrefixed bool) (o auth.StdTx, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgSend(bz []byte, lengthPrefixed bool) (o bank.MsgSend, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgMultiSend(bz []byte, lengthPrefixed bool) (o bank.MsgMultiSend, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgVerifyInvariant(bz []byte, lengthPrefixed bool) (o crisis.MsgVerifyInvariant, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgWithdrawDelegatorReward(bz []byte, lengthPrefixed bool) (o distribution.MsgWithdrawDelegatorReward, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgWithdrawValidatorCommission(bz []byte, lengthPrefixed bool) (o distribution.MsgWithdrawValidatorCommission, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgSetWithdrawAddress(bz []byte, lengthPrefixed bool) (o distribution.MsgSetWithdrawAddress, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeContent(bz []byte, lengthPrefixed bool) (o gov.Content, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgSubmitProposal(bz []byte, lengthPrefixed bool) (o gov.MsgSubmitProposal, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgDeposit(bz []byte, lengthPrefixed bool) (o gov.MsgDeposit, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgVote(bz []byte, lengthPrefixed bool) (o gov.MsgVote, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeTextProposal(bz []byte, lengthPrefixed bool) (o gov.TextProposal, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeSoftwareUpgradeProposal(bz []byte, lengthPrefixed bool) (o gov.SoftwareUpgradeProposal, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgIBCTransfer(bz []byte, lengthPrefixed bool) (o ibc.MsgIBCTransfer, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgIBCReceive(bz []byte, lengthPrefixed bool) (o ibc.MsgIBCReceive, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeParameterChangeProposal(bz []byte, lengthPrefixed bool) (o params.ParameterChangeProposal, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgUnjail(bz []byte, lengthPrefixed bool) (o slashing.MsgUnjail, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgCreateValidator(bz []byte, lengthPrefixed bool) (o staking.MsgCreateValidator, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgEditValidator(bz []byte, lengthPrefixed bool) (o staking.MsgEditValidator, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgDelegate(bz []byte, lengthPrefixed bool) (o staking.MsgDelegate, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgUndelegate(bz []byte, lengthPrefixed bool) (o staking.MsgUndelegate, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgBeginRedelegate(bz []byte, lengthPrefixed bool) (o staking.MsgBeginRedelegate, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBlockchainMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BlockchainMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBcBlockRequestMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BcBlockRequestMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBcBlockResponseMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BcBlockResponseMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBcNoBlockResponseMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BcNoBlockResponseMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBcStatusResponseMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BcStatusResponseMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBcStatusRequestMessage(bz []byte, lengthPrefixed bool) (o tm_blockchain.BcStatusRequestMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeConsensusMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.ConsensusMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeNewRoundStepMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.NewRoundStepMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeNewValidBlockMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.NewValidBlockMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeProposalMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.ProposalMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeProposalPOLMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.ProposalPOLMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeBlockPartMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.BlockPartMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeVoteMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.VoteMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeHasVoteMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.HasVoteMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeVoteSetMaj23Message(bz []byte, lengthPrefixed bool) (o tm_consensus.VoteSetMaj23Message, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeVoteSetBitsMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.VoteSetBitsMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeWALMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.WALMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMsgInfo(bz []byte, lengthPrefixed bool) (o tm_consensus.MsgInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeTimeoutInfo(bz []byte, lengthPrefixed bool) (o tm_consensus.TimeoutInfo, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEndHeightMessage(bz []byte, lengthPrefixed bool) (o tm_consensus.EndHeightMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKey(bz []byte, lengthPrefixed bool) (o tm_crypto.PubKey, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePrivKey(bz []byte, lengthPrefixed bool) (o tm_crypto.PrivKey, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKeyEd25519(bz []byte, lengthPrefixed bool) (o tm_ed25519.PubKeyEd25519, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePrivKeyEd25519(bz []byte, lengthPrefixed bool) (o tm_ed25519.PrivKeyEd25519, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKeySecp256k1(bz []byte, lengthPrefixed bool) (o tm_secp256k1.PubKeySecp256k1, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePrivKeySecp256k1(bz []byte, lengthPrefixed bool) (o tm_secp256k1.PrivKeySecp256k1, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKeyMultisigThreshold(bz []byte, lengthPrefixed bool) (o tm_multisig.PubKeyMultisigThreshold, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEvidenceMessage(bz []byte, lengthPrefixed bool) (o tm_evidence.EvidenceMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEvidenceListMessage(bz []byte, lengthPrefixed bool) (o tm_evidence.EvidenceListMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMempoolMessage(bz []byte, lengthPrefixed bool) (o tm_mempool.MempoolMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeTxMessage(bz []byte, lengthPrefixed bool) (o tm_mempool.TxMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePacket(bz []byte, lengthPrefixed bool) (o tm_conn.Packet, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePacketPing(bz []byte, lengthPrefixed bool) (o tm_conn.PacketPing, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePacketPong(bz []byte, lengthPrefixed bool) (o tm_conn.PacketPong, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePacketMsg(bz []byte, lengthPrefixed bool) (o tm_conn.PacketMsg, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePexMessage(bz []byte, lengthPrefixed bool) (o tm_pex.PexMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePexRequestMessage(bz []byte, lengthPrefixed bool) (o tm_pex.PexRequestMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePexAddrsMessage(bz []byte, lengthPrefixed bool) (o tm_pex.PexAddrsMessage, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeRemoteSignerMsg(bz []byte, lengthPrefixed bool) (o tm_privval.RemoteSignerMsg, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKeyRequest(bz []byte, lengthPrefixed bool) (o tm_privval.PubKeyRequest, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePubKeyResponse(bz []byte, lengthPrefixed bool) (o tm_privval.PubKeyResponse, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeSignVoteRequest(bz []byte, lengthPrefixed bool) (o tm_privval.SignVoteRequest, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeSignedVoteResponse(bz []byte, lengthPrefixed bool) (o tm_privval.SignedVoteResponse, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeSignProposalRequest(bz []byte, lengthPrefixed bool) (o tm_privval.SignProposalRequest, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeSignedProposalResponse(bz []byte, lengthPrefixed bool) (o tm_privval.SignedProposalResponse, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePingRequest(bz []byte, lengthPrefixed bool) (o tm_privval.PingRequest, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodePingResponse(bz []byte, lengthPrefixed bool) (o tm_privval.PingResponse, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeTMEventData(bz []byte, lengthPrefixed bool) (o tm_types.TMEventData, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataNewBlock(bz []byte, lengthPrefixed bool) (o tm_types.EventDataNewBlock, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataNewBlockHeader(bz []byte, lengthPrefixed bool) (o tm_types.EventDataNewBlockHeader, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataTx(bz []byte, lengthPrefixed bool) (o tm_types.EventDataTx, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataRoundState(bz []byte, lengthPrefixed bool) (o tm_types.EventDataRoundState, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataNewRound(bz []byte, lengthPrefixed bool) (o tm_types.EventDataNewRound, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataCompleteProposal(bz []byte, lengthPrefixed bool) (o tm_types.EventDataCompleteProposal, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataVote(bz []byte, lengthPrefixed bool) (o tm_types.EventDataVote, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataValidatorSetUpdates(bz []byte, lengthPrefixed bool) (o tm_types.EventDataValidatorSetUpdates, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEventDataString(bz []byte, lengthPrefixed bool) (o tm_types.EventDataString, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeEvidence(bz []byte, lengthPrefixed bool) (o tm_types.Evidence, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeDuplicateVoteEvidence(bz []byte, lengthPrefixed bool) (o tm_types.DuplicateVoteEvidence, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMockGoodEvidence(bz []byte, lengthPrefixed bool) (o tm_types.MockGoodEvidence, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMockRandomGoodEvidence(bz []byte, lengthPrefixed bool) (o tm_types.MockRandomGoodEvidence, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}

func decodeMockBadEvidence(bz []byte, lengthPrefixed bool) (o tm_types.MockBadEvidence, err error) {
	if lengthPrefixed {
		err = codec.UnmarshalBinaryLengthPrefixed(bz, &o)
	} else {
		err = codec.UnmarshalBinaryBare(bz, &o)
	}
	return
}