import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type { GetHistoryQueryResponse } from "../models/GetHistory";

export function getHistoryQueryOptions<
  TData = GetHistoryQueryResponse,
  TError = unknown,
>(): SWRConfiguration<TData, TError> {
  return {
    fetcher: () => {
      return client<TData, TError>({
        method: "get",
        url: `/history`,
      });
    },
  };
}

/**
 * @summary Get all history tab entries (max 250)
 * @link /history
 */
export function useGetHistory<
  TData = GetHistoryQueryResponse,
  TError = unknown,
>(options?: {
  query?: SWRConfiguration<TData, TError>;
}): SWRResponse<TData, TError> {
  const { query: queryOptions } = options ?? {};

  const query = useSWR<TData, TError, string>(`/history`, {
    ...getHistoryQueryOptions<TData, TError>(),
    ...queryOptions,
  });

  return query;
}
