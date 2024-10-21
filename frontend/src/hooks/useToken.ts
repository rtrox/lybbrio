import { useEffect, useRef, useState, useCallback } from "react";

export interface AccessToken {
  token: string;
  issued_at: string;
  expires_at: string;
}
export interface TokenResponse {
  access_token: AccessToken;
}

const TOKEN_REFRESH_LEEWAY = 1000 * 30; // 30 seconds

export function useToken(
  onTokenInvalid: Function,
  onRefreshRequired: Function
) {
  const accessToken = useRef<string>();
  const { clearAutomaticTokenRefresh, setTokenExpiration } =
    useTokenExpiration(onRefreshRequired);

  const setToken = useCallback(
    (res: TokenResponse) => {
      accessToken.current = res.access_token.token;
      const expirationDate = new Date(res.access_token.expires_at);
      setTokenExpiration(expirationDate);
    },
    [setTokenExpiration]
  );

  const isAuthenticated = useCallback(() => {
    return !!accessToken.current;
  }, []);

  const clearToken = useCallback(
    (shouldClearCookie = true) => {
      const clearRefreshTokenCookie = shouldClearCookie
        ? fetch("auth/logout")
        : Promise.resolve();

      return clearRefreshTokenCookie.finally(() => {
        accessToken.current = "";
        clearAutomaticTokenRefresh();
      });
    },
    [clearAutomaticTokenRefresh]
  );

  const fetcher = useCallback(
    async (url: RequestInfo | URL, options?: RequestInit) => {
      const headers = new Headers(options?.headers);
      if (accessToken.current) {
        headers.set("Authorization", `Bearer ${accessToken.current}`);
      }

      const response = await fetch(url, {
        ...options,
        headers,
      });

      if (response.status === 401 && accessToken.current) {
        clearToken();
        onTokenInvalid();
        return Promise.reject(new Error("Unauthorized"));
      }

      return response;
    },
    []
  );

  return { setToken, clearToken, isAuthenticated, fetcher };
}

export function useTokenExpiration(onTokenRefreshRequired: Function) {
  const clearAutomaticRefresh = useRef<number>();
  const [tokenExpiration, setTokenExpiration] = useState<Date>();

  useEffect(() => {
    if (tokenExpiration instanceof Date && !isNaN(tokenExpiration.valueOf())) {
      const now = new Date();
      const triggerAfterMs =
        tokenExpiration.getTime() - now.getTime() - TOKEN_REFRESH_LEEWAY;
      clearAutomaticRefresh.current = window.setTimeout(async () => {
        onTokenRefreshRequired();
      }, triggerAfterMs);
    }

    return () => {
      window.clearTimeout(clearAutomaticRefresh.current);
    };
  }, [onTokenRefreshRequired, tokenExpiration]);

  const clearAutomaticTokenRefresh = () => {
    window.clearTimeout(clearAutomaticRefresh.current);
    setTokenExpiration(undefined);
  };

  return { setTokenExpiration, clearAutomaticTokenRefresh };
}
