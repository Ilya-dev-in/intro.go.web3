import { TokenListProvider } from '@solana/spl-token-registry';
import { fetchAllDigitalAssetByOwner } from '@metaplex-foundation/mpl-token-metadata'
import axios from 'axios';

export class TokenDescriptionList {
    list;

    async getTokenInfoByMintAsync(umi, mint, owner) {
        if (!this.list) {
            const tokensProvider = await new TokenListProvider().resolve();
            this.list = tokensProvider.getList();
            const assets = await fetchAllDigitalAssetByOwner(umi, owner);
            console.log(assets)
            const tokensInfo = await Promise.all((await fetchAllDigitalAssetByOwner(umi, owner))
                .map(async item => {
                    const token = {
                        address: item.publicKey,
                        mint: item.metadata.mint,
                        symbol: item.metadata.symbol ?? item.metadata.name,
                        decimals: item.mint.decimals
                    }

                    //const metaUriData = await axios.get(item.metadata.uri);
                    token.image = "https://metadata.jito.network/token/jto/image";

                    return token;
                }));
            this.list.push(...tokensInfo)
        }

        return this.list.find(token => token.address === mint || token.mint === mint) ?? { decimals: 1, symbol: null };
    }
}

const tokensProvider = new TokenDescriptionList();
export default tokensProvider;