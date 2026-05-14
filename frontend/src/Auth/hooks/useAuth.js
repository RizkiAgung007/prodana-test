import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../../services/api";

export const useAuth = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setError("");
    setIsLoading(true);

    try {
      const response = await api.post("/login", { email, password });

      const token = response.data.data.token;
      sessionStorage.setItem("token", token);

      navigate("/dashboard");
    } catch (err) {
      setError(err.response?.data?.message || "Email atau password salah");
    } finally {
      setIsLoading(false);
    }
  };

  return {
    email,
    setEmail,
    password,
    setPassword,
    error,
    isLoading,
    handleLogin,
  };
};
