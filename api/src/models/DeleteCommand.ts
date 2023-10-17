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
