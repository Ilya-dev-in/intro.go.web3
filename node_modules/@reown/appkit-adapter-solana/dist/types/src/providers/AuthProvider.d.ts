import type { AnyTransaction, Connection, GetActiveChain, Provider } from '@reown/appkit-utils/solana';
import { ProviderEventEmitter } from './shared/ProviderEventEmitter.js';
import { PublicKey, type SendOptions } from '@solana/web3.js';
import { W3mFrameProvider, type W3mFrameProviderMethods as ProviderAuthMethods } from '@reown/appkit-wallet';
import type { CaipNetwork, ChainNamespace } from '@reown/appkit-common';
export type AuthProviderConfig = {
    getProvider: () => W3mFrameProvider;
    getActiveChain: GetActiveChain;
    getActiveNamespace: () => ChainNamespace | undefined;
    getSession: () => AuthProvider.Session | undefined;
    setSession: (session: AuthProvider.Session | undefined) => void;
    chains: CaipNetwork[];
};
export declare class AuthProvider extends ProviderEventEmitter implements Provider, ProviderAuthMethods {
    readonly name: string;
    readonly type = "AUTH";
    private readonly getProvider;
    private readonly getActiveChain;
    private readonly getActiveNamespace;
    private readonly requestedChains;
    private readonly getSession;
    private readonly setSession;
    constructor({ getProvider, getActiveChain, getActiveNamespace, getSession, setSession, chains }: AuthProviderConfig);
    get publicKey(): PublicKey | undefined;
    get chains(): CaipNetwork[];
    connect(): Promise<string>;
    disconnect(): Promise<void>;
    signMessage(message: Uint8Array): Promise<Uint8Array>;
    signTransaction<T extends AnyTransaction>(transaction: T): Promise<T>;
    signAndSendTransaction<T extends AnyTransaction>(transaction: T, options?: SendOptions): Promise<any>;
    sendTransaction(transaction: AnyTransaction, connection: Connection, options?: SendOptions): Promise<string>;
    signAllTransactions<T extends AnyTransaction[]>(transactions: T): Promise<T>;
    connectEmail: ProviderAuthMethods['connectEmail'];
    connectOtp: ProviderAuthMethods['connectOtp'];
    updateEmail: ProviderAuthMethods['updateEmail'];
    updateEmailPrimaryOtp: ProviderAuthMethods['updateEmailPrimaryOtp'];
    updateEmailSecondaryOtp: ProviderAuthMethods['updateEmailSecondaryOtp'];
    getEmail: ProviderAuthMethods['getEmail'];
    getSocialRedirectUri: ProviderAuthMethods['getSocialRedirectUri'];
    connectDevice: ProviderAuthMethods['connectDevice'];
    connectSocial: ProviderAuthMethods['connectSocial'];
    connectFarcaster: ProviderAuthMethods['connectFarcaster'];
    getFarcasterUri: ProviderAuthMethods['getFarcasterUri'];
    syncTheme: ProviderAuthMethods['syncTheme'];
    syncDappData: ProviderAuthMethods['syncDappData'];
    switchNetwork: ProviderAuthMethods['switchNetwork'];
    private getPublicKey;
    private serializeTransaction;
    private bindEvents;
}
export declare namespace AuthProvider {
    type Session = Awaited<ReturnType<W3mFrameProvider['connect']>>;
}
