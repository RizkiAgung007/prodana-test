import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getUserRole } from "../../services/utils";
import api from "../../services/api";

export const useDashboard = () => {
  const [users, setUsers] = useState([]);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
    role_id: 3,
  });
  const [isEditing, setIsEditing] = useState(false);
  const [selectedId, setSelectedId] = useState(null);
  const navigate = useNavigate();
  const roleId = getUserRole();

  const handleLogout = () => {
    sessionStorage.removeItem("token");
    navigate("/login");
  };

  const fetchUsers = async () => {
    try {
      const response = await api.get("/users");
      setUsers(response.data.data);
    } catch (err) {
      setError("GAgal mengambil data user...");

      if (err.response?.status === 401) {
        handleLogout();
      }
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      if (isEditing) {
        await api.put(`/users?id=${selectedId}`, formData);
        alert("User berhasil diupdate");
      } else {
        await api.post("/users", formData);
        alert("User berhasil ditambah");
      }

      setFormData({ name: "", email: "", password: "", role_id: 3 });
      setIsEditing(false);
      fetchUsers();
    } catch (err) {
      alert(err.response?.data?.message || "Gagal memproses data user");
    }
  };

  const handleEditClick = (user) => {
    setIsEditing(true);
    setSelectedId(user.id);

    let userRoleId = 3;
    if (user.role === "Admin") userRoleId = 1;
    if (user.role === "Editor") userRoleId = 2;

    setFormData({
      name: user.name,
      email: user.email,
      password: "",
      role_id: userRoleId,
    });
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Apakah anda yakin ingin menghapus user ini?")) return;

    try {
      await api.delete(`/users?id=${id}`);
      fetchUsers();
    } catch (err) {
      alert(err.response?.data?.message || "Gagal menghapus user");
    }
  };

  return {
    users,
    error,
    roleId,
    formData,
    setFormData,
    isEditing,
    setIsEditing,
    handleLogout,
    handleDelete,
    handleSubmit,
    handleEditClick,
  };
};
