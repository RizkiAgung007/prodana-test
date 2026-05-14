export const UserForm = ({
  formData,
  setFormData,
  isEditing,
  setIsEditing,
  handleSubmit,
}) => {
  return (
    <div className="mb-8 p-6 bg-white border border-slate-200 rounded-xl shadow-sm">
      <div className="mb-5">
        <h2 className="text-base font-semibold text-slate-900">
          {isEditing ? "Edit Data User" : "Tambah User Baru"}
        </h2>
        <p className="text-sm text-slate-500 mt-1">
          {isEditing
            ? "Perbarui informasi akun di bawah ini."
            : "Masukkan detail untuk membuat akun baru."}
        </p>
      </div>

      <form
        onSubmit={handleSubmit}
        className="grid grid-cols-1 lg:grid-cols-5 gap-4 items-end"
      >
        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-slate-700">
            Nama
          </label>
          <input
            type="text"
            required
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
          />
        </div>

        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-slate-700">
            Email
          </label>
          <input
            type="email"
            required
            value={formData.email}
            onChange={(e) =>
              setFormData({ ...formData, email: e.target.value })
            }
            className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
          />
        </div>

        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-slate-700">
            {isEditing ? "Password Baru" : "Password"}
          </label>
          <input
            type="password"
            required={!isEditing}
            value={formData.password}
            onChange={(e) =>
              setFormData({ ...formData, password: e.target.value })
            }
            placeholder={isEditing ? "Abaikan jika tetap" : ""}
            className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm placeholder:text-slate-400"
          />
        </div>

        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-slate-700">
            Role
          </label>
          <select
            value={formData.role_id}
            onChange={(e) =>
              setFormData({ ...formData, role_id: parseInt(e.target.value) })
            }
            className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
          >
            <option value={1}>Admin</option>
            <option value={2}>Editor</option>
            <option value={3}>Viewer</option>
          </select>
        </div>

        <div className="flex gap-2 lg:col-span-1">
          <button
            type="submit"
            className="w-full cursor-pointer bg-slate-900 hover:bg-slate-800 text-white py-2 px-4 rounded-lg text-sm font-medium transition-colors"
          >
            {isEditing ? "Update" : "Simpan"}
          </button>

          {isEditing && (
            <button
              type="button"
              onClick={() => {
                setIsEditing(false);
                setFormData({ name: "", email: "", password: "", role_id: 3 });
              }}
              className="w-full bg-white border border-slate-200 hover:bg-slate-50 text-slate-700 py-2 px-4 rounded-lg text-sm font-medium transition-colors"
            >
              Batal
            </button>
          )}
        </div>
      </form>
    </div>
  );
};