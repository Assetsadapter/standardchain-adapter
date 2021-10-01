package std_addrdec

import (
	"encoding/hex"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {

	pub, _ := hex.DecodeString("032144da84e7c0037014be1332617ceec15d3561dc209a1d984bf74677a41a63d0")
	addr, _ := Default.AddressEncode(pub)
	t.Logf("addr: %s", addr)
	//	0x5f75ef82839fdc491f15816fce5184f9b65fe0f8
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	addr := "0x5f75ef82839fdc491f15816fce5184f9b65fe0f8"
	hash, _ := Default.AddressDecode(addr)
	t.Logf("hash: %s", hex.EncodeToString(hash))
}

func TestAddressDecoder_checkSum(t *testing.T) {

	addr := "0x855e6f3023e60b46e2a845fe00a233a82f772400"
	t.Logf("checkSumAddr is : %s\n", ChecksumAddmress(addr))
}

// 0x ===>sd
func CustomAddressEncode(address string) string {
	return strings.Replace(address, "0x", "sd", 1)
}

// sd  ====>sd
func CustomAddressDecode(address string) string {
	return strings.Replace(address, "sd", "0x", 1)
}

func TestAddressDecoder_addressCustom(t *testing.T) {
	// wm := testNewWalletManager()
	stAddr := "sd855e6f3023E60B46e2A845Fe00A233a82f772400"
	ethAddr := "0x855e6f3023E60B46e2A845Fe00A233a82f772400"
	deAddr := CustomAddressDecode(stAddr)
	assert.Equal(t, ethAddr, deAddr, "the two address must the same")

	enAddr := CustomAddressEncode(ethAddr)
	assert.Equal(t, enAddr, stAddr, "the two address must the same")

}
