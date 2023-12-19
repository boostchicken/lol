import useSWR from "swr";
import client from "@kubb/swagger-client/client";
import type { SWRConfiguration, SWRResponse } from "swr";
import type {
  ExecuteLolQueryResponse,
  ExecuteLolQueryParams,
  ExecuteLol302,
  ExecuteLol404,
} from "../../models/ExecuteLol";

type ExecuteLolClient = typeof client<
  ExecuteLolQueryResponse,
  ExecuteLol302 | ExecuteLol404,
  never
>;
type ExecuteLol = {
  data: ExecuteLolQueryResponse;
  error: ExecuteLol302 | ExecuteLol404;
  request: never;
  pathParams: never;
  queryParams: ExecuteLolQueryParams;
  headerParams: never;
  response: ExecuteLolQueryResponse;
  client: {
    parameters: Partial<Parameters<ExecuteLolClient>[0]>;
    return: Awaited<ReturnType<ExecuteLolClient>>;
  };
};
export function executeLolQueryOptions<
  TData extends ExecuteLol["response"] = ExecuteLol["response"],
  TError = ExecuteLol["error"],
>(
  params?: ExecuteLol["queryParams"],
  options: ExecuteLol["client"]["parameters"] = {},
): SWRConfiguration<TData, TError> {
  return {
    fetcher: async () => {
      const res = await client<TData, TError>({
        method: "get",
        url: `/lol`,
        params,
        ...options,
      });
      return res.data;
    },
  };
}
/**
 * @description The main entry point of LOL, this is where everything happens
 * @summary Redirect user based on the command provided
 * @link /lol */
export function useExecuteLol<
  TData extends ExecuteLol["response"] = ExecuteLol["response"],
  TError = ExecuteLol["error"],
>(
  params?: ExecuteLol["queryParams"],
  options?: {
    query?: SWRConfiguration<TData, TError>;
    client?: ExecuteLol["client"]["parameters"];
    shouldFetch?: boolean;
  },
): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/lol` as const;
  const query = useSWR<TData, TError, [typeof url, typeof params] | null>(
    shouldFetch ? [url, params] : null,
    {
      ...executeLolQueryOptions<TData, TError>(params, clientOptions),
      ...queryOptions,
    },
  );
  return query;
}
