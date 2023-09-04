import type { SWRConfiguration, SWRResponse } from "swr";
import type { GetLiveConfigQueryResponse } from "../models/GetLiveConfig";
export declare function getLiveConfigQueryOptions<TData = GetLiveConfigQueryResponse, TError = unknown>(): SWRConfiguration<TData, TError>;
/**
 * @description liveconfig has the same return value of config with a contentype of JSON.  However this just exists to make that fool proof
 * @summary Get current configuration in JSON for UI
 * @link /liveconfig
 */
export declare function useGetLiveConfig<TData = GetLiveConfigQueryResponse, TError = unknown>(options?: {
    query?: SWRConfiguration<TData, TError>;
}): SWRResponse<TData, TError>;
