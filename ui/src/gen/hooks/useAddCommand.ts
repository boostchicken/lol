import useSWRMutation from "swr/mutation";
import type {
  SWRMutationConfiguration,
  SWRMutationResponse,
} from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type {
  AddCommandMutationResponse,
  AddCommandPathParams,
  AddCommandQueryParams,
} from "../models/AddCommand";

/**
* @description Add a command to the config and reloads the cache. All strings will be trimmed, please URLEncode your url parameter.

* @summary Add a command  to the config
* @link /add/:command/:type
*/
export function useAddCommand<
  TData = AddCommandMutationResponse,
  TError = unknown,
>(
  command: AddCommandPathParams["command"],
  type: AddCommandPathParams["type"],
  params: AddCommandQueryParams,
  options?: {
    mutation?: SWRMutationConfiguration<TData, TError, string, never, TData>;
  },
): SWRMutationResponse<TData, TError, string> {
  const { mutation: mutationOptions } = options ?? {};

  return useSWRMutation<TData, TError, string>(
    `/add/${command}/${type}`,
    (url) => {
      return client<TData, TError>({
        method: "put",
        url,

        params,
      });
    },
    mutationOptions,
  );
}
