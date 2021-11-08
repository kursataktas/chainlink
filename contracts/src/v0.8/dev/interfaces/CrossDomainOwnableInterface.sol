// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @title CrossDomainOwnableInterface - A contract with helpers for cross-domain contract ownership
interface CrossDomainOwnableInterface {
  function l1Owner() external returns (address);

  function transferL1Ownership(address recipient) external;

  function acceptL1Ownership() external;
}
