FROM node:21.7.3-alpine AS build-stage
WORKDIR /app

# Copy package files
COPY package*.json ./

# Clean install dependencies
RUN rm -rf node_modules && \
    npm cache clean --force && \
    npm install

# Copy the rest of the application
COPY . .

# Ensure vite is in PATH
ENV PATH /app/node_modules/.bin:$PATH

# Set permissions if needed
RUN chmod +x node_modules/.bin/vite

# Expose the development server port
EXPOSE 3000

# Use 0.0.0.0 to make the server accessible outside the container
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

#-------------------- UNCOMMMENT FOR PRODUCTION -------------------------
# # Build the application for production
# RUN npm run build

# # Production stage
# FROM nginx:alpine AS production-stage

# # Copy the built files from the build stage to Nginx
# COPY --from=build-stage /app/dist /usr/share/nginx/html

# # Expose port 80 for the frontend
# EXPOSE 3000

# # Start Nginx
# CMD ["nginx", "-g", "daemon off;"]
