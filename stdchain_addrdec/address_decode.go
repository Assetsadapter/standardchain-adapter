package std_addrdec

import (
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/v2/openwallet"
	"golang.org/x/crypto/sha3"
)

var (
	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
}

//NewAddressDecoder 地址解析器
func NewAddressDecoderV2() *AddressDecoderV2 {
	decoder := AddressDecoderV2{}
	return &decoder
}
func ChecksumAddmress(address string) string {
	address = strings.ToLower(address)
	address = strings.Replace(address, "0x", "", 1)
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(address))
	hash := sha.Sum(nil)
	hashstr := hex.EncodeToString(hash)
	result := []string{"0x"}
	for i, v := range address {
		res, _ := strconv.ParseInt(string(hashstr[i]), 16, 64)
		if res > 7 {
			result = append(result, strings.ToUpper(string(v)))
			continue
		}
		result = append(result, string(v))
	}
	return strings.Join(result, "")
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {
	addr = strings.TrimPrefix(addr, "0x")
	decodeAddr, err := hex.DecodeString(addr)
	if err != nil {
		return nil, err
	}
	return decodeAddr, err

}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	if len(hash) != 32 {
		//公钥hash处理
		publicKey := owcrypt.PointDecompress(hash, owcrypt.ECC_CURVE_SECP256K1)
		hash = owcrypt.Hash(publicKey[1:len(publicKey)], 0, owcrypt.HASH_ALG_KECCAK256)
	}

	//地址添加0x前缀
	address := "0x" + hex.EncodeToString(hash[12:])
	checkSumAddr := ChecksumAddmress(address)
	checkSumAddr = strings.Replace(checkSumAddr, "0x", "sd", 1)
	return checkSumAddr, nil
}

// AddressVerify 地址校验
func (dec *AddressDecoderV2) AddressVerify(address string, opts ...interface{}) bool {
	if address == "" {
		return false
	}

	if strings.Index(address, "sd") != 0 {
		return false
	}

	addrByte, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	if len(addrByte) != 20 {
		return false
	}

	return true
}
