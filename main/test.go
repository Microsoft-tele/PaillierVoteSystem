package main

//func main() {
//	key := CryptoUtils.GetKeysFromJson()
//	m1 := new(big.Int).SetInt64(1)
//	c1, err := paillier.Encrypt(&key.PublicKey, m1.Bytes())
//	fmt.Println("c1:", c1)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	m11 := new(big.Int).SetInt64(1)
//	c11, err := paillier.Encrypt(&key.PublicKey, m11.Bytes())
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("c11:", c11)
//
//	c2 := paillier.AddCipher(&key.PublicKey, c1, c11)
//	fmt.Println("c2:", c2)
//
//	m111, _ := paillier.Decrypt(key, c2)
//	m111int := new(big.Int).SetBytes(m111).String()
//	fmt.Println("m111int", m111int)
//}
