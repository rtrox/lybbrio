import React from "react";
import { useState, useCallback, useEffect, createContext } from "react";
import { TokenResponse, useToken } from "../hooks/useToken";
import { useDebounce } from "../hooks/useDebounce";

enum AuthEvents {
  LOGIN = "login",
  LOGOUT = "logout",
}

export const UnauthorizedError = new Error("Unauthorized");
export const ForbiddenError = new Error("Forbidden");
export const InternalServerError = new Error("Internal Server Error");
export const UnknownError = new Error("Unknown Error");

export interface UserBase {
  name: string;
  email: string;
  password: string;
}

export interface UserPermissions {
  admin: boolean;
  can_create_public: boolean;
  can_edit: boolean;
}

export interface User extends UserBase {
  id: string;
  userPermissions: UserPermissions;
}

interface UserAndTokenResponse extends TokenResponse {
  user: User;
}

export interface AuthContext {
  user: User | null;
  setUser: (user: User) => void;
  loggedIn: boolean;
  login: (username: string, password: string) => Promise<Response | void>;
  logout: () => Promise<Response | void>;
  refreshToken: () => void;
  initialized: boolean;
  fetcher: typeof fetch;
}

const AuthContext = createContext<AuthContext | null>(null);

interface AuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<User | null>(null);
  const loggedIn = !!user;

  const [initialized, setInitialized] = useState(false);
  const debouncedRefresh = useDebounce(refresh, 100);
  const refreshToken = useCallback(refresh, []);
  const debouncedRefreshToken = useCallback(debouncedRefresh, []);

  const onTokenInvalid = () => setUser(null);
  const { setToken, clearToken, isAuthenticated, fetcher } = useToken(
    onTokenInvalid,
    refreshToken
  );

  useEffect(() => {
    refreshToken().finally(() => setInitialized(true));
  }, [refreshToken]);

  useEffect(() => {
    window.addEventListener(
      "storage",
      async (event: WindowEventMap["storage"]) => {
        if (event.key === AuthEvents.LOGOUT && isAuthenticated()) {
          await clearToken(false);
          setUser(null);
        } else if (event.key === AuthEvents.LOGIN && !isAuthenticated()) {
          // Debounce the refresh token to avoid multiple requests
          debouncedRefreshToken();
        }
      }
    );
  }, [clearToken, isAuthenticated, refreshToken]);

  const logout = useCallback(async () => {
    return clearToken().then(() => {
      setUser(null);
      window.localStorage.setItem(AuthEvents.LOGOUT, new Date().toISOString());
    });
  }, [clearToken]);

  const login = useCallback(
    async (username: string, password: string) => {
      return fetcher("/auth/password", {
        method: "POST",
        body: JSON.stringify({ username, password }),
      })
        .then((res) => {
          if (res.ok) {
            return res.json();
          }
          switch (res.status) {
            case 401:
              throw UnauthorizedError;
            case 403:
              throw ForbiddenError;
            case 500:
              throw InternalServerError;
            default:
              throw UnknownError;
          }
          throw new Error("Login failed");
        })
        .then((data: UserAndTokenResponse) => {
          setToken(data);
          setUser(data.user);
          window.localStorage.setItem(
            AuthEvents.LOGIN,
            new Date().toISOString()
          );
          console.log("Logged in. User: ", data.user);
        })
        .catch((error) => {
          console.error("Login failed: ", error);
          throw error;
        });
    },
    [fetcher, setToken]
  );

  async function refresh() {
    const response = await fetcher("/auth/refresh");
    if (response.ok) {
      const res: UserAndTokenResponse = await response.json();
      setUser(res.user);
      setToken(res);
      response.headers.get("X-Request-ID");
      console.log("Token Refreshed.");
    }
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        setUser,
        loggedIn,
        login,
        logout,
        refreshToken,
        initialized,
        fetcher,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = React.useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
