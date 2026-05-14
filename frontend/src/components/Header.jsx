import { Link, useLocation } from "react-router-dom";

export const Header = ({ roleId, handleLogout }) => {
  const location = useLocation();
  const isUserPage = location.pathname === "/dashboard";
  const isProductPage = location.pathname === "/products";

  return (
    <div className="flex flex-col md:flex-row md:justify-between md:items-end mb-8 pb-6 border-b border-slate-200 gap-4">
      <div>
        <h1 className="text-2xl font-bold text-slate-900 tracking-tight">
          {isUserPage ? "Manajemen User" : "Manajemen Produk"}
        </h1>

        <div className="mt-3 flex items-center space-x-6 text-sm">
          <Link
            to="/dashboard"
            className={`${
              isUserPage
                ? "text-slate-900 font-semibold border-b-2 border-slate-900 pb-1"
                : "text-slate-500 hover:text-slate-900 transition-colors"
            }`}
          >
            Data Users
          </Link>

          <Link
            to="/products"
            className={`${
              isProductPage
                ? "text-slate-900 font-semibold border-b-2 border-slate-900 pb-1"
                : "text-slate-500 hover:text-slate-900 transition-colors"
            }`}
          >
            Data Produk
          </Link>
        </div>
      </div>

      <div className="flex items-center gap-4">
        <div className="text-sm text-slate-500 text-right">
          Login sebagai:{" "}
          <span className="font-medium text-indigo-600 bg-indigo-50 px-2 py-1 rounded-md ml-1">
            {roleId === 1 ? "Admin" : roleId === 2 ? "Editor" : "Viewer"}
          </span>
        </div>
        <button
          onClick={handleLogout}
          className="cursor-pointer text-sm font-medium text-slate-600 bg-white border border-slate-200 hover:bg-slate-50 hover:text-red-600 px-4 py-2 rounded-lg transition-all"
        >
          Logout
        </button>
      </div>
    </div>
  );
};