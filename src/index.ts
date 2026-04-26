import { Container, getContainer } from "@cloudflare/containers";

export class MascotsContainer extends Container {
  defaultPort = 8080;
  sleepAfter = "3m";

  // Env vars forwarded into the container process.
  envVars = {
    PORT: "8080",
    D1_DSN: (this.env as Env).D1_DSN,
    CF_ACCESS_CLIENT_ID: (this.env as Env).CF_ACCESS_CLIENT_ID,
    CF_ACCESS_CLIENT_SECRET: (this.env as Env).CF_ACCESS_CLIENT_SECRET,
  };
}

export interface Env {
  MASCOTS: DurableObjectNamespace<MascotsContainer>;
  D1_DSN: string;
  CF_ACCESS_CLIENT_ID: string;
  CF_ACCESS_CLIENT_SECRET: string;
}

export default {
  async fetch(req: Request, env: Env): Promise<Response> {
    const c = getContainer(env.MASCOTS, "singleton");
    return c.fetch(req);
  },
};
