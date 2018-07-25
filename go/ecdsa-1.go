/*
 * Copyright 2015 The mischn Authors
 * This file is part of the mischn library.
 *
**/

// core/types/transaction_signing.go  
func recoverPlain(sighash common.Hash, R, S, Vb *big.Int, homestead bool) (common.Address, error) {  
    V := byte(Vb.Uint64() - 27)  
    if !crypto.ValidateSignatureValues(V, R, S, homestead) {  
        return common.Address{}, ErrInvalidSig  
    }  
// encode signature in uncompressed format  
    r, s := R.Bytes(), S.Bytes()  
    sig := make([]byte, 65)  
    copy(sig[32-len(r):32], r)  
    copy(sig[64-len(s):64], s)  
    sig[64] = V  
// recover the public key from the signature  
    pub, err := crypto.Ecrecover(sighash[:], sig)  
    if err != nil || len(pub) == 0 || pub[0] != 4 {  
        return common.Address{}, err  
    }  
// convert pubKey to Address  
    var addr common.Address  
    copy(addr[:], crypto.Keccak256(pub[1:])[12:])  
    return addr, nil  
}  
