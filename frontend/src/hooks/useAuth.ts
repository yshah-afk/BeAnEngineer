import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { authApi } from "@/lib/api";
import { useAuthStore } from "@/stores/authStore";
import type { LoginPayload, RegisterPayload } from "@/types";

export function useAuth() {
  const { user, isAuthenticated, setAuth, setUser, logout: storeLogout } = useAuthStore();
  const queryClient = useQueryClient();
  const navigate = useNavigate();

  useQuery({
    queryKey: ["auth", "me"],
    queryFn: async () => {
      const userData = await authApi.me();
      setUser(userData);
      return userData;
    },
    enabled: isAuthenticated && !user,
    retry: false,
  });

  const loginMutation = useMutation({
    mutationFn: (payload: LoginPayload) => authApi.login(payload),
    onSuccess: (data) => {
      setAuth(data.user, data.accessToken);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
      navigate("/dashboard");
    },
  });

  const registerMutation = useMutation({
    mutationFn: (payload: RegisterPayload) => authApi.register(payload),
    onSuccess: (data) => {
      setAuth(data.user, data.accessToken);
      queryClient.invalidateQueries({ queryKey: ["auth"] });
      navigate("/dashboard");
    },
  });

  const logout = () => {
    storeLogout();
    queryClient.clear();
    navigate("/");
  };

  return {
    user,
    isAuthenticated,
    login: loginMutation.mutate,
    register: registerMutation.mutate,
    logout,
    isLoggingIn: loginMutation.isPending,
    isRegistering: registerMutation.isPending,
    loginError: loginMutation.error,
    registerError: registerMutation.error,
  };
}
