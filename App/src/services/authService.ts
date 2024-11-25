import axiosInstance from './axiosInstance';

export const signUp = async (userData: { username: string; password: string }) => {
  const response = await axiosInstance.post('/auth/signup', userData);
  return response.data;
};

export const login = async (credentials: { username: string; password: string }) => {
  const response = await axiosInstance.post('/auth/login', credentials);
  return response.data;
};
