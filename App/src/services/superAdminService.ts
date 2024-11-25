import axiosInstance from './axiosInstance';

export const createUser = async (userData: any) => {
  const response = await axiosInstance.post('/super-admin/users', userData);
  return response.data;
};

export const getUsers = async () => {
  const response = await axiosInstance.get('/super-admin/users');
  return response.data;
};

export const getUserById = async (id: string) => {
  const response = await axiosInstance.get(`/super-admin/users/${id}`);
  return response.data;
};

export const updateUser = async (id: string, updatedData: any) => {
  const response = await axiosInstance.put(`/super-admin/users/${id}`, updatedData);
  return response.data;
};

export const deleteUser = async (id: string) => {
  const response = await axiosInstance.delete(`/super-admin/users/${id}`);
  return response.data;
};
