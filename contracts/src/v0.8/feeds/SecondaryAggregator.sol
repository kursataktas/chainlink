// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import {AccessControllerInterface} from "../shared/interfaces/AccessControllerInterface.sol";
import {LinkTokenInterface} from "../shared/interfaces/LinkTokenInterface.sol";
import {PrimaryAggregator} from "./PrimaryAggregator.sol";

contract SecondaryAggregator is PrimaryAggregator {
  constructor(
    LinkTokenInterface link,
    int192 minAnswer_,
    int192 maxAnswer_,
    AccessControllerInterface billingAccessController,
    AccessControllerInterface requesterAccessController,
    uint8 decimals_,
    string memory description_
  )
    PrimaryAggregator(
      link,
      minAnswer_,
      maxAnswer_,
      billingAccessController,
      requesterAccessController,
      decimals_,
      description_
    )
  {}
}
