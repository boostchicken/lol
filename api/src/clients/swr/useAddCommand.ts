import useSWRMutation from "swr/mutation";
import client from "@kubb/swagger-client/client";
import type {
  SWRMutationConfiguration,
  SWRMutationResponse,
} from "swr/mutation";
import type {
  AddCommandMutationResponse,
  AddCommandPathParams,
  AddCommandQueryParams,
} from "../../models/AddCommand";

type AddCommandClient = typeof client<AddCommandMutationResponse, never, never>;
type AddCommand = {
  data: AddCommandMutationResponse;
  error: never;
  request: never;
  pathParams: AddCommandPathParams;
  queryParams: AddCommandQueryParams;
  headerParams: never;
  response: AddCommandMutationResponse;
  client: {
    parameters: Partial<Parameters<AddCommandClient>[0]>;
    return: Awaited<ReturnType<AddCommandClient>>;
  };
};
/**
 * @description Add a command to the config and reloads the cache. All strings will be trimmed, please URLEncode your url parameter.
 * @summary Add a command  to the config
 * @link /add/:command/:type */
export function useAddCommand(
  command: AddCommandPathParams["command"],
  type: AddCommandPathParams["type"],
  params?: AddCommand["queryParams"],
  options?: {
    mutation?: SWRMutationConfiguration<
      AddCommand["response"],
      AddCommand["error"]
    >;
    client?: AddCommand["client"]["parameters"];
    shouldFetch?: boolean;
  },
): SWRMutationResponse<AddCommand["response"], AddCommand["error"]> {
  const {
    mutation: mutationOptions,
    client: clientOptions = {},
    shouldFetch = true,
  } = options ?? {};
  const url = `/add/${command}/${type}` as const;
  return useSWRMutation<
    AddCommand["response"],
    AddCommand["error"],
    [typeof url, typeof params] | null
  >(
    shouldFetch ? [url, params] : null,
    async (_url) => {
      const res = await client<AddCommand["data"], AddCommand["error"]>({
        method: "put",
        url,
        params,
        ...clientOptions,
      });
      return res.data;
    },
    mutationOptions,
  );
}
