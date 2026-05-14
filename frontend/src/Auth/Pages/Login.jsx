import { useAuth } from "../hooks/useAuth";

export const Login = () => {
  const {
    email,
    setEmail,
    password,
    setPassword,
    error,
    isLoading,
    handleLogin,
  } = useAuth();

  return (
    <div className="flex min-h-screen items-center justify-center bg-[#F8FAFC] p-4 font-sans antialiased">
      <div className="w-full max-w-[400px] bg-white border border-slate-200 rounded-2xl shadow-sm p-8">
        {/* Header/Logo Section */}
        <div className="text-center mb-8">
          <h1 className="text-3xl font-extrabold text-slate-900 tracking-tight">
            Prodana
          </h1>
          <p className="text-slate-500 mt-2 text-sm">
            Masuk ke akun Anda untuk melanjutkan
          </p>
        </div>

        {error && (
          <div className="mb-6 p-3 bg-red-50 border border-red-100 text-red-600 rounded-lg text-sm text-center font-medium">
            {error}
          </div>
        )}

        <form onSubmit={handleLogin} className="space-y-5">
          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-slate-700">
              Email
            </label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
              placeholder="nama@email.com"
              required
            />
          </div>

          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-slate-700">
              Password
            </label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-2.5 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm placeholder:text-slate-400"
              placeholder="••••••••"
              required
            />
          </div>

          <div className="pt-2">
            <button
              type="submit"
              disabled={isLoading}
              className={`w-full text-white font-medium py-2.5 rounded-lg transition-colors text-sm
                ${
                  isLoading
                    ? "bg-slate-400 cursor-not-allowed"
                    : "bg-slate-900 hover:bg-slate-800"
                }`}
            >
              {isLoading ? "Memverifikasi..." : "Masuk ke Dashboard"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};