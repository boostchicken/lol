
/**
 * @description Command found and redirecting you to your destination
*/
export type ExecuteLol302 = any | null;

/**
 * @description Command not found
*/
export type ExecuteLol404 = any | null;

export type ExecuteLolQueryParams = {
    /**
     * @description This is your LOL query (eg github boostchicken lol)
     * @type string
    */
    q: string;
};

export type ExecuteLolQueryResponse = any | null;
