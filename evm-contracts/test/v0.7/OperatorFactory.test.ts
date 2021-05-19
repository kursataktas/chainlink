import { contract, setup, helpers, matchers } from '@chainlink/test-helpers'
import { assert } from 'chai'
import { utils } from 'ethers'
import { ContractReceipt } from 'ethers/contract'
import { Operator__factory } from '../../ethers/v0.7/factories/Operator__factory'
import { AuthorizedForwarder__factory } from '../../ethers/v0.7/factories/AuthorizedForwarder__factory'
import { OperatorFactory__factory } from '../../ethers/v0.7/factories/OperatorFactory__factory'

const linkTokenFactory = new contract.LinkToken__factory()
const operatorGeneratorFactory = new OperatorFactory__factory()
const operatorFactory = new Operator__factory()
const forwarderFactory = new AuthorizedForwarder__factory()

let roles: setup.Roles
const provider = setup.provider()

beforeAll(async () => {
  const users = await setup.users(provider)

  roles = users.roles
})

describe('OperatorFactory', () => {
  let link: contract.Instance<contract.LinkToken__factory>
  let operatorGenerator: contract.Instance<OperatorFactory__factory>
  let operator: contract.Instance<Operator__factory>
  let forwarder: contract.Instance<AuthorizedForwarder__factory>

  const deployment = setup.snapshot(provider, async () => {
    link = await linkTokenFactory.connect(roles.defaultAccount).deploy()
    operatorGenerator = await operatorGeneratorFactory
      .connect(roles.defaultAccount)
      .deploy(link.address)
  })

  beforeEach(async () => {
    await deployment()
  })

  it('has a limited public interface', () => {
    matchers.publicAbi(operatorGenerator, [
      'deployNewOperator',
      'deployNewOperatorAndForwarder',
      'deployNewForwarder',
      'deployNewForwarderAndTransferOwnership',
      'getChainlinkToken',
    ])
  })

  describe('#deployNewOperator', () => {
    let receipt: ContractReceipt

    beforeEach(async () => {
      const tx = await operatorGenerator
        .connect(roles.oracleNode)
        .deployNewOperator()

      receipt = await tx.wait()
    })

    it('emits an event', async () => {
      assert.equal(roles.oracleNode.address, receipt.events?.[0].args?.[1])
      assert.equal(receipt?.events?.[0]?.event, 'OperatorCreated')
    })

    it('sets the correct owner', async () => {
      const emittedAddress = helpers.evmWordToAddress(
        receipt.logs?.[0].topics?.[1],
      )

      operator = await operatorFactory
        .connect(roles.defaultAccount)
        .attach(emittedAddress)
      const ownerString = await operator.owner()
      assert.equal(ownerString, roles.oracleNode.address)
    })
  })

  describe('#deployNewOperatorAndForwarder', () => {
    let receipt: ContractReceipt

    beforeEach(async () => {
      const tx = await operatorGenerator
        .connect(roles.oracleNode)
        .deployNewOperatorAndForwarder()

      receipt = await tx.wait()
    })

    it('emits an event recording that the operator was deployed', async () => {
      assert.equal(roles.oracleNode.address, receipt.events?.[0].args?.[1])
      assert.equal(receipt?.events?.[0]?.event, 'OperatorCreated')
    })

    it('emits an event recording that the forwarder was deployed', async () => {
      assert.equal(roles.oracleNode.address, receipt.events?.[0].args?.[1])
      assert.equal(receipt?.events?.[1]?.event, 'AuthorizedForwarderCreated')
    })

    it('sets the correct owner on the operator', async () => {
      operator = await operatorFactory
        .connect(roles.defaultAccount)
        .attach(receipt?.events?.[0]?.args?.[0])
      assert.equal(roles.oracleNode.address, await operator.owner())
    })

    it('sets the operator as the owner of the forwarder', async () => {
      forwarder = await forwarderFactory
        .connect(roles.defaultAccount)
        .attach(receipt?.events?.[1]?.args?.[0])
      const operatorAddress = receipt?.events?.[0]?.args?.[0]
      assert.equal(operatorAddress, await forwarder.owner())
    })
  })

  describe('#deployNewForwarder', () => {
    let receipt: ContractReceipt

    beforeEach(async () => {
      const tx = await operatorGenerator
        .connect(roles.oracleNode)
        .deployNewForwarder()

      receipt = await tx.wait()
    })

    it('emits an event', async () => {
      const emittedOwner = helpers.evmWordToAddress(
        receipt.logs?.[0].topics?.[2],
      )
      assert.equal(emittedOwner, roles.oracleNode.address)
      assert.equal(receipt?.events?.[0]?.event, 'AuthorizedForwarderCreated')
    })

    it('sets the caller as the owner', async () => {
      const emittedAddress = helpers.evmWordToAddress(
        receipt.logs?.[0].topics?.[1],
      )

      forwarder = await forwarderFactory
        .connect(roles.defaultAccount)
        .attach(emittedAddress)
      const ownerString = await forwarder.owner()
      assert.equal(ownerString, roles.oracleNode.address)
    })
  })

  describe('#deployNewForwarderAndTransferOwnership', () => {
    const message = '0x42'
    let receipt: ContractReceipt

    beforeEach(async () => {
      const tx = await operatorGenerator
        .connect(roles.oracleNode)
        .deployNewForwarderAndTransferOwnership(roles.stranger.address, message)
      receipt = await tx.wait()
    })

    it('emits an event', async () => {
      assert.equal(roles.oracleNode.address, receipt.events?.[2].args?.[1])
      assert.equal(receipt?.events?.[2]?.event, 'AuthorizedForwarderCreated')
    })

    it('sets the caller as the owner', async () => {
      forwarder = await forwarderFactory
        .connect(roles.defaultAccount)
        .attach(receipt.events?.[2].args?.[0])
      const ownerString = await forwarder.owner()
      assert.equal(ownerString, roles.oracleNode.address)
    })

    it('proposes a transfer to the recipient', async () => {
      const emittedOwner = helpers.evmWordToAddress(
        receipt.logs?.[0].topics?.[1],
      )
      assert.equal(emittedOwner, roles.oracleNode.address)
      const emittedRecipient = helpers.evmWordToAddress(
        receipt.logs?.[0].topics?.[2],
      )
      assert.equal(emittedRecipient, roles.stranger.address)
    })

    it('proposes a transfer to the recipient with the specified message', async () => {
      const emittedOwner = helpers.evmWordToAddress(
        receipt.logs?.[1].topics?.[1],
      )
      assert.equal(emittedOwner, roles.oracleNode.address)
      const emittedRecipient = helpers.evmWordToAddress(
        receipt.logs?.[1].topics?.[2],
      )
      assert.equal(emittedRecipient, roles.stranger.address)

      const encodedMessage = utils.defaultAbiCoder.encode(['bytes'], [message])
      assert.equal(receipt?.logs?.[1]?.data, encodedMessage)
    })
  })
})
