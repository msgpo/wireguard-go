/* SPDX-License-Identifier: GPL-2.0
 *
 * Copyright (C) 2017-2018 Jason A. Donenfeld <Jason@zx2c4.com>. All Rights Reserved.
 * Copyright (C) 2017-2018 Mathias N. Hall-Andersen <mathias@hall-andersen.dk>.
 */

package main

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/chacha20poly1305"
)

const (
	NoisePublicKeySize  = 32
	NoisePrivateKeySize = 32
)

type (
	NoisePublicKey    [NoisePublicKeySize]byte
	NoisePrivateKey   [NoisePrivateKeySize]byte
	NoiseSymmetricKey [chacha20poly1305.KeySize]byte
	NoiseNonce        uint64 // padded to 12-bytes
)

func loadExactHex(dst []byte, src string) error {
	slice, err := hex.DecodeString(src)
	if err != nil {
		return err
	}
	if len(slice) != len(dst) {
		return errors.New("hex string does not fit the slice")
	}
	copy(dst, slice)
	return nil
}

func (key NoisePrivateKey) IsZero() bool {
	var zero NoisePrivateKey
	return key.Equals(zero)
}

func (key NoisePrivateKey) Equals(tar NoisePrivateKey) bool {
	return subtle.ConstantTimeCompare(key[:], tar[:]) == 1
}

func (key *NoisePrivateKey) FromHex(src string) error {
	return loadExactHex(key[:], src)
}

func (key NoisePrivateKey) ToHex() string {
	return hex.EncodeToString(key[:])
}

func (key *NoisePublicKey) FromHex(src string) error {
	return loadExactHex(key[:], src)
}

func (key NoisePublicKey) ToHex() string {
	return hex.EncodeToString(key[:])
}

func (key NoisePublicKey) IsZero() bool {
	var zero NoisePublicKey
	return key.Equals(zero)
}

func (key NoisePublicKey) Equals(tar NoisePublicKey) bool {
	return subtle.ConstantTimeCompare(key[:], tar[:]) == 1
}

func (key *NoiseSymmetricKey) FromHex(src string) error {
	return loadExactHex(key[:], src)
}

func (key NoiseSymmetricKey) ToHex() string {
	return hex.EncodeToString(key[:])
}
