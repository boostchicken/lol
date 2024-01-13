import useSWR from "swr";
import type { SWRConfiguration, SWRResponse } from "swr";
import client from "@kubb/swagger-client/client";
import type {
  ExecuteLolQueryResponse,
  ExecuteLolQueryParams,
  ExecuteLol302,
  ExecuteLol404,
} from "../../models/ExecuteLol";

export function executeLolQueryOptions<
  TData = ExecuteLolQueryResponse,
  TError = ExecuteLol302 | ExecuteLol404,
>(
  params?: ExecuteLolQueryParams,
  options: Partial<Parameters<typeof client>[0]> = {},
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
 * @link /lol
 */

export function useExecuteLol<
  TData = ExecuteLolQueryResponse,
  TError = ExecuteLol302 | ExecuteLol404,
>(
  params?: ExecuteLolQueryParams,
  options?: {
    query?: SWRConfiguration<TData, TError>;
    client?: Partial<Parameters<typeof client<TData, TError>>[0]>;
    shouldFetch?: boolean;
  },
): SWRResponse<TData, TError> {
  const {
    query: queryOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};

  const url = shouldFetch ? `/lol` : null;
  const query = useSWR<TData, TError, string | null>(url, {
    ...executeLolQueryOptions<TData, TError>(params, clientOptions),
    ...queryOptions,
  });

  return query;
}
