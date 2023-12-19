import type { Config } from "./Config";

export const addCommandPathParamsType = {
    "Alias": "Alias",
    "Redirect": "Redirect",
    "RedirectVarArgs": "RedirectVarArgs"
} as const;
export type AddCommandPathParamsType = (typeof addCommandPathParamsType)[keyof typeof addCommandPathParamsType];
export type AddCommandPathParams = {
    /**
     * @type string
    */
    command: string;
    /**
     * @type string
    */
    type: AddCommandPathParamsType;
};

export type AddCommandQueryParams = {
    /**
     * @type string
    */
    url: string;
};

/**
 * @description Command Added
*/
export type AddCommandMutationResponse = Config;
export namespace AddCommandMutation {
  export type Response = AddCommandMutationResponse;
  export type PathParams = AddCommandPathParams;
  export type QueryParams = AddCommandQueryParams;
}
