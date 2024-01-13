import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type { GetConfigQueryResponse } from "../../models/GetConfig";

export function getConfigQueryOptions<
  TData = GetConfigQueryResponse,
  TError = unknown,
>(
  options: Partial<Parameters<typeof client>[0]> = {},
): SWRConfiguration<TData, TError> {
  return {
    fetcher: async () => {
      const res = await client<TData, TError>({
        method: "get",
        url: `/config`,

        ...options,
      });
      return res.data;
    },
  };
}

/**
 * @summary Get the current config in your format of choice
 * @link /config
 * @deprecated
 */

export function useGetConfig<
  TData = GetConfigQueryResponse,
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

  const url = shouldFetch ? `/config` : null;
  const query = useSWR<TData, TError, string | null>(url, {
    ...getConfigQueryOptions<TData, TError>(clientOptions),
    ...queryOptions,
  });

  return query;
}
