import useSWRMutation from "swr/mutation";
import type {
  SWRMutationConfiguration,
  SWRMutationResponse,
} from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type {
  DeleteCommandMutationResponse,
  DeleteCommandPathParams,
} from "../models/DeleteCommand";

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
    mutation?: SWRMutationConfiguration<TData, TError, string, never, TData>;
  },
): SWRMutationResponse<TData, TError, string> {
  const { mutation: mutationOptions } = options ?? {};

  return useSWRMutation<TData, TError, string>(
    `/delete/${command}`,
    (url) => {
      return client<TData, TError>({
        method: "delete",
        url,
      });
    },
    mutationOptions,
  );
}
