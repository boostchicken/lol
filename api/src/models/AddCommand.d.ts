import type { Config } from "./Config";
export declare const addCommandPathParamsType: {
    readonly Alias: "Alias";
    readonly Redirect: "Redirect";
    readonly RedirectVarArgs: "RedirectVarArgs";
};
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
