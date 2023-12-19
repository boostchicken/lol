import useSWR from "swr";
import client from "@kubb/swagger-client/client";
import type { SWRConfiguration, SWRResponse } from "swr";
import type { GetHistoryQueryResponse } from "../../models/GetHistory";

type GetHistoryClient = typeof client<GetHistoryQueryResponse, never, never>;
type GetHistory = {
  data: GetHistoryQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: GetHistoryQueryResponse;
  client: {
    parameters: Partial<Parameters<GetHistoryClient>[0]>;
    return: Awaited<ReturnType<GetHistoryClient>>;
  };
};
export function getHistoryQueryOptions<
  TData extends GetHistory["response"] = GetHistory["response"],
  TError = GetHistory["error"],
>(
  options: GetHistory["client"]["parameters"] = {},
): SWRConfiguration<TData, TError> {
  return {
    fetcher: async () => {
      const res = await client<TData, TError>({
        method: "get",
        url: `/history`,
        ...options,
      });
      return res.data;
    },
  };
}
/**
 * @summary Get all history tab entries (max 250)
 * @link /history */
export function useGetHistory<
  TData extends GetHistory["response"] = GetHistory["response"],
  TError = GetHistory["error"],
>(options?: {
  query?: SWRConfiguration<TData, TError>;
  client?: GetHistory["client"]["parameters"];
  shouldFetch?: boolean;
}): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/history` as const;
  const query = useSWR<TData, TError, typeof url | null>(
    shouldFetch ? url : null,
    {
      ...getHistoryQueryOptions<TData, TError>(clientOptions),
      ...queryOptions,
    },
  );
  return query;
}
