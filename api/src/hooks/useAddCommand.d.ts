import type { SWRMutationConfiguration, SWRMutationResponse } from "swr/mutation";
import type { AddCommandMutationResponse, AddCommandPathParams, AddCommandQueryParams } from "../models/AddCommand";
/**
* @description Add a command to the config and reloads the cache. All strings will be trimmed, please URLEncode your url parameter.

* @summary Add a command  to the config
* @link /add/:command/:type
*/
export declare function useAddCommand<TData = AddCommandMutationResponse, TError = unknown>(command: AddCommandPathParams["command"], type: AddCommandPathParams["type"], params: AddCommandQueryParams, options?: {
    mutation?: SWRMutationConfiguration<TData, TError, string>;
}): SWRMutationResponse<TData, TError, string>;
