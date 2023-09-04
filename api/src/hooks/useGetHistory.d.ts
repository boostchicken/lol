import type { SWRConfiguration, SWRResponse } from "swr";
import type { GetHistoryQueryResponse } from "../models/GetHistory";
export declare function getHistoryQueryOptions<TData = GetHistoryQueryResponse, TError = unknown>(): SWRConfiguration<TData, TError>;
/**
 * @summary Get all history tab entries (max 250)
 * @link /history
 */
export declare function useGetHistory<TData = GetHistoryQueryResponse, TError = unknown>(options?: {
    query?: SWRConfiguration<TData, TError>;
}): SWRResponse<TData, TError>;
