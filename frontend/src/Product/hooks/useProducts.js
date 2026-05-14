import { useEffect, useState } from "react";
import { getUserRole } from "../../services/utils";
import api from "../../services/api";
import { useNavigate } from "react-router-dom";

export const useProducts = () => {
  const [products, setProducts] = useState([]);
  const [name, setName] = useState("");
  const [price, setPrice] = useState("");
  const [description, setDescription] = useState("");
  const [stock, setStock] = useState("");
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [isAILoading, setIsAILoading] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [selectedId, setSelectedId] = useState(null);

  const roleId = getUserRole();
  const navigate = useNavigate();

  const fetchProducts = async () => {
    try {
      const response = await api.get("/products");
      setProducts(response.data.data || []);
    } catch (err) {
      setError("Gagal mengambil data produk");
      if (err.response?.status === 401) handleLogout();
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);

    try {
      const payload = {
        name: name,
        description: description,
        stock: Number(stock),
        price: Number(price),
      };

      if (isEditing) {
        await api.put(`/products?id=${selectedId}`, payload);
        alert("Produk berhasil diperbarui!");
      } else {
        await api.post("/products", payload);
        alert("Produk berhasil dibuat!");
      }

      setName("");
      setPrice("");
      setStock("");
      setDescription("");
      setIsEditing(false);
      setSelectedId(null);
      fetchProducts();
    } catch (err) {
      alert(err.response?.data?.message || "Gagal memproses produk");
    } finally {
      setIsLoading(false);
    }
  };

  const handleEditClick = (product) => {
    setIsEditing(true);
    setSelectedId(product.id);
    setName(product.name);
    setPrice(product.price);
    setStock(product.stock);
    setDescription(product.description || "");
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Apakah anda yakin ingin menghapus produk ini?"))
      return;
    try {
      await api.delete(`/products?id=${id}`);
      fetchProducts();
    } catch (err) {
      alert(err.response?.data?.message || "Gagal menghapus produk");
    }
  };

  const handleLogout = () => {
    sessionStorage.removeItem("token");
    navigate("/login");
  };

  const handleAIGenerate = async () => {
    if (!name) {
      alert(
        "Silakan isi Nama Produk terlebih dahulu agar AI tahu apa yang harus dideskripsikan!",
      );
      return;
    }

    setIsAILoading(true);
    try {
      const response = await api.post("/generate-desc", {
        product_name: name,
      });

      setDescription(response.data.result);
    } catch (err) {
      alert(
        "Gagal menghubungi AI. Pastikan server Backend dan koneksi internet stabil.",
      );
    } finally {
      setIsAILoading(false);
    }
  };

  return {
    products,
    error,
    roleId,
    name,
    setName,
    price,
    setPrice,
    stock,
    setStock,
    description,
    setDescription, 
    isLoading,
    isAILoading, 
    isEditing,
    setIsEditing,
    handleSubmit,
    handleEditClick,
    handleDelete,
    handleLogout,
    handleAIGenerate, 
  };
};
