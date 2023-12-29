export const configEntriesType = {
  Alias: "Alias",
  Redirect: "Redirect",
  RedirectVarArgs: "RedirectVarArgs",
} as const;
export type ConfigEntriesType =
  (typeof configEntriesType)[keyof typeof configEntriesType];
export type Config = {
  /**
   * @description The host and port to bind too (eg 0.0.0.0:8080)
   * @type string
   */
  Bind: string;
  /**
   * @description Array of all current commands
   * @type array
   */
  Entries: {
    /**
     * @description The command to associate to this url
     * @type string | undefined
     */
    Command?: string;
    /**
     * @description A string following golang printf format for the URL to goto
     * @type string | undefined
     */
    Value?: string;
    /**
     * @description Execution mode for the command.  Must be one of "Alias", "Redirect", "RedirectVarArgs"
     * @type string | undefined
     */
    Type?: ConfigEntriesType;
  }[];
};
