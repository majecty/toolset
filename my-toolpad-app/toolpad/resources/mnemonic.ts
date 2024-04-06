import { toHex } from 'viem';
import { english, generateMnemonic } from 'viem/accounts'
import { mnemonicToAccount } from 'viem/accounts'

/**
 * Toolpad Studio handlers file.
 */

export default async function handler(message: string) {
  return generateMnemonic(english);
}

export async function generateMnemonic_() {
  return generateMnemonic(english);
}

export async function createPrivateKeyFromMnemonic(mnemonic: string, hdpath: string) {
  const account =  mnemonicToAccount(mnemonic, {
    path: hdpath as any,
  });

  return {
    privateKey: toHex(account.getHdKey().privateKey!),
    publicKey: account.publicKey,
  }
}