import UniversalProvider from '@walletconnect/universal-provider';
import type { AnyTransaction, Provider } from '@reown/appkit-utils/solana';
import { ProviderEventEmitter } from './shared/ProviderEventEmitter.js';
import type { SessionTypes } from '@walletconnect/types';
import { Connection, PublicKey, type SendOptions } from '@solana/web3.js';
import type { CaipNetwork } from '@reown/appkit-common';
export type WalletConnectProviderConfig = {
    provider: UniversalProvider;
    chains: CaipNetwork[];
    getActiveChain: () => CaipNetwork | undefined;
};
export declare class WalletConnectProvider extends ProviderEventEmitter implements Provider {
    readonly name = "WalletConnect";
    readonly type = "WALLET_CONNECT";
    readonly icon = "https://imagedelivery.net/_aTEfDRm7z3tKgu9JhfeKA/05338e12-4f75-4982-4e8a-83c67b826b00/md";
    session?: SessionTypes.Struct;
    provider: UniversalProvider;
    private readonly requestedChains;
    private readonly getActiveChain;
    constructor({ provider, chains, getActiveChain }: WalletConnectProviderConfig);
    onUri?: (uri: string) => void;
    get chains(): CaipNetwork[];
    get publicKey(): PublicKey | undefined;
    connect(): Promise<string>;
    disconnect(): Promise<void>;
    signMessage(message: Uint8Array): Promise<Uint8Array>;
    signTransaction<T extends AnyTransaction>(transaction: T): Promise<T>;
    signAndSendTransaction<T extends AnyTransaction>(transaction: T, sendOptions?: SendOptions): Promise<string>;
    sendTransaction(transaction: AnyTransaction, connection: Connection, options?: SendOptions): Promise<string>;
    signAllTransactions<T extends AnyTransaction[]>(transactions: T): Promise<T>;
    private request;
    private get sessionChains();
    private serializeTransaction;
    private getAccount;
    private getRequestedChainsWithDeprecated;
    private getRawRPCParams;
    private checkIfMethodIsSupported;
}
export declare namespace WalletConnectProvider {
    type Request<Params, Result> = {
        params: Params;
        returns: Result;
    };
    type RequestMethods = {
        solana_signMessage: Request<{
            message: string;
            pubkey: string;
        }, {
            signature: string;
        }>;
        solana_signTransaction: Request<{
            transaction: string;
            pubkey: string;
        }, {
            signature: string;
        } | {
            transaction: string;
        }>;
        solana_signAndSendTransaction: Request<{
            transaction: string;
            pubkey: string;
            sendOptions?: SendOptions;
        }, {
            signature: string;
        }>;
        solana_signAllTransactions: Request<{
            transactions: string[];
        }, {
            transactions: string[];
        }>;
    };
    type RequestMethod = keyof RequestMethods;
    type Account = {
        address: string;
        publicKey: Uint8Array;
    };
}
