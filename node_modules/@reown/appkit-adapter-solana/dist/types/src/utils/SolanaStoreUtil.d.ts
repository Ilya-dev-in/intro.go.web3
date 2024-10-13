import type { Connection } from '@reown/appkit-utils/solana';
export interface SolStoreUtilState {
    connection: Connection | null;
}
export declare const SolStoreUtil: {
    state: SolStoreUtilState;
    subscribeKey<K extends "connection">(key: K, callback: (value: SolStoreUtilState[K]) => void): () => void;
    subscribe(callback: (newState: SolStoreUtilState) => void): () => void;
    setConnection(connection: Connection): void;
};
