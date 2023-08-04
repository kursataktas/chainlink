// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IRouterBase} from "./interfaces/IRouterBase.sol";
import {IConfigurable} from "./interfaces/IConfigurable.sol";

import {ConfirmedOwner} from "../../../ConfirmedOwner.sol";

import {Pausable} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/security/Pausable.sol";

abstract contract RouterBase is IRouterBase, Pausable, ITypeAndVersion, ConfirmedOwner {
  // ================================================================
  // |                          Route state                         |
  // ================================================================
  mapping(bytes32 id => address routableContract) private s_route;

  error RouteNotFound(bytes32 id);

  // Use empty bytes to self-identify, since it does not have an id
  bytes32 private constant ROUTER_ID = bytes32(0);

  // ================================================================
  // |                         Proposal state                       |
  // ================================================================

  uint8 private constant MAX_PROPOSAL_SET_LENGTH = 8;

  struct ContractProposalSet {
    bytes32[] ids;
    address[] to;
    uint64 timelockEndBlock;
  }
  ContractProposalSet private s_proposedContractSet;

  event ContractProposed(
    bytes32 proposedContractSetId,
    address proposedContractSetFromAddress,
    address proposedContractSetToAddress,
    uint64 timelockEndBlock
  );

  event ContractUpdated(bytes32 id, address from, address to);

  struct ConfigProposal {
    bytes to;
    uint64 timelockEndBlock;
  }
  mapping(bytes32 id => ConfigProposal) private s_proposedConfig;
  event ConfigProposed(bytes32 id, bytes toBytes);
  event ConfigUpdated(bytes32 id, bytes toBytes);
  error InvalidProposal();
  error IdentifierIsReserved(bytes32 id);

  // ================================================================
  // |                          Timelock state                      |
  // ================================================================
  uint16 private immutable s_maximumTimelockBlocks;
  uint16 private s_timelockBlocks;

  struct TimeLockProposal {
    uint16 from;
    uint16 to;
    uint64 timelockEndBlock;
  }

  TimeLockProposal private s_timelockProposal;

  event TimeLockProposed(uint16 from, uint16 to);
  event TimeLockUpdated(uint16 from, uint16 to);
  error ProposedTimelockAboveMaximum();
  error TimelockInEffect();

  // ================================================================
  // |                       Initialization                         |
  // ================================================================

  constructor(
    address newOwner,
    uint16 timelockBlocks,
    uint16 maximumTimelockBlocks,
    bytes memory selfConfig
  ) ConfirmedOwner(newOwner) Pausable() {
    // Set initial value for the number of blocks of the timelock
    s_timelockBlocks = timelockBlocks;
    // Set maximum number of blocks that the timelock can be
    s_maximumTimelockBlocks = maximumTimelockBlocks;
    // Set the initial configuration for the Router
    s_route[ROUTER_ID] = address(this);
    _updateConfig(selfConfig);
  }

  // ================================================================
  // |                        Route methods                         |
  // ================================================================

  // @inheritdoc IRouterBase
  function getContractById(bytes32 id) public view override returns (address) {
    address currentImplementation = s_route[id];
    if (currentImplementation == address(0)) {
      revert RouteNotFound(id);
    }
    return currentImplementation;
  }

  // @inheritdoc IRouterBase
  function getProposedContractById(bytes32 id) public view override returns (address) {
    // Iterations will not exceed MAX_PROPOSAL_SET_LENGTH
    for (uint8 i = 0; i < s_proposedContractSet.ids.length; ++i) {
      if (id == s_proposedContractSet.ids[i]) {
        return s_proposedContractSet.to[i];
      }
    }
    revert RouteNotFound(id);
  }

  // ================================================================
  // |                 Contract Proposal methods                    |
  // ================================================================

  // @inheritdoc IRouterBase
  function getProposedContractSet()
    external
    view
    override
    returns (uint256 timelockEndBlock, bytes32[] memory ids, address[] memory to)
  {
    timelockEndBlock = s_proposedContractSet.timelockEndBlock;
    ids = s_proposedContractSet.ids;
    to = s_proposedContractSet.to;
    return (timelockEndBlock, ids, to);
  }

  // @inheritdoc IRouterBase
  function proposeContractsUpdate(
    bytes32[] memory proposedContractSetIds,
    address[] memory proposedContractSetAddresses
  ) external override onlyOwner {
    // All arrays must be of equal length and not must not exceed the max length
    uint256 idsArrayLength = proposedContractSetIds.length;
    if (idsArrayLength != proposedContractSetAddresses.length || idsArrayLength > MAX_PROPOSAL_SET_LENGTH) {
      revert InvalidProposal();
    }
    // Iterations will not exceed MAX_PROPOSAL_SET_LENGTH
    for (uint256 i = 0; i < idsArrayLength; ++i) {
      bytes32 id = proposedContractSetIds[i];
      address proposedContract = proposedContractSetAddresses[i];
      if (
        proposedContract == address(0) || // The Proposed address must be a valid address
        s_route[id] == proposedContract // The Proposed address must point to a different address than what is currently set
      ) {
        revert InvalidProposal();
      }
      // Reserved ids cannot be set
      if (id == ROUTER_ID) {
        revert IdentifierIsReserved(id);
      }
    }

    uint64 timelockEndBlock = uint64(block.number + s_timelockBlocks);

    s_proposedContractSet = ContractProposalSet({
      ids: proposedContractSetIds,
      to: proposedContractSetAddresses,
      timelockEndBlock: timelockEndBlock
    });

    // Iterations will not exceed MAX_PROPOSAL_SET_LENGTH
    for (uint256 i = 0; i < proposedContractSetIds.length; ++i) {
      emit ContractProposed({
        proposedContractSetId: proposedContractSetIds[i],
        proposedContractSetFromAddress: s_route[proposedContractSetIds[i]],
        proposedContractSetToAddress: proposedContractSetAddresses[i],
        timelockEndBlock: timelockEndBlock
      });
    }
  }

  // @inheritdoc IRouterBase
  function updateContracts() external override onlyOwner {
    if (block.number < s_proposedContractSet.timelockEndBlock) {
      revert TimelockInEffect();
    }
    // Iterations will not exceed MAX_PROPOSAL_SET_LENGTH
    for (uint256 i = 0; i < s_proposedContractSet.ids.length; ++i) {
      bytes32 id = s_proposedContractSet.ids[i];
      address to = s_proposedContractSet.to[i];
      emit ContractUpdated({id: id, from: s_route[id], to: to});
      s_route[id] = to;
    }

    delete s_proposedContractSet;
  }

  // ================================================================
  // |                   Config Proposal methods                    |
  // ================================================================

  // @dev Must be implemented by inheriting contract
  // Use to set configuration state of the Router
  function _updateConfig(bytes memory config) internal virtual;

  // @inheritdoc IRouterBase
  function proposeConfigUpdateSelf(bytes calldata config) external override onlyOwner {
    s_proposedConfig[ROUTER_ID] = ConfigProposal({
      to: config,
      timelockEndBlock: uint64(block.number + s_timelockBlocks)
    });
    emit ConfigProposed({id: ROUTER_ID, toBytes: config});
  }

  // @inheritdoc IRouterBase
  function updateConfigSelf() external override onlyOwner {
    ConfigProposal memory proposal = s_proposedConfig[ROUTER_ID];
    if (block.number < proposal.timelockEndBlock) {
      revert TimelockInEffect();
    }
    _updateConfig(proposal.to);
    emit ConfigUpdated({id: ROUTER_ID, toBytes: proposal.to});
  }

  // @inheritdoc IRouterBase
  function proposeConfigUpdate(bytes32 id, bytes calldata config) external override onlyOwner {
    s_proposedConfig[id] = ConfigProposal({to: config, timelockEndBlock: uint64(block.number + s_timelockBlocks)});
    emit ConfigProposed({id: id, toBytes: config});
  }

  // @inheritdoc IRouterBase
  function updateConfig(bytes32 id) external override onlyOwner {
    ConfigProposal memory proposal = s_proposedConfig[id];

    if (block.number < proposal.timelockEndBlock) {
      revert TimelockInEffect();
    }

    IConfigurable(getContractById(id)).updateConfig(proposal.to);

    emit ConfigUpdated({id: id, toBytes: proposal.to});
  }

  // ================================================================
  // |                         Timelock methods                     |
  // ================================================================

  // @inheritdoc IRouterBase
  function proposeTimelockBlocks(uint16 blocks) external override onlyOwner {
    if (s_timelockBlocks == blocks) {
      revert InvalidProposal();
    }
    if (blocks > s_maximumTimelockBlocks) {
      revert ProposedTimelockAboveMaximum();
    }
    s_timelockProposal = TimeLockProposal({
      from: s_timelockBlocks,
      to: blocks,
      timelockEndBlock: uint64(block.number + s_timelockBlocks)
    });
  }

  // @inheritdoc IRouterBase
  function updateTimelockBlocks() external override onlyOwner {
    if (block.number < s_timelockProposal.timelockEndBlock) {
      revert TimelockInEffect();
    }
    s_timelockBlocks = s_timelockProposal.to;
  }

  // ================================================================
  // |                     Pausable methods                         |
  // ================================================================

  // @inheritdoc IRouterBase
  function pause() external override onlyOwner {
    _pause();
  }

  // @inheritdoc IRouterBase
  function unpause() external override onlyOwner {
    _unpause();
  }
}
