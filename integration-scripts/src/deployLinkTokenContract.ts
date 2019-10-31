import { createProvider, DEVNET_ADDRESS, deployContract } from './common'
import { LinkTokenFactory } from 'chainlink/dist/src/generated/LinkTokenFactory'
import { Instance } from 'chainlink/dist/src/contract'

export async function deployLinkTokenContract(): Promise<
  Instance<LinkTokenFactory>
> {
  const provider = createProvider()
  const signer = provider.getSigner(DEVNET_ADDRESS)
  if (process.env.LINK_TOKEN_ADDRESS) {
    console.log(
      `LinkToken already deployed at: ${process.env.LINK_TOKEN_ADDRESS}, fetching contract...`,
    )
    const factory = new LinkTokenFactory(signer)
    const linkToken = factory.attach(process.env.LINK_TOKEN_ADDRESS)
    console.log(`Deployed LinkToken at: ${linkToken.address}`)

    return linkToken
  }

  const linkToken = await deployContract({
    Factory: LinkTokenFactory,
    name: 'LinkToken',
    signer,
  })

  return linkToken
}
