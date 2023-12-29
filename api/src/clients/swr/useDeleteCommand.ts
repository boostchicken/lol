import useSWRMutation from "swr/mutation";
import type {
  SWRMutationConfiguration,
  SWRMutationResponse,
} from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type { ResponseConfig } from "@kubb/swagger-client/client";
import type {
  DeleteCommandMutationResponse,
  DeleteCommandPathParams,
} from "../../models/DeleteCommand";

/**
 * @summary Delete a command from the config
 * @link /delete/:command
 */

export function useDeleteCommand<
  TData = DeleteCommandMutationResponse,
  TError = unknown,
>(
  command: DeleteCommandPathParams["command"],
  options?: {
    mutation?: SWRMutationConfiguration<
      ResponseConfig<TData>,
      TError,
      string | null,
      never
    >;
    client?: Partial<Parameters<typeof client<TData, TError>>[0]>;
    shouldFetch?: boolean;
  },
): SWRMutationResponse<ResponseConfig<TData>, TError, string | null, never> {
  const {
    mutation: mutationOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};

  const url = shouldFetch ? `/delete/${command}` : null;
  return useSWRMutation<ResponseConfig<TData>, TError, string | null, never>(
    url,
    (url) => {
      return client<TData, TError>({
        method: "delete",
        url,

        ...clientOptions,
      });
    },
    mutationOptions,
  );
}
