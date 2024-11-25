import axiosInstance from './axiosInstance';

export const getBookings = async () => {
  const response = await axiosInstance.get('/bookings/');
  return response.data;
};

export const getBookingById = async (id: string) => {
  const response = await axiosInstance.get(`/bookings/${id}`);
  return response.data;
};

export const createBooking = async (bookingData: any) => {
  const response = await axiosInstance.post('/bookings/', bookingData);
  return response.data;
};

export const updateBooking = async (id: string, updatedData: any) => {
  const response = await axiosInstance.put(`/bookings/${id}`, updatedData);
  return response.data;
};

export const deleteBooking = async (id: string) => {
  const response = await axiosInstance.delete(`/bookings/${id}`);
  return response.data;
};
