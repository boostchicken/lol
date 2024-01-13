import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type { GetHistoryQueryResponse } from "@boostchicken/lol-api"

export function getHistoryQueryOptions<
  TData = GetHistoryQueryResponse,
  TError = unknown,
>(
  options: Partial<Parameters<typeof client>[0]> = {},
): SWRConfiguration<TData, TError> {
  return {
    fetcher: async () => {
      return client<TData, TError>({
        method: "get",
        url: `/history`,

        ...options,
      }).then((res) => res.data);
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
  client?: Partial<Parameters<typeof client<TData, TError>>[0]>;
  shouldFetch?: boolean;
}): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};

  const url = shouldFetch ? `/history` : null;
  const query = useSWR<TData, TError, string | null>(url, {
    ...getHistoryQueryOptions<TData, TError>(clientOptions),
    ...queryOptions,
  });

  return query;
}
