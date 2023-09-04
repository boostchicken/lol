import type { SWRMutationConfiguration, SWRMutationResponse } from "swr/mutation";
import type { DeleteCommandMutationResponse, DeleteCommandPathParams } from "../models/DeleteCommand";
/**
 * @summary Delete a command from the config
 * @link /delete/:command
 */
export declare function useDeleteCommand<TData = DeleteCommandMutationResponse, TError = unknown>(command: DeleteCommandPathParams["command"], options?: {
    mutation?: SWRMutationConfiguration<TData, TError, string>;
}): SWRMutationResponse<TData, TError, string>;
