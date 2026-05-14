import { useProducts } from "../hooks/useProducts";
import { ProductForm } from "../components/ProductForm";
import { ProductTable } from "../components/ProductTable";
import { Header } from "../../components/Header";
import { Footer } from "../../components/Footer";

export const Products = () => {
  const {
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
    isAILoading,
    isLoading,
    isEditing,
    setIsEditing,
    handleSubmit,
    handleEditClick,
    handleDelete,
    handleLogout,
    handleAIGenerate,
  } = useProducts();

  return (
    <div className="min-h-screen flex flex-col bg-[#F8FAFC] p-4 md:p-8 pb-32 md:pb-24">
      <div className="mx-auto container">
        <Header roleId={roleId} handleLogout={handleLogout} />

        {error && (
          <div className="mb-6 p-4 bg-red-50 border border-red-100 text-red-600 rounded-lg text-sm">
            {error}
          </div>
        )}

        {roleId === 1 && (
          <ProductForm
            name={name}
            setName={setName}
            price={price}
            setPrice={setPrice}
            stock={stock}
            setStock={setStock}
            description={description}
            setDescription={setDescription}
            isEditing={isEditing}
            setIsEditing={setIsEditing}
            isLoading={isLoading}
            isAILoading={isAILoading}
            handleSubmit={handleSubmit}
            handleAIGenerate={handleAIGenerate}
          />
        )}

        <ProductTable
          products={products}
          roleId={roleId}
          handleEditClick={handleEditClick}
          handleDelete={handleDelete}
        />

        <Footer />
      </div>
    </div>
  );
};
