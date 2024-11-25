import axiosInstance from './axiosInstance';

export const createPayment = async (amount: number) => {
  const response = await axiosInstance.post('/payment/create', null, {
    params: { amount },
  });
  return response.data;
};
