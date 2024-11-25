import axiosInstance from './axiosInstance';

export const adminGetBookings = async () => {
  const response = await axiosInstance.get('/admin/bookings');
  return response.data;
};

export const adminCreateBooking = async (bookingData: any) => {
  const response = await axiosInstance.post('/admin/bookings', bookingData);
  return response.data;
};

export const adminUpdateBooking = async (id: string, updatedData: any) => {
  const response = await axiosInstance.put(`/admin/bookings/${id}`, updatedData);
  return response.data;
};

export const adminDeleteBooking = async (id: string) => {
  const response = await axiosInstance.delete(`/admin/bookings/${id}`);
  return response.data;
};
