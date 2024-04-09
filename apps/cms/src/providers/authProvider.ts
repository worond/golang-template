import { AuthProvider } from 'react-admin';
import { UserManager } from 'oidc-client-ts';

import { config } from './configProvider';

const issuer = config.OIDC_ISSUER;
const clientId = config.OIDC_CLIENT_ID;
const redirectUri = config.OIDC_REDIRECT_URI;
const apiUrl = config.API_URL;

const userManager = new UserManager({
  authority: issuer,
  client_id: clientId,
  redirect_uri: redirectUri,
  response_type: 'code',
  scope: 'openid email profile', // Allow to retrieve the email and user name later api side
});

const getProfileFromToken = (tokenJson: string | null) => {
  const token = JSON.parse(tokenJson || '');
  const jwt = JSON.parse(atob(token.access_token.split('.')[1]));

  return { id: 'my-profile', ...jwt };
};

const cleanup = () => {
  // Remove the ?code&state from the URL
  window.history.replaceState(
    {},
    window.document.title,
    window.location.origin
  );
};

export const authProvider: AuthProvider = {
  login: async () => {
    // 1. Redirect to the issuer to ask authentication
    await userManager.signinRedirect();
    return; // Do not return anything, the login is still loading
  },
  logout: () => {
    localStorage.removeItem('token');
    return Promise.resolve();
  },
  checkError: () => {
    localStorage.removeItem('token');
    return Promise.resolve();
  },
  checkAuth: () => {
    const token = localStorage.getItem('token');

    if (!token) {
      return Promise.reject();
    }

    // This is specific to the Google authentication implementation
    const jwt = getProfileFromToken(token);
    const now = new Date();

    return now.getTime() > jwt.exp * 1000
      ? Promise.reject()
      : Promise.resolve();
  },
  getPermissions: () => Promise.resolve(),
  getIdentity: () => {
    const token = window.localStorage.getItem('token');
    const profile = getProfileFromToken(token);

    return Promise.resolve({
      id: profile.sub,
      fullName: profile.name,
      avatar: profile.picture,
    });
  },
  handleCallback: async () => {
    // We came back from the issuer with ?code infos in query params
    const { searchParams } = new URL(window.location.href);
    const code = searchParams.get("code");
    const state = searchParams.get("state");

    // oidc-client uses localStorage to keep a temporary state
    // between the two redirections. But since we need to send it to the API
    // we have to retrieve it manually
    const stateKey = `oidc.${state}`;
    const { code_verifier } = JSON.parse(
      localStorage.getItem(stateKey) || "{}"
    );

    // Transform the code to a token via the API
    const response = await fetch(`${apiUrl}/auth/token`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ code: code, code_verifier }),
    });

    if (!response.ok) {
      cleanup();
      return Promise.reject();
    }

    const token = await response.json();

    localStorage.setItem("token", JSON.stringify(token));
    userManager.clearStaleState();
    cleanup();
    return Promise.resolve();
  },
};
