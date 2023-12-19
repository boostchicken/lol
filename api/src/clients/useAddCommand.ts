import useSWRMutation from "swr/mutation";
import type { SWRMutationConfiguration, SWRMutationResponse } from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type { ResponseConfig } from "@kubb/swagger-client/client";
import type { AddCommandMutationResponse, AddCommandPathParams, AddCommandQueryParams } from "../models/AddCommand";

/**
 * @description Add a command to the config and reloads the cache. All strings will be trimmed, please URLEncode your url parameter.

 * @summary Add a command  to the config
 * @link /add/:command/:type
 */

export function useAddCommand <
  TData = AddCommandMutationResponse, TError = unknown
>(
  command: AddCommandPathParams["command"], type: AddCommandPathParams["type"], params?: AddCommandQueryParams, options?: {
          mutation?: SWRMutationConfiguration<ResponseConfig<TData>, TError, string | null, never>,
          client?: Partial<Parameters<typeof client<TData, TError>>[0]>,
          shouldFetch?: boolean,
        }
): SWRMutationResponse<ResponseConfig<TData>, TError, string | null, never> {
  const { mutation: mutationOptions, client: clientOptions = {}, shouldFetch = true } = options ?? {};
  
  const url = shouldFetch ? `/add/${command}/${type}` : null;
  return useSWRMutation<ResponseConfig<TData>, TError, string | null, never>(
    url,
    (url) => {
      return client<TData, TError>({
        method: "put",
        url,
        
        params,
        
        ...clientOptions,
      })
    },
    mutationOptions
  );
};
