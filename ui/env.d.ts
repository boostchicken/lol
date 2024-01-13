import { KVNamespace } from '@cloudflare/workers-types';
declare global {
    namespace NodeJS {
        interface ProcessEnv {
            [key: string]: string | undefined;
            // The KV Namespace binding type used here comes
            // from `@cloudflare/workers-types`, in order to
            // use it like so, make sure that you have installed
            // the package as a dev dependency and you have added
            // it to your `tsconfig.json` file under
            // `compilerOptions.types`.
            MY_KV: KVNamespace;
        }
    }
}
  
  export {};