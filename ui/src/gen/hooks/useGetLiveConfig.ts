import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type { GetLiveConfigQueryResponse } from "../models/GetLiveConfig";

export function getLiveConfigQueryOptions<
  TData = GetLiveConfigQueryResponse,
  TError = unknown,
>(): SWRConfiguration<TData, TError> {
  return {
    fetcher: () => {
      return client<TData, TError>({
        method: "get",
        url: `/liveconfig`,
      });
    },
  };
}

/**
 * @description liveconfig has the same return value of config with a contentype of JSON.  However this just exists to make that fool proof
 * @summary Get current configuration in JSON for UI
 * @link /liveconfig
 */
export function useGetLiveConfig<
  TData = GetLiveConfigQueryResponse,
  TError = unknown,
>(options?: {
  query?: SWRConfiguration<TData, TError>;
}): SWRResponse<TData, TError> {
  const { query: queryOptions } = options ?? {};

  const query = useSWR<TData, TError, string>(`/liveconfig`, {
    ...getLiveConfigQueryOptions<TData, TError>(),
    ...queryOptions,
  });

  return query;
}
