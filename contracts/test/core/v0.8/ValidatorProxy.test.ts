import { ethers } from "hardhat";
import { publicAbi, constants } from "../../helpers";
import { assert, expect } from "chai";
import { Signer, Contract } from "ethers";
import { Personas, Users, getUsers } from "../../setup";

describe("ValidatorProxy", () => {
  let users: Users;

  let owner: Signer;
  let ownerAddress: string;
  let aggregator: Signer;
  let aggregatorAddress: string;
  let validator: Signer;
  let validatorAddress: string;
  let validatorProxy: Contract;

  beforeEach(async () => {
    users = await getUsers();
    owner = users.personas.Default;
    aggregator = users.contracts.contract1;
    validator = users.contracts.contract2;
    ownerAddress = await owner.getAddress();
    aggregatorAddress = await aggregator.getAddress();
    validatorAddress = await validator.getAddress();

    const vpf = await ethers.getContractFactory("ValidatorProxy", owner);
    validatorProxy = await vpf.deploy(aggregatorAddress, validatorAddress);
    validatorProxy = await validatorProxy.deployed();
  });

  it("has a limited public interface", async () => {
    publicAbi(validatorProxy, [
      // ConfirmedOwner functions
      "acceptOwnership",
      "owner",
      "transferOwnership",
      // ValidatorProxy functions
      "validate",
      "proposeNewAggregator",
      "upgradeAggregator",
      "getAggregators",
      "proposeNewValidator",
      "upgradeValidator",
      "getValidators",
      "typeAndVersion",
    ]);
  });

  describe("#constructor", () => {
    it("should set the aggregator addresses correctly", async () => {
      const response = await validatorProxy.getAggregators();
      assert.equal(response.current, aggregatorAddress);
      assert.equal(response.hasProposal, false);
      assert.equal(response.proposed, constants.ZERO_ADDRESS);
    });

    it("should set the validator addresses conrrectly", async () => {
      const response = await validatorProxy.getValidators();
      assert.equal(response.current, validatorAddress);
      assert.equal(response.hasProposal, false);
      assert.equal(response.proposed, constants.ZERO_ADDRESS);
    });

    it("should set the owner correctly", async () => {
      const response = await validatorProxy.owner();
      assert.equal(response, ownerAddress);
    });
  });

  describe("#proposeNewAggregator", () => {
    let newAggregator: Signer;
    let newAggregatorAddress: string;
    beforeEach(async () => {
      newAggregator = users.contracts.contract3;
      newAggregatorAddress = await newAggregator.getAddress();
    });

    describe("failure", () => {
      it("should only be called by the owner", async () => {
        const stranger = users.contracts.contract4;
        await expect(validatorProxy.connect(stranger).proposeNewAggregator(newAggregatorAddress)).to.be.revertedWith(
          "Only callable by owner",
        );
      });

      it("should revert if no change in proposal", async () => {
        await validatorProxy.proposeNewAggregator(newAggregatorAddress);
        await expect(validatorProxy.proposeNewAggregator(newAggregatorAddress)).to.be.revertedWith("Invalid proposal");
      });

      it("should revert if the proposal is the same as the current", async () => {
        await expect(validatorProxy.proposeNewAggregator(aggregatorAddress)).to.be.revertedWith("Invalid proposal");
      });
    });

    describe("success", () => {
      it("should emit an event", async () => {
        await expect(validatorProxy.proposeNewAggregator(newAggregatorAddress))
          .to.emit(validatorProxy, "AggregatorProposed")
          .withArgs(newAggregatorAddress);
      });

      it("should set the correct address and hasProposal is true", async () => {
        await validatorProxy.proposeNewAggregator(newAggregatorAddress);
        const response = await validatorProxy.getAggregators();
        assert.equal(response.current, aggregatorAddress);
        assert.equal(response.hasProposal, true);
        assert.equal(response.proposed, newAggregatorAddress);
      });

      it("should set a zero address and hasProposal is false", async () => {
        await validatorProxy.proposeNewAggregator(newAggregatorAddress);
        await validatorProxy.proposeNewAggregator(constants.ZERO_ADDRESS);
        const response = await validatorProxy.getAggregators();
        assert.equal(response.current, aggregatorAddress);
        assert.equal(response.hasProposal, false);
        assert.equal(response.proposed, constants.ZERO_ADDRESS);
      });
    });
  });

  describe("#upgradeAggregator", () => {
    describe("failure", () => {
      it("should only be called by the owner", async () => {
        const stranger = users.contracts.contract4;
        await expect(validatorProxy.connect(stranger).upgradeAggregator()).to.be.revertedWith("Only callable by owner");
      });

      it("should revert if there is no proposal", async () => {
        await expect(validatorProxy.upgradeAggregator()).to.be.revertedWith("No proposal");
      });
    });

    describe("success", () => {
      let newAggregator: Signer;
      let newAggregatorAddress: string;
      beforeEach(async () => {
        newAggregator = users.contracts.contract3;
        newAggregatorAddress = await newAggregator.getAddress();
        await validatorProxy.proposeNewAggregator(newAggregatorAddress);
      });

      it("should emit an event", async () => {
        await expect(validatorProxy.upgradeAggregator())
          .to.emit(validatorProxy, "AggregatorUpgraded")
          .withArgs(aggregatorAddress, newAggregatorAddress);
      });

      it("should upgrade the addresses", async () => {
        await validatorProxy.upgradeAggregator();
        const response = await validatorProxy.getAggregators();
        assert.equal(response.current, newAggregatorAddress);
        assert.equal(response.hasProposal, false);
        assert.equal(response.proposed, constants.ZERO_ADDRESS);
      });
    });
  });

  describe("#proposeNewValidator", () => {
    let newValidator: Signer;
    let newValidatorAddress: string;

    beforeEach(async () => {
      newValidator = users.contracts.contract3;
      newValidatorAddress = await newValidator.getAddress();
    });

    describe("failure", () => {
      it("should only be called by the owner", async () => {
        const stranger = users.contracts.contract4;
        await expect(validatorProxy.connect(stranger).proposeNewAggregator(newValidatorAddress)).to.be.revertedWith(
          "Only callable by owner",
        );
      });

      it("should revert if no change in proposal", async () => {
        await validatorProxy.proposeNewValidator(newValidatorAddress);
        await expect(validatorProxy.proposeNewValidator(newValidatorAddress)).to.be.revertedWith("Invalid proposal");
      });

      it("should revert if the proposal is the same as the current", async () => {
        await expect(validatorProxy.proposeNewValidator(validatorAddress)).to.be.revertedWith("Invalid proposal");
      });
    });

    describe("success", () => {
      it("should emit an event", async () => {
        await expect(validatorProxy.proposeNewValidator(newValidatorAddress))
          .to.emit(validatorProxy, "ValidatorProposed")
          .withArgs(newValidatorAddress);
      });

      it("should set the correct address and hasProposal is true", async () => {
        await validatorProxy.proposeNewValidator(newValidatorAddress);
        const response = await validatorProxy.getValidators();
        assert.equal(response.current, validatorAddress);
        assert.equal(response.hasProposal, true);
        assert.equal(response.proposed, newValidatorAddress);
      });

      it("should set a zero address and hasProposal is false", async () => {
        await validatorProxy.proposeNewValidator(newValidatorAddress);
        await validatorProxy.proposeNewValidator(constants.ZERO_ADDRESS);
        const response = await validatorProxy.getValidators();
        assert.equal(response.current, validatorAddress);
        assert.equal(response.hasProposal, false);
        assert.equal(response.proposed, constants.ZERO_ADDRESS);
      });
    });
  });

  describe("#upgradeValidator", () => {
    describe("failure", () => {
      it("should only be called by the owner", async () => {
        const stranger = users.contracts.contract4;
        await expect(validatorProxy.connect(stranger).upgradeValidator()).to.be.revertedWith("Only callable by owner");
      });

      it("should revert if there is no proposal", async () => {
        await expect(validatorProxy.upgradeValidator()).to.be.revertedWith("No proposal");
      });
    });

    describe("success", () => {
      let newValidator: Signer;
      let newValidatorAddress: string;
      beforeEach(async () => {
        newValidator = users.contracts.contract3;
        newValidatorAddress = await newValidator.getAddress();
        await validatorProxy.proposeNewValidator(newValidatorAddress);
      });

      it("should emit an event", async () => {
        await expect(validatorProxy.upgradeValidator())
          .to.emit(validatorProxy, "ValidatorUpgraded")
          .withArgs(validatorAddress, newValidatorAddress);
      });

      it("should upgrade the addresses", async () => {
        await validatorProxy.upgradeValidator();
        const response = await validatorProxy.getValidators();
        assert.equal(response.current, newValidatorAddress);
        assert.equal(response.hasProposal, false);
        assert.equal(response.proposed, constants.ZERO_ADDRESS);
      });
    });
  });

  describe("#validate", () => {
    describe("failure", () => {
      it("reverts when not called by aggregator or proposed aggregator", async () => {
        const stranger = users.contracts.contract5;
        await expect(validatorProxy.connect(stranger).validate(99, 88, 77, 66)).to.be.revertedWith(
          "Not a configured aggregator",
        );
      });

      it("reverts when there is no validator set", async () => {
        const vpf = await ethers.getContractFactory("ValidatorProxy", owner);
        validatorProxy = await vpf.deploy(aggregatorAddress, constants.ZERO_ADDRESS);
        await validatorProxy.deployed();
        await expect(validatorProxy.connect(aggregator).validate(99, 88, 77, 66)).to.be.revertedWith(
          "No validator set",
        );
      });
    });

    describe("success", () => {
      describe("from the aggregator", () => {
        let mockValidator1: Contract;
        beforeEach(async () => {
          const mvf = await ethers.getContractFactory("MockAggregatorValidator", owner);
          mockValidator1 = await mvf.deploy(1);
          mockValidator1 = await mockValidator1.deployed();
          const vpf = await ethers.getContractFactory("ValidatorProxy", owner);
          validatorProxy = await vpf.deploy(aggregatorAddress, mockValidator1.address);
          validatorProxy = await validatorProxy.deployed();
        });

        describe("for a single validator", () => {
          it("calls validate on the validator", async () => {
            await expect(validatorProxy.connect(aggregator).validate(200, 300, 400, 500))
              .to.emit(mockValidator1, "ValidateCalled")
              .withArgs(1, 200, 300, 400, 500);
          });

          it("uses a specific amount of gas", async () => {
            const resp = await validatorProxy.connect(aggregator).validate(200, 300, 400, 500);
            const receipt = await resp.wait();
            assert.equal(receipt.gasUsed.toString(), "35371");
          });
        });

        describe("for a validator and a proposed validator", () => {
          let mockValidator2: Contract;

          beforeEach(async () => {
            const mvf = await ethers.getContractFactory("MockAggregatorValidator", owner);
            mockValidator2 = await mvf.deploy(2);
            mockValidator2 = await mockValidator2.deployed();
            await validatorProxy.proposeNewValidator(mockValidator2.address);
          });

          it("calls validate on the validator", async () => {
            await expect(validatorProxy.connect(aggregator).validate(2000, 3000, 4000, 5000))
              .to.emit(mockValidator1, "ValidateCalled")
              .withArgs(1, 2000, 3000, 4000, 5000);
          });

          it("also calls validate on the proposed validator", async () => {
            await expect(validatorProxy.connect(aggregator).validate(2000, 3000, 4000, 5000))
              .to.emit(mockValidator2, "ValidateCalled")
              .withArgs(2, 2000, 3000, 4000, 5000);
          });

          it("uses a specific amount of gas", async () => {
            const resp = await validatorProxy.connect(aggregator).validate(2000, 3000, 4000, 5000);
            const receipt = await resp.wait();
            assert.equal(receipt.gasUsed.toString(), "45318");
          });
        });
      });

      describe("from the proposed aggregator", () => {
        let newAggregator: Signer;
        let newAggregatorAddress: string;
        beforeEach(async () => {
          newAggregator = users.contracts.contract3;
          newAggregatorAddress = await newAggregator.getAddress();
          await validatorProxy.connect(owner).proposeNewAggregator(newAggregatorAddress);
        });

        it("emits an event", async () => {
          await expect(validatorProxy.connect(newAggregator).validate(555, 666, 777, 888))
            .to.emit(validatorProxy, "ProposedAggregatorValidateCall")
            .withArgs(newAggregatorAddress, 555, 666, 777, 888);
        });
      });
    });
  });
});
