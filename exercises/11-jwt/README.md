# 11 - JWT

## Goal

Implement the two core JWT operations by hand: signing a token with HMAC-SHA256
and verifying one — including detecting tampering and expiration. No external
library: the encoding and the math are the point of this exercise.

## Concepts

- What a JWT is: three base64url segments — header, payload, signature
- HMAC-SHA256 as a shared-secret signature
- Tamper detection: recompute the signature and compare
- The `exp` claim and expiry checks
- Constant-time comparison of signatures

## Syntax Primer

A token is `header.payload.signature`. The signature covers both other
segments, joined by a dot:

```go
signingInput := headerPart + "." + payloadPart
sig := hmac.New(sha256.New, secret)
sig.Write([]byte(signingInput))
signature := base64url(sig.Sum(nil))
```

Claims are a JSON object inside the payload segment:

```go
claims := map[string]any{"sub": "user-1", "exp": now + 3600}
```

## Mental Model

The header says *how* it was signed (`HS256`). The payload carries claims —
assertions about the caller. The signature proves the first two segments were
not modified by anyone who does not know the secret. Verify = split, recompute,
compare, then check the claims you care about (at minimum `exp`).

## Annotated Examples

```go
func expFrom(claims map[string]any) (int64, error) {
	raw, ok := claims["exp"].(float64) // JSON numbers decode as float64
	if !ok {
		return 0, errors.New("missing exp")
	}
	return int64(raw), nil
}
```

## Common Diagnostics

- `signature is invalid` / wrong-secret failures: the secret used to verify
  differs from the signing secret — or you compared signatures with `==`
  instead of a constant-time comparison.
- Token still verifies after tampering: you validated only the payload segment
  without recomputing the signature over the modified input.
- `exp` is never enforced: you decoded but never checked `claims["exp"]`
  against `nowUnix()`.
- `illegal base64 data at input byte`: you used the padded encoding on the
  wrong side — tokens use raw (unpadded) base64url.
- `json: cannot unmarshal number into Go value of type string`: `exp` arrives
  as `float64` after JSON decoding; convert before comparing.

## Exercise

Implement `Verify(secret, token)`:

1. Split the token on `.`; any token that does not split into exactly three
   segments is `ErrMalformed`.
2. Recompute the signature over `parts[0] + "." + parts[1]` with the given
   secret and compare it against `parts[2]` — constant-time, decoded from
   base64url.
3. If the signature does not match, return an error (tampered or wrong secret).
4. Decode the payload segment and parse the claims. Reject a missing `exp`
   and an `exp` in the past with `ErrExpired`.

## Acceptance Criteria

- A token produced by `Sign` verifies and returns its claims.
- `Verify` fails for the wrong secret.
- Changing a single character in the payload makes verification fail.
- Tokens whose `exp` is in the past return `ErrExpired`.

## Hints

- The comparison helper is `hmac.Equal` — it runs in constant time.
- `base64.RawURLEncoding.DecodeString` undoes `encodeJSON`'s encoding.
- Base64 decode can fail: any decode error is a malformed token.
- Claim access: `claims["exp"].(float64)` after JSON decoding — then compare
  with `nowUnix()`.

## Verify

```bash
gofmt -w exercises/11-jwt
go test -tags exercise ./exercises/11-jwt/...
```

The starter ships `Sign` complete, with `Verify` as a stub, so the failing
tests tell you exactly which verification step is missing.

## Reflection Prompts

Why does the signature have to cover the header too? What is the practical
difference between `hmac.Equal` and `==` on `[]byte`? When would `exp` be
insufficient for real authorization?