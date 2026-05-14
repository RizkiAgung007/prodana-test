export const UserTable = ({ users, roleId, handleEditClick, handleDelete }) => {
  const getRoleBadge = (role) => {
    switch (role) {
      case "Admin":
        return "bg-indigo-50 text-indigo-700 ring-1 ring-indigo-200";
      case "Editor":
        return "bg-amber-50 text-amber-700 ring-1 ring-amber-200";
      default:
        return "bg-slate-50 text-slate-600 ring-1 ring-slate-200";
    }
  };

  return (
    <div className="bg-white border border-slate-200 rounded-xl shadow-sm overflow-hidden">
      <div className="overflow-x-auto">
        <table className="w-full text-left text-sm border-collapse">
          <thead>
            <tr className="bg-slate-50/50 border-b border-slate-200">
              <th className="py-3.5 px-6 font-semibold text-slate-600 w-20">
                ID
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 min-w-[150px]">
                Nama
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 min-w-[200px]">
                Email
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 w-32">
                Role
              </th>
              {roleId !== 3 && (
                <th className="py-3.5 px-6 font-semibold text-slate-600 text-right w-40">
                  Aksi
                </th>
              )}
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {users.map((user) => (
              <tr
                key={user.id}
                className="hover:bg-slate-50/80 transition-colors align-middle"
              >
                <td className="py-4 px-6 text-slate-500 font-mono text-xs">
                  #{user.id}
                </td>
                <td className="py-4 px-6">
                  <div className="font-semibold text-slate-900">
                    {user.name}
                  </div>
                </td>
                <td className="py-4 px-6 text-slate-600 italic">
                  {user.email}
                </td>
                <td className="py-4 px-6">
                  <span
                    className={`inline-flex items-center px-2.5 py-1 rounded-full text-[10px] font-bold uppercase tracking-wider ${getRoleBadge(user.role)}`}
                  >
                    {user.role}
                  </span>
                </td>

                {(roleId === 1 || (roleId === 2 && user.role === "Viewer")) && (
                  <td className="py-4 px-6 text-right whitespace-nowrap">
                    <div className="flex justify-end gap-4">
                      <button
                        onClick={() => handleEditClick(user)}
                        className="text-indigo-600 hover:text-indigo-800 font-bold text-xs uppercase tracking-wider transition-colors cursor-pointer"
                      >
                        Edit
                      </button>
                      <button
                        onClick={() => handleDelete(user.id)}
                        className="text-red-500 hover:text-red-700 font-bold text-xs uppercase tracking-wider transition-colors cursor-pointer"
                      >
                        Hapus
                      </button>
                    </div>
                  </td>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {users.length === 0 && (
        <div className="py-12 text-center">
          <div className="inline-flex items-center justify-center w-12 h-12 rounded-full bg-slate-50 mb-3">
            <span className="text-slate-300 text-xl">👤</span>
          </div>
          <p className="text-slate-400 text-sm font-medium">
            Belum ada data user tersedia.
          </p>
        </div>
      )}
    </div>
  );
};
