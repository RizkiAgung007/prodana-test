import { Footer } from "../../components/Footer";
import { Header } from "../../components/Header";
import { UserForm } from "../components/UserForm";
import { UserTable } from "../components/UserTable";
import { useDashboard } from "../hooks/useDashboard";

export const Dashboard = () => {
  const {
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
  } = useDashboard();

  return (
    <div className="min-h-screen flex flex-col bg-[#F8FAFC] p-4 md:p-8 pb-32 md:pb-24">
      <div className="mx-auto container flex-grow">
        <Header roleId={roleId} handleLogout={handleLogout} />

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-100 text-red-600 rounded-lg text-sm">
            {error}
          </div>
        )}

        {roleId !== 3 && (
          <UserForm
            formData={formData}
            setFormData={setFormData}
            isEditing={isEditing}
            setIsEditing={setIsEditing}
            handleSubmit={handleSubmit}
          />
        )}

        <UserTable
          users={users}
          roleId={roleId}
          handleEditClick={handleEditClick}
          handleDelete={handleDelete}
        />

        <Footer />
      </div>
    </div>
  );
};
