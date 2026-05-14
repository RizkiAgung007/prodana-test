export const ProductForm = ({
  name,
  setName,
  price,
  setPrice,
  stock,
  setStock,
  description,
  setDescription,
  isEditing,
  setIsEditing,
  isLoading,
  isAILoading,
  handleSubmit,
  handleAIGenerate,
}) => {
  return (
    <div className="mb-8 p-6 bg-white border border-slate-200 rounded-xl shadow-sm">
      <div className="mb-5">
        <h2 className="text-base font-semibold text-slate-900">
          {isEditing ? "Edit Data Produk" : "Tambah Produk Baru"}
        </h2>
        <p className="text-sm text-slate-500 mt-1">
          {isEditing
            ? "Perbarui informasi produk di bawah ini."
            : "Masukkan detail untuk menambahkan produk ke katalog."}
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-slate-700">
              Nama Produk
            </label>
            <input
              type="text"
              required
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
              placeholder="Ex: Laptop Gaming"
            />
          </div>

          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-slate-700">
              Harga (Rp)
            </label>
            <input
              type="number"
              required
              value={price}
              onChange={(e) => setPrice(e.target.value)}
              className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
              placeholder="Ex: 15000000"
            />
          </div>

          <div className="space-y-1.5">
            <label className="block text-sm font-medium text-slate-700">
              Stok
            </label>
            <input
              type="number"
              required
              value={stock}
              onChange={(e) => setStock(e.target.value)}
              className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm"
              placeholder="Ex: 50"
            />
          </div>
        </div>

        <div className="space-y-1.5 relative">
          <div className="flex justify-between items-end mb-1">
            <label className="block text-sm font-medium text-slate-700">
              Deskripsi Produk
            </label>
            <button
              type="button"
              onClick={handleAIGenerate}
              disabled={isAILoading}
              className="cursor-pointer text-xs font-medium text-indigo-600 bg-indigo-50 hover:bg-indigo-100 px-3 py-1.5 rounded-md transition-colors disabled:opacity-50 flex items-center gap-1"
            >
              {isAILoading ? "Memproses AI..." : "Generate AI"}
            </button>
          </div>
          <textarea
            required
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows="3"
            className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-lg focus:bg-white focus:ring-2 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all outline-none text-sm resize-none"
            placeholder="Tulis deskripsi atau gunakan bantuan AI..."
          ></textarea>
        </div>

        <div className="flex gap-2 pt-2 border-t border-slate-100">
          <button
            type="submit"
            disabled={isLoading}
            className="cursor-pointer bg-slate-900 hover:bg-slate-800 text-white py-2 px-6 rounded-lg text-sm font-medium transition-colors disabled:bg-slate-400"
          >
            {isLoading
              ? "Menyimpan..."
              : isEditing
                ? "Update Produk"
                : "Simpan Produk"}
          </button>

          {isEditing && (
            <button
              type="button"
              onClick={() => {
                setIsEditing(false);
                setName("");
                setPrice("");
                setStock("");
                setDescription("");
              }}
              className="bg-white border border-slate-200 hover:bg-slate-50 text-slate-700 py-2 px-6 rounded-lg text-sm font-medium transition-colors"
            >
              Batal
            </button>
          )}
        </div>
      </form>
    </div>
  );
};
