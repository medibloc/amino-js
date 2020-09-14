import { base64ToBytes, bytesToBase64 } from '@tendermint/belt';
import * as Amino from '../';

/*
 * NOTE:
 * To generate a Tx and encode it, use the following commands. Please don't encode Tx using this project codes. It doesn't test anything.
 * 
 * panaceacli tx did ... --generate-only > tx.json
 * panaceacli tx sign tx.json ... > signed.json
 * panaceacli tx encode signed.json ... > encoded.json
*/

const createDIDTxData = '2ATwYl3uCt8Dpf+YyQopZGlkOnBhbmFjZWE6dGVzdG5ldDpKTVRWS01reWoyTjh6d044WTlKVWcSpQIKHGh0dHBzOi8vd3d3LnczLm9yZy9ucy9kaWQvdjESKWRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnGqcBCi5kaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyNrZXkxEhxTZWNwMjU2azFWZXJpZmljYXRpb25LZXkyMDE4GilkaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyIsaWtDYlhobjhqelBuY25CeEE1N01oRGJrRjdFQ1I2VURaVTlLUTdqckFmTnQiMAouZGlkOnBhbmFjZWE6dGVzdG5ldDpKTVRWS01reWoyTjh6d044WTlKVWcja2V5MRouZGlkOnBhbmFjZWE6dGVzdG5ldDpKTVRWS01reWoyTjh6d044WTlKVWcja2V5MSJAEjwQkRGry4/VVgS9D7griQ+3SPPn25MU3n3kcpfwkuUeKhIixnDSnHdRkgvwV3rFRky/g++gJYr6hc7gF0RTGCoUyM6x8oYJztpiR2E93h0UJvy4kSgSBBDAmgwaagom61rphyEDuqcoEv5RDMTSVWIj1SEdGglFRwOsVlkHr0jEAZKTk9ASQOuZm3EKJXBQJDZrH2Ib8RmvoWJXh5Y7oDrY1QgKDYONEwkXWeVj6BDGSGKGUHz9Mq2CbkAtumu5F1mEZElzsX4=';

const createDIDTx = {
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "did/MsgCreateDID",
        "value": {
          "did": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
          "document": {
            "@context": "https://www.w3.org/ns/did/v1",
            "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
            "verificationMethod": [
              {
                "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1",
                "type": "Secp256k1VerificationKey2018",
                "controller": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
                "publicKeyBase58": "ikCbXhn8jzPncnBxA57MhDbkF7ECR6UDZU9KQ7jrAfNt"
              }
            ],
            "authentication": [
              "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1"
            ]
          },
          "verification_method_id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1",
          "signature": "EjwQkRGry4/VVgS9D7griQ+3SPPn25MU3n3kcpfwkuUeKhIixnDSnHdRkgvwV3rFRky/g++gJYr6hc7gF0RTGA==",
          "from_address": "panacea1er8tru5xp88d5cj8vy7au8g5ym7t3yfgdshcyt"
        }
      }
    ],
    "fee": {
      "amount": null,
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "A7qnKBL+UQzE0lViI9UhHRoJRUcDrFZZB69IxAGSk5PQ"
        },
        "signature": "65mbcQolcFAkNmsfYhvxGa+hYleHljugOtjVCAoNg40TCRdZ5WPoEMZIYoZQfP0yrYJuQC26a7kXWYRkSXOxfg=="
      }
    ],
    "memo": ""
  }
}

const updateDIDTxData = '/QfwYl3uCoQHcR/esAopZGlkOnBhbmFjZWE6dGVzdG5ldDpKTVRWS01reWoyTjh6d044WTlKVWcSygUKHGh0dHBzOi8vd3d3LnczLm9yZy9ucy9kaWQvdjEKHGh0dHBzOi8vdzNpZC5vcmcvc2VjdXJpdHkvdjESKWRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnGqcBCi5kaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyNrZXkxEhxTZWNwMjU2azFWZXJpZmljYXRpb25LZXkyMDE4GilkaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyIsaWtDYlhobjhqelBuY25CeEE1N01oRGJrRjdFQ1I2VURaVTlLUTdqckFmTnQapwEKLmRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnI2tleTISHFNlY3AyNTZrMVZlcmlmaWNhdGlvbktleTIwMTgaKWRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnIixpa0NiWGhuOGp6UG5jbkJ4QTU3TWhEYmtGN0VDUjZVRFpVOUtRN2pyQWZOdCIwCi5kaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyNrZXkxItoBCi5kaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyNrZXkzEqcBCi5kaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyNrZXkzEhxTZWNwMjU2azFWZXJpZmljYXRpb25LZXkyMDE4GilkaWQ6cGFuYWNlYTp0ZXN0bmV0OkpNVFZLTWt5ajJOOHp3TjhZOUpVZyIsaWtDYlhobjhqelBuY25CeEE1N01oRGJrRjdFQ1I2VURaVTlLUTdqckFmTnQaLmRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnI2tleTEiQBI8EJERq8uP1VYEvQ+4K4kPt0jz59uTFN595HKX8JLlHioSIsZw0px3UZIL8Fd6xUZMv4PvoCWK+oXO4BdEUxgqFMjOsfKGCc7aYkdhPd4dFCb8uJEoEgQQwJoMGmoKJuta6YchA7qnKBL+UQzE0lViI9UhHRoJRUcDrFZZB69IxAGSk5PQEkCbWqSV/kewH0Y0BL1QTIBM9ArEgd4OQY4X5Idb1qpHFyitj9rn8/N/g8AWdekiyIxVIUnl6gDE5nwUGvWvIAwf'

const updateDIDTx = {
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "did/MsgUpdateDID",
        "value": {
          "did": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
          "document": {
            "@context": [
              "https://www.w3.org/ns/did/v1",
              "https://w3id.org/security/v1"
            ],
            "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
            "verificationMethod": [
              {
                "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1",
                "type": "Secp256k1VerificationKey2018",
                "controller": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
                "publicKeyBase58": "ikCbXhn8jzPncnBxA57MhDbkF7ECR6UDZU9KQ7jrAfNt"
              },
              {
                "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key2",
                "type": "Secp256k1VerificationKey2018",
                "controller": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
                "publicKeyBase58": "ikCbXhn8jzPncnBxA57MhDbkF7ECR6UDZU9KQ7jrAfNt"
              }
            ],
            "authentication": [
              "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1",
              {
                "id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key3",
                "type": "Secp256k1VerificationKey2018",
                "controller": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
                "publicKeyBase58": "ikCbXhn8jzPncnBxA57MhDbkF7ECR6UDZU9KQ7jrAfNt"
              }
            ]
          },
          "verification_method_id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key1",
          "signature": "EjwQkRGry4/VVgS9D7griQ+3SPPn25MU3n3kcpfwkuUeKhIixnDSnHdRkgvwV3rFRky/g++gJYr6hc7gF0RTGA==",
          "from_address": "panacea1er8tru5xp88d5cj8vy7au8g5ym7t3yfgdshcyt"
        }
      }
    ],
    "fee": {
      "amount": null,
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "A7qnKBL+UQzE0lViI9UhHRoJRUcDrFZZB69IxAGSk5PQ"
        },
        "signature": "m1qklf5HsB9GNAS9UEyATPQKxIHeDkGOF+SHW9aqRxcorY/a5/Pzf4PAFnXpIsiMVSFJ5eoAxOZ8FBr1ryAMHw=="
      }
    ],
    "memo": ""
  }
}

const deactivateDIDTxData = 'sALwYl3uCrcBVPGpQgopZGlkOnBhbmFjZWE6dGVzdG5ldDpKTVRWS01reWoyTjh6d044WTlKVWcSLmRpZDpwYW5hY2VhOnRlc3RuZXQ6Sk1UVktNa3lqMk44endOOFk5SlVnI2tleTIaQBI8EJERq8uP1VYEvQ+4K4kPt0jz59uTFN595HKX8JLlHioSIsZw0px3UZIL8Fd6xUZMv4PvoCWK+oXO4BdEUxgiFMjOsfKGCc7aYkdhPd4dFCb8uJEoEgQQwJoMGmoKJuta6YchA7qnKBL+UQzE0lViI9UhHRoJRUcDrFZZB69IxAGSk5PQEkAbGqLNOCRrzfMsAdGgeqk5qJXCx4boAtme5S3fvOMSZGDR79un0rgMEIsHebs+IkjkxrbwJjtoODdNdl1A73bJ';

const deactivateDIDTx = {
  "type": "auth/StdTx",
  "value": {
    "msg": [
      {
        "type": "did/MsgDeactivateDID",
        "value": {
          "did": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg",
          "verification_method_id": "did:panacea:testnet:JMTVKMkyj2N8zwN8Y9JUg#key2",
          "signature": "EjwQkRGry4/VVgS9D7griQ+3SPPn25MU3n3kcpfwkuUeKhIixnDSnHdRkgvwV3rFRky/g++gJYr6hc7gF0RTGA==",
          "from_address": "panacea1er8tru5xp88d5cj8vy7au8g5ym7t3yfgdshcyt"
        }
      }
    ],
    "fee": {
      "amount": null,
      "gas": "200000"
    },
    "signatures": [
      {
        "pub_key": {
          "type": "tendermint/PubKeySecp256k1",
          "value": "A7qnKBL+UQzE0lViI9UhHRoJRUcDrFZZB69IxAGSk5PQ"
        },
        "signature": "GxqizTgka83zLAHRoHqpOaiVwseG6ALZnuUt37zjEmRg0e/bp9K4DBCLB3m7PiJI5Ma28CY7aDg3TXZdQO92yQ=="
      }
    ],
    "memo": ""
  }
}


describe('Extensions', () => {
    describe('decoding', () => {
        describe('CreateDID Tx', () => {
            it('decodes bytes', () => {
                const bytes = base64ToBytes(createDIDTxData);
                const value = Amino.unmarshalTx(bytes, true);
                expect(value).toMatchObject(createDIDTx);
            });
        });
        describe('UpdateDID Tx', () => {
            it('decodes bytes', () => {
                const bytes = base64ToBytes(updateDIDTxData);
                const value = Amino.unmarshalTx(bytes, true);
                expect(value).toMatchObject(updateDIDTx);
            });
        });
        describe('DeactivateDID Tx', () => {
            it('decodes bytes', () => {
                const bytes = base64ToBytes(deactivateDIDTxData);
                const value = Amino.unmarshalTx(bytes, true);
                expect(value).toMatchObject(deactivateDIDTx);
            });
        });
    });

    describe('encoding', () => {
        describe('CreateDID Tx', () => {
            it('encodes value', () => {
                const bytes = Amino.marshalTx(createDIDTx, true);
                const data  = bytesToBase64(bytes);
                expect(data).toBe(createDIDTxData);
            });
        });
        describe('UpdateDID Tx', () => {
            it('encodes value', () => {
                const bytes = Amino.marshalTx(updateDIDTx, true);
                const data  = bytesToBase64(bytes);
                expect(data).toBe(updateDIDTxData);
            });
        });
        describe('DeactivateDID Tx', () => {
            it('encodes value', () => {
                const bytes = Amino.marshalTx(deactivateDIDTx, true);
                const data  = bytesToBase64(bytes);
                expect(data).toBe(deactivateDIDTxData);
            });
        });
    });
});
