// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnerIsCreator} from "../shared/access/OwnerIsCreator.sol";

import {AccessControllerInterface} from "../shared/interfaces/AccessControllerInterface.sol";
import {AggregatorV2V3Interface} from "../shared/interfaces/AggregatorV2V3Interface.sol";
import {AggregatorValidatorInterface} from "../shared/interfaces/AggregatorValidatorInterface.sol";
import {LinkTokenInterface} from "../shared/interfaces/LinkTokenInterface.sol";
import {OCR2Abstract} from "../shared/ocr2/OCR2Abstract.sol";

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
