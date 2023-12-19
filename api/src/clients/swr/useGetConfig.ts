import useSWR from "swr";
import client from "@kubb/swagger-client/client";
import type { SWRConfiguration, SWRResponse } from "swr";
import type { GetConfigQueryResponse } from "../../models/GetConfig";

type GetConfigClient = typeof client<GetConfigQueryResponse, never, never>;
type GetConfig = {
  data: GetConfigQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: GetConfigQueryResponse;
  client: {
    parameters: Partial<Parameters<GetConfigClient>[0]>;
    return: Awaited<ReturnType<GetConfigClient>>;
  };
};
export function getConfigQueryOptions<
  TData extends GetConfig["response"] = GetConfig["response"],
  TError = GetConfig["error"],
>(
  options: GetConfig["client"]["parameters"] = {},
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
 * @deprecated */
export function useGetConfig<
  TData extends GetConfig["response"] = GetConfig["response"],
  TError = GetConfig["error"],
>(options?: {
  query?: SWRConfiguration<TData, TError>;
  client?: GetConfig["client"]["parameters"];
  shouldFetch?: boolean;
}): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/config` as const;
  const query = useSWR<TData, TError, typeof url | null>(
    shouldFetch ? url : null,
    {
      ...getConfigQueryOptions<TData, TError>(clientOptions),
      ...queryOptions,
    },
  );
  return query;
}
