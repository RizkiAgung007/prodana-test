export const ProductTable = ({
  products,
  roleId,
  handleEditClick,
  handleDelete,
}) => {
  return (
    <div className="bg-white border border-slate-200 rounded-xl shadow-sm overflow-hidden">
      <div className="overflow-x-auto">
        <table className="w-full text-left text-sm border-collapse">
          <thead>
            <tr className="bg-slate-50/50 border-b border-slate-200">
              <th className="py-3.5 px-6 font-semibold text-slate-600 w-20">
                ID
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 min-w-[180px]">
                Nama Produk
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 w-40">
                Harga
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 text-center w-24">
                Stok
              </th>
              <th className="py-3.5 px-6 font-semibold text-slate-600 min-w-[300px] max-w-md">
                Deskripsi
              </th>
              {roleId === 1 && (
                <th className="py-3.5 px-6 font-semibold text-slate-600 text-right w-40">
                  Aksi
                </th>
              )}
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {products.map((product) => (
              <tr
                key={product.id}
                className="hover:bg-slate-50/80 transition-colors align-top" 
              >
                <td className="py-4 px-6 text-slate-500 font-mono text-xs">
                  #{product.id}
                </td>
                <td className="py-4 px-6">
                  <div className="font-semibold text-slate-900 break-words max-w-[200px]">
                    {product.name}
                  </div>
                </td>
                <td className="py-4 px-6 text-slate-700 whitespace-nowrap">
                  <span className="text-slate-400 mr-0.5">Rp</span>
                  <span className="font-medium">
                    {product.price?.toLocaleString("id-ID")}
                  </span>
                </td>
                <td className="py-4 px-6 text-center">
                  <span
                    className={`px-2.5 py-1 rounded-full text-xs font-medium ${
                      product.stock > 10
                        ? "bg-emerald-50 text-emerald-700"
                        : "bg-orange-50 text-orange-700"
                    }`}
                  >
                    {product.stock}
                  </span>
                </td>
                <td className="py-4 px-6">
                  <div
                    className="text-slate-500 leading-relaxed line-clamp-3 overflow-hidden"
                    title={product.description}
                  >
                    {product.description || (
                      <span className="text-slate-300 italic">
                        Tidak ada deskripsi
                      </span>
                    )}
                  </div>
                </td>

                {roleId === 1 && (
                  <td className="py-4 px-6 text-right whitespace-nowrap">
                    <div className="flex justify-end gap-4">
                      <button
                        onClick={() => handleEditClick(product)}
                        className="text-indigo-600 hover:text-indigo-800 font-bold text-xs uppercase tracking-wider transition-colors cursor-pointer"
                      >
                        Edit
                      </button>
                      <button
                        onClick={() => handleDelete(product.id)}
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

      {products.length === 0 && (
        <div className="py-12 text-center">
          <div className="inline-flex items-center justify-center w-12 h-12 rounded-full bg-slate-50 mb-3">
            <span className="text-slate-300 text-xl">📦</span>
          </div>
          <p className="text-slate-400 text-sm font-medium">
            Belum ada data produk tersedia.
          </p>
        </div>
      )}
    </div>
  );
};
