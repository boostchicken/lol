import useSWRMutation from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type {
  SWRMutationConfiguration,
  SWRMutationResponse,
} from "swr/mutation";
import type {
  DeleteCommandMutationResponse,
  DeleteCommandPathParams,
} from "../../models/DeleteCommand";

type DeleteCommandClient = typeof client<
  DeleteCommandMutationResponse,
  never,
  never
>;
type DeleteCommand = {
  data: DeleteCommandMutationResponse;
  error: never;
  request: never;
  pathParams: DeleteCommandPathParams;
  queryParams: never;
  headerParams: never;
  response: DeleteCommandMutationResponse;
  client: {
    parameters: Partial<Parameters<DeleteCommandClient>[0]>;
    return: Awaited<ReturnType<DeleteCommandClient>>;
  };
};
/**
 * @summary Delete a command from the config
 * @link /delete/:command */
export function useDeleteCommand(
  command: DeleteCommandPathParams["command"],
  options?: {
    mutation?: SWRMutationConfiguration<
      DeleteCommand["response"],
      DeleteCommand["error"]
    >;
    client?: DeleteCommand["client"]["parameters"];
    shouldFetch?: boolean;
  },
): SWRMutationResponse<DeleteCommand["response"], DeleteCommand["error"]> {
  const {
    mutation: mutationOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/delete/${command}` as const;
  return useSWRMutation<
    DeleteCommand["response"],
    DeleteCommand["error"],
    typeof url | null
  >(
    shouldFetch ? url : null,
    async (_url) => {
      const res = await client<DeleteCommand["data"], DeleteCommand["error"]>({
        method: "delete",
        url,
        ...clientOptions,
      });
      return res.data;
    },
    mutationOptions,
  );
}
