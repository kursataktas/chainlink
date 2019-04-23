import yargs from 'yargs'
import { Connection } from 'typeorm'
import {
  ChainlinkNode,
  createChainlinkNode,
  deleteChainlinkNode
} from '../entity/ChainlinkNode'
import { closeDbConnection, getDb } from '../database'

const add = async (name: string) => {
  return bootstrap(async (db: Connection) => {
    const [chainlinkNode, secret] = await createChainlinkNode(db, name)
    console.log('created new chainlink node with id %s', chainlinkNode.id)
    console.log('AccessKey', chainlinkNode.accessKey)
    console.log('Secret', secret)
  })
}

const remove = async (name: string) => {
  return bootstrap(async (db: Connection) => {
    deleteChainlinkNode(db, name)
  })
}

const bootstrap = async (cb: any) => {
  const db = await getDb()
  try {
    await cb(db)
  } catch (e) {
    console.error(e)
  }
  try {
    await closeDbConnection()
  } catch (e) {}
}

const _ = yargs
  .usage('Usage: $0 <command> [options]')
  .command({
    command: 'add <name>',
    aliases: 'create',
    describe: 'Add a chainlink node',
    builder: (yargs): any => {
      yargs.positional('name', {
        describe: 'The name of the Chainlink Node to create',
        type: 'string'
      })
    },
    handler: argv => add(argv.name as string)
  })
  .command({
    command: 'delete <name>',
    aliases: 'rm',
    describe: 'Remove a chainlink node',
    builder: (yargs): any => {
      yargs.positional('name', {
        describe: 'The name of the Chainlink Node to remove',
        type: 'string'
      })
    },
    handler: argv => remove(argv.name as string)
  })
  .help('h')
  .alias('h', 'help')
  .demandCommand(1).argv
