import { TokenListProvider } from '@solana/spl-token-registry';
import { fetchAllDigitalAssetByOwner } from '@metaplex-foundation/mpl-token-metadata'

export class TokenDescriptionList {
    list;

    async getTokenInfoByMintAsync(umi, mint, owner) {
        if (!this.list) {
            const tokensProvider = await new TokenListProvider().resolve();
            this.list = tokensProvider.getList();

            const tokensInfo = (await fetchAllDigitalAssetByOwner(umi, owner))
                .map(item => ({
                    address: item.publicKey,
                    mint: item.metadata.mint,
                    symbol: item.metadata.symbol ?? item.metadata.name,
                    decimals: item.mint.decimals
                }));
            this.list.push(...tokensInfo)
        }

        return this.list.find(token => token.address === mint || token.mint === mint) ?? { decimals: 1, symbol: "undefined" };
    }
}

const tokensProvider = new TokenDescriptionList();
export default tokensProvider;