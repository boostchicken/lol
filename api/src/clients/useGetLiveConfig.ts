import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type { GetLiveConfigQueryResponse } from "../models/GetLiveConfig";


export function getLiveConfigQueryOptions <
  TData = GetLiveConfigQueryResponse, TError = unknown
>(
  options: Partial<Parameters<typeof client>[0]> = {}
): SWRConfiguration<TData, TError> {
  return {
    fetcher: () => {
      return client<TData, TError>({
        method: "get",
        url: `/liveconfig`,
        
        
        
        ...options,
      }).then(res => res.data);
    },
  };
};

/**
 * @description liveconfig has the same return value of config with a contentype of JSON.  However this just exists to make that fool proof
 * @summary Get current configuration in JSON for UI
 * @link /liveconfig
 */

export function useGetLiveConfig <TData = GetLiveConfigQueryResponse, TError = unknown>(options?: { 
          query?: SWRConfiguration<TData, TError>,
          client?: Partial<Parameters<typeof client<TData, TError>>[0]>,
          shouldFetch?: boolean,
        }): SWRResponse<TData, TError> {
  const { query: queryOptions, client: clientOptions = {}, shouldFetch = true } = options ?? {};
  
  const url = shouldFetch ? `/liveconfig` : null;
  const query = useSWR<TData, TError, string | null>(url, {
    ...getLiveConfigQueryOptions<TData, TError>(clientOptions),
    ...queryOptions
  });

  return query;
};
