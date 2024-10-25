// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {ConfirmedOwnerWithProposal} from "./ConfirmedOwnerWithProposal.sol";

/**
 * @notice Onchain verification of reports from the offchain reporting protocol
 * @dev For details on its operation, see the offchain reporting protocol design
 * doc, which refers to this contract as simply the "contract".
 */
contract SimpleOCR is ConfirmedOwnerWithProposal {
    error InvalidConfig(string message);

    /**
     * @notice triggers a new run of the offchain reporting protocol
     * @param previousConfigBlockNumber block in which the previous config was set, to simplify historic analysis
     * @param configDigest configDigest of this configuration
     * @param configCount ordinal number of this config setting among all config settings over the life of this contract
     * @param signers ith element is address ith oracle uses to sign a report
     * @param transmitters ith element is address ith oracle uses to transmit a report via the transmit method
     * @param f maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly
     * @param onchainConfig serialized configuration used by the contract (and possibly oracles)
     * @param offchainConfigVersion version of the serialization format used for "offchainConfig" parameter
     * @param offchainConfig serialized configuration used by the oracles exclusively and only passed through the contract
     */
    event ConfigSet(
        uint32 previousConfigBlockNumber,
        bytes32 configDigest,
        uint64 configCount,
        address[] signers,
        address[] transmitters,
        uint8 f,
        bytes onchainConfig,
        uint64 offchainConfigVersion,
        bytes offchainConfig
    );

    constructor() ConfirmedOwnerWithProposal(msg.sender, address(0)) {}

    function typeAndVersion() external pure returns (string memory) {
        return "Simple OCR 1.0.0";
    }

    // Maximum number of oracles the offchain reporting protocol is designed for
    uint256 internal constant MAX_NUM_ORACLES = 31;
    // incremented each time a new config is posted. This count is incorporated
    // into the config digest, to prevent replay attacks.
    uint32 internal s_configCount;
    uint32 internal s_latestConfigBlockNumber; // makes it easier for offchain systems
    // to extract config from logs.

    // Storing these fields used on the hot path in a ConfigInfo variable reduces the
    // retrieval of all of them to a single SLOAD. If any further fields are
    // added, make sure that storage of the struct still takes at most 32 bytes.
    struct ConfigInfo {
        bytes32 latestConfigDigest;
        uint8 f;
        uint8 n;
    }
    ConfigInfo internal s_configInfo;

    // Used for s_oracles[a].role, where a is an address, to track the purpose
    // of the address, or to indicate that the address is unset.
    enum Role {
        // No oracle role has been set for address a
        Unset,
        // Signing address for the s_oracles[a].index'th oracle. I.e., report
        // signatures from this oracle should ecrecover back to address a.
        Signer,
        Transmitter
    }

    struct Oracle {
        uint8 index; // Index of oracle in s_signers/s_transmitters
        Role role; // Role of the address which mapped to this struct
    }

    mapping(address signerOrTransmitter => Oracle) internal s_oracles;

    // s_signers contains the signing address of each oracle
    address[] internal s_signers;

    // s_transmitters contains the transmission address of each oracle,
    // i.e. the address the oracle actually sends transactions to the contract from
    address[] internal s_transmitters;

    /*
     * Config logic
     */

    // Reverts transaction if config args are invalid
    modifier checkConfigValid(
        uint256 numSigners,
        uint256 numTransmitters,
        uint256 f
    ) {
        if (numSigners > MAX_NUM_ORACLES)
            revert InvalidConfig("too many signers");
        if (f == 0) revert InvalidConfig("f must be positive");
        if (numSigners != numTransmitters)
            revert InvalidConfig("oracle addresses out of registration");
        if (numSigners <= 3 * f)
            revert InvalidConfig("faulty-oracle f too high");
        _;
    }

    // solhint-disable-next-line gas-struct-packing
    struct SetConfigArgs {
        address[] signers;
        address[] transmitters;
        uint8 f;
        bytes onchainConfig;
        uint64 offchainConfigVersion;
        bytes offchainConfig;
    }

    /**
    * @notice optionally returns the latest configDigest and epoch for which a
    report was successfully transmitted. Alternatively, the contract may return
    scanLogs set to true and use Transmitted events to provide this information
    to offchain watchers.
    * @return scanLogs indicates whether to rely on the configDigest and epoch
        returned or whether to scan logs for the Transmitted event instead.
    * @return configDigest
    * @return epoch
    */
    function latestConfigDigestAndEpoch()
        external
        view
        virtual
        returns (bool scanLogs, bytes32 configDigest, uint32 epoch)
    {
        return (true, bytes32(0), uint32(0));
    }

    /**
     * @notice sets offchain reporting protocol configuration incl. participating oracles
     * @param _signers addresses with which oracles sign the reports
     * @param _transmitters addresses oracles use to transmit the reports
     * @param _f number of faulty oracles the system can tolerate
     * @param _onchainConfig encoded on-chain contract configuration
     * @param _offchainConfigVersion version number for offchainEncoding schema
     * @param _offchainConfig encoded off-chain oracle configuration
     */
    function setConfig(
        address[] memory _signers,
        address[] memory _transmitters,
        uint8 _f,
        bytes memory _onchainConfig,
        uint64 _offchainConfigVersion,
        bytes memory _offchainConfig
    )
        external
        checkConfigValid(_signers.length, _transmitters.length, _f)
        onlyOwner
    {
        SetConfigArgs memory args = SetConfigArgs({
            signers: _signers,
            transmitters: _transmitters,
            f: _f,
            onchainConfig: _onchainConfig,
            offchainConfigVersion: _offchainConfigVersion,
            offchainConfig: _offchainConfig
        });

        while (s_signers.length != 0) {
            // remove any old signer/transmitter addresses
            uint256 lastIdx = s_signers.length - 1;
            address signer = s_signers[lastIdx];
            address transmitter = s_transmitters[lastIdx];
            delete s_oracles[signer];
            delete s_oracles[transmitter];
            s_signers.pop();
            s_transmitters.pop();
        }

        // Bounded by MAX_NUM_ORACLES in OCR2Abstract.sol
        for (uint256 i = 0; i < args.signers.length; i++) {
            if (args.signers[i] == address(0))
                revert InvalidConfig("signer must not be empty");
            if (args.transmitters[i] == address(0))
                revert InvalidConfig("transmitter must not be empty");
            // add new signer/transmitter addresses
            if (s_oracles[args.signers[i]].role != Role.Unset)
                revert InvalidConfig("repeated signer address");
            s_oracles[args.signers[i]] = Oracle(uint8(i), Role.Signer);
            if (s_oracles[args.transmitters[i]].role != Role.Unset)
                revert InvalidConfig("repeated transmitter address");
            s_oracles[args.transmitters[i]] = Oracle(
                uint8(i),
                Role.Transmitter
            );
            s_signers.push(args.signers[i]);
            s_transmitters.push(args.transmitters[i]);
        }
        s_configInfo.f = args.f;
        uint32 previousConfigBlockNumber = s_latestConfigBlockNumber;
        s_latestConfigBlockNumber = uint32(block.number);
        s_configCount += 1;
        {
            s_configInfo.latestConfigDigest = _configDigestFromConfigData(
                block.chainid,
                address(this),
                s_configCount,
                args.signers,
                args.transmitters,
                args.f,
                args.onchainConfig,
                args.offchainConfigVersion,
                args.offchainConfig
            );
        }
        s_configInfo.n = uint8(args.signers.length);

        emit ConfigSet(
            previousConfigBlockNumber,
            s_configInfo.latestConfigDigest,
            s_configCount,
            args.signers,
            args.transmitters,
            args.f,
            args.onchainConfig,
            args.offchainConfigVersion,
            args.offchainConfig
        );
    }

    function _configDigestFromConfigData(
        uint256 _chainId,
        address _contractAddress,
        uint64 _configCount,
        address[] memory _signers,
        address[] memory _transmitters,
        uint8 _f,
        bytes memory _onchainConfig,
        uint64 _encodedConfigVersion,
        bytes memory _encodedConfig
    ) internal pure returns (bytes32) {
        uint256 h = uint256(
            keccak256(
                abi.encode(
                    _chainId,
                    _contractAddress,
                    _configCount,
                    _signers,
                    _transmitters,
                    _f,
                    _onchainConfig,
                    _encodedConfigVersion,
                    _encodedConfig
                )
            )
        );
        uint256 prefixMask = type(uint256).max << (256 - 16); // 0xFFFF00..00
        uint256 prefix = 0x0001 << (256 - 16); // 0x000100..00
        return bytes32((prefix & prefixMask) | (h & ~prefixMask));
    }

    /**
     * @notice information about current offchain reporting protocol configuration
     * @return configCount ordinal number of current config, out of all configs applied to this contract so far
     * @return blockNumber block at which this config was set
     * @return configDigest domain-separation tag for current config (see __configDigestFromConfigData)
     */
    function latestConfigDetails()
        external
        view
        returns (uint32 configCount, uint32 blockNumber, bytes32 configDigest)
    {
        return (
            s_configCount,
            s_latestConfigBlockNumber,
            s_configInfo.latestConfigDigest
        );
    }

    /**
     * @return list of addresses permitted to transmit reports to this contract
     * @dev The list will match the order used to specify the transmitter during setConfig
     */
    function transmitters() external view returns (address[] memory) {
        return s_transmitters;
    }
}
