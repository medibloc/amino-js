package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	exports := js.Module.Get("exports")

	// @formatter:off
	// Bech32
	exports.Set("encodeBech32", EncodeBech32)
	exports.Set("decodeBech32", DecodeBech32)

	// Basic encoding
	exports.Set("encodeByte",      EncodeByte)
	exports.Set("encodeByteSlice", EncodeByteSlice)
	exports.Set("encodeInt8",      EncodeInt8)
	exports.Set("encodeInt16",     EncodeInt16)
	exports.Set("encodeInt32",     EncodeInt32)
	exports.Set("encodeInt64",     EncodeInt64)
	exports.Set("encodeVarint",    EncodeVarint)
	exports.Set("encodeUint8",     EncodeUint8)
	exports.Set("encodeUint16",    EncodeUint16)
	exports.Set("encodeUint32",    EncodeUint32)
	exports.Set("encodeUint64",    EncodeUint64)
	exports.Set("encodeUvarint",   EncodeUvarint)
	exports.Set("encodeFloat32",   EncodeFloat32)
	exports.Set("encodeFloat64",   EncodeFloat64)
	exports.Set("encodeBool",      EncodeBool)
	exports.Set("encodeString",    EncodeString)
	exports.Set("encodeTime",      EncodeTime)

	// Basic decoding
	exports.Set("decodeByte",      DecodeByte)
	exports.Set("decodeByteSlice", DecodeByteSlice)
	exports.Set("decodeInt8",      DecodeInt8)
	exports.Set("decodeInt16",     DecodeInt16)
	exports.Set("decodeInt32",     DecodeInt32)
	exports.Set("decodeInt64",     DecodeInt64)
	exports.Set("decodeVarint",    DecodeVarint)
	exports.Set("decodeUint8",     DecodeUint8)
	exports.Set("decodeUint16",    DecodeUint16)
	exports.Set("decodeUint32",    DecodeUint32)
	exports.Set("decodeUint64",    DecodeUint64)
	exports.Set("decodeUvarint",   DecodeUvarint)
	exports.Set("decodeFloat32",   DecodeFloat32)
	exports.Set("decodeFloat64",   DecodeFloat64)
	exports.Set("decodeBool",      DecodeBool)
	exports.Set("decodeString",    DecodeString)
	exports.Set("decodeTime",      DecodeTime)

	// Meta
	exports.Set("decodeDisambPrefixBytes", DecodeDisambPrefixBytes)
	exports.Set("nameToDisfix",            NameToDisfix)
	exports.Set("byteSliceSize",           ByteSliceSize)
	exports.Set("uvarintSize",             UvarintSize)
	exports.Set("varintSize",              VarintSize)

	// Typed encoding
	exports.Set("encodeMultiStoreProofOp",              EncodeMultiStoreProofOp)
	exports.Set("encodeIAVLAbsenceOp",                  EncodeIAVLAbsenceOp)
	exports.Set("encodeIAVLValueOp",                    EncodeIAVLValueOp)
	exports.Set("encodePrivKeyLedgerSecp256k1",         EncodePrivKeyLedgerSecp256k1)
	exports.Set("encodeInfo",                           EncodeInfo)
	exports.Set("encodeBIP44Params",                    EncodeBIP44Params)
	exports.Set("encodeLocalInfo",                      EncodeLocalInfo)
	exports.Set("encodeLedgerInfo",                     EncodeLedgerInfo)
	exports.Set("encodeOfflineInfo",                    EncodeOfflineInfo)
	exports.Set("encodeMultiInfo",                      EncodeMultiInfo)
	exports.Set("encodeMsg",                            EncodeMsg)
	exports.Set("encodeTx",                             EncodeTx)
	exports.Set("encodeAccount",                        EncodeAccount)
	exports.Set("encodeVestingAccount",                 EncodeVestingAccount)
	exports.Set("encodeBaseAccount",                    EncodeBaseAccount)
	exports.Set("encodeBaseVestingAccount",             EncodeBaseVestingAccount)
	exports.Set("encodeContinuousVestingAccount",       EncodeContinuousVestingAccount)
	exports.Set("encodeDelayedVestingAccount",          EncodeDelayedVestingAccount)
	exports.Set("encodeStdTx",                          EncodeStdTx)
	exports.Set("encodeMsgSend",                        EncodeMsgSend)
	exports.Set("encodeMsgMultiSend",                   EncodeMsgMultiSend)
	exports.Set("encodeMsgVerifyInvariant",             EncodeMsgVerifyInvariant)
	exports.Set("encodeMsgWithdrawDelegatorReward",     EncodeMsgWithdrawDelegatorReward)
	exports.Set("encodeMsgWithdrawValidatorCommission", EncodeMsgWithdrawValidatorCommission)
	exports.Set("encodeMsgSetWithdrawAddress",          EncodeMsgSetWithdrawAddress)
	exports.Set("encodeContent",                        EncodeContent)
	exports.Set("encodeMsgSubmitProposal",              EncodeMsgSubmitProposal)
	exports.Set("encodeMsgDeposit",                     EncodeMsgDeposit)
	exports.Set("encodeMsgVote",                        EncodeMsgVote)
	exports.Set("encodeTextProposal",                   EncodeTextProposal)
	exports.Set("encodeSoftwareUpgradeProposal",        EncodeSoftwareUpgradeProposal)
	exports.Set("encodeMsgIBCTransfer",                 EncodeMsgIBCTransfer)
	exports.Set("encodeMsgIBCReceive",                  EncodeMsgIBCReceive)
	exports.Set("encodeParameterChangeProposal",        EncodeParameterChangeProposal)
	exports.Set("encodeMsgUnjail",                      EncodeMsgUnjail)
	exports.Set("encodeMsgCreateValidator",             EncodeMsgCreateValidator)
	exports.Set("encodeMsgEditValidator",               EncodeMsgEditValidator)
	exports.Set("encodeMsgDelegate",                    EncodeMsgDelegate)
	exports.Set("encodeMsgUndelegate",                  EncodeMsgUndelegate)
	exports.Set("encodeMsgBeginRedelegate",             EncodeMsgBeginRedelegate)
	exports.Set("encodeBlockchainMessage",              EncodeBlockchainMessage)
	exports.Set("encodeBcBlockRequestMessage",          EncodeBcBlockRequestMessage)
	exports.Set("encodeBcBlockResponseMessage",         EncodeBcBlockResponseMessage)
	exports.Set("encodeBcNoBlockResponseMessage",       EncodeBcNoBlockResponseMessage)
	exports.Set("encodeBcStatusResponseMessage",        EncodeBcStatusResponseMessage)
	exports.Set("encodeBcStatusRequestMessage",         EncodeBcStatusRequestMessage)
	exports.Set("encodeConsensusMessage",               EncodeConsensusMessage)
	exports.Set("encodeNewRoundStepMessage",            EncodeNewRoundStepMessage)
	exports.Set("encodeNewValidBlockMessage",           EncodeNewValidBlockMessage)
	exports.Set("encodeProposalMessage",                EncodeProposalMessage)
	exports.Set("encodeProposalPOLMessage",             EncodeProposalPOLMessage)
	exports.Set("encodeBlockPartMessage",               EncodeBlockPartMessage)
	exports.Set("encodeVoteMessage",                    EncodeVoteMessage)
	exports.Set("encodeHasVoteMessage",                 EncodeHasVoteMessage)
	exports.Set("encodeVoteSetMaj23Message",            EncodeVoteSetMaj23Message)
	exports.Set("encodeVoteSetBitsMessage",             EncodeVoteSetBitsMessage)
	exports.Set("encodeWALMessage",                     EncodeWALMessage)
	exports.Set("encodeMsgInfo",                        EncodeMsgInfo)
	exports.Set("encodeTimeoutInfo",                    EncodeTimeoutInfo)
	exports.Set("encodeEndHeightMessage",               EncodeEndHeightMessage)
	exports.Set("encodePubKey",                         EncodePubKey)
	exports.Set("encodePrivKey",                        EncodePrivKey)
	exports.Set("encodePubKeyEd25519",                  EncodePubKeyEd25519)
	exports.Set("encodePrivKeyEd25519",                 EncodePrivKeyEd25519)
	exports.Set("encodePubKeySecp256k1",                EncodePubKeySecp256k1)
	exports.Set("encodePrivKeySecp256k1",               EncodePrivKeySecp256k1)
	exports.Set("encodePubKeyMultisigThreshold",        EncodePubKeyMultisigThreshold)
	exports.Set("encodeEvidenceMessage",                EncodeEvidenceMessage)
	exports.Set("encodeEvidenceListMessage",            EncodeEvidenceListMessage)
	exports.Set("encodeMempoolMessage",                 EncodeMempoolMessage)
	exports.Set("encodeTxMessage",                      EncodeTxMessage)
	exports.Set("encodePacket",                         EncodePacket)
	exports.Set("encodePacketPing",                     EncodePacketPing)
	exports.Set("encodePacketPong",                     EncodePacketPong)
	exports.Set("encodePacketMsg",                      EncodePacketMsg)
	exports.Set("encodePexMessage",                     EncodePexMessage)
	exports.Set("encodePexRequestMessage",              EncodePexRequestMessage)
	exports.Set("encodePexAddrsMessage",                EncodePexAddrsMessage)
	exports.Set("encodeRemoteSignerMsg",                EncodeRemoteSignerMsg)
	exports.Set("encodePubKeyRequest",                  EncodePubKeyRequest)
	exports.Set("encodePubKeyResponse",                 EncodePubKeyResponse)
	exports.Set("encodeSignVoteRequest",                EncodeSignVoteRequest)
	exports.Set("encodeSignedVoteResponse",             EncodeSignedVoteResponse)
	exports.Set("encodeSignProposalRequest",            EncodeSignProposalRequest)
	exports.Set("encodeSignedProposalResponse",         EncodeSignedProposalResponse)
	exports.Set("encodePingRequest",                    EncodePingRequest)
	exports.Set("encodePingResponse",                   EncodePingResponse)
	exports.Set("encodeTMEventData",                    EncodeTMEventData)
	exports.Set("encodeEventDataNewBlock",              EncodeEventDataNewBlock)
	exports.Set("encodeEventDataNewBlockHeader",        EncodeEventDataNewBlockHeader)
	exports.Set("encodeEventDataTx",                    EncodeEventDataTx)
	exports.Set("encodeEventDataRoundState",            EncodeEventDataRoundState)
	exports.Set("encodeEventDataNewRound",              EncodeEventDataNewRound)
	exports.Set("encodeEventDataCompleteProposal",      EncodeEventDataCompleteProposal)
	exports.Set("encodeEventDataVote",                  EncodeEventDataVote)
	exports.Set("encodeEventDataValidatorSetUpdates",   EncodeEventDataValidatorSetUpdates)
	exports.Set("encodeEventDataString",                EncodeEventDataString)
	exports.Set("encodeEvidence",                       EncodeEvidence)
	exports.Set("encodeDuplicateVoteEvidence",          EncodeDuplicateVoteEvidence)
	exports.Set("encodeMockGoodEvidence",               EncodeMockGoodEvidence)
	exports.Set("encodeMockRandomGoodEvidence",         EncodeMockRandomGoodEvidence)
	exports.Set("encodeMockBadEvidence",                EncodeMockBadEvidence)

	// Typed decoding
	exports.Set("decodeMultiStoreProofOp",              DecodeMultiStoreProofOp)
	exports.Set("decodeIAVLAbsenceOp",                  DecodeIAVLAbsenceOp)
	exports.Set("decodeIAVLValueOp",                    DecodeIAVLValueOp)
	exports.Set("decodePrivKeyLedgerSecp256k1",         DecodePrivKeyLedgerSecp256k1)
	exports.Set("decodeInfo",                           DecodeInfo)
	exports.Set("decodeBIP44Params",                    DecodeBIP44Params)
	exports.Set("decodeLocalInfo",                      DecodeLocalInfo)
	exports.Set("decodeLedgerInfo",                     DecodeLedgerInfo)
	exports.Set("decodeOfflineInfo",                    DecodeOfflineInfo)
	exports.Set("decodeMultiInfo",                      DecodeMultiInfo)
	exports.Set("decodeMsg",                            DecodeMsg)
	exports.Set("decodeTx",                             DecodeTx)
	exports.Set("decodeAccount",                        DecodeAccount)
	exports.Set("decodeVestingAccount",                 DecodeVestingAccount)
	exports.Set("decodeBaseAccount",                    DecodeBaseAccount)
	exports.Set("decodeBaseVestingAccount",             DecodeBaseVestingAccount)
	exports.Set("decodeContinuousVestingAccount",       DecodeContinuousVestingAccount)
	exports.Set("decodeDelayedVestingAccount",          DecodeDelayedVestingAccount)
	exports.Set("decodeStdTx",                          DecodeStdTx)
	exports.Set("decodeMsgSend",                        DecodeMsgSend)
	exports.Set("decodeMsgMultiSend",                   DecodeMsgMultiSend)
	exports.Set("decodeMsgVerifyInvariant",             DecodeMsgVerifyInvariant)
	exports.Set("decodeMsgWithdrawDelegatorReward",     DecodeMsgWithdrawDelegatorReward)
	exports.Set("decodeMsgWithdrawValidatorCommission", DecodeMsgWithdrawValidatorCommission)
	exports.Set("decodeMsgSetWithdrawAddress",          DecodeMsgSetWithdrawAddress)
	exports.Set("decodeContent",                        DecodeContent)
	exports.Set("decodeMsgSubmitProposal",              DecodeMsgSubmitProposal)
	exports.Set("decodeMsgDeposit",                     DecodeMsgDeposit)
	exports.Set("decodeMsgVote",                        DecodeMsgVote)
	exports.Set("decodeTextProposal",                   DecodeTextProposal)
	exports.Set("decodeSoftwareUpgradeProposal",        DecodeSoftwareUpgradeProposal)
	exports.Set("decodeMsgIBCTransfer",                 DecodeMsgIBCTransfer)
	exports.Set("decodeMsgIBCReceive",                  DecodeMsgIBCReceive)
	exports.Set("decodeParameterChangeProposal",        DecodeParameterChangeProposal)
	exports.Set("decodeMsgUnjail",                      DecodeMsgUnjail)
	exports.Set("decodeMsgCreateValidator",             DecodeMsgCreateValidator)
	exports.Set("decodeMsgEditValidator",               DecodeMsgEditValidator)
	exports.Set("decodeMsgDelegate",                    DecodeMsgDelegate)
	exports.Set("decodeMsgUndelegate",                  DecodeMsgUndelegate)
	exports.Set("decodeMsgBeginRedelegate",             DecodeMsgBeginRedelegate)
	exports.Set("decodeBlockchainMessage",              DecodeBlockchainMessage)
	exports.Set("decodeBcBlockRequestMessage",          DecodeBcBlockRequestMessage)
	exports.Set("decodeBcBlockResponseMessage",         DecodeBcBlockResponseMessage)
	exports.Set("decodeBcNoBlockResponseMessage",       DecodeBcNoBlockResponseMessage)
	exports.Set("decodeBcStatusResponseMessage",        DecodeBcStatusResponseMessage)
	exports.Set("decodeBcStatusRequestMessage",         DecodeBcStatusRequestMessage)
	exports.Set("decodeConsensusMessage",               DecodeConsensusMessage)
	exports.Set("decodeNewRoundStepMessage",            DecodeNewRoundStepMessage)
	exports.Set("decodeNewValidBlockMessage",           DecodeNewValidBlockMessage)
	exports.Set("decodeProposalMessage",                DecodeProposalMessage)
	exports.Set("decodeProposalPOLMessage",             DecodeProposalPOLMessage)
	exports.Set("decodeBlockPartMessage",               DecodeBlockPartMessage)
	exports.Set("decodeVoteMessage",                    DecodeVoteMessage)
	exports.Set("decodeHasVoteMessage",                 DecodeHasVoteMessage)
	exports.Set("decodeVoteSetMaj23Message",            DecodeVoteSetMaj23Message)
	exports.Set("decodeVoteSetBitsMessage",             DecodeVoteSetBitsMessage)
	exports.Set("decodeWALMessage",                     DecodeWALMessage)
	exports.Set("decodeMsgInfo",                        DecodeMsgInfo)
	exports.Set("decodeTimeoutInfo",                    DecodeTimeoutInfo)
	exports.Set("decodeEndHeightMessage",               DecodeEndHeightMessage)
	exports.Set("decodePubKey",                         DecodePubKey)
	exports.Set("decodePrivKey",                        DecodePrivKey)
	exports.Set("decodePubKeyEd25519",                  DecodePubKeyEd25519)
	exports.Set("decodePrivKeyEd25519",                 DecodePrivKeyEd25519)
	exports.Set("decodePubKeySecp256k1",                DecodePubKeySecp256k1)
	exports.Set("decodePrivKeySecp256k1",               DecodePrivKeySecp256k1)
	exports.Set("decodePubKeyMultisigThreshold",        DecodePubKeyMultisigThreshold)
	exports.Set("decodeEvidenceMessage",                DecodeEvidenceMessage)
	exports.Set("decodeEvidenceListMessage",            DecodeEvidenceListMessage)
	exports.Set("decodeMempoolMessage",                 DecodeMempoolMessage)
	exports.Set("decodeTxMessage",                      DecodeTxMessage)
	exports.Set("decodePacket",                         DecodePacket)
	exports.Set("decodePacketPing",                     DecodePacketPing)
	exports.Set("decodePacketPong",                     DecodePacketPong)
	exports.Set("decodePacketMsg",                      DecodePacketMsg)
	exports.Set("decodePexMessage",                     DecodePexMessage)
	exports.Set("decodePexRequestMessage",              DecodePexRequestMessage)
	exports.Set("decodePexAddrsMessage",                DecodePexAddrsMessage)
	exports.Set("decodeRemoteSignerMsg",                DecodeRemoteSignerMsg)
	exports.Set("decodePubKeyRequest",                  DecodePubKeyRequest)
	exports.Set("decodePubKeyResponse",                 DecodePubKeyResponse)
	exports.Set("decodeSignVoteRequest",                DecodeSignVoteRequest)
	exports.Set("decodeSignedVoteResponse",             DecodeSignedVoteResponse)
	exports.Set("decodeSignProposalRequest",            DecodeSignProposalRequest)
	exports.Set("decodeSignedProposalResponse",         DecodeSignedProposalResponse)
	exports.Set("decodePingRequest",                    DecodePingRequest)
	exports.Set("decodePingResponse",                   DecodePingResponse)
	exports.Set("decodeTMEventData",                    DecodeTMEventData)
	exports.Set("decodeEventDataNewBlock",              DecodeEventDataNewBlock)
	exports.Set("decodeEventDataNewBlockHeader",        DecodeEventDataNewBlockHeader)
	exports.Set("decodeEventDataTx",                    DecodeEventDataTx)
	exports.Set("decodeEventDataRoundState",            DecodeEventDataRoundState)
	exports.Set("decodeEventDataNewRound",              DecodeEventDataNewRound)
	exports.Set("decodeEventDataCompleteProposal",      DecodeEventDataCompleteProposal)
	exports.Set("decodeEventDataVote",                  DecodeEventDataVote)
	exports.Set("decodeEventDataValidatorSetUpdates",   DecodeEventDataValidatorSetUpdates)
	exports.Set("decodeEventDataString",                DecodeEventDataString)
	exports.Set("decodeEvidence",                       DecodeEvidence)
	exports.Set("decodeDuplicateVoteEvidence",          DecodeDuplicateVoteEvidence)
	exports.Set("decodeMockGoodEvidence",               DecodeMockGoodEvidence)
	exports.Set("decodeMockRandomGoodEvidence",         DecodeMockRandomGoodEvidence)
	exports.Set("decodeMockBadEvidence",                DecodeMockBadEvidence)
	// @formatter:on
}