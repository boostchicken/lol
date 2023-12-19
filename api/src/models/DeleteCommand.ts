import type { Config } from "./Config";

export type DeleteCommandPathParams = {
  /**
   * @type string
   */
  command: string;
};

/**
 * @description Command Deleted
 */
export type DeleteCommandMutationResponse = Config;
export namespace DeleteCommandMutation {
  export type Response = DeleteCommandMutationResponse;
  export type PathParams = DeleteCommandPathParams;
}
