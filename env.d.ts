/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_BASE_HOST: string;
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
