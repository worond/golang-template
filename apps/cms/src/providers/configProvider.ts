const configProvider = {
  requireString(key: string): string {
    const value = import.meta.env[key];
    if (!value) {
      throw new Error(`Required environment variable '${key}' not defined`);
    }
    return value;
  },
};

export const config = {
  API_URL: configProvider.requireString("VITE_API_URL"),
  OIDC_ISSUER: configProvider.requireString("VITE_OIDC_ISSUER"),
  OIDC_CLIENT_ID: configProvider.requireString("VITE_OIDC_CLIENT_ID"),
  OIDC_REDIRECT_URI: configProvider.requireString("VITE_OIDC_REDIRECT_URI"),
};
