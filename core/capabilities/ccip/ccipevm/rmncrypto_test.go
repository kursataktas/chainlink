package ccipevm

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/rmn_remote"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_VerifyRmnReportSignatures(t *testing.T) {
	// NOTE: The following test data (public keys, signatures, ...) are shared from the RMN team.

	onchainRmnRemoteAddr := common.HexToAddress("0x7821bcd6944457d17c631157efeb0c621baa76eb")

	rmnHomeContractConfigDigestHex := "0x785936570d1c7422ef30b7da5555ad2f175fa2dd97a2429a2e71d1e07c94e060"
	rmnHomeContractConfigDigest := common.FromHex(rmnHomeContractConfigDigestHex)
	require.Len(t, rmnHomeContractConfigDigest, 32)
	var rmnHomeContractConfigDigest32 [32]byte
	copy(rmnHomeContractConfigDigest32[:], rmnHomeContractConfigDigest)

	rootHex := "0x48e688aefc20a04fdec6b8ff19df358fd532455659dcf529797cda358e9e5205"
	root := common.FromHex(rootHex)
	require.Len(t, root, 32)
	var root32 [32]byte
	copy(root32[:], root)

	onRampAddr := common.HexToAddress("0x6662cb20464f4be557262693bea0409f068397ed")

	destChainEvmID := int64(4083663998511321420)

	reportData := cciptypes.RMNReport{
		ReportVersionDigest:         cciptypes.Bytes32(crypto.Keccak256Hash([]byte("RMN_V1_6_ANY2EVM_REPORT"))),
		DestChainID:                 cciptypes.NewBigIntFromInt64(destChainEvmID),
		DestChainSelector:           5266174733271469989,
		RmnRemoteContractAddress:    common.HexToAddress("0x3d015cec4411357eff4ea5f009a581cc519f75d3").Bytes(),
		OfframpAddress:              common.HexToAddress("0xc5cdb7711a478058023373b8ae9e7421925140f8").Bytes(),
		RmnHomeContractConfigDigest: rmnHomeContractConfigDigest32,
		LaneUpdates: []cciptypes.RMNLaneUpdate{
			{
				SourceChainSelector: 8258882951688608272,
				OnRampAddress:       onRampAddr.Bytes(),
				MinSeqNr:            9018980618932210108,
				MaxSeqNr:            8239368306600774074,
				MerkleRoot:          root32,
			},
		},
	}

	ctx := tests.Context(t)

	rmnCrypto := NewEVMRMNCrypto()

	r, _ := cciptypes.NewBytes32FromString("0x89546b4652d0377062a398e413344e4da6034ae877c437d0efe0e5246b70a9a1")
	s, _ := cciptypes.NewBytes32FromString("0x95eef2d24d856ccac3886db8f4aebea60684ed73942392692908fed79a679b4e")

	err := rmnCrypto.VerifyReportSignatures(
		ctx,
		[]cciptypes.RMNECDSASignature{{R: r, S: s}},
		reportData,
		[]cciptypes.UnknownAddress{onchainRmnRemoteAddr.Bytes()},
	)
	assert.NoError(t, err)
}

func TestIt(t *testing.T) {
	t.Skipf("local only")
	cl, err := ethclient.Dial("http://localhost:51453")
	require.NoError(t, err)

	rmnRemoteClient, err := rmn_remote.NewRMNRemote(common.HexToAddress("0x68B1D87F95878fE05B998F19b66F4baba5De1aed"), cl)
	require.NoError(t, err)

	cfg, err := rmnRemoteClient.GetVersionedConfig(nil)
	require.NoError(t, err)

	js, _ := json.MarshalIndent(cfg, " ", " ")
	fmt.Println(string(js))

	onRampAddr := common.HexToAddress("0x0000000000000000000000007a2088a1bfc9d81c55368ae168c2c02570cb814f")
	merkleRoot := common.Hex2Bytes("1c61fef7a3dd153943419c1101031316ed7b7a3d75913c34cbe8628033f5924f")

	err = rmnRemoteClient.Verify(
		nil,
		common.HexToAddress("0x09635F643e140090A9A8Dcd712eD6285858ceBef"),
		[]rmn_remote.InternalMerkleRoot{
			{
				SourceChainSelector: 3379446385462418246,
				OnRampAddress:       onRampAddr.Bytes(),
				MinSeqNr:            1,
				MaxSeqNr:            1,
				MerkleRoot:          [32]byte(merkleRoot),
			},
		},
		[]rmn_remote.IRMNRemoteSignature{
			{
				R: [32]byte(common.Hex2Bytes("b67b5590817801a1c8cb4813e6f8bdab205a9284701d1e492593bade6643b29e")),
				S: [32]byte(common.Hex2Bytes("364d64599afd0732368e98bdfb0b1b0efbb455009db6d7c6b08e301fcff20819")),
			},
		},
		big.NewInt(0),
	)
	require.NoError(t, err)
}
