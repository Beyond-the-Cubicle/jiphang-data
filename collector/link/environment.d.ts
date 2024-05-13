declare global {
  namespace NodeJS {
    interface ProcessEnv {
      GG_KEY: string;
      SEOUL_KEY: string;
      DATA_GO_KR_KEY: string;
    }
  }
}

export {};
