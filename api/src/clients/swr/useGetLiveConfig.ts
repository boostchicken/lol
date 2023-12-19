import useSWR from "swr";
import client from "@kubb/swagger-client/client";
import type { SWRConfiguration, SWRResponse } from "swr";
import type { GetLiveConfigQueryResponse } from "../../models/GetLiveConfig";

type GetLiveConfigClient = typeof client<
  GetLiveConfigQueryResponse,
  never,
  never
>;
type GetLiveConfig = {
  data: GetLiveConfigQueryResponse;
  error: never;
  request: never;
  pathParams: never;
  queryParams: never;
  headerParams: never;
  response: GetLiveConfigQueryResponse;
  client: {
    parameters: Partial<Parameters<GetLiveConfigClient>[0]>;
    return: Awaited<ReturnType<GetLiveConfigClient>>;
  };
};
export function getLiveConfigQueryOptions<
  TData extends GetLiveConfig["response"] = GetLiveConfig["response"],
  TError = GetLiveConfig["error"],
>(
  options: GetLiveConfig["client"]["parameters"] = {},
): SWRConfiguration<TData, TError> {
  return {
    fetcher: async () => {
      const res = await client<TData, TError>({
        method: "get",
        url: `/liveconfig`,
        ...options,
      });
      return res.data;
    },
  };
}
/**
 * @description liveconfig has the same return value of config with a contentype of JSON.  However this just exists to make that fool proof
 * @summary Get current configuration in JSON for UI
 * @link /liveconfig */
export function useGetLiveConfig<
  TData extends GetLiveConfig["response"] = GetLiveConfig["response"],
  TError = GetLiveConfig["error"],
>(options?: {
  query?: SWRConfiguration<TData, TError>;
  client?: GetLiveConfig["client"]["parameters"];
  shouldFetch?: boolean;
}): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/liveconfig` as const;
  const query = useSWR<TData, TError, typeof url | null>(
    shouldFetch ? url : null,
    {
      ...getLiveConfigQueryOptions<TData, TError>(clientOptions),
      ...queryOptions,
    },
  );
  return query;
}
